package repository

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

func NewVehicleMapMock() *VehicleMapMock {
	return &VehicleMapMock{}
}

type VehicleMapMock struct {
	mock.Mock
}

func (m *VehicleMapMock) FindAll() (map[int]internal.Vehicle, error) {
	args := m.Called()
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleMapMock) FindByColorAndYear(color string, fabricationYear int) (map[int]internal.Vehicle, error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleMapMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (map[int]internal.Vehicle, error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleMapMock) FindByBrand(brand string) (map[int]internal.Vehicle, error) {
	args := m.Called(brand)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *VehicleMapMock) FindByWeightRange(fromWeight float64, toWeight float64) (map[int]internal.Vehicle, error) {
	args := m.Called(fromWeight, toWeight)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}
