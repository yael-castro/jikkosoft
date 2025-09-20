package main

import (
	"container/list"
	"fmt"
	"github.com/yael-castro/jikkosoft/q11/internal/transaction"
)

func main() {
	const bufferSize = 100
	frequencies := make(map[string]int)

	// READING TRANSACTION BUFFER
	for t := range transaction.Buffer(bufferSize) {
		frequencies[t.ClientID]++
	}

	// FILTERING CLIENT FREQUENCIES TO GET TOP 10 CLIENTS
	const maxCommerces = 10
	clients := list.New()

	for clientID, frequency := range frequencies {
		client := transaction.Client{
			ID:        clientID,
			Frequency: frequency,
		}

		if clients.Len() <= maxCommerces {
			clients.PushBack(client)
			continue
		}

		if frequency > clients.Back().Value.(transaction.Client).Frequency {
			clients.Remove(clients.Back())
			clients.PushBack(client)
		}
	}

	// PRINTING RESULTS
	fmt.Printf("FREQUENCIES\n")
	for clientID, frequency := range frequencies {
		fmt.Printf("CLIENT %v FREQ %v\n", clientID, frequency)
	}

	fmt.Printf("TOP 10 CLIENTS\n")
	for clients.Len() > 0 {
		front := clients.Front()

		client := front.Value.(transaction.Client)
		fmt.Printf("CLIENT %v FREQ %v\n", client.ID, client.Frequency)

		clients.Remove(front)
	}
}
