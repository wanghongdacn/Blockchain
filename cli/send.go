package cli

import (
	"fmt"
	"log"

	"../core"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !core.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !core.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	//TODO: See if wallet has enough to send amount

	tx := core.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := core.NewCoinbaseTX(from, "")
		txs := []*core.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		core.SendTx(core.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}

func (cli *CLI) PopulateWallets(from string, nodeID string) {

	fmt.Printf("Creating core wallets and funding them from wallet %s.\n", from)
	if !core.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}

	//Open Wallets file
	fmt.Println("Opening wallets.")
	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}

	//Get primary blockchain wallet (with all coins)
	fmt.Printf("Accessing wallet %s\n", from)
	wallet := wallets.GetWallet(from)
	fmt.Printf("Getting balance from wallet %s\n", from)
	balance := cli.getBalance(from, nodeID)
	fmt.Printf("Your Primary Blockchain wallet balance is %d\n", balance)

	//Open the blockchain
	fmt.Println("Opening blockchain.")
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	//Create our core wallets (In-House)
	fmt.Println("Creating ICO Wallet.")
	addressICO := wallets.CreateWallet() //Create ICO Wallet
	fmt.Printf("Your new ICO Wallet Address: %s\n", addressICO)

	fmt.Println("Creating DEV Wallet.")
	addressDEV := wallets.CreateWallet() //Create DEV Wallet
	fmt.Printf("Your new DEV Wallet Address: %s\n", addressDEV)

	fmt.Println("Creating OAM Wallet.")
	addressOAM := wallets.CreateWallet() //Create Operations & Marketing Wallet
	fmt.Printf("Your new OAM Wallet Address: %s\n", addressOAM)

	fmt.Println("Creating PLT Wallet.")
	addressPLT := wallets.CreateWallet() //Create Platform Wallet
	fmt.Printf("Your new PLT Wallet Address: %s\n", addressPLT)

	//Save Wallet file
	fmt.Println("Saving Wallets.")
	wallets.SaveToFile(nodeID)
	fmt.Printf("Wallets DB Updated\n")

	//TODO: See if wallet has enough to send amount
	fmt.Println("Creating Wallet Transactions.")
	tx := core.NewUTXOTransaction(&wallet, addressICO, 10000000, &UTXOSet)
	//mine Now
	cbTx := core.NewCoinbaseTX(from, "")
	txs := []*core.Transaction{cbTx, tx}
	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	tx = core.NewUTXOTransaction(&wallet, addressDEV, 500000, &UTXOSet)
	//mine Now
	cbTx = core.NewCoinbaseTX(from, "")
	txs = []*core.Transaction{cbTx, tx}
	newBlock = bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	tx = core.NewUTXOTransaction(&wallet, addressOAM, 1500000, &UTXOSet)
	//mine Now
	cbTx = core.NewCoinbaseTX(from, "")
	txs = []*core.Transaction{cbTx, tx}
	newBlock = bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	tx = core.NewUTXOTransaction(&wallet, addressPLT, 3000000, &UTXOSet)
	//mine Now
	cbTx = core.NewCoinbaseTX(from, "")
	txs = []*core.Transaction{cbTx, tx}
	newBlock = bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	fmt.Println("Success, Blockchain and Core Wallets initialized.")
}
