package main

import (
	"fmt"
	"testing"
)

func TestGetAlerts_Success(t *testing.T) {
	input := AlertInput{
		State: "NY",
	}
	_, out, _ := getAlerts(nil, nil, input)

	fmt.Printf("Output: %+v\n", out)
}
