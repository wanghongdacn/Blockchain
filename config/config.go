package config

//github.com/fatih/structs

// Set these values for your ICO
// TODO: Automate this through User Dashboard->Create ICO ?
//Create ICO Page (www.gwf.io/createico)
//-if Not logged in, login or register account
//-enter company name, Country and upload docs for proof of company
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

//*********************************************************************
//************[DONT CHANGE ANYTHING BELOW THIS LINE]*******************
//*********************************************************************

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

const (
	//NodePort is used for blockchain communication/syncing
	NodePort = "3000"
	//APIPort is used for REST API access to nbode
	RESTPort = "4000"
)

//EmailStruct contains a breakdown of a email
type EmailStruct struct {
	name   string //bob1234
	domain string //gmail.com
}

//PhoneStruct is used to breakdown and store phone numbers
type PhoneStruct struct {
	CountryCode string
	AreaCode    string
	Number      string
	PhoneType   byte
}

//CompanyStruct stores information about the company behind the coin/blockchain
type CompanyStruct struct {
	Name    string
	Country string
	Ceo     string
	Contact string
	Phone   PhoneStruct
}

//PersonNameStruct hold the complete name of a person
type PersonNameStruct struct {
	prefix string //ie. Mr
	first  string //William
	middle string //Blaine
	last   string //Doe
	suffix string //Sr
	GoesBy string //Bob
}

//DateStruct store a date of birth for a person, or start of business for a business
type DateStruct struct {
	year  uint //Greater than current year minus 18
	month uint //1..12
	day   uint //1..31
}

//PersonStruct contains all data pertaining to a individual person
type PersonStruct struct {
	Name  PersonNameStruct
	Email EmailStruct
	Dob   DateStruct
	Phone PhoneStruct
}

//CoinSupplyStruct stores the Total Coin Supply aand the breakdown
type CoinSupplyStruct struct {
	ICO   uint64
	Dev   uint64
	OAM   uint64
	PLT   uint64
	Total uint64
}

//CompliantStruct stores information about AML and KYC states
type CompliantStruct struct {
	AML string
	KYC string
}

//CoinURLStruct stores pre-built URLS for the Coin ICO
type CoinURLStruct struct {
	LandingPage   string
	DashboardPage string
	API           string
	WhitePaper    string
}

//CoinStruct stores all Coin information including Company, Compliantcy
type CoinStruct struct {
	Name      string
	Symbol    string
	Decimals  uint
	Supply    CoinSupplyStruct
	Company   CompanyStruct
	Compliant CompliantStruct
	URLS      CoinURLStruct
}

//CoinSettings Holds all information pertaining to the Coin
var CoinSettings CoinStruct

//GetCoinInfo Sets the Coin to User defined Data
func GetCoinInfo() CoinStruct {
	return CoinStruct{
		CoinName,
		CoinSymbol,
		CoinDecimals,
		CoinSupplyStruct{
			CoinICOSupply,
			CoinDevSupply,
			CoinOAMSupply,
			CoinPLTSupply,
			CoinTotalSupply,
		},
		CompanyStruct{
			CoinCompany,
			CoinCountry,
			CoinCEO,
			CoinContact,
			PhoneStruct{
				"1",
				"928",
				"2463691",
				1,
			},
		},
		CompliantStruct{
			CoinAMLCompliant,
			CoinKYCCompliant,
		},
		CoinURLStruct{
			CoinLandingPage,
			CoinDashboardPage,
			CoinAPI,
			CoinWhitePaper,
		},
	}
}

//SaveCoinInfo saves current coin info to config.json file
func SaveCoinInfo() {}

//LoadCoinInfo loads coin info from config.json file
func LoadCoinInfo() {}
