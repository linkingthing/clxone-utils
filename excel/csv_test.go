package excel

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestWriteAndReadCsv(t *testing.T) {
	FileRootPath = "./"
	contents := [][]string{
		{"H1", "H2", "H3", "H4"},
		{"h", "world", "merge", "merge"},
		{"h", "world", "merge", ""},
		{"aa", `b\b\b`, "", "merge"},
		{"h", "world", "中文", ""},
		{"a", "b&b&b", "c,d,e", ""},
		{"aa", "", "cc;dd;ee", "ff"},
	}

	fileName := fmt.Sprintf("_temp_%s", time.Now().Format("20060102150405"))
	fileName, err := WriteCSVFile(fileName, contents[0], contents[1:])
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(path.Join(FileRootPath, fileName))

	rows, err := ReadCSVFile(fileName)
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
