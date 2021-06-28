package main

import (
	"flag"
	"log"
	"strconv"

	"gopkg.in/zeromq/goczmq.v4"
)

func main() {
	TX_RES_PORT := new(int)
	TX_RES_PORT = flag.Int("winter-tx-response-port", 33334, "Wallet interface tx response port")
	TX_HASH := new(string)
	TX_HASH = flag.String("tx-hash", "", "TX hash")
	flag.Parse()

	// New socket for request
	sock, err := goczmq.NewReq("tcp://127.0.0.1:" + strconv.Itoa(*TX_RES_PORT))
	if err != nil {
		panic(err)
	}

	// send message to Response server
	err = sock.SendFrame([]byte(*TX_HASH), 0)
	if err != nil {
		panic(err)
	}
	log.Print("sent")

	// make slice for response and wait for it from Response server
	res := make([]byte, 32)
	i, err := sock.Read(res)
	if err != nil {
		panic(err)
	}
	log.Print("read", i, string(res))
}
