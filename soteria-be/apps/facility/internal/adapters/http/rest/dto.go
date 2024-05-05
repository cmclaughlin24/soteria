package rest

type CreateFacilityDto struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateFacilityDto struct {
	Name string `json:"name" validate:"required"`
}

type CreateLocationDto struct {
	Code         string `json:"code" validate:"required"`
	Name         string `json:"name" validate:"required"`
	FacilityCode string `json:"facilityCode" vaildate:"required"`
	ParentId     int    `json:"parentId"`
}

type UpdateLocationDto struct {
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required"`
	ParentId int    `json:"parentId"`
}
