package main

import (
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
)

func get_localcoins() []string {
	local_coins, err := os.ReadFile("../localcoins.txt")

	if err != nil {
		fmt.Println("No local coins found")
		os.Exit(1)
	}

	return strings.Split(strings.Trim(string(local_coins), "\n"), "\n")
}

func set_localcoins(coins_to_be_setted []string) {
	/*
	   Here we are joining the coins in a single string
	   with a new line between them.
	*/
	var bytes_to_write = []byte(strings.Join(coins_to_be_setted, "\n"))

	/*
	   if we have stuff to write, we also want to add
	   another new line character (10 in ASCII) in the
	   end of the string

	   in case there are no stuff to write, we don't want
	   to add it
	*/
	if len(bytes_to_write) != 0 {
		bytes_to_write = append(bytes_to_write, 10)
	}

	if err := os.WriteFile("../localcoins.txt", bytes_to_write, 600); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func add_coins(ids []string) {
	localcoins := get_localcoins()
	/*
	   If the following condition is true
	   it means that the file is empty and
	   because we don't want to add an empty
	   string in localcoins, we are setting
	   localcoins an empty string array
	*/
	if localcoins[0] == "" {
		localcoins = []string{}
	}

	var ids_to_be_added = []string{}
	for _, id := range ids {
		if !slices.Contains(localcoins, id) {
			ids_to_be_added = append(ids_to_be_added, strings.ToLower(id))
		}
	}

	set_localcoins(append(localcoins, ids_to_be_added...))
}

func remove_coins(ids []string) {
	localcoins := get_localcoins()
	var coins_to_be_setted = []string{}

	for _, lc := range localcoins {
		if !slices.Contains(ids, lc) {
			coins_to_be_setted = append(coins_to_be_setted, lc)
		}
	}
	set_localcoins(coins_to_be_setted)
}

func test_ids(ids []string) *string {
	client := http.Client{}

	for _, id := range ids {
		req, _ := http.NewRequest("GET", API+strings.ToLower(id)+API_params, nil)
		res, _ := client.Do(req)

		if res.StatusCode == http.StatusNotFound {
			return &id
		}
	}

	return nil
}
