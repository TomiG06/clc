package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "time"
)

var (
    local = false
    coins_as_args = -1
    help = false
)

var HelpCommand = []string{"--help", "-h"}

func Contains(str string, slice []string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}

func main() {
    argc := len(os.Args)
    coins := []string{};

    if argc == 1 {
        help = true
    } else {
        var lwr string

        for i, v := range os.Args {
            if v[0] != '-' { continue }

            lwr = strings.ToLower(v)

            if Contains(lwr, HelpCommand) {
                help = true
            } else if lwr == "-l" {
                local = true
            } else if lwr == "-c" {
                coins_as_args = i
            } else {
                log.Fatalf("Invalid flag '%v'\nType 'clc --help' for more info\n", os.Args[i])
            }
        }
    }

    if help {
        fmt.Println("usage: 'clc [-l, -c [ids]]'")
        return
    }

    if local {
        local_coins, err := os.ReadFile("../localcoins.txt")

        if err != nil { log.Fatalln("No local coins found") }

        coins = strings.Split(strings.Trim(string(local_coins), " \n"), ",")
    }

    if coins_as_args > 0 {

        for i := coins_as_args + 1; i < argc && os.Args[i][0] != '-'; i++ {
            if !Contains(os.Args[i], coins) {
                coins = append(coins, os.Args[i])
            }
        }
    }

    for _, coin := range coins {
        go FetchAndDisplay(coin)
        time.Sleep(time.Millisecond * 10)
    }

    time.Sleep(time.Second * 1)
}

