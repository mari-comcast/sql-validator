package main

import (
	"fmt"
	"os"

	"github.com/comcast/sql-validator/utils"
)

func main() {
	command, err := utils.InterpretCommand(os.Args)
	if err != nil {
		panic(err.Error())
	}

	sqlData, err := utils.ReadFile(command.SqlFile)
	if err != nil {
		panic(err.Error())
	}

	txtData, err := utils.ReadFile(command.TxtFile)
	if err != nil {
		panic(err.Error())
	}

	isValid := utils.CompareSQLValues(sqlData, txtData)
	if !isValid {
		fmt.Println("SQL values and the given values doesn't match")
	} else {
		fmt.Println("It's a valid SQL and values. You can proceed further !")
	}
}
