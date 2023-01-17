//The data transfer object(DTO) handles the incoming request body and outgoing request body without exposing any sensitive information.
package dto

//This is the request DTO for the guest model.
type GuestReqDto struct {
	Name               string `json:"name,omitempty"`
	Table_ID           int    `json:"table_id,omitempty"`
	Acompanying_Guests int    `json:"accompanying_guests"`
	TimeArrived        string `json:"time_arrived,omitempty"`
}

//This is the response DTO for the guest model.
type GuestResDto struct {
	Id                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Table_ID           int    `json:"table_id,omitempty"`
	Acompanying_Guests int    `json:"accompanying_guests"`
	TimeArrived        string `json:"time_arrived,omitempty"`
}
