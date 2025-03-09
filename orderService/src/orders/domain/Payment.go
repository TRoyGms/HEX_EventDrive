package domain

import "math/rand"

type Payment struct {
    ID          int
    Name        string
    Amount      int
    Description string
}

func NewPayment(name string, amount int, description string) *Payment {
    return &Payment{
        ID:          generateID(),
        Name:        name,
        Amount:      amount,
        Description: description,
    }
}

func generateID() int {
    return rand.Intn(100000)
}
