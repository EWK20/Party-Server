package service_test

import (
	"errors"
	"testing"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Initialising a mock object to act as the repo
type MockTableRepo struct {
	tableMock mock.Mock
}

// This is a mock function, it simulates the behaviour of the real function, so any arguments that will be passed in and what will be returned
func (m *MockTableRepo) FindAll() ([]model.Table, error) {
	args := m.tableMock.Called()
	if args.Error(1) != nil {
		return []model.Table{}, args.Error(1)
	}

	return args.Get(0).([]model.Table), nil
}

func (m *MockTableRepo) FindById(id int) (model.Table, error) {
	args := m.tableMock.Called(id)
	if args.Error(1) != nil {
		return model.Table{}, args.Error(1)
	}

	return args.Get(0).(model.Table), nil
}

func (m *MockTableRepo) Save(table model.Table) (model.Table, error) {
	args := m.tableMock.Called(table)
	if args.Error(1) != nil {
		return model.Table{}, args.Error(1)
	}
	return table, nil
}

func (m *MockTableRepo) Update(table model.Table) (model.Table, error) {
	args := m.tableMock.Called(table)
	if args.Error(1) != nil {
		return model.Table{}, args.Error(1)
	}
	return table, nil
}

func TestTableFindAllSuccess(t *testing.T) {

	testObj := new(MockTableRepo)

	var resDto dto.TableResDto
	var resDtoArr []dto.TableResDto

	expectedRes := []model.Table{
		{
			Id:       1,
			Capacity: 10,
		},
		{
			Id:       2,
			Capacity: 2,
		},
		{
			Id:       3,
			Capacity: 2,
		},
		{
			Id:       4,
			Capacity: 5,
		},
		{
			Id:       5,
			Capacity: 15,
		},
	}

	testObj.tableMock.On("FindAll").Return(expectedRes, nil)

	testRes, err := testObj.FindAll()

	for _, v := range testRes {
		resDto.Id = v.Id
		resDto.Capacity = v.Capacity

		resDtoArr = append(resDtoArr, resDto)
	}

	testObj.tableMock.AssertExpectations(t)

	assert.Equal(t, 5, len(resDtoArr))
	assert.Equal(t, 15, resDtoArr[4].Capacity)
	assert.Nil(t, err)
}

func TestTableFindAllError(t *testing.T) {
	testObj := new(MockTableRepo)

	var resDto dto.TableResDto
	var resDtoArr []dto.TableResDto

	testObj.tableMock.On("FindAll").Return([]model.Table{}, errors.New("Could not find tables"))

	testRes, err := testObj.FindAll()

	for _, v := range testRes {
		resDto.Id = v.Id
		resDto.Capacity = v.Capacity

		resDtoArr = append(resDtoArr, resDto)
	}

	testObj.tableMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not find tables", err.Error())
	assert.Equal(t, 0, len(resDtoArr))
}

func TestTableFindByIdSuccess(t *testing.T) {
	testObj := new(MockTableRepo)

	var resDto dto.TableResDto

	testId := 1

	expectedRes := model.Table{Id: testId, Capacity: 10}

	testObj.tableMock.On("FindById", testId).Return(expectedRes, nil)

	testRes, err := testObj.FindById(testId)

	resDto.Id = testRes.Id
	resDto.Capacity = testRes.Capacity

	testObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, resDto.Id)
	assert.Equal(t, 10, resDto.Capacity)
}

func TestTableFindByIdError(t *testing.T) {
	testObj := new(MockTableRepo)

	var resDto dto.TableResDto
	testId := 1

	testObj.tableMock.On("FindById", testId).Return(model.Table{}, errors.New("Could not find table"))

	testRes, err := testObj.FindById(testId)

	resDto.Id = testRes.Id
	resDto.Capacity = testRes.Capacity

	testObj.tableMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not find table", err.Error())
	assert.Equal(t, dto.TableResDto{}, resDto)
}

func TestTableSaveSuccess(t *testing.T) {
	testObj := new(MockTableRepo)

	var res dto.TableResDto
	table := model.Table{Id: 1, Capacity: 15}

	testObj.tableMock.On("Save", table).Return(table, nil)

	testRes, err := testObj.Save(table)

	res.Id = testRes.Id
	res.Capacity = testRes.Capacity

	testObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, res.Id)
	assert.Equal(t, 15, res.Capacity)
}

func TestTableSaveError(t *testing.T) {
	testObj := new(MockTableRepo)

	var res dto.TableResDto
	table := model.Table{Id: 1, Capacity: 15}

	testObj.tableMock.On("Save", table).Return(model.Table{}, errors.New("Could not create table"))

	testRes, err := testObj.Save(table)

	res.Id = testRes.Id
	res.Capacity = testRes.Capacity

	testObj.tableMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not create table", err.Error())
	assert.Equal(t, dto.TableResDto{}, res)
}

func TestTableCheckSpaceSuccess(t *testing.T) {
	testObj := new(MockTableRepo)

	space := 0

	expectedRes := []model.Table{
		{
			Id:       1,
			Capacity: 10,
		},
		{
			Id:       2,
			Capacity: 2,
		},
		{
			Id:       3,
			Capacity: 2,
		},
		{
			Id:       4,
			Capacity: 5,
		},
		{
			Id:       5,
			Capacity: 15,
		},
	}

	testObj.tableMock.On("FindAll").Return(expectedRes, nil)

	testRes, err := testObj.FindAll()

	for _, v := range testRes {
		space += v.Capacity
	}

	testObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 34, space)
}

func TestTableCheckSpaceError(t *testing.T) {
	testObj := new(MockTableRepo)

	testObj.tableMock.On("FindAll").Return([]model.Table{}, errors.New("Could not find tables"))

	testRes, err := testObj.FindAll()

	testObj.tableMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not find tables", err.Error())
	assert.Equal(t, []model.Table{}, testRes)
}
