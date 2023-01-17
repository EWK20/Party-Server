//This package will hold all the tests for the service layer
package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Initialising a mock object to act as the repo
type MockGuestRepo struct {
	guestMock mock.Mock
}

// These are mock functions, they simulate the behaviour of the real function, so any arguments that will be passed in and what will be returned

func (m *MockGuestRepo) FindAll() ([]model.Guest, error) {
	args := m.guestMock.Called()
	if args.Error(1) != nil {
		return []model.Guest{}, args.Error(1)
	}

	return args.Get(0).([]model.Guest), nil
}

func (m *MockGuestRepo) FindByName(name string) (model.Guest, error) {
	args := m.guestMock.Called(name)
	if args.Error(1) != nil {
		return model.Guest{}, args.Error(1)
	}

	return args.Get(0).(model.Guest), nil
}

func (m *MockTableRepo) FindTableById(id int) (model.Table, error) {
	args := m.tableMock.Called(id)
	if args.Error(1) != nil {
		return model.Table{}, args.Error(1)
	}

	return args.Get(0).(model.Table), nil
}

func (m *MockGuestRepo) Save(guest model.Guest) (model.Guest, error) {
	args := m.guestMock.Called(guest)
	if args.Error(1) != nil {
		return model.Guest{}, args.Error(1)
	}

	return args.Get(0).(model.Guest), nil
}

func (m *MockGuestRepo) UpdateGuest(guest model.Guest) (model.Guest, error) {
	args := m.guestMock.Called(guest)
	if args.Error(1) != nil {
		return model.Guest{}, args.Error(1)
	}

	return args.Get(0).(model.Guest), nil
}

func (m *MockTableRepo) UpdateTable(table model.Table) (model.Table, error) {
	args := m.tableMock.Called(table)
	if args.Error(1) != nil {
		return model.Table{}, args.Error(1)
	}

	return args.Get(0).(model.Table), nil
}

func (m *MockGuestRepo) GetArrivedGuests() ([]model.Guest, error) {
	args := m.guestMock.Called()
	if args.Error(1) != nil {
		return []model.Guest{}, args.Error(1)
	}

	return args.Get(0).([]model.Guest), nil
}

func (m *MockGuestRepo) Delete(guest model.Guest) error {
	args := m.guestMock.Called(guest)
	if args.Error(0) != nil {
		return args.Error(0)
	}

	return nil
}

// This will test a success use case when trying to retrieve all guests
func TestGuestFindAllSuccess(t *testing.T) {
	testObj := new(MockGuestRepo)

	var res dto.GuestResDto
	var resArr []dto.GuestResDto

	expectedRes := []model.Guest{
		{
			Id:                 1,
			Name:               "Echez",
			Table_ID:           3,
			Acompanying_Guests: 2,
			TimeArrived:        "",
		},
		{
			Id:                 2,
			Name:               "John",
			Table_ID:           1,
			Acompanying_Guests: 6,
			TimeArrived:        "19:30",
		},
		{
			Id:                 3,
			Name:               "Hannah",
			Table_ID:           6,
			Acompanying_Guests: 10,
			TimeArrived:        "22:16",
		},
	}

	testObj.guestMock.On("FindAll").Return(expectedRes, nil)

	testRes, err := testObj.FindAll()

	for _, v := range testRes {
		res.Id = v.Id
		res.Name = v.Name
		res.Table_ID = v.Table_ID
		res.Acompanying_Guests = v.Acompanying_Guests

		resArr = append(resArr, res)
	}

	testObj.guestMock.AssertExpectations(t)

	assert.Equal(t, 1, resArr[0].Id)
	assert.Equal(t, "Echez", resArr[0].Name)
	assert.Equal(t, 3, resArr[0].Table_ID)
	assert.Equal(t, 2, resArr[0].Acompanying_Guests)
	assert.Equal(t, 3, len(testRes))
	assert.Nil(t, err)

}

func TestGuestFindAllError(t *testing.T) {
	testObj := new(MockGuestRepo)

	testObj.guestMock.On("FindAll").Return([]model.Guest{}, errors.New("Could not retrieve guests"))

	testRes, err := testObj.FindAll()

	testObj.guestMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not retrieve guests", err.Error())
	assert.Equal(t, []model.Guest{}, testRes)

}

