# Cli Boilerplate

A boiler plate project for the popular `github.com/urfave/cli` previously `codegangsta/cli`.

## Install

`go get github.com/robtec/cli-boilerplate`

or fork this repo

## Build or Install using make

* `make install` - adds the binary to `$GOPATH/bin`
* `make build` - builds the binary

The `Makefile` builds the binary and adds the short git commit id and the version to the final build.

## Customise

The binary is named `cli`, the parent directory of `main.go` in `cmd/cli/`
You can rename the directory to change the name of the final binary

## Extras

Checkout the `github.com/robtec/cli-boilderplate/config` package for saving configs to `$HOME/.cli`
