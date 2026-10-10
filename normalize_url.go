package main

import (
	"fmt"
	u "net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	parsedURL, err := u.Parse(s)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")
	return fullPath, nil
}
