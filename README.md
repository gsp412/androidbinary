androidbinary
=====

Android binary file parser

## High Level API

### Parse APK files

``` go
package main

import (
    "github.com/gsp412/androidbinary/apk"
)

func main() {
    pkg, _ := apk.OpenFile("your-android-app.apk")
    defer pkg.Close()

    icon, _ := pkg.Icon(nil) // returns the icon of APK as image.Image
    
    pkgName := pkg.PackageName() // returns the package name.
    
    // values-zh-rCN
    config := &androidbinary.ResTableConfig{
        Language: [2]uint8{'z', 'h'},
        Country:  [2]uint8{'C', 'N'},
    }

    label, _ := apk.Label(config) // returns the APK lable.
    
    versionCode := apk.VersionCode() // returns the version code.
    
    versionName := apk.VersionName() // returns the version name.
    
    publicKey := apk.PublicKey() // returns the cret public key.
}
```

## Low Level API

### Parse XML binary

``` go
package main

import (
    "encoding/xml"

    "github.com/gsp412/androidbinary"
    "github.com/gsp412/androidbinary/apk"
)

func main() {
    f, _ := os.Open("AndroidManifest.xml")
    xml, _ := androidbinary.NewXMLFile(f)
    reader := xml.Reader()

    // read XML from reader
    var manifest apk.Manifest
    data, _ := ioutil.ReadAll(reader)
    xml.Unmarshal(data, &manifest)
}
```

### Parse Resource files

``` go
package main

import (
    "fmt"
    "github.com/gsp412/androidbinary"
)

func main() {
    f, _ := os.Open("resources.arsc")
    rsc, _ := androidbinary.NewTableFile(f)
    resource, _ := rsc.GetResource(androidbinary.ResID(0xCAFEBABE), nil)
    fmt.Println(resource)
}
```

### Parse Cert files

``` go
package main

import (
    "fmt"
    "github.com/gsp412/androidbinary"
)

func main() {
    b, _ := ioutil.ReadFile("CERT.RSA")

    c, _ := androidbinary.NewCertFile(b)

    fmt.Println(c.PublicKey)
}
```

## License

This software is released under the MIT License, see LICENSE.
