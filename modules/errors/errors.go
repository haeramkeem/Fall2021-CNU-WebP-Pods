package errors

import (
	"pods/modules/logging"
)

func Check(err error) error {
    if err != nil {
        logging.Warn(err)
        return err
    }
    return nil
}
