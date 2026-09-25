package queries

import (
	"net/url"
	"strings"
)

type Query map[string][]string

func Parse(rawURL string) (Query, error) {
	mapedQueries := make(Query)
	findIndex := strings.Index(rawURL, "?")
	if findIndex == -1 {
		return mapedQueries, nil
	}

	queryString := rawURL[findIndex+1:]

	if queryString == "" {
		return mapedQueries, nil
	}

	queryParts := strings.Split(queryString, "&")

	for _, part := range queryParts {
		equalIndex := strings.Index(part, "=")
		if equalIndex == -1 || equalIndex == 0 {
			continue
		}

		key, err := decodeQueryValue(part[:equalIndex])
		if err != nil {
			return mapedQueries, err
		}

		value, err := decodeQueryValue(part[equalIndex+1:])
		if err != nil {
			return mapedQueries, err
		}

		if value == "" {
			continue
		}

		mapedQueries[key] = append(mapedQueries[key], value)
	}

	return mapedQueries, nil
}

func decodeQueryValue(value string) (string, error) {
	return url.QueryUnescape(value)
}
