package ch2iowriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func FileWrite(path string) error {
	float := 0.1
	integer := 1
	str := "foo"
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, "File contents: %f, %d, %s", float, integer, str)
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func WriteCSV(w io.Writer, records [][]string) error {
	csvWriter := csv.NewWriter(w)
	for _, r := range records {
		if err := csvWriter.Write(r); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}
	return nil
}