func TestGuestSaveSuccess(t *testing.T) {
	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	var guestResDto dto.GuestResDto

	tableTestId := 1

	expectedTableRes := model.Table{Id: tableTestId, Capacity: 10}

	tableTestObj.tableMock.On("FindTableById", tableTestId).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(tableTestId)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 10, tableRes.Capacity)

	guest := model.Guest{Id: 3, Name: "John", Table_ID: tableRes.Id, Acompanying_Guests: 8}

	guestTestObj.guestMock.On("Save", guest).Return(guest, nil)

	guestRes, err := guestTestObj.Save(guest)

	guestResDto.Name = guestRes.Name
	guestResDto.Acompanying_Guests = guestRes.Acompanying_Guests

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, "John", guestResDto.Name)
	assert.Equal(t, 8, guestResDto.Acompanying_Guests)

}

func TestGuestSaveTooManyGuestsError(t *testing.T) {
	guestTestObj := new(MockGuestRepo)
	var guestResDto dto.GuestResDto

	request := dto.GuestReqDto{Name: "John", Table_ID: 1, Acompanying_Guests: 12}

	//Find table by Id
	table := model.Table{Id: 1, Capacity: 10}

	guest := model.Guest{Id: 3, Name: "John", Table_ID: 1, Acompanying_Guests: 8}

	// Check if table capacity is less than the amount of guests attending, throw error
	if table.Capacity < (request.Acompanying_Guests + 1) {
		guestTestObj.guestMock.On("Save", guest).Return(model.Guest{}, errors.New("There are too many guests"))

	}

	guestRes, err := guestTestObj.Save(guest)

	guestResDto.Name = guestRes.Name
	guestResDto.Acompanying_Guests = guestRes.Acompanying_Guests

	guestTestObj.guestMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "There are too many guests", err.Error())
	assert.Equal(t, dto.GuestResDto{}, guestResDto)
}

func TestGuestSaveError(t *testing.T) {
	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	var guestResDto dto.GuestResDto

	//Find table by id
	tableTestId := 1

	expectedTableRes := model.Table{Id: tableTestId, Capacity: 10}

	tableTestObj.tableMock.On("FindTableById", tableTestId).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(tableTestId)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 10, tableRes.Capacity)

	// Attempting to create the guest, throw error
	guest := model.Guest{Id: 3, Name: "John", Table_ID: tableRes.Id, Acompanying_Guests: 8}

	guestTestObj.guestMock.On("Save", guest).Return(model.Guest{}, errors.New("Could not create guest"))

	guestRes, err := guestTestObj.Save(guest)

	guestResDto.Name = guestRes.Name
	guestResDto.Acompanying_Guests = guestRes.Acompanying_Guests

	guestTestObj.guestMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not create guest", err.Error())
	assert.Equal(t, dto.GuestResDto{}, guestResDto)
}

func TestGuestCheckinSuccess(t *testing.T) {

	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	guestRequest := dto.GuestReqDto{Name: "Hannah", Table_ID: 3, Acompanying_Guests: 7}

	//Find the guest
	expectedGuestRes := model.Guest{Id: 1, Name: "Hannah", Table_ID: 3, Acompanying_Guests: 5}

	guestTestObj.guestMock.On("FindByName", guestRequest.Name).Return(expectedGuestRes, nil)

	guest, err := guestTestObj.FindByName(guestRequest.Name)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, guest.Id)
	assert.Equal(t, "Hannah", guest.Name)
	assert.Equal(t, 3, guest.Table_ID)
	assert.Equal(t, 5, guest.Acompanying_Guests)

	// Find the guest's table
	expectedTableRes := model.Table{Id: 3, Capacity: 10}

	tableTestObj.tableMock.On("FindTableById", guest.Table_ID).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(guest.Table_ID)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 10, tableRes.Capacity)

	// Update table capacity
	newExpectedTableRes := model.Table{Id: 3, Capacity: 3}

	tableRes.Capacity -= (guestRequest.Acompanying_Guests + 1)

	tableTestObj.tableMock.On("UpdateTable", tableRes).Return(newExpectedTableRes, nil)

	newTableRes, err := tableTestObj.UpdateTable(tableRes)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 3, newTableRes.Id)
	assert.Equal(t, 3, newTableRes.Capacity)

	// Update the old accompanying guest number and log time of arrival

	newExpectedGuestRes := model.Guest{Id: 1, Name: "Hannah", Table_ID: 3, Acompanying_Guests: 7, TimeArrived: time.Now().Format("15:04")}

	guest.Acompanying_Guests = guestRequest.Acompanying_Guests
	guest.TimeArrived = time.Now().Format("15:04")

	guestTestObj.guestMock.On("UpdateGuest", guest).Return(newExpectedGuestRes, nil)

	newGuest, err := guestTestObj.UpdateGuest(guest)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, newGuest.Id)
	assert.Equal(t, "Hannah", newGuest.Name)
	assert.Equal(t, 3, newGuest.Table_ID)
	assert.Equal(t, 7, newGuest.Acompanying_Guests)
	assert.Equal(t, time.Now().Format("15:04"), newGuest.TimeArrived)

}

