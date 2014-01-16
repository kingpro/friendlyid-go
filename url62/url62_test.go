package url62

import (
	"testing"
)

func Test_FromUUID(t *testing.T) {
	u, _ := FromUUID("6ba7b814-9dad-11d1-80b4-00c04fd430c8")

	if u != "3h8Pgh03y0pa2W6ltuUfZ6" {
		t.Error("Not valid")
	}
}

func Test_ToUUID(t *testing.T) {
	u, _ := ToUUID("3h8Pgh03y0pa2W6ltuUfZ6")

	if u != "6ba7b814-9dad-11d1-80b4-00c04fd430c8" {
		t.Error("Not valid")
	}
}

func Test_Wrong_ToUUID(t *testing.T) {
	_, err := ToUUID("3h8Pgh0")

	if err == nil {
		t.Error("Not valid")
	}
}
