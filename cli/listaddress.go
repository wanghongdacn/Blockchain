package cli

import (
	"fmt"
	"log"

	"github.com/NlaakStudios/Blockchain/core"
)

//ListAddresses show all addresses to stdout
func (cli *CLI) ListAddresses() {
	wallets, err := core.NewWallets(cli.NodePort)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
