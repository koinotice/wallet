package main

import (
	"github.com/nats-io/go-nats"
	"log"
	"runtime"
)

func main() {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()
	type Order struct {
		RoomId  int
		OrderId int
	}

	if _, err := ec.Subscribe("CloseBetOrderByRoomIdOrderId", func(s *Order) {
		log.Printf("Stock: %v - Price: %v", s.RoomId, s.OrderId)

	}); err != nil {
		log.Fatal(err)
	}



	select {}
	runtime.Goexit()
}
