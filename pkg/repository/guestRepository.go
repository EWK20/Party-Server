package repository

import (
	"log"

	"github.com/getground/tech-tasks/backend/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GuestRepository interface {
	FindAll() ([]model.Guest, error)
	FindByName(name string) (model.Guest, error)
	Save(guest model.Guest) (model.Guest, error)
	Update(guest model.Guest) error
	GetArrivedGuests() ([]model.Guest, error)
	Delete(guest model.Guest) error
}

type guestDatabase struct {
	connection *gorm.DB
}

func NewGuestRepository() GuestRepository {
	dsn := "user:password@tcp(host.docker.internal:3306)/getground?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		//If connection fails then throw error
		panic("Could not connect to db")
	}

	log.Println("Database connection is successful")

	db.AutoMigrate(&model.Guest{})

	return &guestDatabase{
		connection: db,
	}

}

func (db *guestDatabase) FindAll() ([]model.Guest, error) {
	var guests []model.Guest
	if err := db.connection.Set("gorm:auto_preload", true).Find(&guests).Error; err != nil {
		return guests, err
	}
	return guests, nil
}

func (db *guestDatabase) FindByName(name string) (model.Guest, error) {
	var guest model.Guest
	if err := db.connection.Where(&model.Guest{Name: name}).First(&guest).Error; err != nil {
		return guest, err
	}
	return guest, nil
}

func (db *guestDatabase) Save(guest model.Guest) (model.Guest, error) {

	if err := db.connection.Create(&guest).Error; err != nil {
		return guest, err
	}
	return guest, nil
}

func (db *guestDatabase) Update(guest model.Guest) error {
	if err := db.connection.Save(&guest).Error; err != nil {
		log.Println("Checkin Service - Could not create guest")
		return err
	}
	return nil
}

func (db *guestDatabase) GetArrivedGuests() ([]model.Guest, error) {
	var guests []model.Guest
	if err := db.connection.Not("time_arrived = ?", "").Find(&guests).Error; err != nil {
		log.Println("Get Arrived Guests Service - Could not retrieve guests")
		return guests, err
	}
	return guests, nil
}

func (db *guestDatabase) Delete(guest model.Guest) error {
	if err := db.connection.Delete(&guest).Error; err != nil {
		log.Println("Checkout Service - Could not delete guest")
		return err
	}
	return nil
}
