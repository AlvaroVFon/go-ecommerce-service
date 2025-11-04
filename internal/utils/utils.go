// Package utils provides utility functions for common tasks.
package utils

import (
	"strconv"
)

func ParsePaginationParams(pageStr, limitStr string, defaultLimit, maxLimit int) (page int, limit int) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, limitErr := strconv.Atoi(limitStr)
	if limitErr != nil || limit < 1 {
		limit = defaultLimit
	} else if limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}
