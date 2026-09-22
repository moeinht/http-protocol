package setGenerateHeader

import (
	"errors"
	"github.com/moeinht/http-protocol/request/header"
	"strings"
)

func Generate(name string, value string) (header.Header, error) {
	newHeader := header.Header{}

	name = strings.Trim(name, " ")
	value = strings.Trim(value, " ")

	if name == "" || value == "" {
		return newHeader, errors.New("The Value And Name Must Be Assinged")
	}

	newHeader.Name = name
	newHeader.Value = value

	return newHeader, nil
}
