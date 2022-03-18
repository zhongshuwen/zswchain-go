package ecc

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"math/big"

	"github.com/zhongshuwen/gmsm/sm2"
	"github.com/zhongshuwen/zswchain-go/libbsuite/btcutil/base58"
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

func DecodeWIFBytes(wif string) (*DecodedWIF, error) {
	decoded := base58.Decode(wif)
	decodedLen := len(decoded)
	var compress bool

	// Length of base58 decoded WIF must be 32 bytes + an optional 1 byte
	// (0x01) if compressed, plus 1 byte for netID + 4 bytes of checksum.
	switch decodedLen {
	case 1 + PrivKeyBytesLen + 1 + 4:
		if decoded[33] != compressMagic {
			return nil, ErrMalformedPrivateKey
		}
		compress = true
	case 1 + PrivKeyBytesLen + 4:
		compress = false
	default:
		return nil, ErrMalformedPrivateKey
	}

	// Checksum is first four bytes of double SHA256 of the identifier byte
	// and privKey.  Verify this matches the final 4 bytes of the decoded
	// private key.
	var tosum []byte
	if compress {
		tosum = decoded[:1+PrivKeyBytesLen+1]
	} else {
		tosum = decoded[:1+PrivKeyBytesLen]
	}
	cksum := DoubleHashWIF(tosum)[:4]
	if !bytes.Equal(cksum, decoded[decodedLen-4:]) {
		return nil, ErrChecksumMismatch
	}

	netId := decoded[0]
	privKeyBytes := decoded[1 : 1+PrivKeyBytesLen]

	return &DecodedWIF{privKeyBytes, compress, netId}, nil
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
	encodeLen := 1 + PrivKeyBytesLen + 4
	if compressed {
		encodeLen++
	}

	a := make([]byte, 0, encodeLen)
	a = append(a, netID)
	// Pad and append bytes manually, instead of using Serialize, to
	// avoid another call to make.
	a = paddedAppend(PrivKeyBytesLen, a, wifBytes)
	if compressed {
		a = append(a, compressMagic)
	}
	cksum := DoubleHashWIF(a)[:4]
	a = append(a, cksum...)
	return base58.Encode(a)
}
