package api

import (

	//"github.com/NlaakStudios/Blockchain/cli"
	"fmt"
	"net/http"

	"github.com/NlaakStudios/Blockchain/config"

	"github.com/gin-gonic/gin"
)

// InitBlockchainAPI starts the built in RestAPI Router
func InitBlockchainAPI() {
	fmt.Printf("Starting RESTAPI %s\n", ":4000")

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
	go router.Run(":3001")
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
