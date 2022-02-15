package database

import (
	"fmt"
	"gomicro/ent"
)

// get all account
func GetAllTransfers() ([]*ent.Transfers, error) {
	tx, err := Conn.Tx(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return nil, err
	}
	resp, err := tx.Transfers.Query().All(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return resp, err
	}
	tx.Commit()
	return resp, err
}

// get an account by id
func GetTransferById(id int) (*ent.Transfers, error) {
	tx, err := Conn.Tx(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return nil, err
	}
	resp, err := tx.Transfers.Get(ctx, id)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return resp, err
	}
	tx.Commit()
	return resp, err
}

// delete an account by id
func DeleteTransfer(id int) error {
	tx, err := Conn.Tx(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return err
	}
	//pgErr, ok :=  err.(*pq.Error)
	err = tx.Transfers.DeleteOneID(id).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return err
	}
	tx.Commit()
	return err
}

func TransferMoney(to_id int, from_id int, amount int64) (*ent.Transfers, error) {
	tx, err := Conn.Tx(ctx)
	if err != nil {
		return nil, err
	}
	transaction, err := tx.Transfers.Create().
		SetFromAccountID(from_id).
		SetToAccountID(to_id).
		SetAmount(amount).Save(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return nil, err
	}

	_, err = tx.Accounts.UpdateOneID(from_id).SetBalance(-amount).Save(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return nil, err
	}

	_, err = tx.Accounts.UpdateOneID(to_id).SetBalance(amount).Save(ctx)
	//pgErr, ok :=  err.(*pq.Error)
	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println(err)
		err = rollback(tx)
		return nil, err
	}
	tx.Commit()
	return transaction, err
}
