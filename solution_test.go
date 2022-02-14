package solution

import (
	"testing"

	"github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	msg := GetMessage()
	if msg != emoji.Sprint("Hello :world_map:!") {
		t.Fatalf("msg is not equals %s", "Hello 🗺️!")
	}
}
