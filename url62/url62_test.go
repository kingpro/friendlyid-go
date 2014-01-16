package url62

import (
	"testing"
)

func TestFromUUID(t *testing.T) {
	u := FromUUID("6ba7b814-9dad-11d1-80b4-00c04fd430c8")

	if u != "3h8Pgh03y0pa2W6ltuUfZ6" {
		t.Error("Not valid")
	}
}
