package security_test

import (
	"time"

	enc "github.com/named-data/ndnd/std/encoding"
)

// This file provides testing constants for the security package.

// Name: /ndn/alice/KEY/cK%1D%A4%E1%5B%91%CF
// SigType: Ed25519
const KEY_ALICE = `
BroHGwgDbmRuCAVhbGljZQgDS0VZCAhY3Lb6ZymkghQDGAEJFTAwLgIBADAFBgMr
ZXAEIgQgG8Z7YkpBxVDqIFIm6EJlfCujheiW0262kJUj5vkmLDkWIhsBBRwdBxsI
A25kbggFYWxpY2UIA0tFWQgIWNy2+mcppIIXQB19BDvrahM3DR1DV7ESosKW4uzE
Z27QIGFgKR4LEuflvnSZZGFRKFxTUF5S2f/ZO/4B4NoxrF1ZOHD9NCWTkwI=
`

const KEY_ALICE_PEM = `-----BEGIN NDN KEY-----
Name: /ndn/alice/KEY/X%DC%B6%FAg%29%A4%82
SigType: Ed25519

BroHGwgDbmRuCAVhbGljZQgDS0VZCAhY3Lb6ZymkghQDGAEJFTAwLgIBADAFBgMr
ZXAEIgQgG8Z7YkpBxVDqIFIm6EJlfCujheiW0262kJUj5vkmLDkWIhsBBRwdBxsI
A25kbggFYWxpY2UIA0tFWQgIWNy2+mcppIIXQB19BDvrahM3DR1DV7ESosKW4uzE
Z27QIGFgKR4LEuflvnSZZGFRKFxTUF5S2f/ZO/4B4NoxrF1ZOHD9NCWTkwI=
-----END NDN KEY-----
`

// Name: /ndn/KEY/%27%C4%B2%2A%9F%7B%81%27/ndn/v=1651246789556
// SigType: ECDSA-SHA256
// Validity: 2022-04-29 15:39:50 +0000 UTC - 2026-12-31 23:59:59 +0000 UTC
const CERT_ROOT = `
Bv0BSwcjCANuZG4IA0tFWQgIJ8SyKp97gScIA25kbjYIAAABgHX6c7QUCRgBAhkE
ADbugBVbMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPuDnW4oq0mULLT8PDXh0
zuBg+0SJ1yPC85jylUU+hgxX9fDNyjlykLrvb1D6IQRJWJHMKWe6TJKPUhGgOT65
8hZyGwEDHBYHFAgDbmRuCANLRVkICCfEsiqfe4En/QD9Jv0A/g8yMDIyMDQyOVQx
NTM5NTD9AP8PMjAyNjEyMzFUMjM1OTU5/QECKf0CACX9AgEIZnVsbG5hbWX9AgIV
TkROIFRlc3RiZWQgUm9vdCAyMjA0F0gwRgIhAPYUOjNakdfDGh5j9dcCGOz+Ie1M
qoAEsjM9PEUEWbnqAiEApu0rg9GAK1LNExjLYAF6qVgpWQgU+atPn63Gtuubqyg=
`

const CERT_ROOT_PEM = `-----BEGIN NDN CERT-----
Name: /ndn/KEY/%27%C4%B2%2A%9F%7B%81%27/ndn/v=1651246789556
SigType: Sha256WithEcdsa
SignerKey: /ndn/KEY/%27%C4%B2%2A%9F%7B%81%27
Validity: 2022-04-29 15:39:50 +0000 UTC - 2026-12-31 23:59:59 +0000 UTC

Bv0BSwcjCANuZG4IA0tFWQgIJ8SyKp97gScIA25kbjYIAAABgHX6c7QUCRgBAhkE
ADbugBVbMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPuDnW4oq0mULLT8PDXh0
zuBg+0SJ1yPC85jylUU+hgxX9fDNyjlykLrvb1D6IQRJWJHMKWe6TJKPUhGgOT65
8hZyGwEDHBYHFAgDbmRuCANLRVkICCfEsiqfe4En/QD9Jv0A/g8yMDIyMDQyOVQx
NTM5NTD9AP8PMjAyNjEyMzFUMjM1OTU5/QECKf0CACX9AgEIZnVsbG5hbWX9AgIV
TkROIFRlc3RiZWQgUm9vdCAyMjA0F0gwRgIhAPYUOjNakdfDGh5j9dcCGOz+Ie1M
qoAEsjM9PEUEWbnqAiEApu0rg9GAK1LNExjLYAF6qVgpWQgU+atPn63Gtuubqyg=
-----END NDN CERT-----
`

// Key names for the above keys and certs
var KEY_ALICE_NAME, _ = enc.NameFromStr("/ndn/alice/KEY/X%DC%B6%FAg%29%A4%82")
var KEY_ROOT_NAME, _ = enc.NameFromStr("/ndn/KEY/%27%C4%B2%2A%9F%7B%81%27")

// Issuer name for testing
var ISSUER = enc.NewGenericComponent("myissuer")

// Validity times for testing
var T1 = time.Date(2022, 4, 29, 15, 39, 50, 0, time.UTC)
var T2 = time.Date(2089, 12, 31, 23, 59, 59, 0, time.UTC)
