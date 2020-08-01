package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := stdin.Text()
	items := strings.Fields(input)
	fmt.Println(items[1])
	exec.Command("open", items[1]).Run()
}
