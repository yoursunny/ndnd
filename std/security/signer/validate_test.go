package signer_test

import (
	"encoding/base64"
	"testing"

	enc "github.com/named-data/ndnd/std/encoding"
	spec "github.com/named-data/ndnd/std/ndn/spec_2022"
	sig "github.com/named-data/ndnd/std/security/signer"
	tu "github.com/named-data/ndnd/std/utils/testutils"
	"github.com/stretchr/testify/require"
)

// testValidateSelfSigned tests the self-signed certificate validation.
// These certificates are generated by the ndn-cxx library, and are used for
// checking the interoperability of the validation logic.
func testValidateSelfSigned(t *testing.T, certB64 string) {
	certWire := tu.NoErr(base64.StdEncoding.DecodeString(certB64))

	// Helper to test the signature validation result
	test := func(result bool) {
		certData, sigCov, err := spec.Spec{}.ReadData(enc.NewBufferReader(certWire))
		require.NoError(t, err)
		require.Equal(t, result, tu.NoErr(sig.ValidateData(certData, sigCov, certData)))
	}

	// Test with valid signature
	test(true)

	// Tamper with the signature (guess)
	certWire[len(certWire)-10] ^= 0x01
	test(false)
}

// TestEd25519ValidateInterop tests the Ed25519 self-signed certificate validation.
func TestEd25519ValidateInterop(t *testing.T) {
	tu.SetT(t)
	testValidateSelfSigned(t, `
Bv0BCgc1CAxFZDI1NTE5LWRlbW8IA0tFWQgQNWE2MTVkYjdjZjA2MDNiNQgEc2Vs
ZjYIAAABgQD8AY0UCRgBAhkEADbugBUsMCowBQYDK2VwAyEAQxUZBL+3I3D4oDIJ
tJvuCTguHM7AUbhlhA/wu8ZhrkwWVhsBBRwnByUIDEVkMjU1MTktZGVtbwgDS0VZ
CBA1YTYxNWRiN2NmMDYwM2I1/QD9Jv0A/g8xOTcwMDEwMVQwMDAwMDD9AP8PMjAy
MjA1MjZUMTUyODQ0F0DAAWCZzxQSCAV0tluFDry5aT1b+EgoYgT1JKxbKVb/tINx
M43PFy/2hDe8j61PuYD9tCah0TWapPwfXWi3fygA`)
}

// TestEccValidateInterop tests the ECC self-signed certificate validation.
func TestEccValidateInterop(t *testing.T) {
	tu.SetT(t)
	testValidateSelfSigned(t, `
Bv0BLAcrCANuZG4IBWFsaWNlCANLRVkICH+xyxHFjoVZCARzZWxmNggAAAGUjIqk
mBQJGAECGQQANu6AFVswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATMnLNLJ+sy
AuZ0XYKJskUyVDxBLl5SzljivyaUVizG+yKLePyrmDE+VE8rXNhzENqTxuJNXB7X
RRLCzfAwjsxOFkwbAQMcHQcbCANuZG4IBWFsaWNlCANLRVkICH+xyxHFjoVZ/QD9
Jv0A/g8yMDI1MDEyMlQwNTQ0NDP9AP8PMjA0NTAxMTdUMDU0NDQzF0cwRQIgPBz3
hoMiTZzX/cSamxah0qaXNkveGkZqqao2nLQnC/sCIQCGf7akPnoYFMq40sRV4nHF
deWX6c79riyoiuURUu0Vhw==`)
}

// TestRsaValidateInterop tests the RSA self-signed certificate validation.
func TestRsaValidateInterop(t *testing.T) {
	tu.SetT(t)
	testValidateSelfSigned(t, `
Bv0CsAcpCANuZG4IA3JzYQgDS0VZCAgPSfvRdJCkjggEc2VsZjYIAAABlIyNmh0U
CRgBAhkEADbugBX9ASYwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDP
EGpk1NbeNeoLjL3JtKXa5MKAix82Qq58QXo2Y+a5NDCr0bFoATcg/KxvAWErZ6oc
q85D5RLE2TmtMH4H3gctBTKlKTwwG/JHyKGXPSs8UbQ31HxSlSQEqL9JNLL9oPwp
dsR1YfBK89nTMVg7w/e3T1wGCehWImr4JH1X8tjUa6DqCMWsMolVcf7E2XkocTdJ
WtmuxkBullqWvvIBAKHuNQSwamMoSlX9WFeWgPUbRRaxTZ5RNVdLHbsgR1CrZWDH
q3ewb1+AkwQbFdRTjuGlgpdpxqQDvWXAicwDlSW8PbodaKbeoXg9extvsS8n+/nx
IeNM2BDE2B3JsMBeMXsDAgMBAAEWShsBARwbBxkIA25kbggDcnNhCANLRVkICA9J
+9F0kKSO/QD9Jv0A/g8yMDI1MDEyMlQwNTQ3NTf9AP8PMjA0NTAxMTdUMDU0NzU3
F/0BAJ565DvyxAztpdnSCKTYMrhvwzN7+kuUmQMJzTPKloVIeFFbEwXmZkrrUbSQ
iSGvWiyNMYdJ9daUnmr2MqKfY0T9X0Qso/Ri1A/veO1l3dy+9X2Bwpz+pbrmrRXH
RAxCSnqQCN7b5rBVzWxcAG1JA/FUmOPMhaOdVjuMjK08a5Q5kJJU+AtIWLn2ljvL
pg0fJD/j1RB5KfGnu0dPDoGVwd2Tt1ODUvheg49LPwcTH/XoWJLJ0qhC6xfFT3ph
1Nto5tUCxLGwU5N9jlah96YNRy6f+1tZX+6v6SOOj9tVQZBXX+/3baK/U7Z0uFg/
kkmngSpSseV5W0LXjiRx+4BUOFE=`)
}
