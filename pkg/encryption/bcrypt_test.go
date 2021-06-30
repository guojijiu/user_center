package encryption

import (
	"fmt"
	"testing"
)

func TestBcryptHash(t *testing.T) {
	raw := "zxcvbnm,./"
	hashV, err := BcryptHash(raw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hashV)
}

func TestBcryptCheck(t *testing.T) {
	//raw := "zxcvbnm,./"
	//hashV := "$2y$04$1oIQ8nqyLOR.WWz8kAT9eup2M3wBEfo0hWfiT4kKeUDjJO/GaWvP."
	//isMatch := tool.BcryptCheck(raw, hashV)
	//fmt.Println(isMatch)
	//assert.Equal(t, true, isMatch)
	var testCaseVss = []struct {
		raw       string
		hashV     string
		expectedV bool
	}{
		{"zxcvbnm,./", "$2y$04$1oIQ8nqyLOR.WWz8kAT9eup2M3wBEfo0hWfiT4kKeUDjJO/GaWvP.", true},
		{"w6CZcV", "$2y$10$HEHg8tyLc1SaRCUwbuJrn.Re1lRexThk7.115YQp4l17JfT8DJxO6", true},
		{"s123456", "$2y$10$au7aF9ur5YtQHxMcqSWZZeNXduPdXB5JhyFsZXF4S1onIUB2wrrQq", true},
		{"zxcvbnm,./123456", "$2y$10$ATnVdwiIv8zesz9bztxFU.Ko9R6sYqQ6yIhkwmeWqSgtw8UeTqS4i", true},
	}
	for _, testCaseV := range testCaseVss {
		isMatch := BcryptCheck(testCaseV.raw, testCaseV.hashV)
		if !isMatch {
			t.Errorf("BcryptCheck(%s, %s) = %t want %t \n", testCaseV.raw, testCaseV.hashV, isMatch, testCaseV.expectedV)
		}
	}
}
