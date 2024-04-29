package domain

type Facility struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"-"`
	UpdatedBy string `json:"-"`
}
