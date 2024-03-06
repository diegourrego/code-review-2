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
