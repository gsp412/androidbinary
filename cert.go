package androidbinary

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/gsp412/androidbinary/libs/pkcs7"
)

// CertFile.
type CertFile struct {
	PublicKey string
	Cert      *x509.Certificate
}

func NewCertFile(b []byte) (c *CertFile, err error) {
	c = new(CertFile)

	_pkcs7, err := pkcs7.Parse(b)
	if err != nil {
		return
	}

	if len(_pkcs7.Certificates) == 0 {
		return nil, errors.New("public key not exist")
	}

	c.Cert = _pkcs7.Certificates[0]

	pk_bytes, err := x509.MarshalPKIXPublicKey(c.Cert.PublicKey)
	if err != nil {
		return
	}

	block := &pem.Block{
		Type  : "RSA PUBLIC KEY",
		Bytes :  pk_bytes,
	}

	c.PublicKey = string(pem.EncodeToMemory(block))

	return
}
