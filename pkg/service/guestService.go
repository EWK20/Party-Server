//This package holds all the business logic for the application
package service

import (
	"log"
	"time"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/model"
	"github.com/getground/tech-tasks/backend/pkg/repository"
)

//The guest service
type GuestService interface {
	FindAll() ([]dto.GuestResDto, error)
	Save(req dto.GuestReqDto) (dto.GuestResDto, error)
	Checkin(req dto.GuestReqDto) (dto.GuestResDto, error)
	Checkout(name string) error
	GetArrivedGuests() ([]dto.GuestResDto, error)
}

type guestService struct {
	guestRepository repository.GuestRepository
	tableRepository repository.TableRepository
}

func NewGuestService(guestRepo repository.GuestRepository, tableRepo repository.TableRepository) GuestService {
	return &guestService{
		guestRepository: guestRepo,
		tableRepository: tableRepo,
	}
}

//This function will call the table repository to retrieve all of the tables, then it maps the table entity to the response data transfer object
func (service *guestService) FindAll() ([]dto.GuestResDto, error) {

	var res dto.GuestResDto
	var resArr []dto.GuestResDto

	//This query runs -> SELECT * FROM `guest`
	guests, err := service.guestRepository.FindAll()
	if err != nil {
		log.Println("Get Tables Service - Could not find tables")
		return resArr, err
	}

	//Looping through the guests arr and mapping it to the response dto (data transfer object)
	for _, v := range guests {
		res.Id = v.Id
		res.Name = v.Name
		res.Table_ID = v.Table_ID
		res.Acompanying_Guests = v.Acompanying_Guests

		resArr = append(resArr, res)
	}

	return resArr, nil
}

func (service *guestService) Save(req dto.GuestReqDto) (dto.GuestResDto, error) {

	var guest model.Guest
	var res dto.GuestResDto

	id := req.Table_ID

	//* Retrieves the specified table by Id
	table, err := service.tableRepository.FindById(id)
	if err != nil {
		log.Println("Create Guest Service - Could not find specified table")
		return res, err
	}

	//* Added 1 to accompnaying guests because it will then include the main guest
	//* If the capacity of the table is smaller than the amount of people coming, then throw an error
	if table.Capacity < (req.Acompanying_Guests + 1) {
		log.Println("Create Guest Service - There are too many guests")
		return res, err
	}

	guest.Name = req.Name
	guest.Table_ID = req.Table_ID
	guest.Acompanying_Guests = req.Acompanying_Guests

	//* Create the guest
	//* This query runs -> INSERT INTO `guest` (`name`,`table_id`,`acompanying_guests`) VALUES ('sara',5,9)
	newGuest, err := service.guestRepository.Save(guest)
	if err != nil {
		log.Println("Create Guest Service - Could not create guest")
		return res, err
	}

	res.Name = newGuest.Name
	res.Acompanying_Guests = newGuest.Acompanying_Guests

	return res, nil
}

func (service *guestService) Checkin(req dto.GuestReqDto) (dto.GuestResDto, error) {
	var res dto.GuestResDto

	var newTable model.Table

	// Find specified guest
	guest, err := service.guestRepository.FindByName(req.Name)
	if err != nil {
		log.Println("Checkin Service - Could not find guest")
		return res, err
	}

	// Find the guest's table
	table, err := service.tableRepository.FindById(guest.Table_ID)
	if err != nil {
		log.Println("Checkin Service - Could not find specified table")
		return res, err
	}

	// Added 1 to accompnaying guests because it will then include the main guest
	// If the capacity of the table is smaller than the actual amount of people coming, then throw an error
	if table.Capacity < (req.Acompanying_Guests + 1) {
		log.Println("Checkin Service - There are too many guests")
		return res, err
	}

	// Reduce the listed capcity once guest actually arrives
	newTable.Id = table.Id
	newTable.Capacity = table.Capacity - (req.Acompanying_Guests + 1)

	// Save the new table capacity
	// This query runs ->  UPDATE `table` SET `capacity`=1 WHERE `id` = 5
	err = service.tableRepository.Update(newTable)
	if err != nil {
		log.Println("Checkin Service - Could not update table")
		return res, err
	}

	// Update the old accompanying guest number and log time of arrival
	guest.Acompanying_Guests = req.Acompanying_Guests
	guest.TimeArrived = time.Now().Format("15:04")

	// Save the guest details
	// This query runs -> UPDATE `guest` SET `name`='john',`table_id`=2,`acompanying_guests`=1 WHERE `id` = 1
	err = service.guestRepository.Update(guest)
	if err != nil {
		log.Println("Checkin Service - Could not create guest")
		return res, err
	}

	// Map the new guest object to the response dto
	res.Name = guest.Name

	return res, nil
}

func (service *guestService) Checkout(name string) error {
	var newTable model.Table

	// Find guest by name
	// This query runs -> SELECT * FROM `guest` WHERE `guest`.`name` = 'sara' ORDER BY `guest`.`id` LIMIT 1
	guest, err := service.guestRepository.FindByName(name)
	if err != nil {
		log.Println("Checkout Service - Could not find guest")
		return err
	}

	// Find the guest's table
	table, err := service.tableRepository.FindById(guest.Table_ID)
	if err != nil {
		log.Println("Checkout Service - Could not find specified table")
		return err
	}

	// Adding the available space back to the table
	newTable.Id = table.Id
	newTable.Capacity = table.Capacity + (guest.Acompanying_Guests + 1)

	// Save the new table capacity
	// This query runs -> UPDATE `table` SET `capacity`=1 WHERE `id` = 5
	err = service.tableRepository.Update(newTable)
	if err != nil {
		log.Println("Checkout Service - Could not update table")
		return err
	}

	// Soft Delete - There would be a boolean property in the struct for the guest to indicate if they have left the party.
	// Would be updated with an UPDATE query
	// Then the GET methods would be changed so that they only retrieve guests that "have arrived"

	// Hard Delete - Removes guest from database
	// This query runs -> DELETE FROM `guest` WHERE `guest`.`id` = 2
	err = service.guestRepository.Delete(guest)
	if err != nil {
		log.Println("Checkout Service - Could not delete guest")
		return err
	}

	return nil
}

func (service *guestService) GetArrivedGuests() ([]dto.GuestResDto, error) {
	var res dto.GuestResDto
	var resArr []dto.GuestResDto

	//* This query runs -> SELECT * FROM `guest` WHERE NOT time_arrived = ''
	guests, err := service.guestRepository.GetArrivedGuests()
	if err != nil {
		log.Println("Get Arrived Guests Service - Could not retrieve guests")
		return nil, err
	}

	for _, v := range guests {
		res.Name = v.Name
		res.Acompanying_Guests = v.Acompanying_Guests
		res.TimeArrived = v.TimeArrived

		resArr = append(resArr, res)
	}

	return resArr, nil
}
