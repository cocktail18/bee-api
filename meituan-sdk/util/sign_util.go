package util

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

// 计算请求的签名，具体规则见 https://developer.meituan.com/docs/biz/comm-dev-isv-sign-rule
func CalculateSign(param map[string]string, signKey string) (string, error) {
	sortedString := getSortedString(param, signKey)
	return sha1Encode(sortedString)
}

func getSortedString(param map[string]string, signKey string) string {
	keys := make([]string, 0, len(param))
	for key, value := range param {
		//剔除value为空的key
		if value != "" {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString(signKey)

	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(param[k])
	}

	return sb.String()
}

// string计算sha1
func sha1Encode(data string) (string, error) {
	t := sha1.New()
	_, err := io.WriteString(t, data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}
