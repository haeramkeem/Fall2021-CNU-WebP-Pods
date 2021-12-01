package errors

import (
	"pods/modules/logging"
)

func Check(err error) {
    if err != nil {
        logging.Fatal(err)
    }
}
