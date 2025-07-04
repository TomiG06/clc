package main

import (
	"fmt"
	"os"
	"slices"
	"sync"
)

var (
	local         = false
	coins_as_args = -1
	add           = false
	remove        = false
)

func parse(args []string) {
	if args[1] == "--add" {
		add = true
		return
	}

	if args[1] == "--remove" {
		remove = true
		return
	}

	for i, v := range args {
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

func main() {
	var coins []string
	args := os.Args
	argc := len(args)

	parse(args)

	if add {
		fmt.Println("Testing id validity...")
		if id := test_ids(args[2:]); id != nil {
			fmt.Printf("ID '%v' is invalid\n", *id)
			os.Exit(1)
		}
		add_coins(args[2:])
		return
	}

	if remove {
		remove_coins(args[2:])
		return
	}

	if local {
		coins = get_localcoins()
	}

	if coins_as_args > 0 {
		for i := coins_as_args + 1; i < argc && args[i][0] != '-'; i++ {
			if !slices.Contains(coins, args[i]) {
				coins = append(coins, args[i])
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
