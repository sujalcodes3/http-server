package util

import (
    "fmt"
    "os"
)

func HandleError(err error, while string) {
	if err != nil {
		er := fmt.Errorf("Error while %s, err: %s\n", while, err.Error())
		HandleError(er, fmt.Sprintf("while displaying error for %s", while))
		os.Exit(1)
	}
}

