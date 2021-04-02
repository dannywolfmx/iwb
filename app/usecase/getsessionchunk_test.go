package usecase

import (
	"testing"

	"github.com/dannywolfmx/iwb/app/domain/entity"
	"github.com/dannywolfmx/iwb/app/domain/service"
	mock_app "github.com/dannywolfmx/iwb/app/repository/mock"
	"github.com/golang/mock/gomock"
)

// TODO
// 	*
//

func TestGetSessionChunk(t *testing.T) {
	//TEST TO VALIDATE POSITION IN RANGE
	//Arrange
	//
	servPosition := service.PositionService{}
	testUserPosition := entity.Position{
		X: service.MaxUserPositionX,
		Y: service.MaxUserPositionY,
	}

	testSession := entity.NewSession("test", "123456")

	testChunkPosition := servPosition.CalculateChunkPosition(testUserPosition)
	testSession.UserPosition = testUserPosition

	testChunk := entity.NewChunk()

	//MOCK REPOSITORY
	c := gomock.NewController(t)
	defer c.Finish()
	m := mock_app.NewMockChunkRepository(c)

	//testPosition -> (chunk: testchunk, error: nil)
	m.EXPECT().Get(gomock.Eq(testChunkPosition)).Return(testChunk, nil).MaxTimes(1)

	//Act
	getSessionChunk := NewGetSessionChunk(m, servPosition)
	chunk, err := getSessionChunk.Execute(testSession)
	//Assert

	if err != nil {
		t.Fatalf("Error on get session chunk execute: %s", err)
	}

	if chunk == nil {
		t.Fatal("Chunk nil on get session chunk execute")
	}

	//TEST TO VALIDATE POSITION OUT OF RANGE

	//Arrange
	//
	testUserPosition = entity.Position{
		X: service.MaxUserPositionX + 1,
		Y: service.MaxUserPositionY + 1,
	}

	testSession = entity.NewSession("test", "123456")

	testChunkPosition = servPosition.CalculateChunkPosition(testUserPosition)
	testSession.UserPosition = testUserPosition

	testChunk = entity.NewChunk()

	//testPosition -> (chunk: testchunk, error: nil)
	m.EXPECT().Get(gomock.Eq(testChunkPosition)).Return(testChunk, nil).MaxTimes(0)

	//Act
	getSessionChunk = NewGetSessionChunk(m, servPosition)
	chunk, err = getSessionChunk.Execute(testSession)
	//Assert

	if err == nil {
		t.Fatal("Expect an error from the usecase")
	}

	if chunk != nil {
		t.Fatal("Chunk need to be nil on get session chunk execute")
	}
}
