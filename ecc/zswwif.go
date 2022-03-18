package ecc

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/zhongshuwen/gmsm/sm2"
	"github.com/zhongshuwen/zswchain-go/libbsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type DecodedWIF struct {
	data       []byte
	compressed bool
	netId      byte
}

// ErrMalformedPrivateKey describes an error where a WIF-encoded private
// key cannot be decoded due to being improperly formatted.  This may occur
// if the byte length is incorrect or an unexpected magic number was
// encountered.
var ErrMalformedPrivateKey = errors.New("malformed private key")

// ErrChecksumMismatch describes an error where decoding failed due to
// a bad checksum.
var ErrChecksumMismatch = errors.New("checksum mismatch")

// compressMagic is the magic byte used to identify a WIF encoding for
// an address created from a compressed serialized public key.
const compressMagic byte = 0x01
const PrivKeyBytesLen = 32

// 32byte
func zeroByteSlice() []byte {
	return []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}
}

// DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
func DoubleHashWIF(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func DoubleHashWIFSuffix(b []byte, suffix string) []byte {
	first := sha256.Sum256(append(b, []byte(suffix)...))
	second := sha256.Sum256(first[:])
	return second[:]
}
func Ripe160SuffixStringChecksum(b []byte, suffix string) []byte {
	h := ripemd160.New()
	_, _ = h.Write(b) // this implementation has no error path

	// FIXME: this seems to be only rolled out to the `SIG_` things..
	// proper support for importing `EOS` keys isn't rolled out into `dawn4`.
	_, _ = h.Write([]byte(suffix)) // conditionally ?
	sum := h.Sum(nil)
	return sum[:4]
}

func DecodeWIFBytes(wif string, suffix string) ([]byte, error) {
	decoded := base58.Decode(wif)
	decodedLen := len(decoded)
	println("decodedLength=%d, hex=%s", decodedLen, hex.EncodeToString(decoded))
	if decodedLen != 36 {
		return nil, ErrMalformedPrivateKey
	}
	cksum := Ripe160SuffixStringChecksum(decoded[:PrivKeyBytesLen], suffix)

	if !bytes.Equal(cksum, decoded[decodedLen-4:]) {
		return nil, ErrChecksumMismatch
	}

	return decoded[:PrivKeyBytesLen], nil
}
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}

func ReadSM2PrivateKeyFromBytes(d []byte) (*sm2.PrivateKey, error) {
	c := sm2.P256Sm2()
	k := new(big.Int).SetBytes(d)
	params := c.Params()
	one := new(big.Int).SetInt64(1)
	n := new(big.Int).Sub(params.N, one)
	if k.Cmp(n) >= 0 {
		return nil, errors.New("privateKey's d is overflow.")
	}
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}
func EncodeWifSM2PrivateKey(netID byte, wifBytes []byte, compressed bool) string {
	// Precalculate size.  Maximum number of bytes before base58 encoding
	// is one byte for the network, 32 bytes of private key, possibly one
	// extra byte if the pubkey is to be compressed, and finally four
	// bytes of checksum.
	cksum := Ripe160SuffixStringChecksum(wifBytes, "GM")
	encodeLen := PrivKeyBytesLen + 4
	a := make([]byte, 0, encodeLen)
	a = append(a, wifBytes...)
	a = append(a, cksum...)
	return base58.Encode(a)
}
