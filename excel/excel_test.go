package excel

import (
	"fmt"
	"github.com/linkingthing/cement/slice"
	"os"
	"path"
	"testing"
	"time"
)

func TestWriteAndReadExcel(t *testing.T) {
	FileRootPath = "./"
	contents := [][]string{
		{"H1", "H2", "H3", "H4"},
		{"h", "world", "merge", "merge"},
		{"h", "world", "merge", ""},
		{"aa", "b/b/b", ""},
		{"aa", `b\b\b`, "", "merge"},
		{"h", "world", "中文", ""},
		{"a", "b&b&b", "c,d,e", ""},
		{"l", "m:n::q"},
		{"aa", "", "cc;dd;ee", "ff"},
		{"x", "y", "x\ny\nee"},
	}

	dropList := []string{"h", "x", "a", "l", "aa"}
	ops := []Operate{
		{
			Range: Range{
				TopLeftAxis:     Coordinate{ColNo: 1, RowNo: 2},
				BottomRightAxis: Coordinate{ColNo: 1, RowNo: len(contents)},
			},
			Type:  OperateSetDropList,
			Extra: dropList,
		},
		{
			Range: Range{
				TopLeftAxis:     Coordinate{ColNo: 3, RowNo: 2},
				BottomRightAxis: Coordinate{ColNo: 4, RowNo: 5},
			},
			Type: OperateMergeCell,
		},
	}

	fileName := fmt.Sprintf("_temp_%s", time.Now().Format("20060102150405"))
	fileName, err := WriteExcelFile(fileName, contents[0], contents[1:], ops...)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(path.Join(FileRootPath, fileName))

	rows, err := ReadExcelFile(fileName)
	if err != nil {
		t.Fatal(err)
	}

	topLeft := ops[1].Range.TopLeftAxis
	bottomRight := ops[1].Range.BottomRightAxis
	mergeValue := contents[topLeft.RowNo-1][topLeft.ColNo-1]
	contents[topLeft.RowNo][bottomRight.ColNo-1] = mergeValue
	contents[topLeft.RowNo+1][bottomRight.ColNo-2] = mergeValue
	contents[topLeft.RowNo+2][bottomRight.ColNo-2] = contents[topLeft.RowNo-1][topLeft.ColNo-1]

	for i, row := range rows {
		if i > 0 && len(row) > 0 && slice.SliceIndex(dropList, row[0]) < 0 {
			t.Fatalf("%s not in drop list %q", row[0], dropList)
		}

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
