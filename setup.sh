#!/bin/bash

cd src
make

cd ../
touch localcoins.txt
sudo ln -sf $(pwd)/clc /usr/local/bin/clc
