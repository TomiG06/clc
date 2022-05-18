package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

var (
    local = false
    coins_as_args = -1
    help = false
    add = false
    remove = false
)

func main() {
    argc := len(os.Args)
    coins := []string{};

    if argc == 1 {
        os.Args = append(os.Args, "--help")
    }

    lwr := strings.ToLower(os.Args[1])

    if lwr == "--help" {
        help = true
    } else if lwr == "--add" {
        add = true
    } else if lwr == "--remove" {
        remove = true
    } else {
        var lwr string

        for i, v := range os.Args {
            if v[0] != '-' && i != 1 { continue }

            lwr = strings.ToLower(v)

            if lwr == "-l" {
                local = true
            } else if lwr == "-c" {
                coins_as_args = i
            } else {
                fmt.Printf("Invalid flag '%v'\nType 'clc --help' for more info\n", v)
                os.Exit(1)
            }
        }
    }

    if help {
        fmt.Println("usage: 'clc [--help, --remove, --add] [-l, -c] [id's]'")
        return
    }

    if add {
        fmt.Println("Testing id validity...")

        if id := test_ids(os.Args[2:]); id != nil {
            fmt.Printf("ID '%v' is invalid\n", *id)
            os.Exit(1)
        }

        add_coins(os.Args[2:])
        return
    }

    if remove {
        remove_coins(os.Args[2:])
        return
    }

    if local {
        coins = get_localcoins()
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
    }

    time.Sleep(time.Second * 1)
}

