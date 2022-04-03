package connection

type Connection struct {
	// Connection params
	IsSecure string `yaml:"is_secure" mapstructure:"is_secure" default:"yes"`
	// Scheme -> wss is encrypted connection!
	//Scheme string `yaml:"scheme" mapstructure:"scheme" default:"wss://"`
	Host string `yaml:"host" mapstructure:"host" default:"localhost"`
	// Port -> the default is 30001, it's the SSL one!
	Port uint16 `yaml:"port" mapstructure:"port" default:"30001"`
	//UriPath string `yaml:"uri_path" mapstructure:"uri_path" default:"/"`
	// AcceptCertificate -> Accept Self Signed Certificate, temporarily we will accept... but later we should disable
	// And add specific certificates, which will be embedded in each of the Applications!!! TODO: ADD CUSTOM CERTS IN APP!
	// TODO: disable Accept Cert, and accept only the one embedded or signed by the hostname!!!
	AcceptCertificate string `yaml:"accept_certificate" mapstructure:"accept_certificate" default:"yes"`

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
}
