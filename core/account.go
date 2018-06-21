package core

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

//UserAccountStruct hold basic user information
type UserAccountStruct struct {
	ID       uint64
	Email    string
	Password string
	Wallets  []Wallet
}

const dbAccountsFile = "Accounts_%s.db"

// AccountsDB implements interactions with a DB
type AccountsDB struct {
	Tip []byte
	DB  *bolt.DB
}

//AccountsDBExists ...
func AccountsDBExists(nodeID string, halt bool, failedMessage string) {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if dbExists(dbFile) == false {
		fmt.Println(failedMessage)
		if halt == true {
			os.Exit(1)
		}
	}
}
