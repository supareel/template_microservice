package database

import (
	"context"
	"gomicro/ent"
	"gomicro/util"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func connectTestDB() {
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASS", "secret")
	os.Setenv("DB_SERVER", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "simple_bank")
	os.Setenv("DB_SSL", "disable")

	ConnectToDB()
}

func createRandomAccount(t *testing.T) (*ent.Accounts, *ent.Accounts, error) {
	connectTestDB()

	testData := &ent.Accounts{
		Owner:    util.RandomString(10),
		Balance:  util.RandomInt(100, 10000),
		Currency: util.RandomCurrency(),
	}

	resp, err := CreateAccount(testData)
	return testData, resp, err
}

func clearAccountsTable() {
	connectTestDB()
	_, err := Conn.Accounts.Delete().Exec(context.Background())

	if err != nil {
		os.Exit(1)
	}

}

func TestCreateAccount(t *testing.T) {
	clearAccountsTable()
	testData, resp, _ := createRandomAccount(t)
	require.Equal(t, testData.Owner, resp.Owner)
	require.Equal(t, testData.Balance, resp.Balance)
	require.Equal(t, testData.Currency, resp.Currency)
	require.NotZero(t, resp.ID)
	require.NotZero(t, resp.Balance)

}

func TestGetAccount(t *testing.T) {
	clearAccountsTable()
	_, account1, _ := createRandomAccount(t)
	account2, _ := GetAccountById(account1.ID)

	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccountBalance(t *testing.T) {
	clearAccountsTable()
	_, account1, _ := createRandomAccount(t)
	new_balance := util.RandomInt(300, 10000)
	account2, _ := UpdateAccountBalance(account1.ID, new_balance)

	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, new_balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	clearAccountsTable()
	_, account1, _ := createRandomAccount(t)
	err := DeleteAccount(account1.ID)
	require.NoError(t, err)

	account2, err := GetAccountById(account1.ID)
	require.EqualError(t, err, "ent: accounts not found")
	require.Empty(t, account2)
}

func TestGetAllAccounts(t *testing.T) {
	clearAccountsTable()
	_, account1, _ := createRandomAccount(t)
	_, account2, _ := createRandomAccount(t)

	accountList, err := GetAllAccounts()
	require.NoError(t, err)

	require.Equal(t, 2, len(accountList))
	require.NotEmpty(t, account1, accountList[0])
	require.NotEmpty(t, account2, accountList[1])
}
