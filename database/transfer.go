package database

import (
	"fmt"
	"gomicro/constants"
	"gomicro/ent"
)

// get all account
func GetAllTransfers() ([]*ent.Transfers, constants.DBOperationStatus, error) {
	tx, err := Conn.Tx(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}
	resp, err := tx.Transfers.Query().All(ctx)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return resp, constants.ERROR, err
	}
	tx.Commit()
	return resp, constants.QUERIED, err
}

// get an account by id
func GetTransferById(id int) (*ent.Transfers, constants.DBOperationStatus, error) {
	tx, err := Conn.Tx(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}
	resp, err := tx.Transfers.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return resp, constants.ERROR, err
	}
	tx.Commit()
	return resp, constants.QUERIED, err
}

// delete an account by id
func DeleteTransfer(id int) (constants.DBOperationStatus, error) {
	tx, err := Conn.Tx(ctx)
	if err != nil {
		return constants.ERROR, err
	}
	err = tx.Transfers.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return constants.ERROR, err
	}
	tx.Commit()
	return constants.DELETED, err
}


func TransferMoney(to_id int, from_id int, amount int64) (*ent.Transfers, constants.DBOperationStatus, error) {
	tx, err := Conn.Tx(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}
	transaction, err := tx.Transfers.Create().
	SetFromAccountID(from_id).
	SetToAccountID(to_id).
	SetAmount(amount).Save(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}

	_, err = tx.Accounts.UpdateOneID(from_id).SetBalance(-amount).Save(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}

	_, err = tx.Accounts.UpdateOneID(to_id).SetBalance(amount).Save(ctx)
	if err != nil {
		return nil, constants.ERROR, err
	}

	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return nil, constants.ERROR, err
	}
	tx.Commit()
	return transaction, constants.UPDATED, err
}