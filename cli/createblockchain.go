package cli

import (
	"fmt"
	"log"

	"github.com/NlaakStudios/Blockchain/core"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	fmt.Printf("Creating new blockchain with primary wallet address of %s.\n", address)
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
		// TODO: Create Primary Wallet (address)
		//cli.createWallet(nodeID)
	}

	//utils.CreateDirIfNotExist("data")

	bc := core.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
