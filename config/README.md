# config
--
    import "\github.com\NlaakStudios\Blockchain\config"


## Usage

```go
const (
	//CoinName is your Coin or Tokens ICO name. ie Bitcoin Copper ICO
	CoinName = "GWF ICO Coin"
	//CoinSymbol is the 3-4 uppercase letters used to identify your coin in the exchange/listings
	CoinSymbol = "GWFI"
	//CoinPurpose is the reason for your new coin or token. What will it privide. Summarize here.
	CoinPurpose = "This blockchain and coin supply is being used for the GWF ICO and will be swapped out 1:1 for Spartacus Token when the actual chains development is complete."
	//CoinAMLCompliant shows if your ICO is AML compliant. Yes|No
	CoinAMLCompliant = "Yes"
	//CoinKYCCompliant shows if your ICO is KYC compliant. Yes|No
	CoinKYCCompliant = "Yes"
	//CoinCompany is the registered name for your company
	CoinCompany = "Nlaak Studios, LLC"
	//CoinCountry is the country in which you compan is registered
	CoinCountry = "Cayman Islands"
	//CoinCEO is the CEO or Owner of the registered company
	CoinCEO = "Andrew Donelson"
	//CoinContact = is the valid direct email to the {CoinCEO} and is also the Account Email
	CoinContact = "gwf@nlaak.com"
	//CoinSubsidy is the Number of coins given to all new addresses
	CoinSubsidy = 0
	//CoinICOSupply is the total number of coins sold in ICO
	CoinICOSupply = uint64(10000000) //10M ICO
	//CoinDevSupply is the total numbner of coins reserved for developers (initial bonus)
	CoinDevSupply = uint64(500000) //500k Dev Team
	//CoinOAMSupply is the total numnber of coins reserved for Operating and Marketing Costs (min 6 months)
	CoinOAMSupply = uint64(1500000) //1.5M Operations & Marketing
	//CoinPLTSupply is the total number of coins reserved for working capital/misc expenses not forseen
	CoinPLTSupply = uint64(3000000) //3M Platform, Misc, Petty Cash, Capital, etc.

	/* INTERNAL USE ONLY */
	//12M, 10M ICO, 400k Staff, 1.6M Operations & Marketing
	CoinTotalSupply = uint64(CoinICOSupply + CoinDevSupply + CoinOAMSupply + CoinPLTSupply)
	CoinDecimals    = 10
	//CoinLandingPage is the URL to your landing page (We handle this)
	//format: https://www/gwf.io/ico/{CoinSymbol}/
	CoinLandingPage = "https://www/gwf.io/ico/%s"
	//CoinWhitePapere is the URL to your whitepaper (We handle this)
	//format: https://www/gwf.io/ico/{CoinSymbol}/
	CoinWhitePaper = "https://www/gwf.io/ico/%s/whitepaper"
	//CoinDashboardPage is the URL to your users dashboard page (We handle this)
	//format: https://www/gwf.io/ico/{CoinSymbol}/dashboard
	CoinDashboardPage = "https://www/gwf.io/ico/%s/dashboard"
	//CoinAPI is the URL to your API endpoints (We handle this)
	//format: https://www/gwf.io/ico/api/{CoinSymbol}/{EndPoint}
	CoinAPI = "https://www/gwf.io/ico/api/%s/%s"
	//FilePathBlockchain is the complete path to the ICO's blockchain database
	//Format: data/ico/{CoinSymbol}/blockchain-{CoinPort}.db
	FilePathBlockchain = "data/ico/%s/blockchain-%s.db"
	//FilePathWallets is the complete path to the ICO's wallets file
	//Format: data/ico/{CoinSymbol}/wallets-{CoinPort}.dat
	FilePathWallets = "data/ico/%s/wallets-%s.dat"
)
```
Set these values for your ICO TODO: Automate this through User Dashboard->Create
ICO ? Create ICO Page (www.gwf.io/createico) -if Not logged in, login or
register account -enter company name, Country and upload docs for proof of
company

```go
const (
	//PhoneTypeUnknown represents a defaul unknown phone type
	PhoneTypeUnknown = byte(0)
	//PhoneTypeMobile represents a Mobile or Cell phone number
	PhoneTypeMobile = byte(1)
	//PhoneTypeHome represents a home phone number
	PhoneTypeHome = byte(2)
	//PhoneTypeBusiness represents a business phone number
	PhoneTypeBusiness = byte(3)
)
```

#### func  LoadCoinInfo

```go
func LoadCoinInfo()
```
LoadCoinInfo loads coin info from config.json file

#### func  SaveCoinInfo

```go
func SaveCoinInfo()
```
SaveCoinInfo saves current coin info to config.json file

#### type CoinStruct

```go
type CoinStruct struct {
	Name      string
	Symbol    string
	Decimals  uint
	Supply    CoinSupplyStruct
	Company   CompanyStruct
	Compliant CompliantStruct
	URLS      CoinURLStruct
}
```

CoinStruct stores all Coin information including Company, Compliantcy

```go
var CoinSettings CoinStruct
```
CoinSettings Holds all information pertaining to the Coin

#### func  GetCoinInfo

```go
func GetCoinInfo() CoinStruct
```
GetCoinInfo Sets the Coin to User defined Data

#### type CoinSupplyStruct

```go
type CoinSupplyStruct struct {
	ICO   uint64
	Dev   uint64
	OAM   uint64
	PLT   uint64
	Total uint64
}
```

CoinSupplyStruct stores the Total Coin Supply aand the breakdown

#### type CoinURLStruct

```go
type CoinURLStruct struct {
	LandingPage   string
	DashboardPage string
	API           string
	WhitePaper    string
}
```

CoinURLStruct stores pre-built URLS for the Coin ICO

#### type CompanyStruct

```go
type CompanyStruct struct {
	Name    string
	Country string
	Ceo     string
	Contact string
	Phone   phone.PhoneStruct
}
```

CompanyStruct stores information about the company behind the coin/blockchain

#### type CompliantStruct

```go
type CompliantStruct struct {
	AML string
	KYC string
}
```

CompliantStruct stores information about AML and KYC states

#### type DateStruct

```go
type DateStruct struct {
}
```

DateStruct store a date of birth for a person, or start of business for a
business

#### type EmailStruct

```go
type EmailStruct struct {
}
```

EmailStruct contains a breakdown of a email

#### type PersonNameStruct

```go
type PersonNameStruct struct {
	GoesBy string //Bob
}
```

PersonNameStruct hold the complete name of a person

#### type PersonStruct

```go
type PersonStruct struct {
	Name  PersonNameStruct
	Email EmailStruct
	Dob   DateStruct
	Phone phone.PhoneStruct
}
```

PersonStruct contains all data pertaining to a individual person
