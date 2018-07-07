package cli

import (
	"fmt"
	"net/http"

	"github.com/NlaakStudios/Blockchain/config"

	"github.com/gin-gonic/gin"
)

//API is the main client structure
type API struct {
	RESTPort string
}

// InitRESTAPI starts the built in RestAPI Router
func (api *API) InitRESTAPI() {
	fmt.Printf("Starting REST on port %s...", api.RESTPort)

	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/", defaultEndpoint)
		apiv1.POST("/balance", getBalanceEndpoint)
		apiv1.POST("/submit", submitEndpoint)
		apiv1.POST("/read", readEndpoint)
	}

	// Start and run the server
	go router.Run(fmt.Sprintf("localhost:%s", api.RESTPort))
	fmt.Printf("Success.\n")
}

/********************************************************************
 *** Blockchain API V1 PUBLIC Functions
 ********************************************************************/
func defaultEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "GWF Blockchain API",
		"version": config.Version(),
		"public":  "usage",
		"/":       "default, show version and endpoints",
		"/wallet/balance/{address}":         "Get balance of address",
		"/wallet/history/{address}":         "Get transaction history of address",
		"/wallets/list":                     "list of all addresses",
		"authenticated":                     "usage",
		"/wallet/send/{amount}/{from}/{to}": "Send amount from to",
	})
}

func getBalanceEndpoint(c *gin.Context) {
	//balance := cli.GetBalance("{address}")
	c.JSON(http.StatusOK, gin.H{
		"message": "getBalance",
	})
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Submit",
	})
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Read",
	})
}

/********************************************************************
 *** Blockchain API V1 AUTHENTICATED Functions
 ********************************************************************/
