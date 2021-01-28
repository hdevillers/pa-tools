package sam

import (
	"errors"
	"regexp"
)

func CigarConsume(s string) (int, error) {
	// First, check if the CIGAR is valid
	valid, err := regexp.MatchString(`^[0-9MIDNSHP=X]+$`, s)
	if err != nil {
		return -1, errors.New("Failed to validate the CIGAR.")
	}
	if valid {
		// Split the CIGAR into operations
		c := 1
		return c, nil
	} else {
		return -1, errors.New("Wrong CIGAR format.")
	}
}