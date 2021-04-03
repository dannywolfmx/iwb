package usecase

import (
	"testing"
	"time"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	mock_app "github.com/dannywolfmx/iwb/app/repository/mock"
	"github.com/golang/mock/gomock"
)

// TODO
// Get a session and update them last chunk update attribute
// Validad failure cases
func TestUpdateLastChunkUpdate(t *testing.T) {
	//Arrage
	testTimeNow := time.Now()
	testSession := entity.NewSession("test", "12345")
	testReturnedSession := &entity.Session{
		UserPosition:    testSession.UserPosition,
		Account:         testSession.Account,
		LastChunkUpdate: testTimeNow,
		Token:           testSession.Token,
	}
	//Act
	c := gomock.NewController(t)
	defer c.Finish()

	m := mock_app.NewMockSessionRepository(c)

	//testSession -> (error:nil)
	m.EXPECT().Update(gomock.Eq(testSession)).Return(nil)
	updateLastChunk := NewUpdateLastChunkUpdate(m)
	//Assert
	//
	session, err := updateLastChunk.Execute(testSession, testTimeNow)

	if err != nil {
		t.Fatalf("An error when tried to get an updated session %s", err)
	}

	if *session != *testReturnedSession {
		t.Fatalf("Unexpected returnded session: %+v, expected: %+v", session, testReturnedSession)
	}
}
