package domain

type Facility struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedBy string `json:"-"`
	UpdatedBy string `json:"-"`
}
