package ecc

import (
	"fmt"

	"github.com/zhongshuwen/gmsm/sm2"
	"github.com/zhongshuwen/zswchain-go/libbsuite/btcutil/base58"
)

type innerGMSignature struct {
}

func newInnerGMSignature() innerSignature {
	return &innerGMSignature{}
}

// verify checks the signature against the pubKey. `hash` is a sha256
// hash of the payload to verify.
func (s *innerGMSignature) verify(content []byte, hash []byte, pubKey PublicKey) bool {
	sigLength := int(content[33+1]) + 2
	if (len(content) - 33) < sigLength {
		return false
	}
	pubKeyInst := sm2.Decompress(pubKey.Content)
	return pubKeyInst.VerifyDigest(hash, content[33:(sigLength+33)])
}

func (s *innerGMSignature) publicKey(content []byte, hash []byte) (out PublicKey, err error) {

	sigLength := int(content[33+1]) + 2
	if (len(content) - 33) < sigLength {
		return out, fmt.Errorf("invalid signature length")
	}
	pubKeyInst := sm2.Decompress(content[0:33])
	if pubKeyInst.VerifyDigest(hash, content[33:(sigLength+33)]) {
		return PublicKey{
			Curve:   CurveGM,
			Content: content[0:33],
			inner:   &innerGMPublicKey{},
		}, nil
	} else {
		return out, fmt.Errorf("invalid signature in recovery")
	}
}

func (s innerGMSignature) string(content []byte) string {
	checksum := ripemd160checksumHashCurve(content, CurveGM)
	buf := append(content[:], checksum...)
	return "SIG_GM_" + base58.Encode(buf)
}

func (s innerGMSignature) signatureMaterialSize() *int {
	return gmSignatureDataSize
}
