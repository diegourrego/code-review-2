package service

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

func NewVehicleDefaultMock() *VehicleDefaultMock {
	return &VehicleDefaultMock{}
}

type VehicleDefaultMock struct {
	mock.Mock
}

func (m *VehicleDefaultMock) FindByColorAndYear(color string, fabricationYear int) (map[int]internal.Vehicle, error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (map[int]internal.Vehicle, error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleDefaultMock) AverageMaxSpeedByBrand(brand string) (float64, error) {
	args := m.Called(brand)
	return args.Get(0).(float64), args.Error(1)
}

func (m *VehicleDefaultMock) AverageCapacityByBrand(brand string) (int, error) {
	args := m.Called(brand)
	return args.Get(0).(int), args.Error(1)
}

func (m *VehicleDefaultMock) SearchByWeightRange(query internal.SearchQuery, ok bool) (map[int]internal.Vehicle, error) {
	args := m.Called(query, ok)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
