package excel

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	UTF8BOM = "\xEF\xBB\xBF"
)

func WriteCSVFile(fileName string, tableHeader []string, contents [][]string) (string, error) {
	fileName = fileName + FileSuffixCsv
	filepath, err := NormalizeFilepath(fileName)
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("create csv file %s failed: %s", filepath, err.Error())
	}

	defer file.Close()
	if _, err = file.WriteString(UTF8BOM); err != nil {
		return "", err
	}
	w := csv.NewWriter(file)
	if err = w.Write(tableHeader); err != nil {
		return "", fmt.Errorf("write table header to csv file %s failed: %s", filepath, err.Error())
	}

	if err = w.WriteAll(contents); err != nil {
		return "", fmt.Errorf("write data to csv file %s failed: %s", filepath, err.Error())
	}

	w.Flush()
	return fileName, nil
}

func ReadCSVFile(fileName string) ([][]string, error) {
	filepath, err := NormalizeFilepath(fileName)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(b))
	header, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	} else if !utf8.ValidString(header[0]) {
		r = csv.NewReader(transform.NewReader(bytes.NewReader(b), simplifiedchinese.GB18030.NewDecoder()))
		if header, err = r.Read(); err != nil {
			return nil, err
		}
	}

	contents, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	header[0] = strings.TrimLeft(header[0], UTF8BOM)
	return append([][]string{header}, contents...), nil
}
