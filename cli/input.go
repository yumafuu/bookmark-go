package cli

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput(indicateText string, s *string) {
	fmt.Printf("%v > ", indicateText)
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	*s = stdin.Text()
}
