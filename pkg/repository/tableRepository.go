// The repository layer is used to access and interact with the database.

package repository

import (
	"log"

	"github.com/getground/tech-tasks/backend/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TableRepository interface {
	FindAll() ([]model.Table, error)
	FindById(id int) (model.Table, error)
	Save(table model.Table) error
	Update(table model.Table) error
	Delete(table model.Table) error
}

type tableDatabase struct {
	connection *gorm.DB
}

func NewTableRepository() TableRepository {
	dsn := "user:password@tcp(host.docker.internal:3306)/getground?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		//If connection fails then throw error
		panic("Could not connect to db")
	}

	log.Println("Database connection is successful")

	db.AutoMigrate(&model.Table{})

	return &tableDatabase{
		connection: db,
	}

}

func (db *tableDatabase) FindAll() ([]model.Table, error) {
	var tables []model.Table
	if err := db.connection.Find(&tables).Error; err != nil {
		return tables, err
	}

	return tables, nil
}

func (db *tableDatabase) FindById(id int) (model.Table, error) {
	var table model.Table
	if err := db.connection.Find(&table, id).Error; err != nil {
		return table, err
	}

	return table, nil
}

func (db *tableDatabase) Save(table model.Table) error {
	if err := db.connection.Create(&table).Error; err != nil {
		return err
	}
	return nil
}

func (db *tableDatabase) Update(table model.Table) error {
	if err := db.connection.Save(&table).Error; err != nil {
		return err
	}
	return nil
}

func (db *tableDatabase) Delete(table model.Table) error {
	if err := db.connection.Delete(&table).Error; err != nil {
		return err
	}
	return nil
}
