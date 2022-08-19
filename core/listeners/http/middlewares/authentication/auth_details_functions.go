package authentication

import "github.com/google/uuid"

func (a *AuthDetails) GetUserID() string {
	return a.UserDetails.UserID
}

// GetUserAsUUID -> handle error by yourself!
func (a *AuthDetails) GetUserAsUUID() (uuid.UUID, error) {
	return uuid.Parse(a.UserDetails.UserID)
}

func (a *AuthDetails) UserUUID() uuid.UUID {
	id, _ := a.GetUserAsUUID()
	return id
}

func (a *AuthDetails) GetFullName() string {
	return a.UserDetails.GetFullName()
}

func (a *AuthDetails) GetRole() string {
	return a.UserDetails.GetRole()
}

func (a *AuthDetails) GetUserType() string {
	return a.UserDetails.GetUserType()
}

//-------------------------------------\\

func (a *AuthDetails) GetDeviceID() string {
	return a.DeviceDetails.DeviceID
}

// GetDeviceAsUUID -> handle error by yourself!
func (a *AuthDetails) GetDeviceAsUUID() (uuid.UUID, error) {
	return uuid.Parse(a.DeviceDetails.DeviceID)
}

func (a *AuthDetails) DeviceUUID() uuid.UUID {
	id, _ := a.GetDeviceAsUUID()
	return id
}
