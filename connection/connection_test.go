package connection_test

import (
	"strings"
	"testing"

	"github.com/m-lab/traceroute-caller/connection"
)

func TestUUID(t *testing.T) {
	conn := connection.Connection{Cookie: "1be3"}
	tmp, err := conn.UUID()
	s := strings.Split(tmp, "_")
	if err != nil || len(s) != 3 || s[2] != "0000000000001BE3" {
		t.Error("Make uuid from cookie incorrect")
	}
}