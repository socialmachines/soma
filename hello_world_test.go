package soma

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if output := HelloWorld(); output != "Hello World" {
		t.Fatal(output)
	}
}
