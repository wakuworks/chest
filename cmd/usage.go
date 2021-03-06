package cmd

import (
	"fmt"
)

func Usage() {
	fmt.Println(`
Usage:
    chest usage        Show this message
    chest put <file>   Put specified config file in the chest
    chest take <file>  Take specified config file from the chest into the project root
    chest list         Show list of config files in the chest
    chest open         Place config files from the chest into the project root as symbolic links
    chest close        Remove symbolic links created by "open"
`)
}
