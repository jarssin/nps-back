package database_test

import (
	"context"
	"testing"

	"github.com/jarssin/nps-back/internal/infra/database"
)

func TestConnectAndClose(t *testing.T) {
	ctx := context.Background()
	db, err := database.Connect(ctx)
	if err != nil {
		t.Skip("MongoDB not available for test")
	}
	if db != nil {
		err = db.Close(ctx)
		if err != nil {
			t.Errorf("expected nil error on close, got %v", err)
		}
	}
}
