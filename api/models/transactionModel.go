package models

import (
	"context"
	"log"

	"github.com/golang-demos/ecommerce-basic/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	WalletID  primitive.ObjectID `json:"wallet_id" bson:"wallet_id"`
	Amount    float32            `json:"amount" bson:"amount"`
	TransType string             `json:"trans_type" bson:"trans_type"`
}

func MakeTransaction(wallet Wallet, TransType string, Amount float32) (*Transaction, bool) {
	var transaction Transaction

	transaction.UserID = wallet.UserID
	transaction.WalletID = wallet.ID
	transaction.Amount = Amount
	transaction.TransType = TransType

	result, err := database.TransactionColllection.InsertOne(context.Background(), transaction)
	log.Print(err)
	if err != nil {
		return new(Transaction), false
	}
	if ObjId, ok := result.InsertedID.(primitive.ObjectID); ok {
		transaction.ID = ObjId
	}
	return &transaction, true
}
