package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/linkingthing/cement/slice"
)

const (
	UTF8BOM            = "\xEF\xBB\xBF"
	TimeFormat         = "2006-01-02 15:04:05"
	CSVFileSuffix      = ".csv"
	UploadDirectoryKey = "directory"
	UploadFileKey      = "path"
	UploadFileName     = "filename"
	FileResourceName   = "files"
	IgnoreAuditLog     = "ignoreAuditLog"
)

var FileRootPath = "/opt/files"

const (
	ActionNameImportCSV           = "importcsv"
	ActionNameExportCSV           = "exportcsv"
	ActionNameExportCSVTemplate   = "exportcsvtemplate"
	ActionNameImportIPCSV         = "importipcsv"
	ActionNameExportIPCSV         = "exportipcsv"
	ActionNameExportIPCSVTemplate = "exportipcsvtemplate"
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

func (result *ImportResult) FlushResult(fileName string, tableHeader []string) error {
	if len(result.FailedData) > 0 {
		failedFilePath, err := WriteCSVFile(fileName, tableHeader, result.FailedData)
		if err != nil {
			return err
		}
		result.FailedFile = failedFilePath
	}
	return nil
}

func WriteCSVFile(fileName string, tableHeader []string, contents [][]string) (string, error) {
	fileName = fileName + CSVFileSuffix
	filepath := path.Join(FileRootPath, fileName)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", fmt.Errorf("create csv file %s failed: %s", filepath, err.Error())
	}

	defer file.Close()
	if _, err = file.WriteString(UTF8BOM); err != nil {
		return "", err
	}
	w := csv.NewWriter(file)
	if err := w.Write(tableHeader); err != nil {
		return "", fmt.Errorf("write table header to csv file %s failed: %s", filepath, err.Error())
	}

	if err := w.WriteAll(contents); err != nil {
		return "", fmt.Errorf("write data to csv file %s failed: %s", filepath, err.Error())
	}

	w.Flush()
	return fileName, nil
}

func ReadCSVFile(fileName string) ([][]string, error) {
	if strings.Contains(fileName, "../") {
		return nil, fmt.Errorf("file name invalid with path traversal attacks")
	}

	file, err := os.Open(path.Join(FileRootPath, fileName))
	if err != nil {
		return nil, err
	}

	defer file.Close()
	contents, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	if len(contents) > 1 && len(contents[0]) > 1 {
		firstField := strings.TrimLeft(contents[0][0], UTF8BOM)
		contents[0] = append([]string{firstField}, contents[0][1:]...)
	}

	return contents, nil
}

func IsSpaceField(field string) bool {
	for _, r := range field {
		if unicode.IsSpace(r) == false {
			return false
		}
	}

	return true
}

func ParseTableHeader(tableHeaderFields, validTableHeaderFields, mandatoryFields []string) ([]string, error) {
	var headerFields []string
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

	fields := make([]string, 0)
	emptyFieldCnt := 0
	missingMandatory := false
	for i, field := range tableFields {
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

func CreateUploadFolder(folderName string) error {
	if _, err := os.Stat(path.Join(FileRootPath, folderName)); os.IsNotExist(err) {
		if err := os.Mkdir(path.Join(FileRootPath, folderName), 0777); err != nil {
			return fmt.Errorf("createFolder %s failed:%s ", folderName, err.Error())
		}
	}

	return nil
}
