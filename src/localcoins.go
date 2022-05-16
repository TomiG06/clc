package main

import(
    "fmt"
    "os"
    "strings"
)

func Contains(str string, slice []string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}


func get_localcoins() []string {
    local_coins, err := os.ReadFile("../localcoins.txt")

    if err != nil {
        fmt.Println("No local coins found")
        os.Exit(1)
    }

    return strings.Split(strings.Trim(string(local_coins), "\n"), "\n")
}

func set_localcoins(coins_to_be_setted []string) {
    if err := os.WriteFile("../localcoins.txt", []byte(strings.Join(coins_to_be_setted, "\n")), 600); err != nil {
        fmt.Print(err)
        os.Exit(1)
    }
}

func add_coins(ids []string) {
    localcoins := get_localcoins()
    var ids_to_be_added = []string{}
    for _, id := range ids {
        if !Contains(id, localcoins) {
            ids_to_be_added = append(ids_to_be_added, strings.ToLower(id))
        }
    }

    set_localcoins(append(localcoins, ids_to_be_added...))
}

func remove_coins(ids []string) {
    localcoins := get_localcoins()
    var coins_to_be_setted = []string{}

    for _, lc := range localcoins {
        if !Contains(lc, ids) {
            coins_to_be_setted = append(coins_to_be_setted, lc)
        }
    }
    set_localcoins(coins_to_be_setted)
}

