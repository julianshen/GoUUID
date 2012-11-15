package uuid

import (
	"testing"
)

func TestVersion4(t *testing.T) {
	id, err := RandomUUID()
	if err != nil {
		t.Fatal(err)
		return
	}

	str := id.String()

	id2, err := UUIDFromString(str)

	switch {
	case err != nil:
		t.Fatal(err)
	case !id.Equals(id2):
		t.Fatal("Wrong UUID generated")
	case id.version() != 4:
		t.Fatalf("Should be version 4 but get Version %d", id.version())
	}

	return
}
