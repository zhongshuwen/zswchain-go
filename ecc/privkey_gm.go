package ecc

import (
	"bytes"
	"fmt"

	"crypto/rand"

	"github.com/zhongshuwen/gmsm/sm2"
)

type innerGMPrivateKey struct {
	privKey *sm2.PrivateKey
}

func (k *innerGMPrivateKey) publicKey() PublicKey {
	return PublicKey{Curve: CurveGM, Content: sm2.Compress(&k.privKey.PublicKey), inner: &innerGMPublicKey{}}
}

func (k *innerGMPrivateKey) sign(hash []byte) (out Signature, err error) {
	if len(hash) != 32 {
		return out, fmt.Errorf("hash should be 32 bytes")
	} // 从文件读取数据
	signedBytes, err := k.privKey.SignDigest(rand.Reader, hash, nil) // 签名
	if err != nil {
		return out, err
	}
	outBytes := sm2.Compress(&k.privKey.PublicKey)
	outBytes = append(outBytes, signedBytes...)
	if len(outBytes) < 105 {
		outBytes = append(outBytes, bytes.Repeat([]byte{0}, 105-len(outBytes))...)
	}

	return Signature{Curve: CurveGM, Content: outBytes, inner: &innerGMSignature{}}, nil

	/*	compactSig, err := k.privKey.SignCanonical(btcec.S256(), hash)

		if err != nil {
			return out, fmt.Errorf("canonical, %s", err)
		}

		return Signature{Curve: CurveGM, Content: compactSig, inner: &innerGMSignature{}}, nil
	*/
}
func (k *innerGMPrivateKey) string() string {
	return "PVT_GM_" + EncodeWifSM2PrivateKey(0x80, k.privKey.D.Bytes(), true)
}
