package crypto

import "fmt"

var invalidVersion = fmt.Errorf("invalid version")

func InvalidVersion() error {
	return invalidVersion
}
