# clc

## Description
Fetches prices from the web and displays them in the command line.
I built this because I wanted to check some cryptocurrency prices while I am working on stuff in the command line

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

## Local Coins
Let's say you are interested in checking the prices of coins x, y, z and a, but you don't want to pass them to the program every time you run it. 
This is why local coins exist. The only thing you have to do is to add them to your local coins by running `clc --add x y z a` command and every time you want to check them just run `clc -l`. In case you want to remove let's say coin y, the only thing you have to do is to run the `clc --remove y` command

