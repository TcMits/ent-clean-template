package testutils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	_ "go.beyondstorage.io/services/memory"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
)

func GetMemmoryStorager(t *testing.T) types.Storager {
	t.Helper()
	memoryName := uuid.NewString()
	storager, err := services.NewStoragerFromString("memory:///" + memoryName)
	require.NoError(t, err)
	return storager
}
