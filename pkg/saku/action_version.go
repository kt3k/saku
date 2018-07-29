package saku

import "fmt"

func actionVersion() error {
	fmt.Printf("saku@%s\n", Version)

	return nil
}
