package domain

type Location struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	Name           string `json:"name"`
	FacilityCode   string `json:"facilityCode"`
	LocationTypeId int    `json:"locationTypeId"`
	ParentId       int    `json:"parentId"`
	CreatedBy      string `json:"-"`
	UpdatedBy      string `json:"-"`
}
