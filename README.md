# cli
--
    import "\github.com\NlaakStudios\Blockchain\cli"


## Usage

#### type CLI

```go
type CLI struct {
	CoinInfo config.CoinStruct
}
```

CLI is the main client structure

#### func (*CLI) PopulateWallets

```go
func (cli *CLI) PopulateWallets(from string, nodeID string)
```

#### func (*CLI) Run

```go
func (cli *CLI) Run()
```
Run parses command line arguments and processes commands
