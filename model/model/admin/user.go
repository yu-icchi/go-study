package admin

type User struct {}

func (u *User) Collection() string {
	return "AdminUser"
}
