package dto

//This is the request DTO for the table model.
type TableReqDto struct {
	Capacity int `json:"capacity"`
}

//This is the response DTO for the table model.
type TableResDto struct {
	Id       int `json:"id,omitempty"`
	Capacity int `json:"capacity"`
}
