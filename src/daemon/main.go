package daemon

import (
	"fmt"

	"github.com/ttacon/chalk"
)

// Start does nothing for now.
func Start() {
	fmt.Println(chalk.Red.Color("The daemon is not ready yet."))
}
