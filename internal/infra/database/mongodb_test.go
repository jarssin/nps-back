package database_test

import (
	"testing"

	"github.com/jarssin/nps-back/internal/infra/database"
)

func TestConnectAndClose(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Skip("MongoDB not available for test")
	}
	if db != nil {
		err = db.Close()
		if err != nil {
			t.Errorf("expected nil error on close, got %v", err)
		}
	}
}
