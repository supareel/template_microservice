package database

import (
	"fmt"
	"gomicro/constants"
	"gomicro/ent"
)

// create an account
func CreateTransfer(data *ent.Transfers) (*ent.Transfers, constants.DBOperationStatus, error) {
	resp, err := Conn.Transfers.Create().
  SetFromAccoutID(data.FromAccoutID).
  SetToAccountID(data.ToAccountID).
  SetAmount(data.Amount).
  Save(ctx)
	
  if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.CREATED, err
}

// get all account
func GetAllTransfers() ([]*ent.Transfers, constants.DBOperationStatus, error) {
	resp, err := Conn.Transfers.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.QUERIED, err
}

// get an account by id
func GetTransferById(id int) (*ent.Transfers, constants.DBOperationStatus, error) {
	resp, err := Conn.Transfers.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		return resp, constants.ERROR, err
	}
	return resp, constants.QUERIED, err
}

// delete an account by id
func DeleteTransfer(id int) (constants.DBOperationStatus, error) {
	err := Conn.Transfers.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return constants.ERROR, err
	}
	return constants.DELETED, err
}