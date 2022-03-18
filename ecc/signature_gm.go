package ecc

import (
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
	pubKeyInst := sm2.Decompress(pubKey.Content)
	return pubKeyInst.VerifyDigest(hash, content[33:])
}

func (s *innerGMSignature) publicKey(content []byte, hash []byte) (out PublicKey, err error) {

	if err != nil {
		return out, err
	}

	return PublicKey{
		Curve:   CurveGM,
		Content: content[0:33],
		inner:   &innerGMPublicKey{},
	}, nil
}

func (s innerGMSignature) string(content []byte) string {
	checksum := ripemd160checksumHashCurve(content, CurveGM)
	buf := append(content[:], checksum...)
	return "SIG_GM_" + base58.Encode(buf)
}

func (s innerGMSignature) signatureMaterialSize() *int {
	return signatureDataSize
}
