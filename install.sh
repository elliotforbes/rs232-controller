#!/bin/bash

build_and_run() {
    GOOS=linux GOARCH=amd64 go build -o rs232-controller

    ./rs232-controller &
}

main() {
    sudo apt-get update
    sudo apt-get -y upgrade
    wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz
    sudo tar -xvf go1.13.3.linux-amd64.tar.gz
    sudo mv go /usr/local

    build_and_run

}

main