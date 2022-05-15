#!/bin/bash

cd src
go build main.go coin.go

cd ../
sudo ln -sf $(pwd)/clc /usr/local/bin/clc
