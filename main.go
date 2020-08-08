package main

import (
	"fmt"
	"log"
)

func main() {
	v, err := FileVault("test", FileVaultOptions{CreateNew: true})
	if err != nil {
		log.Fatal(err)
	}

	err = v.Set("password", "f", "Gbola")
	if err != nil {
		log.Fatal(err)

	}
	value, err := v.Get("password", "Gbola")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

}
