package main

import (
	"fmt"
	"github.com/dgraph-io/badger"
	abciserver "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"os/signal"
	"syscall"
	"tendermint-practice/kvstore/app"
)

func main() {
	// open db
	db, err := badger.Open(badger.DefaultOptions("store").WithTruncate(true))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open badger db: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// set options
	logger := log.MustNewDefaultLogger(log.LogFormatJSON, log.LogLevelInfo, false)
	kvStoreApp := app.NewKVStoreApp(db)

	server := abciserver.NewSocketServer("tcp://127.0.0.1:26658", kvStoreApp)
	server.SetLogger(logger)

	// start server
	err = server.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error starting socket server: %v", err)
		os.Exit(1)
	}
	defer server.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
