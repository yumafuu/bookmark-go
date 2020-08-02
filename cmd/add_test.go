package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func testAdd(t *testing.T) {
	add(addCmd, []string{})
	stdin := bytes.NewBufferString("foo\n")
	fmt.Println(stdin)
}
