package growel

import (
	"fmt"
	"strings"
)

func ParsePath(path string) ([]string, map[string]string) {
	var query_str string
	path, query_str = extractQueryStr(path)

	parts := splitPath(path)

	if len(parts) == 0 {
		return parts, map[string]string{}
	}

	query := parseQueryStr(query_str)

	return parts, query
}

func splitPath(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

func extractQueryStr(str string) (path string, querystr string) {
	str_slice := strings.Split(str, "?")
	leng := len(str_slice)
	if leng > 2 {
		fmt.Println("Error parsing path: multiple '?' found.")
		// return anyway
	} else if leng == 1 {
		// no query string found
		return str_slice[0], ""
	}
	return str_slice[0], str_slice[1]
}

func parseQueryStr(str string) map[string]string {
	query := make(map[string]string)
	if str == "" {
		return query
	}
	pairs := strings.Split(str, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			fmt.Println("Error parsing query string: invalid key-value pair.")
			continue
		}
		query[kv[0]] = kv[1]
	}
	return query
}
