package ecc

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestK1PrivateToPublic(t *testing.T) {
	wif := "5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss"
	privKey, err := NewPrivateKey(wif)
	require.NoError(t, err)

	pubKey := privKey.PublicKey()

	pubKeyString := pubKey.String()
	assert.Equal(t, PublicKeyPrefixCompat+"859gxfnXyUriMgUeThh1fWv3oqcpLFyHa3TfFYC4PK2HqhToVM", pubKeyString)
}

func TestPrefixedK1PrivateToPublic(t *testing.T) {
	wif := "PVT_K1_5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss"
	privKey, err := NewPrivateKey(wif)
	require.NoError(t, err)

	pubKey := privKey.PublicKey()

	pubKeyString := pubKey.String()
	assert.Equal(t, PublicKeyPrefixCompat+"859gxfnXyUriMgUeThh1fWv3oqcpLFyHa3TfFYC4PK2HqhToVM", pubKeyString)
}

func TestR1PrivateToPublic(t *testing.T) {
	encodedPrivKey := "PVT_R1_2o5WfMRU4dTp23pbcbP2yn5MumQzSMy3ayNQ31qi5nUfa2jdWC"
	_, err := NewPrivateKey(encodedPrivKey)
	require.NoError(t, err)

	// FIXME: Actual retrieval of publicKey from privateKey for R1 is not done yet, disable this check
	// pubKey := privKey.PublicKey()

	//pubKeyString := pubKey.String()
	//assert.Equal(t, "PUB_R1_0000000000000000000000000000000000000000000000", pubKeyString)
}

func TestGMPrivateToPublic(t *testing.T) {
	encodedPrivKey := "PVT_GM_98aQksvjx9xpD2xtSECphUA3NykuZxmX89ZL45byCbJLFwmcH"
	privKey, err := NewPrivateKey(encodedPrivKey)
	require.NoError(t, err)

	// FIXME: Actual retrieval of publicKey from privateKey for R1 is not done yet, disable this check
	pubKey := privKey.PublicKey()

	pubKeyString := pubKey.String()
	assert.Equal(t, "PUB_GM_7gyUV7Q8YF59EAsWgoQcYKdrYy5NEGFhCzx5pdJFtd7rURgaic", pubKeyString)
}

func TestGMPrivateToPrivate(t *testing.T) {
	encodedPrivKey := "PVT_GM_98aQksvjx9xpD2xtSECphUA3NykuZxmX89ZL45byCbJLFwmcH"
	privKey, err := NewPrivateKey(encodedPrivKey)
	require.NoError(t, err)

	// FIXME: Actual retrieval of publicKey from privateKey for R1 is not done yet, disable this check
	privKeyStringAfter := privKey.String()

	assert.Equal(t, encodedPrivKey, privKeyStringAfter)
}

func TestNewPublicKeyAndSerializeCompress(t *testing.T) {
	// Copied test from eosjs(-.*)?
	key, err := NewPublicKey(PublicKeyPrefixCompat + "6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV")
	require.NoError(t, err)
	assert.Equal(t, "02c0ded2bc1f1305fb0faac5e6c03ee3a1924234985427b6167ca569d13df435cf", hex.EncodeToString(key.Content))
}

func TestNewRandomPrivateKey(t *testing.T) {
	key, err := NewRandomPrivateKey()
	require.NoError(t, err)
	// taken from eosiojs-ecc:common.test.js:12
	assert.Regexp(t, "^5[HJK].*", key.String())
}

func TestPrivateKeyValidity(t *testing.T) {
	tests := []struct {
		in    string
		valid bool
	}{
		{"5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss", true},
		{"5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjsm", false},
	}

	for _, test := range tests {
		_, err := NewPrivateKey(test.in)
		if test.valid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
			assert.Equal(t, "checksum mismatch", err.Error())
		}
	}
}

func TestPublicKeyValidity(t *testing.T) {
	tests := []struct {
		in  string
		err error
	}{
		{PublicKeyPrefixCompat + "859gxfnXyUriMgUeThh1fWv3oqcpLFyHa3TfFYC4PK2HqhToVM", nil},
		{"MMM859gxfnXyUriMgUeThh1fWv3oqcpLFyHa3TfFYC4PK2HqhToVM", fmt.Errorf(`public key should start with "PUB_K1_", "PUB_R1_", "PUB_WA_", "PUB_GM_" or the old "` + PublicKeyPrefixCompat + `"`)},
		{PublicKeyPrefixCompat + "859gxfnXyUriMgUeThh1fWv3oqcpLFyHa3TfFYC4PK2HqhTo", fmt.Errorf("public key checksum failed, found 0e2e1094 but expected 169c2652")},
	}

	for idx, test := range tests {
		_, err := NewPublicKey(test.in)
		if test.err == nil {
			assert.NoError(t, err, fmt.Sprintf("test %d with key %q", idx, test.in))
		} else {
			assert.Error(t, err)
			assert.Equal(t, test.err.Error(), err.Error())
		}
	}
}

