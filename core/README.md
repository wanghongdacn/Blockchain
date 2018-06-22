# core
--
    import "\github.com\NlaakStudios\Blockchain\core"


## Usage

```go
var KnownNodes = []string{}
```

#### func  AccountsDBExists

```go
func AccountsDBExists(nodeID string, halt bool, failedMessage string)
```
AccountsDBExists ...

#### func  HashPubKey

```go
func HashPubKey(pubKey []byte) []byte
```
HashPubKey hashes public key

#### func  SendTx

```go
func SendTx(addr string, tnx *Transaction)
```

#### func  StartServer

```go
func StartServer(nodeID, minerAddress string)
```
StartServer starts a node

#### func  ValidateAddress

```go
func ValidateAddress(address string) bool
```
ValidateAddress check if address if valid

#### type AccountsDB

```go
type AccountsDB struct {
	Tip []byte
	DB  *bolt.DB
}
```

AccountsDB implements interactions with a DB

#### type Block

```go
type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}
```

Block represents a block in the blockchain

#### func  DeserializeBlock

```go
func DeserializeBlock(d []byte) *Block
```
DeserializeBlock deserializes a block

#### func  NewBlock

```go
func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block
```
NewBlock creates and returns Block

#### func  NewGenesisBlock

```go
func NewGenesisBlock(coinbase *Transaction) *Block
```
NewGenesisBlock creates and returns genesis Block

#### func (*Block) HashTransactions

```go
func (b *Block) HashTransactions() []byte
```
HashTransactions returns a hash of the transactions in the block

#### func (*Block) Serialize

```go
func (b *Block) Serialize() []byte
```
Serialize serializes the block

#### type Blockchain

```go
type Blockchain struct {
	Tip []byte
	DB  *bolt.DB
}
```

Blockchain implements interactions with a DB

#### func  CreateBlockchain

```go
func CreateBlockchain(address, nodeID string) *Blockchain
```
CreateBlockchain creates a new blockchain DB

#### func  NewBlockchain

```go
func NewBlockchain(nodeID string) *Blockchain
```
NewBlockchain creates a new Blockchain with genesis Block

#### func (*Blockchain) AddBlock

```go
func (bc *Blockchain) AddBlock(block *Block)
```
AddBlock saves the block into the blockchain

#### func (*Blockchain) FindTransaction

```go
func (bc *Blockchain) FindTransaction(ID []byte) (Transaction, error)
```
FindTransaction finds a transaction by its ID

#### func (*Blockchain) FindUTXO

```go
func (bc *Blockchain) FindUTXO() map[string]TXOutputs
```
FindUTXO finds all unspent transaction outputs and returns transactions with
spent outputs removed

#### func (*Blockchain) GetBestHeight

```go
func (bc *Blockchain) GetBestHeight() int
```
GetBestHeight returns the height of the latest block

#### func (*Blockchain) GetBlock

```go
func (bc *Blockchain) GetBlock(blockHash []byte) (Block, error)
```
GetBlock finds a block by its hash and returns it

#### func (*Blockchain) GetBlockHashes

```go
func (bc *Blockchain) GetBlockHashes() [][]byte
```
GetBlockHashes returns a list of hashes of all the blocks in the chain

#### func (*Blockchain) Iterator

```go
func (bc *Blockchain) Iterator() *BlockchainIterator
```
Iterator returns a BlockchainIterat

#### func (*Blockchain) MineBlock

```go
func (bc *Blockchain) MineBlock(transactions []*Transaction) *Block
```
MineBlock mines a new block with the provided transactions

#### func (*Blockchain) SignTransaction

```go
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey)
```
SignTransaction signs inputs of a Transaction

#### func (*Blockchain) VerifyTransaction

```go
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool
```
VerifyTransaction verifies transaction input signatures

#### type BlockchainIterator

```go
type BlockchainIterator struct {
}
```


#### func (*BlockchainIterator) Next

```go
func (i *BlockchainIterator) Next() *Block
```

#### type MerkleNode

```go
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}
```

MerkleNode represent a Merkle tree node

#### func  NewMerkleNode

```go
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode
```
NewMerkleNode creates a new Merkle tree node

#### type MerkleTree

```go
type MerkleTree struct {
	RootNode *MerkleNode
}
```

MerkleTree represent a Merkle tree

#### func  NewMerkleTree

```go
func NewMerkleTree(data [][]byte) *MerkleTree
```
NewMerkleTree creates a new Merkle tree from a sequence of data

#### type ProofOfWork

```go
type ProofOfWork struct {
}
```

ProofOfWork represents a proof-of-work

#### func  NewProofOfWork

```go
func NewProofOfWork(b *Block) *ProofOfWork
```
NewProofOfWork builds and returns a ProofOfWork

#### func (*ProofOfWork) Run

```go
func (pow *ProofOfWork) Run() (int, []byte)
```
Run performs a proof-of-work

#### func (*ProofOfWork) Validate

```go
func (pow *ProofOfWork) Validate() bool
```
Validate validates block's PoW

#### type TXInput

```go
type TXInput struct {
	Txid      []byte // Transaction ID
	Vout      int    // TODO: float32
	Signature []byte
	PubKey    []byte
}
```


#### func (*TXInput) UsesKey

```go
func (in *TXInput) UsesKey(pubKeyHash []byte) bool
```

#### type TXOutput

```go
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}
```

TXOutput represents a transaction output