func TestGuestCheckinTooManyGuestsError(t *testing.T) {
	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	var resDto dto.GuestResDto

	guestRequest := dto.GuestReqDto{Name: "Hannah", Table_ID: 3, Acompanying_Guests: 12}

	//Find the guest
	expectedGuestRes := model.Guest{Id: 1, Name: "Hannah", Table_ID: 3, Acompanying_Guests: 5}

	guestTestObj.guestMock.On("FindByName", guestRequest.Name).Return(expectedGuestRes, nil)

	guest, err := guestTestObj.FindByName(guestRequest.Name)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, guest.Id)
	assert.Equal(t, "Hannah", guest.Name)
	assert.Equal(t, 3, guest.Table_ID)
	assert.Equal(t, 5, guest.Acompanying_Guests)

	// Find the guest's table
	expectedTableRes := model.Table{Id: 3, Capacity: 10}

	tableTestObj.tableMock.On("FindTableById", guest.Table_ID).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(guest.Table_ID)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 10, tableRes.Capacity)

	if tableRes.Capacity < (guestRequest.Acompanying_Guests + 1) {
		guestTestObj.guestMock.On("UpdateGuest", guest).Return(model.Guest{}, errors.New("There are too many guests"))
	}

	newGuest, err := guestTestObj.UpdateGuest(guest)

	resDto.Name = newGuest.Name

	guestTestObj.guestMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "There are too many guests", err.Error())
	assert.Equal(t, dto.GuestResDto{}, resDto)

}

func TestGuestCheckoutSuccess(t *testing.T) {
	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	name := "Hannah"

	//Find the guest
	expectedGuestRes := model.Guest{Id: 1, Name: "Hannah", Table_ID: 3, Acompanying_Guests: 7}

	guestTestObj.guestMock.On("FindByName", name).Return(expectedGuestRes, nil)

	guest, err := guestTestObj.FindByName(name)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, guest.Id)
	assert.Equal(t, "Hannah", guest.Name)
	assert.Equal(t, 3, guest.Table_ID)
	assert.Equal(t, 7, guest.Acompanying_Guests)

	// Find table by Id
	tableTestId := 3

	expectedTableRes := model.Table{Id: tableTestId, Capacity: 2}

	tableTestObj.tableMock.On("FindTableById", tableTestId).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(tableTestId)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 2, tableRes.Capacity)

	// Update table capacity
	tableRes.Capacity += (guest.Acompanying_Guests + 1)

	tableTestObj.tableMock.On("UpdateTable", tableRes).Return(tableRes, nil)

	newTableRes, err := tableTestObj.UpdateTable(tableRes)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 3, newTableRes.Id)
	assert.Equal(t, 10, newTableRes.Capacity)

	// Delete guest from db
	guestTestObj.guestMock.On("Delete", guest).Return(nil)

	err = guestTestObj.Delete(guest)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)

}

func TestGuestCheckoutError(t *testing.T) {
	tableTestObj := new(MockTableRepo)
	guestTestObj := new(MockGuestRepo)

	name := "Hannah"

	//Find the guest
	expectedGuestRes := model.Guest{Id: 1, Name: "Hannah", Table_ID: 3, Acompanying_Guests: 7}

	guestTestObj.guestMock.On("FindByName", name).Return(expectedGuestRes, nil)

	guest, err := guestTestObj.FindByName(name)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 1, guest.Id)
	assert.Equal(t, "Hannah", guest.Name)
	assert.Equal(t, 3, guest.Table_ID)
	assert.Equal(t, 7, guest.Acompanying_Guests)

	// Find table by Id
	tableTestId := 3

	expectedTableRes := model.Table{Id: tableTestId, Capacity: 2}

	tableTestObj.tableMock.On("FindTableById", tableTestId).Return(expectedTableRes, nil)

	tableRes, err := tableTestObj.FindTableById(tableTestId)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, expectedTableRes, tableRes)
	assert.Equal(t, 2, tableRes.Capacity)

	// Update table capacity
	tableRes.Capacity += (guest.Acompanying_Guests + 1)

	tableTestObj.tableMock.On("UpdateTable", tableRes).Return(tableRes, nil)

	newTableRes, err := tableTestObj.UpdateTable(tableRes)

	tableTestObj.tableMock.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, 3, newTableRes.Id)
	assert.Equal(t, 10, newTableRes.Capacity)

	// Delete guest from db
	guestTestObj.guestMock.On("Delete", guest).Return(errors.New("Could not delete guest"))

	err = guestTestObj.Delete(guest)

	guestTestObj.guestMock.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, "Could not delete guest", err.Error())
}
