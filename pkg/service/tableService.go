package service

import (
	"log"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/model"
	"github.com/getground/tech-tasks/backend/pkg/repository"
)

type TableService interface {
	FindAll() ([]dto.TableResDto, error)
	FindById(id int) (dto.TableResDto, error)
	Save(req dto.TableReqDto) (dto.TableResDto, error)
	CheckSpace() (int, bool)
}

type tableService struct {
	tableRepository repository.TableRepository
}

func NewTableService(tableRepo repository.TableRepository) TableService {
	return &tableService{
		tableRepository: tableRepo,
	}
}

func (service *tableService) FindAll() ([]dto.TableResDto, error) {
	var res dto.TableResDto
	var resArr []dto.TableResDto

	tables, err := service.tableRepository.FindAll()
	if err != nil {
		log.Println("Get Tables Service - Could not find tables")
		return resArr, err
	}

	for _, v := range tables {
		res.Id = v.Id
		res.Capacity = v.Capacity

		resArr = append(resArr, res)
	}

	return resArr, nil
}

func (service *tableService) FindById(id int) (dto.TableResDto, error) {
	var res dto.TableResDto

	table, err := service.tableRepository.FindById(id)
	if err != nil {
		log.Println("Get Table By Id Service - Could not find table")
		return res, err
	}

	res.Id = table.Id
	res.Capacity = table.Capacity

	return res, nil
}

func (service *tableService) Save(req dto.TableReqDto) (dto.TableResDto, error) {
	var table model.Table
	var res dto.TableResDto

	table.Capacity = req.Capacity

	err := service.tableRepository.Save(table)
	if err != nil {
		log.Println("Create Table Service - Could not create table")
		return res, err
	}

	res.Id = table.Id
	res.Capacity = table.Capacity

	return res, nil
}

func (service *tableService) CheckSpace() (int, bool) {

	space := 0

	tables, err := service.tableRepository.FindAll()
	if err != nil {
		log.Println("Available Space Service - Could not get tables")
		return 0, false
	}

	for _, v := range tables {
		space += v.Capacity
	}

	return space, true
}
