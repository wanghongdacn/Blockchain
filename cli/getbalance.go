package cli

import (
	"fmt"
	"log"

	"github.com/NlaakStudios/Blockchain/core"
	"github.com/NlaakStudios/Blockchain/utils"
)

func (cli *CLI) showBalance(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

func (cli *CLI) getBalance(address, nodeID string) int {

	//fmt.Println("getBalance(%s, %s)", address, nodeID)
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	//fmt.Printf("Balance of '%s': %d\n", address, balance)
	return balance
}
