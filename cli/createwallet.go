package cli

import (
	"fmt"

	"github.com/NlaakStudios/Blockchain/core"
)

func (cli *CLI) createWallet() {
	wallets, _ := core.NewWallets(cli.NodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(cli.NodeID)

	fmt.Printf("Your new address: %s\n", address)
}
