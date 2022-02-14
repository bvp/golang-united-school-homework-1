package solution

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestGetMessage(t *testing.T) {
	msg := GetMessage()
	if msg != emoji.Sprint("Hello :world_map:!") {
		t.Errorf("error")
	}
}
