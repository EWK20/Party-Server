package model

// Creating table model
type Table struct {
	Id       int `json:"id" gorm:"primaryKey"`
	Capacity int `json:"capacity"`
}

func (u *Table) TableName() string {
	// custom table name, this is default
	return "table"
}
