package database

import (
	"fmt"
	"gomicro/ent"
)

// create an account
func CreateAccount(data *ent.Accounts) (*ent.Accounts, error) {
	resp, err := Conn.Accounts.Create().SetOwner(data.Owner).
		SetBalance(data.Balance).
		SetCurrency(data.Currency).Save(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, err
}

// get all account
func GetAllAccounts() ([]*ent.Accounts, error) {
	resp, err := Conn.Accounts.Query().All(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, err
}

// get an account by id
func GetAccountById(id int) (*ent.Accounts, error) {
	resp, err := Conn.Accounts.Get(ctx, id)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, err
}

// update an account by id
func UpdateAccountBalance(id int, balance int64) (*ent.Accounts, error) {
	resp, err := Conn.Accounts.UpdateOneID(id).SetBalance(balance).Save(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	return resp, err
}

// delete an account by id
func DeleteAccount(id int) error {
	err := Conn.Accounts.DeleteOneID(id).Exec(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
