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

func TestHandlerVehicle_FindByBrandAndYearRange(t *testing.T) {
	t.Run("success - case 01: should returns a list with vehicles", func(t *testing.T) {
		// Arrange
		carBrand := "Chevrolet"
		startYear := 1995
		endYear := 1997
		sv := service.NewVehicleDefaultMock()
		expectedResult := map[int]internal.Vehicle{
			11: {
				Id: 11,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Chevrolet",
					Model:           "G-Series 2500",
					Registration:    "9292",
					Color:           "Mauv",
					FabricationYear: 1996,
					Capacity:        3,
					MaxSpeed:        239,
					FuelType:        "gas",
					Transmission:    "manual",
					Weight:          152.87,
					Dimensions: internal.Dimensions{
						Height: 50.84,
						Length: 0,
						Width:  216.53,
					},
				},
			},
			14: {
				Id: 14,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Chevrolet",
					Model:           "Suburban 2500",
					Registration:    "051",
					Color:           "Pink",
					FabricationYear: 1997,
					Capacity:        5,
					MaxSpeed:        173,
					FuelType:        "gas",
					Transmission:    "automatic",
					Weight:          65.95,
					Dimensions: internal.Dimensions{
						Height: 40.51,
						Length: 0,
						Width:  135.28,
					},
				},
			},
		}
		sv.On("FindByBrandAndYearRange", carBrand, startYear, endYear).Return(expectedResult, nil)
		hd := NewHandlerVehicle(sv)
		hdFunc := hd.FindByBrandAndYearRange()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/brand/Chevrolet/between/1995/1997", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("brand", carBrand)
		chiCtx.URLParams.Add("start_year", "1995")
		chiCtx.URLParams.Add("end_year", "1997")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{
		   "data": {
		       "11": {
		           "Id": 11,
		           "Brand": "Chevrolet",
		           "Model": "G-Series 2500",
		           "Registration": "9292",
		           "Color": "Mauv",
		           "FabricationYear": 1996,
		           "Capacity": 3,
		           "MaxSpeed": 239,
		           "FuelType": "gas",
		           "Transmission": "manual",
		           "Weight": 152.87,
		           "Height": 50.84,
		           "Length": 0,
		           "Width": 216.53
		       },
		       "14": {
		           "Id": 14,
		           "Brand": "Chevrolet",
		           "Model": "Suburban 2500",
		           "Registration": "051",
		           "Color": "Pink",
		           "FabricationYear": 1997,
		           "Capacity": 5,
		           "MaxSpeed": 173,
		           "FuelType": "gas",
		           "Transmission": "automatic",
		           "Weight": 65.95,
		           "Height": 40.51,
		           "Length": 0,
		           "Width": 135.28
		       }
		   },
		   "message": "vehicles found"
		}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		sv.AssertExpectations(t)

	})

	t.Run("failure- case01: should returns a not found error", func(t *testing.T) {
		// Arrange
		carBrand := "Chevrolet"
		startYear := 2023
		endYear := 2024
		sv := service.NewVehicleDefaultMock()
		expectedResult := map[int]internal.Vehicle{}
		sv.On("FindByBrandAndYearRange", carBrand, startYear, endYear).Return(expectedResult, internal.ErrRepositoryNotFound)
		hd := NewHandlerVehicle(sv)
		hdFunc := hd.FindByBrandAndYearRange()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/brand/Chevrolet/between/2023/2024", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("brand", carBrand)
		chiCtx.URLParams.Add("start_year", "2023")
		chiCtx.URLParams.Add("end_year", "2024")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"status":"Not Found","message":"No vehicles were found for the provided criteria"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
		sv.AssertExpectations(t)
	})
}

func TestHandlerVehicle_AverageMaxSpeedByBrand(t *testing.T) {
	t.Run("success case 01: should returns max average speed", func(t *testing.T) {
		// Arrange
		carBrand := "Chevrolet"
		sv := service.NewVehicleDefaultMock()
		sv.On("AverageMaxSpeedByBrand", carBrand).Return(164.5, nil)

		hd := NewHandlerVehicle(sv)
		hdFunc := hd.AverageMaxSpeedByBrand()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/average_speed/brand/Chevrolet", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("brand", carBrand)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{
		   "data": 164.5,
		   "message": "average max speed found"
		}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})

	t.Run("failure case 01: should returns an error for an unknown brand", func(t *testing.T) {
		// Arrange
		carBrand := "Mercedes"
		sv := service.NewVehicleDefaultMock()
		sv.On("AverageMaxSpeedByBrand", carBrand).Return(0.0, internal.ErrServiceNoVehicles)

		hd := NewHandlerVehicle(sv)
		hdFunc := hd.AverageMaxSpeedByBrand()

		// Act
		req := httptest.NewRequest(http.MethodGet, "/vehicles/average_speed/brand/Mercedes", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("brand", carBrand)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusNotFound
		expectedBody := `{"status":"Not Found","message":"vehicles not found"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json"},
		}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}
