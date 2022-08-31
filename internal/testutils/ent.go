package testutils

import (
	"context"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func GetSqlite3TestClient(ctx context.Context, t *testing.T) *ent.Client {
	memoryName := uuid.NewString()
	client, err := ent.Open("sqlite3", "file:"+memoryName+"?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	require.NoError(t, client.Schema.Create(ctx))
	return client
}
