package foobargo

import (
	"fmt"
	"strings"
)

// returns upper cased shout
func Echo(shout string) string {
	return fmt.Sprintf("%s", strings.ToUpper(shout))
}
