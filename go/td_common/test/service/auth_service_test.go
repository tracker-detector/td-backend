package service_test

import (
	"context"

	"testing"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/mocks"
	"github.com/stretchr/testify/suite"
)

func TestAuthService(t *testing.T) {
	suite.Run(t, &AuthServiceTest{})
}

type AuthServiceTest struct {
	suite.Suite
	authService       *service.AuthService
	userRepo          *mocks.UserRepository
	encrytpionService *mocks.IEncryptionService
}

func (suite *AuthServiceTest) SetupTest() {
	suite.userRepo = new(mocks.UserRepository)
	suite.encrytpionService = new(mocks.IEncryptionService)
	suite.authService = service.NewAuthService(suite.userRepo, suite.encrytpionService)
}

func (suite *AuthServiceTest) TestShouldValidateBearerToken() {
	// given
	ctx := context.Background()
	key := "someKey"
	token := "Bearer someBeaererToken"
	tokenValue := "someBeaererToken"
	user := &model.UserData{
		ID:    "someId",
		Role:  model.ADMIN,
		Email: "test@test.com",
		Key:   key,
	}
	suite.userRepo.On("FindAll", ctx).Return([]*model.UserData{user}, nil)
	suite.encrytpionService.On("CompareHashAndPassword", key, tokenValue).Return(true)

	// when
	valid, err := suite.authService.ValidateBearerToken(ctx, token, true)

	// then
	suite.True(valid, "The auth service should allow access")
	suite.Nil(err)
	suite.True(suite.encrytpionService.AssertNumberOfCalls(suite.T(), "CompareHashAndPassword", 1))
	suite.True(suite.userRepo.AssertNumberOfCalls(suite.T(), "FindAll", 1))
}

func (suite *AuthServiceTest) TestShouldInvalidateBearerTokenRouteOnlyForAdmins() {
	// given
	ctx := context.Background()
	key := "someKey"
	token := "Bearer someBeaererToken"
	tokenValue := "someBeaererToken"
	user := &model.UserData{
		ID:    "someId",
		Role:  model.CLIENT,
		Email: "test@test.com",
		Key:   key,
	}
	suite.userRepo.On("FindAll", ctx).Return([]*model.UserData{user}, nil)
	suite.encrytpionService.On("CompareHashAndPassword", key, tokenValue).Return(true)

	// when
	valid, err := suite.authService.ValidateBearerToken(ctx, token, true)

	// then
	suite.False(valid, "The auth service should not allow access")
	suite.Nil(err)
	suite.True(suite.encrytpionService.AssertNumberOfCalls(suite.T(), "CompareHashAndPassword", 1))
	suite.True(suite.userRepo.AssertNumberOfCalls(suite.T(), "FindAll", 1))
}

func (suite *AuthServiceTest) TestShouldInvalidateBearerTokenWhenInWrongFormat() {
	// given
	ctx := context.Background()
	token := "asd someBeaererToken"

	// when
	valid, err := suite.authService.ValidateBearerToken(ctx, token, true)

	// then
	suite.False(valid, "The auth service should not allow access")
	suite.NotNil(err)
	suite.True(suite.encrytpionService.AssertNumberOfCalls(suite.T(), "CompareHashAndPassword", 0))
	suite.True(suite.userRepo.AssertNumberOfCalls(suite.T(), "FindAll", 0))
}

func (suite *AuthServiceTest) TestShouldInvalidateBearerTokenWhenCredentialsWrong() {
	// given
	ctx := context.Background()
	key := "someKey"
	token := "Bearer someBeaererToken"
	tokenValue := "someBeaererToken"
	user := &model.UserData{
		ID:    "someId",
		Role:  model.ADMIN,
		Email: "test@test.com",
		Key:   key,
	}
	suite.userRepo.On("FindAll", ctx).Return([]*model.UserData{user}, nil)
	suite.encrytpionService.On("CompareHashAndPassword", key, tokenValue).Return(false)

	// when
	valid, err := suite.authService.ValidateBearerToken(ctx, token, true)

	// then
	suite.False(valid, "The auth service should not allow access")
	suite.Nil(err)
	suite.True(suite.encrytpionService.AssertNumberOfCalls(suite.T(), "CompareHashAndPassword", 1))
	suite.True(suite.userRepo.AssertNumberOfCalls(suite.T(), "FindAll", 1))
}
