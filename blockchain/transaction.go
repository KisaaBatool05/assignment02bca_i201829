package blockchain

import "time"

type Transaction struct {
	Sender   string
	Receiver string
	Amount   string
	DateTime time.Time
}

func NewTransaction(sender, receiver, amount string) *Transaction {
	transaction := Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		DateTime: time.Now(),
	}

	return &transaction
}
