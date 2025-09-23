package forecast

import (
	"fmt"
	"testing"
)

func TestGetAlerts_Success(t *testing.T) {
	input := Input{
		State: "NY",
	}
	_, out, _ := GetAlerts(nil, nil, input)

	fmt.Printf("Output: %+v\n", out)
}
