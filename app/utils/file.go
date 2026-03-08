package utils

import "os"

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DeleteFile(filename string) error {
	return os.Remove(filename)
}

func CreateFile(filename string, force bool) error {
	if force {
		DeleteFile(filename)
	}

	_, err := os.Create(filename)
	if err != nil {
		return err
	}

	return nil
}

func CreateDirectory(dirname string, recursive bool) error {
	if recursive {
		return os.MkdirAll(dirname, os.ModePerm)
	}
	return os.Mkdir(dirname, os.ModePerm)
}
