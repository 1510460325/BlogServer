package util

import (
	"strconv"
	"strings"
)

func SplitToIds(strIds string) []uint32{
	strList := strings.Split(strIds, ",")
	ids := make([]uint32, 0, len(strList))
	for _, str := range strList {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, uint32(id))
		}
	}
	return ids
}
