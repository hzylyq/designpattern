package pagemaker

import (
	"os"
)

type dataBase struct {
}

func NewDataBase() *dataBase {
	return &dataBase{}
}

func (db *dataBase) Properties(dbName string) error {
	fileName := dbName + ".txt"

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		return err
	}

}
