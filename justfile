
test:
    go build
    sudo cp ./aniba /usr/bin
    clear
    @time aniba

run:
    go run .

install:
    go build
    sudo cp ./aniba /usr/bin

build:
    go build
