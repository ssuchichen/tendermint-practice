package kvstore

import (
	"fmt"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type KVStoreApp struct{}

var _ abcitypes.Application = (*KVStoreApp)(nil)

func NewKVStoreApp() *KVStoreApp {
	return &KVStoreApp{}
}

func (app KVStoreApp) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	fmt.Println(">> Info")
	return abcitypes.ResponseInfo{}
}

func (app KVStoreApp) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	fmt.Println(">> Query")
	return abcitypes.ResponseQuery{}
}

func (app KVStoreApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	fmt.Println(">> CheckTx")
	return abcitypes.ResponseCheckTx{}
}

func (app KVStoreApp) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	fmt.Println(">> InitChain")
	return abcitypes.ResponseInitChain{}
}

func (app KVStoreApp) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	fmt.Println(">> BeginBlock")
	return abcitypes.ResponseBeginBlock{}
}

func (app KVStoreApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	fmt.Println(">> DeliverTx")
	return abcitypes.ResponseDeliverTx{}
}

func (app KVStoreApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	fmt.Println(">> EndBlock")
	return abcitypes.ResponseEndBlock{}
}

func (app KVStoreApp) Commit() abcitypes.ResponseCommit {
	fmt.Println(">> Commit")
	return abcitypes.ResponseCommit{}
}

func (app KVStoreApp) ListSnapshots(req abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (app KVStoreApp) OfferSnapshot(req abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (app KVStoreApp) LoadSnapshotChunk(req abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (app KVStoreApp) ApplySnapshotChunk(req abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}
