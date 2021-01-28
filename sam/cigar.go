package sam

import (
	"errors"
	"regexp"
	"strconv"
)

func CigarConsume(s string) (int, error) {
	// First, check if the CIGAR is valid
	valid, err := regexp.MatchString(`^[0-9MIDNSHP=X]+$`, s)
	if err != nil {
		return -1, errors.New("Failed to validate the CIGAR.")
	}
	if valid {
		// Split the CIGAR into operations
		re := regexp.MustCompile(`([0-9]+)([MIDNSHP=X])`)
		ops := re.FindAllStringSubmatch(s, -1)
		if ops == nil {
			return -1, errors.New("No operation found in CIGAR.")
		}
		refc := 0 // Consumed bases on the reference
		quec := 0 // consumed bases on the query
		for _, op := range ops {
			// Convert the number of consumed bases into int
			c, _ := strconv.Atoi(op[1])
			switch op[2] {
			case "M", "=", "X":
				// Comsume both
				refc += c
				quec += c
				break
			case "I", "S":
				// Consume query
				quec += c
				break
			case "D","N":
				// Consume reference
				refc += c
				break
			case "H", "P":
				// Do not consume
				break
			}
		}
		return refc, nil
	} else {
		return -1, errors.New("Wrong CIGAR format.")
	}
}