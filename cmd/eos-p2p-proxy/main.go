package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/zhongshuwen/zswchain-go/p2p"
	"github.com/streamingfast/logging"
)

var peer1 = flag.String("peer1", "localhost:9876", "peer 1")
var peer2 = flag.String("peer2", "localhost:2222", "peer 2")
var chainID = flag.String("chain-id", "308cae83a690640be3726a725dde1fa72a845e28cfc63f28c3fa0a6ccdb6faf0", "peer 1")
var showLog = flag.Bool("v", false, "show detail log")

func main() {
	flag.Parse()

	fmt.Println("P2P Proxy")

	if *showLog {
		logging.Set(logging.MustCreateLogger(), "github.com/zhongshuwen/zswchain-go/p2p")
	}
	defer p2p.SyncLogger()

	cID, err := hex.DecodeString(*chainID)
	if err != nil {
		log.Fatal(err)
	}

	proxy := p2p.NewProxy(
		p2p.NewOutgoingPeer(*peer1, "zsw-proxy", nil),
		p2p.NewOutgoingPeer(*peer2, "zsw-proxy", &p2p.HandshakeInfo{
			ChainID: cID,
		}),
	)

	//proxy := p2p.NewProxy(
	//	p2p.NewOutgoingPeer("localhost:9876", chainID, "zsw-proxy", false),
	//	p2p.NewIncommingPeer("localhost:1111", chainID, "zsw-proxy"),
	//)

	//proxy := p2p.NewProxy(
	//	p2p.NewIncommingPeer("localhost:2222", "zsw-proxy"),
	//	p2p.NewIncommingPeer("localhost:1111", "zsw-proxy"),
	//)

	proxy.RegisterHandler(p2p.StringLoggerHandler)
	fmt.Println(proxy.ConnectAndStart())

}
