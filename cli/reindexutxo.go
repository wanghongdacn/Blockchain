package cli

import (
	"fmt"

	"github.com/NlaakStudios/Blockchain/core"
)

//ReIndexUTXO redoes the chain index
func (cli *CLI) ReIndexUTXO() {
	bc := core.NewBlockchain(cli.NodePort)
	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
