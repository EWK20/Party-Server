//The model shows the structure of the data that will be interacted with through the repository.
package model

// Creating guest model
type Guest struct {
	Id                 int    `json:"id" gorm:"primaryKey"`
	Name               string `json:"name"`
	Table_ID           int    `json:"table_id"`
	Table              Table  `gorm:"foreignKey:Table_ID;references:Id"`
	Acompanying_Guests int    `json:"accompanying_guests"`
	TimeArrived        string `json:"time_arrived"`
}

func (u *Guest) TableName() string {
	// custom table name, this is default
	return "guest"
}
