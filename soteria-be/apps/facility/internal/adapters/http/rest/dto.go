package rest

type CreateFacilityDto struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateFacilityDto struct {
	Name string `json:"name" validate:"required"`
}
