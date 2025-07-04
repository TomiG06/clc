# clc

[![Go Report Card](https://goreportcard.com/badge/github.com/TomiG06/clc)](https://goreportcard.com/report/github.com/TomiG06/clc)

## Description
Fetches prices from the web and displays them in the command line.
Its purpose of existance is for someone to be able to check the real time prices of their favourite cryptocurrencies without having to open a browser tab

More about the API [here](https://www.coingecko.com/en/api/documentation)

## Setup
```
$ git clone https://github.com/TomiG06/clc.git
$ cd clc
$ ./setup.sh
```

## Commands
* `-l`          fetch local coins
* `-c`          fetch coins entered
* `--add`       add coins entered to local coins
* `--remove`    remove coins entered from local coins
* `--list`      display localcoins

## Local Coins
Let's say you are interested in checking the prices of coins x, y and z, but you don't want to pass them to the program every time you run it. 
This is why local coins exist. The only thing you have to do is to add them to your local coins by running `clc --add x y z` command and every time you want to check them just run `clc -l`. In case you want to remove let's say coin y, the only thing you have to do is to run the `clc --remove y` command

