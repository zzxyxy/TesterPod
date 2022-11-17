package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const ERRORMESSAGE = "the delay must be given in a format like 1s or 1000ms"

func ParseDuration(raw string) (time.Duration, error) {
	re := regexp.MustCompile(`^(\d+)(s|ms)$`)

	if !re.Match([]byte(raw)) {
		return 0, fmt.Errorf(ERRORMESSAGE)
	}

	parts := re.FindStringSubmatch(raw)
	delay, _ := strconv.ParseInt(parts[1], 10, 64)

	var multiplier time.Duration
	if parts[2] == "s" {
		multiplier = time.Second
	} else {
		multiplier = time.Millisecond
	}

	return time.Duration(delay) * multiplier, nil
}
