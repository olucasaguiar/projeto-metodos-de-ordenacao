package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Report struct {
	filename string
}

func NewReport(basedir string, filename string, header []string, force bool) (*Report, error) {
	err := CreateDirectory(basedir, true)
	if err != nil {
		return nil, err
	}

	reportPath := fmt.Sprintf("%s%c%s", basedir, os.PathSeparator, filename)
	if FileExists(reportPath) && !force {
		return &Report{filename: reportPath}, nil
	}

	if !FileExists(reportPath) || force {
		err = CreateFile(reportPath, force)
		if err != nil {
			return nil, err
		}
	}

	report := &Report{filename: reportPath}
	report.WriteHeader(header...)
	return report, nil

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
