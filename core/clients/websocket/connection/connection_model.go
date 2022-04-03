package connection

import "github.com/KyaXTeam/go-core/v2/core/helpers/_struct"

// AuthByToken -> this is the library default auth type which is set in the Header
const AuthByToken = 1

// AuthByBearerToken -> this is the protocol standard -> which is set in the Header
const AuthByBearerToken = 2

// AuthBasic -> // Username and Passsword -> which is set in the Header
const AuthBasic = 3

// AuthByGETParamToken -> same Token but attached as GET Param -> which is set as Http GET Param
// Encrypted format of this auth type is not necessary, because the attacker will anyway get the auth token
const AuthByGETParamToken = 4

// Other authentication types are not indicated here
// More complex authentications should be done before using this!

// TODO: Security:
// for better authentication, it's necessary that the token be changed more often, or a special protocol should be used
// for best authentication

type AuthOptions struct {
	// AuthType -> this is the authentication type
	AuthType int `yaml:"auth_type" mapstructure:"auth_type" default:"1"`
	// Token -> it's used by AuthByToken & AuthByBearerToken
	Token string `yaml:"token" mapstructure:"token" default:""`
	// Username -> it's used by AuthBasic
	Username string `yaml:"username" mapstructure:"username" default:""`
	// Password -> it's used by AuthBasic
	Password string `yaml:"password" mapstructure:"password" default:""`
}

type Connection struct {
	// Connection params
	IsSecure string `yaml:"is_secure" mapstructure:"is_secure" default:"yes"`
	// Scheme -> wss is encrypted connection!
	Scheme string `yaml:"scheme" mapstructure:"scheme" default:"wss://"`
	Host   string `yaml:"host" mapstructure:"host" default:"localhost"`
	// Port -> the default is 0, in this case no port will be used!
	Port    uint16 `yaml:"port" mapstructure:"port" default:"0"`
	UriPath string `yaml:"uri_path" mapstructure:"uri_path" default:"/"`
	// AcceptCertificate -> Accept Self Signed Certificate
	AcceptCertificate string `yaml:"accept_certificate" mapstructure:"accept_certificate" default:"no"`

	// TODO: we should also add here CUUSTOM CERTIFICATES FOR BETTER SECURITY AND AUTHORITY
	// TODO: it's very easy to be an authorized certified owner and become a Man in Middle!

	// RequestHeader ->  // We don't need to have this exported in the config!
	// Header that's being sent to the Server
	// We will not be using Request here, if the user want's to modify, he will receive a callback
	// where he can handle
	// RequestHeader http.Header `yaml:"-"`

	// MaxRetries 0 - means no retries, -1 - means infinite, >0 - means a specific nr of retries:
	// How many times it should retry to connect to same host. It's better to not set to infinite because
	// if you have multiple connections, the infinite param it's better to set on websocket client itself
	MaxRetries int16 `yaml:"max_retries" mapstructure:"max_retries" default:"3"`
	// RetryTimeout -> 5 Seconds by default
	RetryTimeout uint16 `yaml:"retry_timeout" mapstructure:"retry_timeout" default:"5"`

	// Authentication
	// EnableAuth -> If authentication is needed
	EnableAuth string `yaml:"enable_auth" mapstructure:"enable_auth" default:"no"`
	// AuthType -> What type of authentication is needed
	AuthOptions AuthOptions `yaml:"auth_options" mapstructure:"auth_options"`
}

func DefaultConfig(connObj *Connection) (Connection, error) {
	if connObj == nil {
		connObj = &Connection{}
	}
	var _err error
	// Set the default values for the object!
	_err = _struct.SetDefaultValues(connObj)
	if _err != nil {
		return *connObj, _err
	}
	// Setting logger defaults
	_err = _struct.SetDefaultValues(&connObj.AuthOptions)
	if _err != nil {
		return *connObj, _err
	}

	return *connObj, _err
}
