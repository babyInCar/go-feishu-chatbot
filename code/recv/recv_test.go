package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {

	out, err := SendMsg("hello", "12312312")
	if err != nil {
		t.Error(err)
	}

	if out != "Hi, What can I do for you!" {
		t.Error("Hello/hello should be Hi, What can I do for you!")
	}
}

func TestCalc2(t *testing.T) {

	out, err := SendMsg("thanks", "23412313")
	if err != nil {
		t.Error(err)
	}

	if out != "It's my pleasure" {
		t.Error("Thanks/thanks should be It's my pleasure")
	}
}

func TestCalc3(t *testing.T) {
	// Pay success
	out, err := SendMsg("success", "12312412")
	if err != nil {
		t.Error(err)
	}

	if out != "Congratulations, Pay success!" {
		t.Error("pay success should reply Congratulations, Pay success!")
	}
}
