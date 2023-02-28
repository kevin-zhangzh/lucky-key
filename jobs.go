package lucky

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kevin-zhangzh/lucky-key/schema"
	"math/big"
)

func (l *Lucky) runJobs() {
	l.scheduler.Every(5).Second().SingletonMode().Do(l.genLuckKey)

	l.scheduler.StartAsync()
}

func (l *Lucky) genLuckKey() {
	log.Info("start to find")
	priKey, err := crypto.GenerateKey()
	priKeyByte := crypto.FromECDSA(priKey)
	publicKey := priKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	bal, err := l.cli.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Error("get balance failed", "err", err)
	}
	if bal.Cmp(big.NewInt(0)) > 0 {
		err = l.wdb.InsertAsset(schema.Asset{
			EccKey:  hexutil.Encode(priKeyByte)[2:],
			Address: address.Hex(),
			Balance: bal.String(),
		})
		if err != nil {
			log.Error("insert to db failed", "err", err)
			return
		}
		log.Info("find an asset", "address", address.Hex())
	}
}
