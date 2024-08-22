package kit

import (
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

// GetDBPropertyChildIds 倒序排列
func GetDBPropertyChildIds(str string) string {
	if str == "" {
		return ""
	}
	propPair := make([]string, 0)
	strArr := strings.Split(str, ",")
	for _, s := range strArr {
		if s == "" {
			continue
		}
		propPair = append(propPair, s)
	}
	sort.Slice(propPair, func(i, j int) bool {
		return propPair[i] < propPair[j]
	})
	return strings.Join(propPair, ",") + ","
}

func GetPropIdsByStr(str string) []int64 {
	propIds := make([]int64, 0)
	strArr := strings.Split(str, ",")
	for _, s := range strArr {
		if s == "" {
			continue
		}
		sArr := strings.Split(s, ":")
		if len(sArr) != 2 {
			continue
		}
		propIds = append(propIds, cast.ToInt64(sArr[0]), cast.ToInt64(sArr[1]))
	}
	return lo.Uniq(propIds)
}
