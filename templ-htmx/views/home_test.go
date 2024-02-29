package views

import (
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	comp, err := compontentToString(Home())
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(comp, "Car Show Example App") {
		t.Errorf("Expected Car Show Example App  got '%s' ", comp)
	}
}
