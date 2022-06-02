package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

//we are only interested in these characteristics of every coin
type Coin struct {
    name string
    price float64
    change float64
}

const (
    /*
        colors are used in order to print the changes

        red:    negative change
        green:  positive or zero change
    */
    red_color string = "\033[31m"
    green_color string = "\033[32m"
    color_reset string = "\033[0m"

    //Between these strings goes the coin ID
    API string = "https://api.coingecko.com/api/v3/coins/"
    API_params string = "?localization=false&tickers=false&developer_data=false&sparkline=false"
)

func FetchAndDisplay(coin_id string) {
    Client := http.Client{}
    var data map[string]interface{}
    var coin Coin

    req, _ := http.NewRequest("GET", API + strings.ToLower(coin_id) + API_params, nil)
    req.Header.Set("Content-type", "application/json")

    //Fetch data
    res, _ := Client.Do(req)

    /*
        Check if response was completed successfully
        if a 404 error occured then the id is not valid
    */
    if res.StatusCode == http.StatusNotFound {
        fmt.Printf("Invalid id '%v'\n", coin_id)
        os.Exit(1)
    }

    body, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()

    json.Unmarshal(body, &data)
    coin.change = data["market_data"].(map[string]interface{})["price_change_percentage_24h"].(float64)
    coin.name = data["name"].(string)
    coin.price = data["market_data"].(map[string]interface{})["current_price"].(map[string]interface{})["usd"].(float64)

    var color string

    //Check for change and set the color
    if coin.change < 0 {
        color = red_color
    } else {
        color = green_color
    }

    fmt.Printf("%15v %15.6f %v%10.1f%%%v\n", strings.ToUpper(coin.name), coin.price, color, coin.change, color_reset)
}

