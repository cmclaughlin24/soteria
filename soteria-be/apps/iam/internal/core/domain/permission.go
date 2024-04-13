package domain

type Permission struct {
	Id       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func NewPermission(id, resource, action string) *Permission {
	return &Permission{
		Id:       id,
		Resource: resource,
		Action:   action,
	}
}
