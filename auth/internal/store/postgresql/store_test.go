package postgresql

import (
	"fmt"
	"testing"

	"github.com/likoscp/Advanced-Programming-2/auth/internal/configs"
)

func TestConnectDB(t *testing.T) {
	config := &configs.ConfigDB{
		Host:     "localhost",
		Addr:     "5432",
		User:     "admin",
		Password: "admin",
		Name:     "comics",
	}

	_, err := NewStore(config)
	fmt.Println(err)
	if err != nil {
		t.Fatalf("cannot connect to db, err: %v", err)
	}

}
