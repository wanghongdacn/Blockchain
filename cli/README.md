# cli
--
    import "\github.com\NlaakStudios\Blockchain\cli"


## Usage

#### func  InitBlockchainAPI

```go
func InitBlockchainAPI()
```
InitBlockchainAPI starts the built in RestAPI Router

#### type CLI

```go
type CLI struct {
	CoinInfo config.CoinStruct
}
```


#### func (*CLI) PopulateWallets

```go
func (cli *CLI) PopulateWallets(from string, nodeID string)
```

#### func (*CLI) Run

```go
func (cli *CLI) Run()
```
Run parses command line arguments and processes commands
