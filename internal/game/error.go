package game

import (
	"fmt"
)

type levelError struct {
	level int
	name  string
}

func (l levelError) Error() string {
	return fmt.Sprintf("level error : %s (%d)", l.name, l.level)
}
