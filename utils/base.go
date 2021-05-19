package utils

import (
	"fmt"
)

func checkErr(err error, des string) error {
	if err != nil {
		_ = fmt.Sprintf("%s Error ... %s", des, err)
		// panic(err)
		return err
	}
	return nil
}
