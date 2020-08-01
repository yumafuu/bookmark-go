package cli

import (
	"bufio"
	"fmt"
	"os"
)

func WaitingInput(indicateText string, v *string) {
	fmt.Printf("%v > ", indicateText)
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	*v = stdin.Text()
}