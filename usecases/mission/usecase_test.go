package mission

import (
	"context"
	"errors"
	"idnmedia/repositories"
	mission "idnmedia/repositories/mission"
	missionMock "idnmedia/repositories/mission/mocks"
	"idnmedia/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	ctx = utils.SetCtxPLayer(ctx, &utils.CtxPLayer{
		Id:    1,
		Email: "mock@mock.com",
	})
	mock := missionMock.NewRepository(t)
	uc := NewUsecase(mock)
	req := MissionEntity{
		Title:          "test",
		Description:    "description",
		GoldBounty:     10,
		DeadlineSecond: 60,
	}
	expectedRes := MissionEntity{
		Id:             1,
		Title:          "test",
		Description:    "description",
		GoldBounty:     10,
		DeadlineSecond: 60,
	}
	model := mission.MissionModel{
		Title:          "test",
		Description:    "description",
		GoldBounty:     10,
		DeadlineSecond: 60,
		BaseModel: repositories.BaseModel{
			CreatedBy: "mock@mock.com",
			UpdatedBy: "mock@mock.com",
		},
	}
	expectedErr := errors.New("error")
	t.Run("Success Case", func(t *testing.T) {
		mock.On("Create", ctx, model).Return(1, nil).Once()
		res, err := uc.Create(ctx, &req)
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})

	t.Run("Error Case Create DB", func(t *testing.T) {
		mock.On("Create", ctx, model).Return(0, expectedErr).Once()
		_, err := uc.Create(ctx, &req)
		assert.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
