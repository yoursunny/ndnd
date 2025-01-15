package crypto

import (
	"crypto/hmac"
	"crypto/sha256"

	enc "github.com/named-data/ndnd/std/encoding"
	"github.com/named-data/ndnd/std/ndn"
)

// HmacSigner is a Data signer that uses a provided HMAC key.
type HmacSigner struct {
	key []byte
}

func (signer *HmacSigner) Type() ndn.SigType {
	return ndn.SignatureHmacWithSha256
}

func (*HmacSigner) KeyLocator() enc.Name {
	return nil
}

func (*HmacSigner) EstimateSize() uint {
	return 32
}

func (signer *HmacSigner) Sign(covered enc.Wire) ([]byte, error) {
	mac := hmac.New(sha256.New, signer.key)
	for _, buf := range covered {
		_, err := mac.Write(buf)
		if err != nil {
			return nil, enc.ErrUnexpected{Err: err}
		}
	}
	return mac.Sum(nil), nil
}

func (*HmacSigner) Public() ([]byte, error) {
	return nil, ndn.ErrNoPubKey
}

// NewHmacSigner creates a Data signer that uses DigestSha256.
func NewHmacSigner(key []byte) ndn.Signer {
	return &HmacSigner{key}
}

func CheckHmacSig(sigCovered enc.Wire, sigValue []byte, key []byte) bool {
	mac := hmac.New(sha256.New, []byte(key))
	for _, buf := range sigCovered {
		_, err := mac.Write(buf)
		if err != nil {
			return false
		}
	}
	return hmac.Equal(mac.Sum(nil), sigValue)
}
