package cli

import (
	"fmt"

	"github.com/NlaakStudios/Blockchain/core"
)

func (cli *CLI) createWallet() {
	wallets, _ := core.NewWallets(cli.NodePort)
	address := wallets.CreateWallet()
	wallets.SaveToFile(cli.NodePort)

	fmt.Printf("Your new address: %s\n", address)
}
