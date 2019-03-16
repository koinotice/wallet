package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
 	"github.com/koinotice/bee/contracts"
	"github.com/nats-io/go-nats"
	"koinotice.eth/contracts"
	"math/big"
	"runtime"

	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type BetOrder struct {
	RoomId  int
	OrderId int
}

func main() {


	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	// with no 0x
	greeterAddress := "0Ab0487f30bf0C4E5B479258a31f926e7B548dBe"

	// with no 0x
	priv := "117bbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872c2af"

	key, err := crypto.HexToECDSA(priv)

	/**
	 * Connecting to contract at an address
	 */

	contractAddress := common.HexToAddress(greeterAddress)
	instance, err := hashdice.NewHashDice(contractAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)

	// not sure why I have to set this when using testrpc
	// var nonce int64 = 0
	// auth.Nonce = big.NewInt(nonce)

	/**
	 * Calling contract method
	 */


	/**
	 * Events
	 */

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	abiPath, _ := filepath.Abs("./contracts/HashDice.abi")
	file, err := ioutil.ReadFile(abiPath)

	if err != nil {
		fmt.Println("Failed to read file:", err)
	}

	_, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			//var greetEvent struct {
			//	Name  string
			//	Count *big.Int
			//}
			//
			//err = hashdiceAbi.Unpack(&greetEvent, "_Greet", log.Data)
			//
			//if err != nil {
			//	fmt.Println("Failed to unpack:", err)
			//}

			fmt.Println("Contract:", log.Address.Hex())
			fmt.Println("Name:", log.Data)
			//fmt.Println("Count:", greetEvent.Count)
		}
	}
	nc, err := nats.Connect("127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	if _, err := ec.Subscribe("CloseBetOrderByRoomIdOrderId", func(order *BetOrder) {
		log.Printf("RoomId: %v - OrderId: %v", order.RoomId, order.OrderId)

		RoomId := new(big.Int).SetUint64(uint64(order.RoomId))
		OrderId := new(big.Int).SetUint64(uint64(order.OrderId))

		tx, err := instance.CloseBetOrder(auth, RoomId, OrderId)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	}); err != nil {
		log.Fatal(err)
	}

	select {}
	runtime.Goexit()
}

func getTxSign(order *BetOrder) {
	/**
	 * Connecting to provider
	 */
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	// with no 0x
	greeterAddress := "0Ab0487f30bf0C4E5B479258a31f926e7B548dBe"

	// with no 0x
	priv := "117bbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872c2af"

	key, err := crypto.HexToECDSA(priv)

	/**
	 * Connecting to contract at an address
	 */

	contractAddress := common.HexToAddress(greeterAddress)
	instance, err := hashdice.NewHashDice(contractAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)

	// not sure why I have to set this when using testrpc
	// var nonce int64 = 0
	// auth.Nonce = big.NewInt(nonce)

	/**
	 * Calling contract method
	 */
	RoomId := new(big.Int).SetUint64(uint64(order.RoomId))
	OrderId := new(big.Int).SetUint64(uint64(order.OrderId))

	tx, err := instance.CloseBetOrder(auth, RoomId, OrderId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	/**
	 * Events
	 */

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	abiPath, _ := filepath.Abs("./contracts/HashDice.abi")
	file, err := ioutil.ReadFile(abiPath)

	if err != nil {
		fmt.Println("Failed to read file:", err)
	}

	_, err = abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			//var greetEvent struct {
			//	Name  string
			//	Count *big.Int
			//}
			//
			//err = hashdiceAbi.Unpack(&greetEvent, "_Greet", log.Data)
			//
			//if err != nil {
			//	fmt.Println("Failed to unpack:", err)
			//}

			fmt.Println("Contract:", log.Address.Hex())
			fmt.Println("Name:", log.Data)
			//fmt.Println("Count:", greetEvent.Count)
		}
	}

}
