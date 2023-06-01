package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

const (
	DefaultSheet = "Sheet1"
)

type OperateType uint8

const (
	OperateSetDropList OperateType = iota + 1
	OperateMergeCell
)

func (ot OperateType) String() string {
	switch ot {
	case OperateSetDropList:
		return "set-drop-list"
	case OperateMergeCell:
		return "merge-cell"
	default:
		return fmt.Sprintf("<%d>", ot)
	}
}

type Coordinate struct {
	ColNo int
	RowNo int
}

func (cr Coordinate) ToCellName() (string, error) {
	return excelize.CoordinatesToCellName(cr.ColNo, cr.RowNo)
}

type Range struct {
	TopLeftAxis     Coordinate
	BottomRightAxis Coordinate
}

func (r Range) ToCellRange() (topLeftCell, bottomRightCell string, err error) {
	topLeftCell, err = r.TopLeftAxis.ToCellName()
	if err != nil {
		return
	}

	bottomRightCell, err = r.BottomRightAxis.ToCellName()
	if err != nil {
		return
	}

	return
}

type Operate struct {
	Range

	Type  OperateType
	Extra interface{}
}

type ExcelParser struct {
	*excelize.File

	create   bool
	filepath string
}

func NewExcelParser(fileName string, create bool) (*ExcelParser, error) {
	fpath, err := NormalizeFilepath(fileName)
	if err != nil {
		return nil, err
	}

	var fp *excelize.File
	if create {
		fp = excelize.NewFile()
	} else {
		fp, err = excelize.OpenFile(fpath)
		if err != nil {
			return nil, fmt.Errorf("open %s failed: only support format of XLSX", fileName)
		}
	}

	return &ExcelParser{
		File:     fp,
		create:   create,
		filepath: fpath,
	}, nil
}

func (parser *ExcelParser) Close() error {
	if parser.create {
		return parser.SaveAs(parser.filepath)
	}
	return parser.File.Close()
}

func (parser *ExcelParser) ReadAll(sheet string) ([][]string, error) {
	rows, err := parser.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	mergeCells, err := parser.GetMergeCells(sheet)
	if err != nil {
		return nil, err
	}

	for _, mc := range mergeCells {
		startCol, startRow, err := excelize.CellNameToCoordinates(mc.GetStartAxis())
		if err != nil {
			return nil, err
		}

		endCol, endRow, err := excelize.CellNameToCoordinates(mc.GetEndAxis())
		if err != nil {
			return nil, err
		}

		value := mc.GetCellValue()
		for i := startRow - 1; i < endRow && i < len(rows); i++ {
			for j := startCol - 1; j < endCol && j < len(rows[i]); j++ {
				rows[i][j] = value
			}
		}
	}

	return rows, nil
}

func (parser *ExcelParser) WriteAll(sheet string, contents [][]string) error {
	for i := range contents {
		if err := parser.SetSheetRow(sheet, fmt.Sprintf("A%d", i+1), &contents[i]); err != nil {
			return fmt.Errorf("write row %d failed: %v", i+1, err)
		}
	}

	return nil
}

func (parser *ExcelParser) ApplyOperates(sheet string, operates []Operate) (err error) {
	for _, op := range operates {
		switch op.Type {
		case OperateSetDropList:
			err = parser.setDropList(sheet, op)
		case OperateMergeCell:
			err = parser.mergeCells(sheet, op)
		default:
			err = fmt.Errorf("unsupported operate type of %s", op.Type)
		}

		if err != nil {
			return
		}
	}
	return
}

func (parser *ExcelParser) setDropList(sheet string, op Operate) error {
	options, _ := op.Extra.([]string)
	if len(options) == 0 {
		return fmt.Errorf("invalid drop options: %v", op.Extra)
	}

	topLeftCell, bottomRightCell, err := op.ToCellRange()
	if err != nil {
		return err
	}

	dv := excelize.NewDataValidation(true)
	dv.Sqref = fmt.Sprintf("%s:%s", topLeftCell, bottomRightCell)
	dv.SetDropList(options)
	return parser.AddDataValidation(sheet, dv)
}

func (parser *ExcelParser) mergeCells(sheet string, op Operate) error {
	topLeftCell, bottomRightCell, err := op.ToCellRange()
	if err != nil {
		return err
	}

	return parser.MergeCell(sheet, topLeftCell, bottomRightCell)
}

func MakeDropListOperate(opRange Range, dropList []string) Operate {
	return Operate{
		Range: opRange,
		Type:  OperateSetDropList,
		Extra: dropList,
	}
}

func MakeMergeCellOperate(opRange Range) Operate {
	return Operate{
		Range: opRange,
		Type:  OperateMergeCell,
	}
}

func WriteExcelFile(fileName string, header []string, contents [][]string, ops ...Operate) (string, error) {
	fileName = fileName + FileSuffixExcel
	parser, err := NewExcelParser(fileName, true)
	if err != nil {
		return "", err
	}

	if err = parser.WriteAll(DefaultSheet, append([][]string{header}, contents...)); err != nil {
		return "", err
	}

	if err = parser.ApplyOperates(DefaultSheet, ops); err != nil {
		return "", err
	}

	return fileName, parser.Close()
}

func ReadExcelFile(fileName string, sheets ...string) ([][]string, error) {
	parser, err := NewExcelParser(fileName, false)
	if err != nil {
		return nil, err
	}
	defer parser.Close()

	if len(sheets) == 0 {
		sheets = append(sheets, parser.GetSheetName(0))
	}

	var contents [][]string
	for _, sheet := range sheets {
		rows, err := parser.ReadAll(sheet)
		if err != nil {
			return contents, fmt.Errorf("read sheet %s error: %v", sheet, err)
		}
		contents = append(contents, rows...)
	}
	return contents, nil
}
