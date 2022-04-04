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

func CompressReal(a *sm2.PublicKey) []byte {
	data := sm2.Compress(a)
	if data[0] == 1 {
		data[0] = 3
	} else {
		data[0] = 2
	}
	return data
}

func DecompressReal(a []byte) *sm2.PublicKey {
	aCopy := make([]byte, len(a))
	copy(aCopy, a)
	if aCopy[0] == 3 {
		aCopy[0] = 1
	} else {
		aCopy[0] = 0
	}

	return sm2.Decompress(aCopy)
}

func (k *innerGMPrivateKey) publicKey() PublicKey {
	return PublicKey{Curve: CurveGM, Content: CompressReal(&k.privKey.PublicKey), inner: &innerGMPublicKey{}}
}

func (k *innerGMPrivateKey) sign(hash []byte) (out Signature, err error) {
	if len(hash) != 32 {
		return out, fmt.Errorf("hash should be 32 bytes")
	} // 从文件读取数据
	signedBytes, err := k.privKey.SignDigest(rand.Reader, hash, nil) // 签名
	if err != nil {
		return out, err
	}
	outBytes := CompressReal(&k.privKey.PublicKey)
	outBytes = append(outBytes, signedBytes...)
	if len(outBytes) < 105 {
		outBytes = append(outBytes, bytes.Repeat([]byte{0}, 105-len(outBytes))...)
	}

	return Signature{Curve: CurveGM, Content: outBytes, inner: &innerGMSignature{}}, nil

}
func (k *innerGMPrivateKey) string() string {
	return "PVT_GM_" + EncodeWifSM2PrivateKey(0x80, k.privKey.D.Bytes(), true)
}
