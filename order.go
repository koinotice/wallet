package main

import (
	"context"
	"fmt"
	"koinotice.eth/contracts"

	"github.com/koinotice/bee/contracts"
	"github.com/nats-io/go-nats"
	"log"
	"math/big"
	"runtime"

	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Order struct {
	RoomId  int
	OrderId int
}

//var  nonce   =0;
func main() {
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


	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("9a82accbc285082eebde4a69249738b5aedccb2008e14d2591bd0801de695515")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)


	address := common.HexToAddress("0xa7DBa6053A0D631177340E8061bc12F5009bA453")
	instance, err := hashdice.NewHashDice(address, client)

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	if _, err := ec.Subscribe("CloseBetOrderByRoomIdOrderId", func(order *Order) {
		log.Printf("RoomId: %v - OrderId: %v", order.RoomId, order.OrderId)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		//nonce++
		auth.Nonce = big.NewInt(int64(nonce))
		auth.GasPrice = gasPrice


		RoomId := new(big.Int).SetUint64(uint64(order.RoomId))
		OrderId := new(big.Int).SetUint64(uint64(order.OrderId))

		tx, err := instance.CloseBetOrder(auth, RoomId, OrderId)
		if err != nil {
			log.Println(err)
		}else{
		//	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
			fmt.Printf("tx sent: %s", tx ) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

		}

	}); err != nil {
		log.Fatal(err)
	}

	select {}
	 runtime.Goexit()
}

