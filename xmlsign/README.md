# xmlsign

XML Digital Signatures implemented in pure Go.

## Installation

Install `xmlsign` into your `$GOPATH` using `go get`:

```
$ go get github.com/govenue/encoding/xmlsign
```

## Usage

### Signing

```go
package main

import (
    "github.com/govenue/encoding/xmlsign"
    "github.com/govenue/encoding/xmltree"
)

func main() {
    // Generate a key and self-signed certificate for signing
    randomKeyStore := dsig.RandomKeyStoreForTest()
    ctx := dsig.NewDefaultSigningContext(randomKeyStore)
    elementToSign := &xmltree.Element{
        Tag: "ExampleElement",
    }
    elementToSign.CreateAttr("ID", "id1234")

    // Sign the element
    signedElement, err := ctx.SignEnveloped(elementToSign)
    if err != nil {
        panic(err)
    }

    // Serialize the signed element. It is important not to modify the element
    // after it has been signed - even pretty-printing the XML will invalidate
    // the signature.
    doc := xmltree.NewDocument()
    doc.SetRoot(signedElement)
    str, err := doc.WriteToString()
    if err != nil {
        panic(err)
    }

    println(str)
}
```

### Signature Validation

```go
// Validate an element against a root certificate
func validate(root *x509.Certificate, el *xmltree.Element) {
    // Construct a signing context with one or more roots of trust.
    ctx := dsig.NewDefaultValidationContext(&dsig.MemoryX509CertificateStore{
        Roots: []*x509.Certificate{root},
    })

    // It is important to only use the returned validated element.
    // See: https://www.w3.org/TR/xmldsig-bestpractices/#check-what-is-signed
    validated, err := ctx.Validate(el)
    if err != nil {
        panic(err)
    }

    doc := xmltree.NewDocument()
    doc.SetRoot(validated)
    str, err := doc.WriteToString()
    if err != nil {
        panic(err)
    }

    println(str)
}
```
