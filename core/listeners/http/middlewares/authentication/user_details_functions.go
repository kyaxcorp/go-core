package authentication

func (u *UserDetails) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *UserDetails) GetRole() string {
	return u.Role
}

func (u *UserDetails) GetUserType() string {
	return u.UserType
}
