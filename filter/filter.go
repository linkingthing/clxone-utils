package filter

import (
	"strings"

	restresource "github.com/linkingthing/gorest/resource"
)

const (
	FilterNameName       = "name"
	FilterNameComment    = "comment"
	FilterNameCreateTime = "create_time"
	IgnoreAuditLog       = "ignoreAuditLog"
	SymbolComma          = ","

	FilterNameTimeFrom = "from"
	FilterNameTimeTo   = "to"
	TimeFromSuffix     = " 00:00"
	TimeToSuffix       = " 23:59"
	TimeFormatYMD      = "2006-01-02"
	TimeFormatYMDHM    = "2006-01-02 15:04"
)

func GenStrConditionsFromFilters(filters []restresource.Filter, filterNames ...string) map[string]interface{} {
	if len(filters) == 0 {
		return nil
	}

	conditions := make(map[string]interface{})
	for _, filterName := range filterNames {
		if value, ok := GetFilterValueWithEqModifierFromFilters(filterName, filters); ok {
			conditions[filterName] = value
		}
	}

	return conditions
}

func GetFilterValueWithEqModifierFromFilters(filterName string, filters []restresource.Filter) (string, bool) {
	for _, filter := range filters {
		if filter.Name == filterName && filter.Modifier == restresource.Eq {
			if len(filter.Values) == 1 && strings.TrimSpace(filter.Values[0]) != "" {
				return filter.Values[0], true
			}
			break
		}
	}

	return "", false
}

func GenLikeConditionsFromFilters(filters []restresource.Filter, filterNames ...string) map[string]interface{} {
	if len(filters) == 0 {
		return nil
	}
	conditions := make(map[string]interface{})
	for _, filterName := range filterNames {
		if value, ok := GetFilterValueWithLikeModifierFromFilters(filterName, filters); ok {
			conditions[filterName] = value
		}
	}
	return conditions
}

func GetFilterValueWithLikeModifierFromFilters(filterName string, filters []restresource.Filter) (string, bool) {
	for _, filter := range filters {
		if filter.Name == filterName && filter.Modifier == restresource.Like {
			if len(filter.Values) == 1 && strings.TrimSpace(filter.Values[0]) != "" {
				return filter.Values[0], true
			}
			break
		}
	}
	return "", false
}

func GetFilterValueWithEqModifierFromFilter(filter restresource.Filter) (string, bool) {
	if filter.Modifier == restresource.Eq && len(filter.Values) == 1 &&
		strings.TrimSpace(filter.Values[0]) != "" {
		return filter.Values[0], true
	}

	return "", false
}
