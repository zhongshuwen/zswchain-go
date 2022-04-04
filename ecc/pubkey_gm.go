package ecc

import (
	"fmt"

	"github.com/zhongshuwen/zswchain-go/libbsuite/btcd/btcec"
)

type innerGMPublicKey struct {
}

func newInnerGMPublicKey() innerPublicKey {
	return &innerGMPublicKey{}
}

func (p *innerGMPublicKey) key(content []byte) (*btcec.PublicKey, error) {

	return nil, fmt.Errorf("sm2 does not support")
}

func (p *innerGMPublicKey) prefix() string {
	return PublicKeyGMPrefix
}

func (p *innerGMPublicKey) keyMaterialSize() *int {
	return publicKeyDataSize
}
