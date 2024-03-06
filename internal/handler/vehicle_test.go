package handler

import (
	"app/internal"
	"app/internal/service"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("success - case01: should returns a list of vehicles and a 200 status code", func(t *testing.T) {
		// Arrange
		carColor := "Orange"
		carFabricationYear := 1995
		foundCar := internal.Vehicle{
			Id: 25,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "Land Rover",
				Model:           "Discovery",
				Registration:    "03178",
				Color:           "Orange",
				FabricationYear: 1995,
				Capacity:        4,
				MaxSpeed:        175,
				FuelType:        "diesel",
				Transmission:    "manual",
				Weight:          293.77,
				Dimensions: internal.Dimensions{
					Height: 47.17,
					Length: 0,
					Width:  198.33,
				},
			},
		}

		expectedResult := map[int]internal.Vehicle{
			foundCar.Id: foundCar,
		}
		sv := service.NewVehicleDefaultMock()
		sv.On("FindByColorAndYear", carColor, carFabricationYear).Return(expectedResult, nil)

		hd := NewHandlerVehicle(sv)
		hdFunc := hd.FindByColorAndYear()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/color/Orange/year/1995", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("color", "Orange")
		chiCtx.URLParams.Add("year", "1995")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{
		   "data": {
		       "25": {
		           "Id": 25,
		           "Brand": "Land Rover",
		           "Model": "Discovery",
		           "Registration": "03178",
		           "Color": "Orange",
		           "FabricationYear": 1995,
		           "Capacity": 4,
		           "MaxSpeed": 175,
		           "FuelType": "diesel",
		           "Transmission": "manual",
		           "Weight": 293.77,
		           "Height": 47.17,
		           "Length": 0,
		           "Width": 198.33
		       }
		   },
		   "message": "vehicles found"
		}`
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})

	t.Run("failure - case01: should returns a not found message and 404 status code", func(t *testing.T) {
		// Arrange
		carColor := "Cyan"
		carFabricationYear := 2023

		expectedResult := map[int]internal.Vehicle{}
		sv := service.NewVehicleDefaultMock()

		sv.On("FindByColorAndYear", carColor, carFabricationYear).Return(expectedResult, internal.ErrRepositoryNotFound)

		hd := NewHandlerVehicle(sv)
		hdFunc := hd.FindByColorAndYear()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/color/Cyan/year/2023", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("color", carColor)
		chiCtx.URLParams.Add("year", "2023")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"status":"Not Found","message":"No vehicles were found for the provided criteria"}`
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		sv.AssertExpectations(t)
	})
}
