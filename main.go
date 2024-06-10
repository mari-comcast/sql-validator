package main

import (
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

	utils.CompareSQLValues(sqlData, txtData)
}
