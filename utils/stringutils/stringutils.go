package stringutils

import (
	"errors"
	"strings"
)

func IsEmpty(s string) bool {
	return s == "" || len(s) == 0 || strings.TrimSpace(s) == ""
}

func BuildErrorFromMessages(msgs []string) error {
	return errors.New(strings.Join(msgs, ", "))
}
