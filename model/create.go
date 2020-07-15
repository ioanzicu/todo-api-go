package model

import "fmt"

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", "Ioan", "Go to sleep at 22:00")
	defer insertQ.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}