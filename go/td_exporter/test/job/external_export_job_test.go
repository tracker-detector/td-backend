package job_test

import (
	"fmt"
	"io"
	"os"

	"testing"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/mocks"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/job"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestExternalExportJob(t *testing.T) {
	suite.Run(t, &ExternalExportJobTest{})
}

type ExternalExportJobTest struct {
	suite.Suite
	internalJob    *job.ExternalExportJob
	requestRepo    *mocks.RequestRepository
	storageService *mocks.IStorageService
}

func (suite *ExternalExportJobTest) SetupTest() {
	suite.requestRepo = new(mocks.RequestRepository)
	suite.storageService = new(mocks.IStorageService)
	suite.internalJob = job.NewExternalExportJob(suite.requestRepo, suite.storageService)
}

func (suite *ExternalExportJobTest) TestExecute_Success() {
	// given
	requests := testsupport.LoadRequestJson()
	extractorFilePath := "./testdata/exporter/exporter204.js"
	file, _ := os.Open(extractorFilePath)
	exporter := &model.Exporter{
		ID:                   "someId1",
		Name:                 "JsExtractor204",
		Description:          "someDescription",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: &extractorFilePath,
	}
	suite.Len(requests, 10)
	suite.requestRepo.On("StreamByDataset", mock.Anything, "").Return(testsupport.CreateResultsChannel(requests), testsupport.CreateErrorChannel([]error{}))
	suite.storageService.On("PutObject", mock.Anything, "", "JsExtractor204_EasyPrivacy_.csv.gz", mock.Anything, int64(-1), "application/gzip").Run(func(args mock.Arguments) {
		fileName := args.Get(2).(string)
		pr := args.Get(3).(io.Reader)
		file, _ := os.Create(fileName)
		io.Copy(file, pr)
	}).Return(nil)
	suite.storageService.On("GetObject", mock.Anything, "", extractorFilePath).Return(file, nil)

	// when
	metrics := suite.internalJob.Execute(exporter, "EasyPrivacy", "")
	// then
	fmt.Println(metrics.Error)
	suite.Assertions.Equal(10, metrics.Total)
	suite.Assertions.Equal(9, metrics.NonTracker)
	suite.Assertions.Equal(1, metrics.Tracker)
	expectedCsv := testsupport.LoadFile("./testdata/requests/expected_encoding.csv")
	actualCsv := testsupport.LoadGzFile("./JsExtractor204_EasyPrivacy_.csv.gz")
	suite.Assertions.Equal(expectedCsv, actualCsv)
	os.Remove("./JsExtractor204_EasyPrivacy_.csv.gz")
}
