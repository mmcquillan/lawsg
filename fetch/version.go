package fetch

import (
	"fmt"
)

func Version(version string) {
	if version == "" {
		version = "(compiled)"
	}
	fmt.Println("")
	fmt.Println("lawsg", version)
	fmt.Println("")
}
