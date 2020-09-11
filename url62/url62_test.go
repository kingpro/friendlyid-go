package url62

import (
	"testing"
)

func Test_FromUUID(t *testing.T) {
	u, _ := FromUUID("796e28ae-f497-0143-f797-52169b36be94")

	if u != "3h8Pgh03y0pa2W6ltuUfZ6" {
		t.Error("Not valid")
	}
}

func Test_ToUUID(t *testing.T) {
	u, _ := ToUUID("3h8Pgh03y0pa2W6ltuUfZ6")

	if u != "796e28ae-f497-0143-f797-52169b36be94" {
		t.Error("Not valid")
	}
}

func Test_Wrong_ToUUID(t *testing.T) {
	u, _ := ToUUID("3h8Pgh0")

	if u != "00000000-0000-0000-0000-0030e021f32a" {
		t.Error("Not valid")
	}
}

func Test_ManyZeroes_UUID(t *testing.T) {
	ref := "00000000-0000-4000-8000-000000000000"
	u, err := FromUUID(ref)
	if u != "1VgEh72lXvTXkG" {
		t.Error("Not valid")
	}

	d, err := ToUUID(u)
	if err != nil || d != ref {
		t.Error("Not valid")
	}
}
