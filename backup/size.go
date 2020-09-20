package backup

import (
	"errors"
	"strconv"
	"strings"
)

func ParseSize(sizeStr string) (int64, error) {
	sizeStr = strings.ToLower(sizeStr)
	unitIdx := len(sizeStr) - 1
	if len(sizeStr) == 0 {
		return 0.0, errors.New("invalid size")
	}
	if sizeStr[unitIdx] == 'b' {
		unitIdx -= 1
	}

	var size float64 = 0.0
	var err error
	if val, ok := sizeMap[sizeStr[unitIdx]]; ok {
		size, err = strconv.ParseFloat(sizeStr[0:unitIdx], 64)
		size = size * val
	} else {
		size, err = strconv.ParseFloat(sizeStr[0:unitIdx+1], 64)
	}
	if err != nil {
		return 0, err
	}

	if size < 0 {
		return 0, errors.New("size may not be negative")
	}

	return int64(size), nil
}
