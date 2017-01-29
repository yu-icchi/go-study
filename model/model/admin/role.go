package admin

type Role struct {}

func (r *Role) Collection() string {
	return "AdminRole"
}

