package cli

import (
	"fmt"
	"log"

	"github.com/NlaakStudios/Blockchain/api"
	"github.com/NlaakStudios/Blockchain/core"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {

	cli.NodeID = nodeID
	fmt.Printf("Starting node %s...", nodeID)
	fmt.Printf("Miner address detected. ")
	if len(minerAddress) > 0 {
		if core.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Failed: Wrong miner address!")
		}
	}
	core.StartServer(cli.NodeID, minerAddress)
	fmt.Printf("Success.\n")
	
	//Start GoRoutine for RestAPI
	api.InitBlockchainAPI()
}
