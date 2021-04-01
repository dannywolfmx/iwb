package usecase

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	mock_app "github.com/dannywolfmx/iwb/app/repository/mock"
	"github.com/golang/mock/gomock"
)

func TestGenerateSession(t *testing.T) {

	//Setup
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	testToken := "123456"
	testGenerateTokenFunction := func(user *entity.User) (string, error) {
		return testToken, nil
	}

	testUser := &entity.User{Name: "Test"}
	testSession := entity.NewSession(testUser.Name, testToken)

	repo := mock_app.NewMockSessionRepository(ctrl)

	//on save check if the passed session is the same as testSession
	repo.EXPECT().Save(gomock.Eq(testSession)).Return(nil)

	service := NewGenerateSession(repo, testGenerateTokenFunction)

	if _, err := service.Execute(testUser); err != nil {
		t.Fatalf("Error on execute service: %s", err)
	}
}
