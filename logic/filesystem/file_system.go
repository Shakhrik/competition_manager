package filesystem

import (
	"os"
)

func GetOrCreateFile(filename string) (*os.File, error) {
	var file *os.File
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		file, err = os.Open(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

}
