package excel

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/linkingthing/cement/slice"
)

type FileType uint8

const (
	FileTypeExcel FileType = iota
	FileTypeCsv
)

func (ft FileType) String() string {
	switch ft {
	case FileTypeExcel:
		return "excel"
	case FileTypeCsv:
		return "csv"
	default:
		return fmt.Sprintf("<%d>", ft)
	}
}

const (
	FileSuffixExcel = ".xlsx"
	FileSuffixCsv   = ".csv"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

const (
	ActionNameImport           = "import"
	ActionNameExport           = "export"
	ActionNameExportTemplate   = "export_template"
	ActionNameImportIP         = "import_ip"
	ActionNameExportIP         = "export_ip"
	ActionNameExportIPTemplate = "export_ip_template"
)

type ImportFile struct {
	Name string `json:"name"`
}

type ExportFile struct {
	Path string `json:"path"`
}

type ImportResult struct {
	Total      int        `json:"total"`
	Success    int        `json:"success"`
	Failed     int        `json:"failed"`
	FailedFile string     `json:"failedFile"`
	FailedData [][]string `json:"-"`
	FileType   FileType   `json:"-"`
}

func (result *ImportResult) InitData(total int) {
	result.Total = total
	result.Success = result.Total
	result.Failed = 0
}

func (result *ImportResult) AddFailedData(data []string) {
	result.Failed++
	result.Success--
	result.FailedData = append(result.FailedData, data)
}

func (result *ImportResult) FlushResult(fileName string, tableHeader []string) (err error) {
	if len(result.FailedData) == 0 {
		return nil
	}

	switch result.FileType {
	case FileTypeExcel:
		result.FailedFile, err = WriteExcelFile(fileName, tableHeader, result.FailedData)
	case FileTypeCsv:
		result.FailedFile, err = WriteCSVFile(fileName, tableHeader, result.FailedData)
	default:
		return fmt.Errorf("unsupported file type of %s", result.FileType)
	}

	return
}

func ParseTableHeader(tableHeaderFields, validTableHeaderFields, mandatoryFields []string) ([]string, error) {
	headerFields := make([]string, 0, len(tableHeaderFields))
	mandatoryFieldCnt := 0
	for _, field := range tableHeaderFields {
		field = strings.Trim(field, "\r\n ")
		if slice.SliceIndex(validTableHeaderFields, field) == -1 {
			return nil, fmt.Errorf("the file table header field %s is invalid", field)
		} else if slice.SliceIndex(mandatoryFields, field) != -1 {
			mandatoryFieldCnt += 1
		}
		headerFields = append(headerFields, field)
	}

	if mandatoryFieldCnt != len(mandatoryFields) {
		return nil, fmt.Errorf("the file must contains mandatory field %v", mandatoryFields)
	}

	return headerFields, nil
}

func ParseTableFields(tableFields, tableHeaderFields, mandatoryFields []string) ([]string, bool, bool) {
	if len(tableFields) == 0 {
		return nil, true, true
	}

	// ensure every row's length is same as the header's
	for i := len(tableHeaderFields) - len(tableFields); i > 0; i-- {
		tableFields = append(tableFields, "")
	}

	fields := make([]string, 0)
	emptyFieldCnt := 0
	missingMandatory := false
	for i, field := range tableFields {
		if i >= len(tableHeaderFields) {
			break
		}

		if IsSpaceField(field) {
			if slice.SliceIndex(mandatoryFields, tableHeaderFields[i]) != -1 {
				missingMandatory = true
			}
			emptyFieldCnt += 1
			fields = append(fields, "")
		} else {
			field = strings.TrimRight(field, "\r\n ")
			fields = append(fields, field)
		}
	}

	return fields, missingMandatory, emptyFieldCnt == len(tableFields)
}

func IsSpaceField(field string) bool {
	for _, r := range field {
		if unicode.IsSpace(r) == false {
			return false
		}
	}

	return true
}
