package cli

import (
	"fmt"
	"log"

	"github.com/NlaakStudios/Blockchain/core"
)

//CreateBlockchain creates an new with an associated master wallet
func (cli *CLI) CreateBlockchain(address string) {
	fmt.Printf("Creating new blockchain with primary wallet address of %s.\n", address)
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
		// TODO: Create Primary Wallet (address)
		//cli.createWallet(nodeID)
	}

	//utils.CreateDirIfNotExist("data")

	bc := core.CreateBlockchain(address, cli.NodePort)
	defer bc.DB.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
