package main

import (
	"fmt"
	"os"
)

var osExit = os.Exit

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	osExit(0)
	return nil
}