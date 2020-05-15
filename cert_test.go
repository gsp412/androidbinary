package androidbinary

import (
	"fmt"
	"io/ioutil"
	"testing"

	"crypto/md5"
)

const EXPECT_PUBLIC_KEY = `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsaK77ndL/WZSt7j/jEHh
Fhwwc/K9DfYkn5X8rs4LCMrrOY+qkRFmgu8HojBtqhuuATwaZ6m3ZeJFxH41d6NN
6TgE4KBJvyeb7F3FWUOF36LIdmR9yD2q6ReJwfHfctKd9oFxo+iz9rwk/WjMv/yM
yFQ7AUWenc0DqY3u2Zjq46j2qBUeRGtWZN2PUyGb58s/qlWMdEbT7DoF0VXMQczn
h3OAOgXJbcaddJTvl6iDV3bpbJX1OLrnQ04boSpLhJd3VDxyG0qQYyYq3aHyEFec
BATwQm4Un4H2DztrPh79xDStQaFdmK1BAvtBziMmcKiC7Q0Pt+c188yBdCkUcIUM
iQIDAQAB
-----END RSA PUBLIC KEY-----
`

func TestNewCertFile(t *testing.T) {
	b, _ := ioutil.ReadFile("testdata/CERT.RSA")

	c, err := NewCertFile(b)
	if nil != err {
		t.Error(err.Error())
	}
	println("START", c.PublicKey, "END")

	println("START", EXPECT_PUBLIC_KEY, "END")


	fmt.Printf("%x\n", md5.Sum([]byte(c.PublicKey)))

	fmt.Printf("%x\n", md5.Sum([]byte(EXPECT_PUBLIC_KEY)))

	if c.PublicKey != EXPECT_PUBLIC_KEY {
		t.Error("Does not match expectations")
	}

}
