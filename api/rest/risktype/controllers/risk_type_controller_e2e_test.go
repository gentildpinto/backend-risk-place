package risktype_controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	risk_type_controller "github.com/risk-place-angola/backend-risk-place/api/rest/risktype/controllers"
	"github.com/risk-place-angola/backend-risk-place/domain/entities"
	"github.com/risk-place-angola/backend-risk-place/domain/repository/mocks"
	"github.com/risk-place-angola/backend-risk-place/usecase/risktype"
	"github.com/stretchr/testify/assert"
)

func TestRiskTypeControllers(t *testing.T) {
	t.Run("should return 201 when create a risk type", func(t *testing.T) {

		e := echo.New()

		data := entities.RiskType{
			ID:          "0c1baa42-3909-4bdb-837f-a80e68232ecd",
			Name:        "Assalto",
			Description: "Assalto a mão armada",
		}

		jsonData, _ := json.Marshal(data)

		res := httptest.NewRequest("POST", "/api/v1/risk/risktype", bytes.NewBuffer(jsonData))
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRiskTypeRepository := mocks.NewMockRiskTypeRepository(ctrl)
		mockRiskTypeRepository.EXPECT().Save(gomock.Any()).Return(nil)

		riskTypeUseCase := risktype.NewRiskTypeUseCase(mockRiskTypeRepository)
		riskTypeController := risk_type_controller.NewRiskTypeController(riskTypeUseCase)

		if assert.NoError(t, riskTypeController.RiskTypeCreateController(ctx)) {
			assert.Equal(t, http.StatusCreated, rec.Code, "error status code != 201")
		}
	})

	t.Run("should return 200 when update a risk type", func(t *testing.T) {
		e := echo.New()

		data := &entities.RiskType{
			ID:          "0c1baa42-3909-4bdb-837f-a80e68232ecd",
			Name:        "Assaltam",
			Description: "Assalto a mão armada",
		}

		jsonData, _ := json.Marshal(data)

		res := httptest.NewRequest("PUT", "/api/v1/risk/risktype/:id", bytes.NewBuffer(jsonData))
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("0c1baa42-3909-4bdb-837f-a80e68232ecd")
		ctx.SetPath("/api/v1/risk/risktype/:id")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRiskTypeRepository := mocks.NewMockRiskTypeRepository(ctrl)
		mockRiskTypeRepository.EXPECT().FindByID(gomock.Any()).Return(data, nil)
		mockRiskTypeRepository.EXPECT().Update(gomock.Any()).Return(nil)

		riskTypeUseCase := risktype.NewRiskTypeUseCase(mockRiskTypeRepository)
		riskTypeController := risk_type_controller.NewRiskTypeController(riskTypeUseCase)

		if assert.NoError(t, riskTypeController.RiskTypeUpdateController(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code, "error status code != 200")
		}

	})

	t.Run("should return 200 when delete a risk type", func(t *testing.T) {
		e := echo.New()

		data := &entities.RiskType{
			ID:          "0c1baa42-3909-4bdb-837f-a80e68232ecd",
			Name:        "Assaltam",
			Description: "Assalto a mão armada",
		}

		res := httptest.NewRequest("DELETE", "/api/v1/risk/risktype/:id", nil)
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("0c1baa42-3909-4bdb-837f-a80e68232ecd")
		ctx.SetPath("/api/v1/risk/risktype/:id")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRiskTypeRepository := mocks.NewMockRiskTypeRepository(ctrl)
		mockRiskTypeRepository.EXPECT().FindByID(gomock.Any()).Return(data, nil)
		mockRiskTypeRepository.EXPECT().Delete(gomock.Any()).Return(nil)

		riskTypeUseCase := risktype.NewRiskTypeUseCase(mockRiskTypeRepository)
		riskTypeController := risk_type_controller.NewRiskTypeController(riskTypeUseCase)

		if assert.NoError(t, riskTypeController.RiskTypeDeleteController(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code, "error status code != 200")
		}

	})

	t.Run("should return 200 when find a risk type by id", func(t *testing.T) {
		e := echo.New()

		data := &entities.RiskType{
			ID:          "0c1baa42-3909-4bdb-837f-a80e68232ecd",
			Name:        "Assaltam",
			Description: "Assalto a mão armada",
		}

		res := httptest.NewRequest("GET", "/api/v1/risk/risktype/:id", nil)
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)
		ctx.SetParamNames("id")
		ctx.SetParamValues("0c1baa42-3909-4bdb-837f-a80e68232ecd")
		ctx.SetPath("/api/v1/risk/risktype/:id")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRiskTypeRepository := mocks.NewMockRiskTypeRepository(ctrl)
		mockRiskTypeRepository.EXPECT().FindByID(gomock.Any()).Return(data, nil)

		riskTypeUseCase := risktype.NewRiskTypeUseCase(mockRiskTypeRepository)
		riskTypeController := risk_type_controller.NewRiskTypeController(riskTypeUseCase)

		if assert.NoError(t, riskTypeController.RiskTypeFindByIdController(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code, "error status code != 200")
		}

	})

	t.Run("should return 200 when find all risk types", func(t *testing.T) {
		e := echo.New()

		data := []*entities.RiskType{
			{
				ID:          "0c1baa42-3909-4bdb-837f-a80e68232ecd",
				Name:        "Assaltam",
				Description: "Assalto a mão armada",
			},
			{
				ID:          "hsy36q83-4009-1bcc-837f-a80937253648",
				Name:        "Inundação",
				Description: "Inundação devido a chuvas",
			},
		}

		res := httptest.NewRequest("GET", "/api/v1/risk/risktype", nil)
		res.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		ctx := e.NewContext(res, rec)
		ctx.SetPath("/api/v1/risk/risktype")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRiskTypeRepository := mocks.NewMockRiskTypeRepository(ctrl)
		mockRiskTypeRepository.EXPECT().FindAll().Return(data, nil)

		riskTypeUseCase := risktype.NewRiskTypeUseCase(mockRiskTypeRepository)
		riskTypeController := risk_type_controller.NewRiskTypeController(riskTypeUseCase)

		if assert.NoError(t, riskTypeController.RiskTypeFindAllController(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code, "error status code != 200")
		}

	})

}
