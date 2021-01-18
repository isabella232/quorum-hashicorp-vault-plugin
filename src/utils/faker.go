package utils

import (
	"math/rand"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/entities"
	"github.com/ethereum/go-ethereum/common"
)

func FakeETHAccount() *entities.ETHAccount {
	return &entities.ETHAccount{
		Address:             common.HexToAddress(randHexString(12)).String(),
		PublicKey:           common.HexToHash(randHexString(12)).String(),
		CompressedPublicKey: common.HexToHash(randHexString(12)).String(),
		Namespace:           "_",
	}
}

func FakeZksAccount() *entities.ZksAccount {
	return &entities.ZksAccount{
		Algorithm: entities.ZksAlgorithmEDDSA,
		Curve: entities.ZksCurveBN256,
		PublicKey: "0x022d15d1be3a2459f4dfdca5b6ec3d255107592c4f231952fcf2b0cb7d77ec05",
		PrivateKey: "0x022d15d1be3a2459f4dfdca5b6ec3d255107592c4f231952fcf2b0cb7d77ec05d2751c7f7db7277404b1cd0d83ed1480bef16ac8f502d90283048aa64bb872d6d2795be44a31796693f1084a1f07f9b94c3dcbde35780291fcb0e2e3eeed5c65",
		Namespace: "_",
	}
}

func randHexString(n int) string {
	var letterRunes = []rune("abcdef0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}