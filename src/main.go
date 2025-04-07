package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	local         = false
	coins_as_args = -1
	add           = false
	remove        = false
)

func main() {
	argc := len(os.Args)
	coins := []string{}

	if os.Args[1] == "--add" {
		add = true
	} else if os.Args[1] == "--remove" {
		remove = true
	} else {
		for i, v := range os.Args {
			if v[0] != '-' && i != 1 {
				continue
			}

			if v == "-l" {
				local = true
			} else if v == "-c" {
				coins_as_args = i
			} else {
				fmt.Printf("Invalid flag '%v'\nType 'clc --help' for more info\n", v)
				os.Exit(1)
			}
		}
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

	var wg sync.WaitGroup

	for _, coin := range coins {
		wg.Add(1)
		go FetchAndDisplay(coin, &wg)
	}

	wg.Wait()
}
