package service

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVehicleDefaultMock_FindByColorAndYear(t *testing.T) {
	t.Run("success - case 01: should returns a list of cars", func(t *testing.T) {
		// Arrange
		rp := repository.NewVehicleMapMock()
		carColor := "Orange"
		carYear := 1995
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

		rp.On("FindByColorAndYear", carColor, carYear).Return(expectedResult, nil)

		sv := NewServiceVehicleDefault(rp)

		// Act
		vehiclesFound, errorFound := sv.FindByColorAndYear(carColor, carYear)

		// Assert
		require.Equal(t, expectedResult, vehiclesFound)
		require.Nil(t, errorFound)
		rp.AssertExpectations(t)
	})

	t.Run("failure - case 01: should returns an error when there's no vehicles under criteria", func(t *testing.T) {
		// Arrange
		rp := repository.NewVehicleMapMock()
		carColor := "Cyan"
		carYear := 2022

		expectedResult := map[int]internal.Vehicle{}

		rp.On("FindByColorAndYear", carColor, carYear).Return(expectedResult, internal.ErrRepositoryNotFound)

		sv := NewServiceVehicleDefault(rp)

		// Act
		vehiclesFound, errorFound := sv.FindByColorAndYear(carColor, carYear)

		// Assert
		require.Equal(t, expectedResult, vehiclesFound)
		require.EqualError(t, errorFound, "repository: No vehicles were found for the provided criteria")
		rp.AssertExpectations(t)
	})
}

func TestVehicleDefaultMock_FindByBrandAndYearRange(t *testing.T) {
	t.Run("success - case01: should returns a list of cars", func(t *testing.T) {
		// Arrange
		rp := repository.NewVehicleMapMock()
		carBrand := "Orange"
		startYear := 1995
		endYear := 2000

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

		rp.On("FindByBrandAndYearRange", carBrand, startYear, endYear).Return(expectedResult, nil)

		sv := NewServiceVehicleDefault(rp)

		// Act
		vehiclesFound, errorFound := sv.FindByBrandAndYearRange(carBrand, startYear, endYear)

		// Assert
		require.Equal(t, expectedResult, vehiclesFound)
		require.Nil(t, errorFound)
		rp.AssertExpectations(t)
	})

	t.Run("failure - case 01: should returns an error when there's no vehicles under criteria", func(t *testing.T) {
		// Arrange
		rp := repository.NewVehicleMapMock()
		carBrand := "Mercedes"
		startYear := 2020
		endYear := 2021

		expectedResult := map[int]internal.Vehicle{}

		rp.On("FindByBrandAndYearRange", carBrand, startYear, endYear).Return(expectedResult, internal.ErrRepositoryNotFound)

		sv := NewServiceVehicleDefault(rp)

		// Act
		vehiclesFound, errorFound := sv.FindByBrandAndYearRange(carBrand, startYear, endYear)

		// Assert
		require.Equal(t, expectedResult, vehiclesFound)
		require.EqualError(t, errorFound, "repository: No vehicles were found for the provided criteria")
		rp.AssertExpectations(t)
	})
}
