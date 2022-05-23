#!/bin/bash

cd src
go build main.go coin.go localcoins.go

cd ../
touch localcoins.txt
sudo ln -sf $(pwd)/clc /usr/local/bin/clc
