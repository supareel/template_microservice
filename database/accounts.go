package database

import (
	"context"
	"fmt"
	"gomicro/constants"
	"gomicro/ent"
)

var ctx = context.Background()

// create an account
func CreateAccount(data ent.Accounts) (*ent.Accounts, constants.DBOperationStatus) {
	resp, err := Conn.Accounts.Create().SetOwner(data.Owner).
		SetBalance(data.Balance).
		SetCurrency(data.Currency).Save(ctx)

	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR
	}
	return resp, constants.CREATED
}


// get all account
func GetAllAccounts() ([]*ent.Accounts, constants.DBOperationStatus) {
	resp, err := Conn.Accounts.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR
	}
	return resp, constants.QUERIED
}

// get an account by id
func GetAccountById(id int) (*ent.Accounts, constants.DBOperationStatus) {
	resp, err := Conn.Accounts.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR
	}
	return resp, constants.QUERIED
}


// update an account by id
func UpdateAccountBalance(id int, balance int64) (*ent.Accounts, constants.DBOperationStatus) {
	resp, err := Conn.Accounts.UpdateOneID(id).SetBalance(balance).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR
	}
	return resp, constants.UPDATED
}


// delete an account by id
func DeleteAccountBalance(id int) (constants.DBOperationStatus) {
	err := Conn.Accounts.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return constants.ERROR
	}
	return constants.DELETED
}