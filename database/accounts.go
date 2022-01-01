package database

import (
	"fmt"
	"gomicro/constants"
	"gomicro/ent"
)

// create an account
func CreateAccount(data *ent.Accounts) (*ent.Accounts, constants.DBOperationStatus, error) {
	resp, err := Conn.Accounts.Create().SetOwner(data.Owner).
		SetBalance(data.Balance).
		SetCurrency(data.Currency).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.CREATED, err
}


// get all account
func GetAllAccounts() ([]*ent.Accounts, constants.DBOperationStatus, error) {
	resp, err := Conn.Accounts.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.QUERIED, err
}

// get an account by id
func GetAccountById(id int) (*ent.Accounts, constants.DBOperationStatus, error) {
	resp, err := Conn.Accounts.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.QUERIED, err
}


// update an account by id
func UpdateAccountBalance(id int, balance int64) (*ent.Accounts, constants.DBOperationStatus, error) {
	resp, err := Conn.Accounts.UpdateOneID(id).SetBalance(balance).Save(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.UPDATED, err
}


// delete an account by id
func DeleteAccount(id int) (constants.DBOperationStatus, error) {
	err := Conn.Accounts.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return constants.ERROR, err
	}
	return constants.DELETED, err
}