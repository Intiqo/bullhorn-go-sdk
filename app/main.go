package main

import (
	"fmt"
	"os"

	"github.com/adwitiyaio/bullhorn-go-sdk"
)

func main() {
	params := &bullhorn.AuthParams{
		ClientId:          "736821b7-a654-4001-a78d-05384152732d",
		ClientSecret:      "7xxkcNc6MNiW73WASArWMObv",
		Username:          "colleago.admin",
		Password:          "cjHVwnbmU!pRY3N",
		AuthenticationUrl: "https://auth-west9.bullhornstaffing.com",
		LoginUrl:          "https://rest-west9.bullhornstaffing.com",
		RestTokenTTL:      "1440",
	}
	client, err := bullhorn.NewClient(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	query := "status = 'Active Account'"
	_, data, err := client.SearchEntity(bullhorn.ClientCorporationEntity, query,
		bullhorn.QueryOptions{Fields: []string{"*"}})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(data.([]bullhorn.ClientCorporation)[0].Name)
}
