package ini

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	ini, err := ReadIni("example_ini/simple.conf")
	if err != nil {
		t.Error(err)
	}

	if ini["Blah"]["One"] != "1" {
		t.Errorf("Expected One to be 1 but got '%s'", ini["Blah"]["One"])
	}
}

func TestStringOutput(t *testing.T) {
	ini, err := ReadIni("example_ini/tiny.conf")
	if err != nil {
		t.Error(err)
	}

	output := fmt.Sprint(ini)
	expectedStrings := []string{
		"Section One__Card Count=3",
		"Section One__Other Attribute=1",
		"Lots of Words__Active=Hello There",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Could not find '%s' in '%s'", expected, output)
		}
	}
}
