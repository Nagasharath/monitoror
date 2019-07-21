package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/monitoror/monitoror/models/tiles"
	. "github.com/monitoror/monitoror/monitorable/port"
	"github.com/monitoror/monitoror/monitorable/port/mocks"
	"github.com/monitoror/monitoror/monitorable/port/models"

	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
)

func TestUsecase_CheckPort_Success(t *testing.T) {
	// Init
	mockRepo := new(mocks.Repository)
	mockRepo.On("OpenSocket", AnythingOfType("string"), AnythingOfType("int")).Return(nil)
	usecase := NewPortUsecase(mockRepo)

	// Params
	param := &models.PortParams{
		Hostname: "test.com",
		Port:     1234,
	}

	// Expected
	eTile := tiles.NewHealthTile(PortTileType)
	eTile.Label = fmt.Sprintf("%s:%d", param.Hostname, param.Port)
	eTile.Status = tiles.SuccessStatus

	// Test
	rTile, err := usecase.Port(param)

	if assert.NoError(t, err) {
		assert.Equal(t, eTile, rTile)
		mockRepo.AssertNumberOfCalls(t, "OpenSocket", 1)
		mockRepo.AssertExpectations(t)
	}
}

func TestUsecase_CheckPort_Fail(t *testing.T) {
	// Init
	mockRepo := new(mocks.Repository)
	mockRepo.On("OpenSocket", AnythingOfType("string"), AnythingOfType("int")).Return(errors.New("port error"))
	usecase := NewPortUsecase(mockRepo)

	// Params
	param := &models.PortParams{
		Hostname: "test.com",
		Port:     1234,
	}

	// Expected
	eTile := tiles.NewHealthTile(PortTileType)
	eTile.Label = fmt.Sprintf("%s:%d", param.Hostname, param.Port)
	eTile.Status = tiles.FailedStatus

	// Test
	rTile, err := usecase.Port(param)

	if assert.NoError(t, err) {
		assert.Equal(t, eTile, rTile)
		mockRepo.AssertNumberOfCalls(t, "OpenSocket", 1)
		mockRepo.AssertExpectations(t)
	}
}