package authentication

func (u *UserDetails) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *UserDetails) GetRoleStr() string {
	return u.Role.(string)
}

func (u *UserDetails) GetRole() interface{} {
	return u.Role
}

func (u *UserDetails) GetUserTypeStr() string {
	return u.UserType.(string)
}

func (u *UserDetails) GetUserType() interface{} {
	return u.UserType
}
