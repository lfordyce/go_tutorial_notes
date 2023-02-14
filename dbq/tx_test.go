package dbq

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

func _TestTransaction(t *testing.T) {
	conn, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	provider := NewTxProvider(conn)
	if err := provider.Tx(ctx, func(txContext TxContext) error {

		tx := FromCtxOr(ctx, conn)
		if _, err := tx.ExecContext(ctx, "", ""); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func CreateProfile(ctx context.Context) {
	//tx := FromCtxOr(ctx, )
}
