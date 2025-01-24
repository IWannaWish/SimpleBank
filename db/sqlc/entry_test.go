package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: sql.NullInt64{
			Int64: account.ID,
			Valid: true,
		},
		Amount: account.Balance,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID.Int64, entry.AccountID.Int64)
	require.Equal(t, arg.AccountID.Valid, entry.AccountID.Valid)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	// Создаем 10 записей для указанного аккаунта
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: sql.NullInt64{
				Int64: account.ID,
				Valid: true,
			},
			Amount: int64(i + 1), // Разные суммы для уникальности записей
		}
		_, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
	}

	// Устанавливаем параметры для выборки
	arg := ListEntriesParams{
		AccountID: sql.NullInt64{
			Int64: account.ID,
			Valid: true,
		},
		Limit:  5, // Берем первые 5 записей
		Offset: 0, // Начинаем с первой записи
	}

	// Вызов ListEntries
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, int(arg.Limit)) // Должно быть ровно 5 записей

	// Проверяем каждую запись
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID.Int64) // ID аккаунта совпадает
		require.True(t, entry.AccountID.Valid)              // Поле AccountID валидно
		require.NotZero(t, entry.ID)                        // Проверяем ID записи
		require.NotZero(t, entry.CreatedAt)                 // Должна быть заполнена дата создания
	}

	// Проверяем выборку с Offset
	arg.Offset = 5 // Берем оставшиеся 5 записей
	entries2, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries2, 5) // Должно быть ровно 5 записей

	// Проверяем, что записи разные
	require.NotEqual(t, entries, entries2) // Первая и вторая выборки должны отличаться
}
