package main

import (
	"fmt"
	code "gitee.com/windcoder/qingyucms/internal/pkg/qycms-error-code"
	errors "gitee.com/windcoder/qingyucms/internal/pkg/qygo-errors"
)

func main() {
	if err := getUser(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
func getUser() error {
	if err := queryDataBase(); err != nil {
		return errors.Wrap(err, "get user failed.")
	}
	return nil
}
func queryDataBase() error {
	return errors.WithCode(code.ErrDatabase, "user 'Lingfei Kong' not found.")
}