#### func  NewTXOutput

```go
func NewTXOutput(value int, address string) *TXOutput
```
NewTXOutput create a new TXOutput

#### func (*TXOutput) IsLockedWithKey

```go
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool
```
IsLockedWithKey checks if the output can be used by the owner of the pubkey

#### func (*TXOutput) Lock

```go
func (out *TXOutput) Lock(address []byte)
```
Lock signs the output

#### type TXOutputs

```go
type TXOutputs struct {
	Outputs []TXOutput
}
```

TXOutputs collects TXOutput

#### func  DeserializeOutputs

```go
func DeserializeOutputs(data []byte) TXOutputs
```
DeserializeOutputs deserializes TXOutputs

#### func (TXOutputs) Serialize

```go
func (outs TXOutputs) Serialize() []byte
```
Serialize serializes TXOutputs

#### type Transaction

```go
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
```

Transaction represents a Bitcoin transaction

#### func  DeserializeTransaction

```go
func DeserializeTransaction(data []byte) Transaction
```
DeserializeTransaction deserializes a transaction

#### func  NewBlockchainTX

```go
func NewBlockchainTX(to, data string) *Transaction
```
NewBlockchainTX creates a new transaction to deposit total supply into the
primary blockchain wallet

#### func  NewCoinbaseTX

```go
func NewCoinbaseTX(to, data string) *Transaction
```
NewCoinbaseTX creates a new coinbase transaction

#### func  NewUTXOTransaction

```go
func NewUTXOTransaction(wallet *Wallet, to string, amount int, UTXOSet *UTXOSet) *Transaction
```
NewUTXOTransaction creates a new transaction

#### func (*Transaction) Hash

```go
func (tx *Transaction) Hash() []byte
```
Hash returns the hash of the Transaction

#### func (Transaction) IsCoinbase

```go
func (tx Transaction) IsCoinbase() bool
```
IsCoinbase checks whether the transaction is coinbase

#### func (Transaction) Serialize

```go
func (tx Transaction) Serialize() []byte
```
Serialize returns a serialized Transaction

#### func (*Transaction) Sign

```go
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction)
```
Sign signs each input of a Transaction

#### func (Transaction) String

```go
func (tx Transaction) String() string
```
String returns a human-readable representation of a transaction

#### func (*Transaction) TrimmedCopy

```go
func (tx *Transaction) TrimmedCopy() Transaction
```
TrimmedCopy creates a trimmed copy of Transaction to be used in signing

#### func (*Transaction) Verify

```go
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool
```
Verify verifies signatures of Transaction inputs

#### type UTXOSet

```go
type UTXOSet struct {
	Blockchain *Blockchain
}
```

UTXOSet represents UTXO set

#### func (UTXOSet) CountTransactions

```go
func (u UTXOSet) CountTransactions() int
```
CountTransactions returns the number of transactions in the UTXO set

#### func (UTXOSet) FindSpendableOutputs

```go
func (u UTXOSet) FindSpendableOutputs(pubkeyHash []byte, amount int) (int, map[string][]int)
```
FindSpendableOutputs finds and returns unspent outputs to reference in inputs

#### func (UTXOSet) FindUTXO

```go
func (u UTXOSet) FindUTXO(pubKeyHash []byte) []TXOutput
```
FindUTXO finds UTXO for a public key hash

#### func (UTXOSet) Reindex

```go
func (u UTXOSet) Reindex()
```
Reindex rebuilds the UTXO set

#### func (UTXOSet) Update

```go
func (u UTXOSet) Update(block *Block)
```
Update updates the UTXO set with transactions from the Block The Block is
considered to be the tip of a blockchain

#### type UserAccountStruct

```go
type UserAccountStruct struct {
	ID       uint64
	Email    string
	Password string
	Wallets  []Wallet
}
```

UserAccountStruct hold basic user information

#### type Wallet

```go
type Wallet struct {
	UserID uint64 //Just Added

	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}
```

Wallet stores private and public keys

#### func  NewWallet

```go
func NewWallet() *Wallet
```
NewWallet creates and returns a Wallet

#### func  NewWalletFull

```go
func NewWalletFull(userid uint64, label string) *Wallet
```
NewWalletFull creates a wallet address and assigns the passed in userid and
label and returns it

#### func (Wallet) GetAddress

```go
func (w Wallet) GetAddress() []byte
```
GetAddress returns wallet address

#### type Wallets

```go
type Wallets struct {
	Wallets map[string]*Wallet
}
```

Wallets stores a collection of wallets

#### func  NewWallets

```go
func NewWallets(nodeID string) (*Wallets, error)
```
NewWallets creates Wallets and fills it from a file if it exists

#### func (*Wallets) CreateWallet

```go
func (ws *Wallets) CreateWallet() string
```
CreateWallet adds a Wallet to Wallets

#### func (*Wallets) GetAddresses

```go
func (ws *Wallets) GetAddresses() []string
```
GetAddresses returns an array of addresses stored in the wallet file

#### func (Wallets) GetWallet

```go
func (ws Wallets) GetWallet(address string) Wallet
```
GetWallet returns a Wallet by its address

#### func (*Wallets) LoadFromFile

```go
func (ws *Wallets) LoadFromFile(nodeID string) error
```
LoadFromFile loads wallets from the file

#### func (Wallets) SaveToFile

```go
func (ws Wallets) SaveToFile(nodeID string)
```
SaveToFile saves wallets to a file
