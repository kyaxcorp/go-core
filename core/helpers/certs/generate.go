package certs

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/kyaxcorp/go-core/core/helpers/_struct/defaults"
	"github.com/kyaxcorp/go-core/core/helpers/err"
	"github.com/kyaxcorp/go-core/core/helpers/file"
	"github.com/kyaxcorp/go-core/core/helpers/filesystem"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

// Defaults
/*var (
	host       = flag.String("host", "", "Comma-separated hostnames and IPs to generate a certificate for")
	validFrom  = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
	validFor   = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
	isCA       = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
	rsaBits    = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
	ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	ed25519Key = flag.Bool("ed25519", false, "Generate an Ed25519 key")
)*/
type CertGeneration struct {
	Host         string        `default:""`                    // "Comma-separated hostnames and IPs to generate a certificate for"
	ValidFrom    string        `default:"Jan 1 00:00:00 2021"` // Creation date formatted as Jan 1 15:04:05 2011
	ValidForDays time.Duration `default:"36500"`               // "Duration that certificate is valid for"
	IsCA         bool          `default:"false"`               // "whether this cert should be its own Certificate Authority"
	RSABits      int           `default:"2048"`                // "Size of RSA key to generate. Ignored if --ecdsa-curve is set"
	EcdsaCurve   string        `default:""`                    // "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521"
	ED25519Key   bool          `default:"false"`               // "Generate an Ed25519 key"

	// By default if they are empty... they'll be saved in the default directory!
	KeyPath  string `default:""`
	CertPath string `default:""`
}

func GenerateCerts(scope string, options *CertGeneration) error {
	var _err error
	_err = defaults.Set(options)
	if _err != nil {
		return _err
	}

	savePath := GetCertsFullPathByScope(scope)
	var certPath string
	if options.CertPath == "" {
		certPath = savePath + filesystem.DirSeparator() + "cert.pem"
		options.CertPath = certPath
	} else {
		certPath = options.CertPath
	}

	var keyPath string
	if options.KeyPath == "" {
		keyPath = savePath + filesystem.DirSeparator() + "cert.key"
		options.KeyPath = keyPath
	} else {
		keyPath = options.KeyPath
	}

	if len(options.Host) == 0 {
		return err.New(0, "Missing required host parameter")
	}

	var priv interface{}
	switch options.EcdsaCurve {
	case "":
		if options.ED25519Key {
			_, priv, _err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, _err = rsa.GenerateKey(rand.Reader, options.RSABits)
		}
	case "P224":
		priv, _err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, _err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, _err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, _err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		//log.Fatalf("Unrecognized elliptic curve: %q", options.EcdsaCurve)
		return err.New(0, fmt.Sprintf("Unrecognized elliptic curve: %q", options.EcdsaCurve))
	}
	if _err != nil {
		// log.Fatalf("Failed to generate private key: %v", _err)
		return err.New(0, fmt.Sprintf("Failed to generate private key: %v", _err))
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	var notBefore time.Time
	if len(options.ValidFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, _err = time.Parse("Jan 2 15:04:05 2006", options.ValidFrom)
		if _err != nil {
			// log.Fatalf("Failed to parse creation date: %v", _err)
			return err.New(0, fmt.Sprintf("Failed to parse creation date: %v", _err))
		}
	}

	notAfter := notBefore.Add(options.ValidForDays * time.Hour * 24)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _err := rand.Int(rand.Reader, serialNumberLimit)
	if _err != nil {
		// log.Fatalf("Failed to generate serial number: %v", _err)
		return err.New(0, fmt.Sprintf("Failed to generate serial number: %v", _err))
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(options.Host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if options.IsCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, _err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if _err != nil {
		// log.Fatalf("Failed to create certificate: %v", _err)
		return err.New(0, fmt.Sprintf("Failed to create certificate: %v", _err))
	}

	certOut, _err := os.Create(certPath)
	if _err != nil {
		//log.Fatalf("Failed to open cert.pem for writing: %v", _err)
		return err.New(0, fmt.Sprintf("Failed to open cert.pem for writing: %v", _err))

	}
	if er := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); er != nil {
		//log.Fatalf("Failed to write data to cert.pem: %v", err)
		return err.New(0, fmt.Sprintf("Failed to write data to cert.pem: %v", er))
	}
	if er := certOut.Close(); er != nil {
		//log.Fatalf("Error closing cert.pem: %v", err)
		return err.New(0, fmt.Sprintf("Error closing cert.pem: %v", er))
	}
	//log.Print("wrote cert.pem\n")

	keyOut, _err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if _err != nil {
		//log.Fatalf("Failed to open key.pem for writing: %v", _err)
		return err.New(0, fmt.Sprintf("Failed to open key.pem for writing: %v", _err))
	}
	privBytes, _err := x509.MarshalPKCS8PrivateKey(priv)
	if _err != nil {
		//log.Fatalf("Unable to marshal private key: %v", _err)
		return err.New(0, fmt.Sprintf("Unable to marshal private key: %v", _err))
	}
	if er := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); er != nil {
		//log.Fatalf("Failed to write data to key.pem: %v", err)
		return err.New(0, fmt.Sprintf("Failed to write data to key.pem: %v", er))
	}
	if er := keyOut.Close(); er != nil {
		//log.Fatalf("Error closing key.pem: %v", err)
		return err.New(0, fmt.Sprintf("Error closing key.pem: %v", er))
	}
	//log.Print("wrote key.pem\n")
	// Check if files exist

	if !file.Exists(options.CertPath) {
		return err.New(0, "cert file doesn't exist", options.CertPath)
	}
	if !file.Exists(options.KeyPath) {
		return err.New(0, "key file doesn't exist", options.KeyPath)
	}

	return nil
}
