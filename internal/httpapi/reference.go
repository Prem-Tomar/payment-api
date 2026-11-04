package httpapi

import "strings"

func normalizeReference(value string, required bool) (string, bool) {
	reference := strings.TrimSpace(value)

	if required && reference == "" {
		return "", false
	}

	return reference, true
}