func TestK1Signature(t *testing.T) {
	wif := "5KYZdUEo39z3FPrtuX2QbbwGnNP5zTd7yyr2SC1j299sBCnWjss"
	privKey, err := NewPrivateKey(wif)
	require.NoError(t, err)

	cnt := []byte("hi")
	digest := sigDigest([]byte{}, cnt, nil)
	signature, err := privKey.Sign(digest)
	require.NoError(t, err)

	assert.True(t, signature.Verify(digest, privKey.PublicKey()))
}

func TestGMSignature(t *testing.T) {
	wif := "PVT_GM_98aQksvjx9xpD2xtSECphUA3NykuZxmX89ZL45byCbJLFwmcH"
	privKey, err := NewPrivateKey(wif)
	require.NoError(t, err)

	cnt := []byte("hi")
	digest := sigDigest([]byte{}, cnt, nil)
	signature, err := privKey.Sign(digest)
	require.NoError(t, err)

	assert.True(t, signature.Verify(digest, privKey.PublicKey()))
}

func TestGMSignature2(t *testing.T) {
	wif := "PVT_GM_98aQksvjx9xpD2xtSECphUA3NykuZxmX89ZL45byCbJLFwmcH"
	privKey, err := NewPrivateKey(wif)
	require.NoError(t, err)

	//cnt := []byte("hi")
	//digest := sigDigest([]byte{}, cnt, nil)
	digest, err := hex.DecodeString("19c020ada6b096cb2510be0ee88b82c0b56ce94eb7e12e1d6b545dfe96f0a903")

	signature, err := privKey.Sign(digest)
	require.NoError(t, err)

	assert.True(t, signature.Verify(digest, privKey.PublicKey()))
	//assert.Equal(t, "blah", signature.String())
}
func TestR1Signature(t *testing.T) {
	encodedPrivKey := "PVT_R1_2o5WfMRU4dTp23pbcbP2yn5MumQzSMy3ayNQ31qi5nUfa2jdWC"
	privKey, err := NewPrivateKey(encodedPrivKey)
	require.NoError(t, err)

	cnt := []byte("hi")
	digest := sigDigest([]byte{}, cnt, nil)
	_, err = privKey.Sign(digest)
	assert.Error(t, err)
	assert.Equal(t, "R1 not supported", err.Error())
}

func TestDecodeSM2PEMPublicKey(t *testing.T) {
	encodedSm2PemPublicKey := `
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEVCZLZTdGzfogF3RKdz/8SXSNU3Zq
LDrwFWSWWbiOlacoZI9DzcEj8//lPbhy0AGb50F2u9ZO8LSxk8QNPEffXg==
-----END PUBLIC KEY-----
	`
	zswPublicKeyString, err := SM2PemToZSWPublicKeyString([]byte(encodedSm2PemPublicKey))
	assert.NoError(t, err, "unexpected error decoding SM2 public key pem")
	assert.Equal(t, "PUB_GM_5XYqnUzbW8MXx5gJbY7vcs6tZixXFp9HV3LgQNgwqx5bGduFHc", zswPublicKeyString)
	realPublicKey, err := NewPublicKey(zswPublicKeyString)
	assert.NoError(t, err, "unexpected error decoding SM2 public key string")

	digest, err := base64.StdEncoding.DecodeString("Av0fkx5xAjZE2X2iVfhGmdB0BmcptRrT72QOGurJzx4=")
	assert.NoError(t, err, "error decoding base64 digest")
	signatureSimple, err := base64.StdEncoding.DecodeString("MEQCIFwpeQpe1H4jfKwJoqE3SmfBzlPRx+dsKzHY85BUYjEZAiBpxpMyYpztygFDcVe8H1SDpVUkMWZHsrK2I3hO1rsmNQ==")
	assert.NoError(t, err, "error decoding simple pm2 base64 signature")
	sigData := []byte{byte(CurveGM)}
	sigData = append(sigData, realPublicKey.Content[0:33]...)
	sigData = append(sigData, signatureSimple...)
	assert.LessOrEqual(t, len(sigData), 106, "signature too large!")
	if len(sigData) < 106 {
		sigData = append(sigData, bytes.Repeat([]byte{0}, 106-len(sigData))...)
	}
	assert.Equal(t, len(sigData), 106, "invalid size for sig data must be 106 bytes include the front type byte")
	signatureInstance, err := NewSignatureFromData(sigData)
	assert.NoError(t, err, "unable to create signature from data %w", err)
	assert.Equal(t, signatureInstance.String(), "SIG_GM_J75M5JQwBaAdfUjzkoTnSz3eL7xrw4wzcAy4pC93okraWfZuZLSzsFBYN9Jz5vhxEPjRcJfdDhYyDLT6gUDjyuSRkkiEB9ucS2rtJWmjpKrtaJNnKMRzCo8ZUUD9HMAaVmxcbdcgoLyhvfi5UXCb")
	assert.Equal(t, true, signatureInstance.Verify(digest, realPublicKey), "verification of the signature failed for this public key and digest")
	recoveredPubKey, err := signatureInstance.PublicKey(digest)
	assert.NoError(t, err, "error recovering public key from signature instance")
	assert.Equal(t, realPublicKey.String(), recoveredPubKey.String(), "public key recovered from the signature does not match our original public key")

}
