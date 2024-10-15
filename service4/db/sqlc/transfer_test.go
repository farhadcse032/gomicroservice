package db

import (
	"context"
	"microServices/service4/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        int64(util.RandomInt(100, 1000)),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
func TestCreateTransfer(t *testing.T) {
	payer := CreateAccount(t)
	receiver := CreateAccount(t)
	createRandomTransfer(t, payer, receiver)
}

func TestGetTransfer(t *testing.T) {
	payer := CreateAccount(t)
	receiver := CreateAccount(t)
	transfer1 := createRandomTransfer(t, payer, receiver)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

func TestListTransfers(t *testing.T) {
	payer := CreateAccount(t)
	receiver := CreateAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, payer, receiver)
	}
	arg := ListTransfersParams{
		FromAccountID: payer.ID,
		ToAccountID:   receiver.ID,
		Limit:         5,
		Offset:        0,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, len(transfers), 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
