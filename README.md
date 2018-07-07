# cli
--
    import "\github.com\NlaakStudios\Blockchain\cli"


## Usage

#### type API

```go
type API struct {
	RESTPort string
}
```

API is the main client structure

#### func (*API) InitRESTAPI

```go
func (api *API) InitRESTAPI()
```
InitRESTAPI starts the built in RestAPI Router

#### type CLI

```go
type CLI struct {
	NodePort string
	CoinInfo config.CoinStruct
}
```

CLI is the main client structure

#### func (*CLI) GetBalance

```go
func (cli *CLI) GetBalance(address string) int
```

#### func (*CLI) PopulateWallets

```go
func (cli *CLI) PopulateWallets(from string)
```

#### func (*CLI) Run

```go
func (cli *CLI) Run()
```
Run parses command line arguments and processes commands
