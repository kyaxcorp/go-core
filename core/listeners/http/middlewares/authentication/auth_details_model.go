package authentication

import "time"

type AuthDetails struct {
	DeviceDetails    DeviceDetails
	UserDetails      UserDetails
	AuthTokenDetails AuthTokenDetails
}

// DeviceDetails -> These are the details
type DeviceDetails struct {
	// Main
	//     uint64
	// By using string we can set any type of identification...
	DeviceID     string
	DeviceUUID   string
	CreatedDate  time.Time
	UpdatedDate  time.Time
	Timezone     string
	IsAuthorized bool
	PushToken    string

	// Secondary
	PlatformType    string
	Name            string
	PlatformVersion string
	AppVersion      string
}

type UserDetails struct {
	//UserID    uint64
	// By using string we can set any type of identification...
	UserID    string
	Email     string
	FirstName string
	LastName  string
	Username  string
	IsActive  bool
	// UserType -> admin, client etc...
	UserType string
	// Role -> SuperAdmin, Admin, SalesMan etc...
	Role string

	// Secondary
	Phone1 string
	Phone2 string
}

type AuthTokenDetails struct {
	TokenID string
	// By using string we can set any type of identification...
	Token       string
	CreatedDate time.Time
	ExpireDate  time.Time
	TTL         uint64
}
