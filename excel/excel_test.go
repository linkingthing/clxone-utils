package excel

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestWriteAndReadExcel(t *testing.T) {
	FileRootPath = "./"
	contents := [][]string{
		{"H1", "H2", "H3"},
		{"hello", "world", "中文"},
		{"a", "b&b&b", "c,d,e"},
		{"l", "m:n::q"},
		{"aa", "", "cc;dd;ee", "ff"},
		{"x", "y", "x\ny\nee"},
	}

	fileName := fmt.Sprintf("_temp_%s", time.Now().Format("20060102150405"))
	fileName, err := WriteExcelFile(fileName, contents[0], contents[1:])
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(path.Join(FileRootPath, fileName))

	rows, err := ReadExcelFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	for i, row := range rows {
		target := contents[i]
		for j := 0; j < len(row); j++ {
			if target[j] != row[j] {
				t.Fatalf("%d:%d: want: %q, got: %q", i, j, target, row)
			}
		}
	}
}

func TestParseTableFields(t *testing.T) {
	header := []string{"h1", "h2", "h3"}
	rows := [][]string{
		{"x"},
		{"hello", "world", "中文"},
		{"a", "b&b&b", "c,d,e"},
		{"l", "m:n::q"},
		{"y1", "y2", "y3", "y4", "y5"},
	}

	result := [][]string{
		{"x", "", ""},
		{"hello", "world", "中文"},
		{"a", "b&b&b", "c,d,e"},
		{"l", "m:n::q", ""},
		{"y1", "y2", "y3"},
	}

	for i, row := range rows {
		fields, _, _ := ParseTableFields(row, header, nil)
		if len(fields) != len(result[i]) {
			t.Fatalf("row: %d, want: %q, got: %q", i, result[i], fields)
		}
		for j := 0; j < len(fields); j++ {
			if fields[j] != result[i][j] {
				t.Fatalf("file:%d:%d: want: %s, got: %s",
					i, j, result[i][j], fields[j])
			}
		}
	}
}
