package validator

import (
	"errors"
	"strings"
)

/*
Validate checks whether a Stellar public address is valid.
It returns true when the address appears valid.
Otherwise it returns false and a descriptive error.
*/

func Validate(address string) (bool, error) {
	address = strings.TrimSpace(address)

	if address == "" {
		return false, errors.New("address cannot be empty")
	}

	if !strings.HasPrefix(address, "G") {
		return false, errors.New("stellar public address must start with 'G'")
	}

	if len(address) != 56 {
		return false, errors.New("stellar public address must be exactly 56 characters")
	}

	return true, nil
}