package converter

import "github.com/Tracking-Detector/td-backend/go/td_common/model"

type ReduceType string
type ReduceFunction func([]model.RequestDataLabel) bool

const (
	OR           ReduceType = "or"
	AND          ReduceType = "and"
	EASY_PRIVACY ReduceType = "EasyPrivacy"
	EASY_LIST    ReduceType = "EasyList"
	HUMAN        ReduceType = "Human"
)

func IsValidReduceType(input string) bool { // TODO make this some kind if service
	validReduceTypes := []ReduceType{OR, AND, EASY_PRIVACY, EASY_LIST, HUMAN}
	for _, rt := range validReduceTypes {
		if ReduceType(input) == rt {
			return true
		}
	}
	return false
}

func orFunction(labels []model.RequestDataLabel) bool {
	isTracking := false
	for _, label := range labels {
		isTracking = isTracking || label.IsLabeled
	}
	if isTracking {
		return true
	}
	return false
}

func andFunction(labels []model.RequestDataLabel) bool {
	isTracking := true
	for _, label := range labels {
		isTracking = isTracking && label.IsLabeled
	}
	if isTracking {
		return true
	}
	return false
}

func easyPrivacyFunction(labels []model.RequestDataLabel) bool {
	for _, label := range labels {
		if label.Blocklist == "EasyPrivacy" {
			if label.IsLabeled {
				return true
			} else {
				return false
			}
		}

	}
	return false
}

func easyListFunction(labels []model.RequestDataLabel) bool {
	for _, label := range labels {
		if label.Blocklist == "EasyList" {
			if label.IsLabeled {
				return true
			} else {
				return false
			}
		}

	}
	return false
}

func humanFunction(labels []model.RequestDataLabel) bool {
	for _, label := range labels {
		if label.Blocklist == "Human" {
			if label.IsLabeled {
				return true
			} else {
				return false
			}
		}

	}
	return false
}

func ConvertRequestModel(request *model.RequestData, reducer ReduceType) *model.ReducedRequestData {
	if request == nil {
		return nil
	}
	if len(request.Labels) == 0 {
		return nil
	}

	var reducerFn ReduceFunction
	switch reducer {
	case OR:
		reducerFn = orFunction
	case AND:
		reducerFn = andFunction
	case EASY_PRIVACY:
		reducerFn = easyPrivacyFunction
	case EASY_LIST:
		reducerFn = easyListFunction
	case HUMAN:
		reducerFn = humanFunction
	}

	reducedModel := &model.ReducedRequestData{}
	reducedModel.DocumentId = request.DocumentId
	reducedModel.DocumentLifecycle = request.DocumentLifecycle
	reducedModel.FrameId = request.FrameId
	reducedModel.FrameType = request.FrameType
	reducedModel.Initiator = request.Initiator
	reducedModel.Method = request.Method
	reducedModel.ParentFrameId = request.ParentFrameId
	reducedModel.RequestHeaders = request.RequestHeaders
	reducedModel.RequestId = request.RequestId
	reducedModel.Response = request.Response
	reducedModel.Success = request.Success
	reducedModel.TabId = request.TabId
	reducedModel.TimeStamp = request.TimeStamp
	reducedModel.Type = request.Type
	reducedModel.URL = request.URL
	reducedModel.Tracker = reducerFn(request.Labels)
	return reducedModel
}
