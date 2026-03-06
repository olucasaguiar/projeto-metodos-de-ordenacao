package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Report struct {
	filename string
}

func NewReport(basedir string, filename string, force bool) (*Report, error) {
	err := CreateDirectory(basedir, true)
	if err != nil {
		return nil, err
	}

	reportPath := fmt.Sprintf("%s%c%s", basedir, os.PathSeparator, filename)
	err = CreateFile(reportPath, force)
	if err != nil {
		return nil, err
	}
	return &Report{filename: reportPath}, nil
}

func (r *Report) WriteHeader(headers ...string) error {
	file, err := os.OpenFile(r.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write(headers)
}

func (r *Report) WriteLines(lines ...[]string) error {
	file, err := os.OpenFile(r.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(lines)
}
