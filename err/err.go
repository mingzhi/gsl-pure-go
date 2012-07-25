package err

import (
	"fmt"
)

type Error struct {
	Message string
	Status  int
}

func (err Error) Error() string {
	return fmt.Sprintf("%s, with status %d", err.Message, err.Status)
}
