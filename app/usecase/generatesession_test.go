package usecase

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	mock_app "github.com/dannywolfmx/iwb/app/repository/mock"
	"github.com/golang/mock/gomock"
)

func TestGenerateSession(t *testing.T) {

	//Setup test variables
	testToken := "123456"
	testGenerateTokenFunction := func(user *entity.User) (string, error) {
		return testToken, nil
	}

	testValidateUser := func(user *entity.User) bool {
		return true
	}
	testUser := &entity.User{Name: "Test"}
	testSession := entity.NewSession(testUser.Name, testToken)

	//Setup mock
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	repo := mock_app.NewMockSessionRepository(ctrl)

	//on save check if the passed session is the same as testSession
	repo.EXPECT().Save(gomock.Eq(testSession)).Return(nil)

	//Run the usecase test
	service := NewGenerateSession(repo, testGenerateTokenFunction, testValidateUser)

	if _, err := service.Execute(testUser); err != nil {
		t.Fatalf("Error on execute service: %s", err)
	}
}
