package app

import (
	"bytes"
	"fmt"
	"github.com/dgraph-io/badger"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type KVStoreApp struct {
	db           *badger.DB
	currentBatch *badger.Txn
}

var _ abcitypes.Application = (*KVStoreApp)(nil)

func NewKVStoreApp(db *badger.DB) *KVStoreApp {
	return &KVStoreApp{db: db}
}

func (app *KVStoreApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	fmt.Println(">> Info")
	return abcitypes.ResponseInfo{}
}

func (app *KVStoreApp) Query(req abcitypes.RequestQuery) (resQuery abcitypes.ResponseQuery) {
	fmt.Println(">> Query")

	resQuery.Key = req.Data
	err := app.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(resQuery.Key)
		if err != nil {
			if err != badger.ErrKeyNotFound {
				return err
			}

			resQuery.Log = "not exists"
			return nil
		}
		return item.Value(func(val []byte) error {
			resQuery.Log = "exists"
			resQuery.Value = val
			return nil
		})
	})
	if err != nil {
		panic(err)
	}

	return abcitypes.ResponseQuery{}
}

func (app *KVStoreApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	fmt.Println(">> CheckTx")

	code := app.isValid(req.Tx)
	return abcitypes.ResponseCheckTx{Code: code, GasWanted: 1}
}

func (app *KVStoreApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	fmt.Println(">> InitChain")
	return abcitypes.ResponseInitChain{}
}

func (app *KVStoreApp) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	fmt.Println(">> BeginBlock")

	app.currentBatch = app.db.NewTransaction(true)
	return abcitypes.ResponseBeginBlock{}
}

func (app *KVStoreApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	fmt.Println(">> DeliverTx")

	code := app.isValid(req.Tx)
	if code != 0 {
		return abcitypes.ResponseDeliverTx{Code: code}
	}

	parts := bytes.Split(req.Tx, []byte("="))
	key, value := parts[0], parts[1]
	err := app.currentBatch.Set(key, value)
	if err != nil {
		panic(err)
	}

	return abcitypes.ResponseDeliverTx{}
}

func (app *KVStoreApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	fmt.Println(">> EndBlock")
	return abcitypes.ResponseEndBlock{}
}

func (app *KVStoreApp) Commit() abcitypes.ResponseCommit {
	fmt.Println(">> Commit")

	app.currentBatch.Commit()
	return abcitypes.ResponseCommit{Data: []byte{}}
}

func (app *KVStoreApp) ListSnapshots(req abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (app *KVStoreApp) OfferSnapshot(req abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (app *KVStoreApp) LoadSnapshotChunk(req abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (app *KVStoreApp) ApplySnapshotChunk(req abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}

// isValid checks whether tx is valid.
// 格式错误,返回1
// 已经存在,返回2
// 正确,返回0
func (app *KVStoreApp) isValid(tx []byte) (code uint32) {
	// check format
	parts := bytes.Split(tx, []byte("="))
	if len(parts) != 2 {
		return 1
	}

	key, value := parts[0], parts[1]

	// check if the same key=value already exists
	err := app.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			if err != badger.ErrKeyNotFound { // 其他错误
				return err
			}
			// key不存在
			return nil
		}

		// key存在
		return item.Value(func(val []byte) error {
			if bytes.Equal(val, value) {
				code = 2
			}
			return nil
		})
	})

	if err != nil {
		panic(err)
	}

	return code
}
