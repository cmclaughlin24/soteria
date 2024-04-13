package domain

type User struct {
	Id              string           `json:"id"`
	Name            string           `json:"name"`
	Email           string           `json:"email"`
	PhoneNumber     string           `json:"phoneNumber"`
	Password        string           `json:"-"`
	DeliveryMethods []string         `json:"deliveryMethods"`
	TimeZone        string           `json:"timeZone"`
	Permissions     []UserPermission `json:"permissions"`
}

func NewUser(id, name, email, phoneNumber, password string, deliveryMethods []string, timeZone string) *User {
	return &User{
		Id:              id,
		Name:            name,
		Email:           email,
		PhoneNumber:     phoneNumber,
		Password:        password,
		DeliveryMethods: deliveryMethods,
		TimeZone:        timeZone,
	}
}

func (u *User) AddPermission(permission UserPermission) {
	u.Permissions = append(u.Permissions, permission)
}

type UserPermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
