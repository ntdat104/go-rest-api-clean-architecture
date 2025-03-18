package utils

import (
	"fmt"
	"strings"

	"github.com/lithammer/shortuuid/v4"
)

func GenerateKey(prefix string, values ...interface{}) string {
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%s:%s", prefix, strings.Join(parts, ":"))
}

func GenerateUUID() string {
	return shortuuid.New()
}
