package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    argc := len(os.Args)
    coins := []string{};

    //Check if locals are wanted
    if argc > 1 && strings.ToLower(os.Args[1]) != "-l" {

        for x := 1; x < argc; x++ {
            coins = append(coins, os.Args[x])
        }

    } else {

        local_coins, err := os.ReadFile("../localcoins.txt")

        if err != nil {
            fmt.Println("No local coins found")
            os.Exit(1)
        }

        coins = strings.Split(strings.Trim(string(local_coins), " \n"), ",")

    }

    for _, coin := range coins {
        go FetchAndDisplay(coin)
    }

    time.Sleep(time.Second * 1)
}

