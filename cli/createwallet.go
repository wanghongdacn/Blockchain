package cli

import (
	"fmt"

	"github.com/NlaakStudios/Blockchain/core"
)

//CreateWallet creates a new wallet address
func (cli *CLI) CreateWallet() {
	wallets, _ := core.NewWallets(cli.NodePort)
	address := wallets.CreateWallet()
	wallets.SaveToFile(cli.NodePort)

	fmt.Printf("Your new address: %s\n", address)
}
