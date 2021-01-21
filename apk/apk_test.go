package apk

import (
	_ "image/jpeg"
	_ "image/png"
	"testing"

	"github.com/gsp412/androidbinary"
	"github.com/stretchr/testify/assert"
)

const EXPECT_PUBLIC_KEY = `-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCR+z5LObMtjaUQOhpAoFpP2eka
xTA5KcwTdNnNeIzdh1izCk6nYOsP42KNhKfVbv4JdvFRgxjne+dW0jUUM8Ank/wS
H3gcbE0Y0//2XvKFuwGT5278riXmvmzfZlCEBa4DTPGCcMzOISX2dDGMH70NbE3y
clqOCytA6nnzYrs8gwIDAQAB
-----END RSA PUBLIC KEY-----
`

func TestParseAPKFile(t *testing.T) {
	apk, err := OpenFile("testdata/chineselunar.apk")
	if !assert.NoError(t, err) {
		return
	}
	defer apk.Close()

	icon, err := apk.Icon(nil)
	assert.NoError(t, err)
	assert.NotNil(t, icon)

	config := &androidbinary.ResTableConfig{
		Language: [2]uint8{'z', 'h'},
		Country:  [2]uint8{'C', 'N'},
	}

	label, err := apk.Label(config)
	assert.NoError(t, err)
	assert.Equal(t, "中国农历", label)
	t.Log("app label:", label)

	manifest := apk.Manifest()
	assert.Equal(t, manifest.SDK.Min.MustInt32(), int32(3))

	mainActivity, err := apk.MainActivity()
	assert.NoError(t, err)
	assert.Equal(t, ".ChineseLunar", mainActivity)

	packageName := apk.PackageName()
	assert.Equal(t, "since2006.apps.chineselunar", packageName)
	t.Log("package name:", packageName)

	versionCode := apk.VersionCode()
	assert.Equal(t, int32(1), versionCode)
	t.Log("version code:", versionCode)

	versionName := apk.VersionName()
	assert.Equal(t, "1.0", versionName)
	t.Log("version name:", versionName)

	publicKey := apk.PublicKey()
	assert.Equal(t, EXPECT_PUBLIC_KEY, publicKey)
	t.Log(publicKey)

}
