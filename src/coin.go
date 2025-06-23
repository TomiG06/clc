package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// we are only interested in these characteristics of every coin
type Coin struct {
	name   string
	price  float64
	change float64
}

const (
	/*
	   colors are used in order to print the changes

	   red:    negative change
	   green:  positive or zero change
	*/
	red_color   string = "\033[31m"
	green_color string = "\033[32m"
	color_reset string = "\033[0m"

	//Between these strings goes the coin ID
	API        string = "https://api.coingecko.com/api/v3/coins/"
	API_params string = "?localization=false&tickers=false&developer_data=false&sparkline=false"
)

// Print TooManyRequests message once
var once sync.Once = sync.Once{}

func Display(coin *Coin) {
	var color string
	if coin.change < 0 {
		color = red_color
	} else {
		color = green_color
	}

	fmt.Printf(
		"%15v %15.2f %v%10.1f%%%v\n",
		strings.ToUpper(coin.name),
		coin.price,
		color,
		coin.change,
		color_reset,
	)
}

func FetchAndDisplay(coin_id string, wg *sync.WaitGroup) {
	defer wg.Done()

	Client := http.Client{}
	var data map[string]interface{}
	var coin Coin

	req, _ := http.NewRequest("GET", API+strings.ToLower(coin_id)+API_params, nil)
	req.Header.Set("Content-type", "application/json")

	//Fetch data
	res, _ := Client.Do(req)
	defer res.Body.Close()

	/*
		   Check if response was completed successfully
		   if a 404 error occurs then the id is not valid
		   if a 429 error occurs then there were too many requests
		   (we are using the free api version)

	       return is cleaner, even if everyone will end up doing it
	*/
	if res.StatusCode == http.StatusNotFound {
		fmt.Printf("Invalid id '%v'\n", coin_id)
		return
	} else if res.StatusCode == http.StatusTooManyRequests {
		once.Do(func() {
			fmt.Println("Too many requests. Try again in a couple of minutes.")
		})
		return
	}

	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &data)
	coin.change = data["market_data"].(map[string]interface{})["price_change_percentage_24h"].(float64)
	coin.name = data["name"].(string)
	coin.price = data["market_data"].(map[string]interface{})["current_price"].(map[string]interface{})["usd"].(float64)

	Display(&coin)
}
