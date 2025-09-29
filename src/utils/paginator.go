package utils

import (
	"blog-api/src/shared"
	"strconv"
	"strings"
)

func validateSortDir(input []string) shared.Direction {
	if len(input) <= 0 {
		return "DESC"
	}
	up := strings.ToUpper(input[0])
	if up == "ASC" {
		return "ASC"
	}
	return "DESC"
}

func validateIntWithDefault(input []string, defaultNum int) int {
	if len(input) <= 0 {
		return defaultNum
	}
	num, err := strconv.Atoi(input[0])
	if err != nil || num <= 0 {
		return defaultNum
	}
	return num
}

func ParsePaginationParams(values map[string][]string, sortBySetter func([]string) string) shared.PaginationParams {
	sortBy := sortBySetter(values["sortBy"])
	sortDirection := validateSortDir(values["sortDirection"])
	pageSize := validateIntWithDefault(values["pageSize"], 10)
	pageNumber := validateIntWithDefault(values["pageNumber"], 1)
	return shared.PaginationParams{sortBy, sortDirection, pageSize, pageNumber}
}
