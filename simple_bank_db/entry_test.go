package simple_bank_db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wanderlust-create/adventure-life-simple-bank/util"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)


	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	createRandomEntry(t, account1)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntrys(t *testing.T) {
	account := createRandomAccount(t)
	for i := 1; i <= 5; i++ {
		createRandomEntry(t, account)
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID:   account.ID,
		Limit:       5,
		Offset:      5,
	}
	entrys, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, entry := range entrys {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
