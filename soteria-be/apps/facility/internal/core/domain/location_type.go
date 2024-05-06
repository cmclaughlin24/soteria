package domain

type LocationType struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	EnableChildren bool   `json:"enableChildren"`
	CreatedBy      string `json:"-"`
	UpdatedBy      string `json:"-"`
}
