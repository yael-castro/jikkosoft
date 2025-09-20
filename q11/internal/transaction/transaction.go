package transaction

import (
	"math/rand/v2"
	"strconv"
	"time"
)

const bufferLimit = 50

type Client struct {
	ID        string
	Frequency int
}
type Transaction struct {
	CreatedAt time.Time
	ClientID  string
	Amount    float64
}

// Buffer returns a buffer that generates n fake transactions
func Buffer(n int) <-chan Transaction {
	transactionCh := make(chan Transaction, bufferLimit)

	go func() {
		defer close(transactionCh)

		for range n {
			random := rand.N(int64(30))
			clientID := strconv.FormatInt(random, 10)

			t := Transaction{
				Amount:    1_500,
				ClientID:  clientID,
				CreatedAt: time.Now(),
			}

			transactionCh <- t
		}
	}()

	return transactionCh
}
