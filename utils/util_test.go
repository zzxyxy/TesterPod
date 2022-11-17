package utils

import (
	"testing"
	"time"
)

func TestDurationParserHappy(t *testing.T) {
	res, err := ParseDuration("6ms")
	if err != nil || res != 6*time.Millisecond {
		t.Fail()
	}

	res, err = ParseDuration("2s")
	if err != nil || res != 2*time.Second {
		t.Fail()
	}
}
