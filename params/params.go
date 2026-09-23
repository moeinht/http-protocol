package params

import (
	"strings"
)

type Params map[string]string

func Match(routePath, requestPath string) (bool, Params) {
	requestParts := strings.Split(requestPath, "/")
	routeParts := strings.Split(routePath, "/")

	if len(routeParts) != len(requestParts) {
		return false, nil
	}

	params := make(Params)
	for i := 0; i < len(routeParts); i++ {
		requestPart := requestParts[i]
		routePart := routeParts[i]

		if routePart == requestPart {
			continue
		}

		if strings.HasPrefix(routePart, ":") {
			key := strings.Trim(routePart, ":")

			if key == "" {
				return false, nil
			}

			params[key] = requestPart
			continue
		}
		return false, nil
	}

	return true, params
}
