package main

import (
	"fmt"
	"os"
)

var osExit = os.Exit

func commandExit(c *config, locationArea string) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	osExit(0)
	return nil
}