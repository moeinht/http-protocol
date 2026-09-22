package response

import (
	"github.com/moeinht/http-protocol/request/header"
	"github.com/moeinht/http-protocol/response/setGenerateHeader"
	statusline "github.com/moeinht/http-protocol/response/status-line"
	"strconv"
	"strings"
)

type Response struct {
	StatusLine statusline.ResponseLine
	Headers    []header.Header
	Body       []byte
}

func (rs *Response) GetHeader(name string) (string, bool) {
	for _, h := range rs.Headers {
		if strings.EqualFold(h.Name, name) {
			return h.Value, true
		}
	}

	return "", false
}

func (rs *Response) SetHeader(name string, value string) error {
	generatedHeader, err := setGenerateHeader.Generate(name, value)
	if err != nil {
		return err
	}

	rs.Headers = append(rs.Headers, generatedHeader)

	return nil
}

func (rs *Response) SetBody(body []byte) error {
	rs.Body = body

	contentLength := strconv.Itoa(len(body))

	return rs.SetHeader("Content-Length", contentLength)
}

func (rs *Response) Serialize() ([]byte, error) {
	responseByte := []byte{}
	emptyLine := "\r\n"
	statusLineResponse := rs.StatusLine.Version + " " + strconv.FormatInt(int64(rs.StatusLine.Status), 10) + " " + rs.StatusLine.Reason + emptyLine
	headersResponse := ""
	for _, value := range rs.Headers {
		headerValue := value.Name + ": " + value.Value
		headersResponse += headerValue + emptyLine
	}
	headersResponse += emptyLine

	responseByte = append(responseByte, []byte(statusLineResponse)...)
	responseByte = append(responseByte, []byte(headersResponse)...)
	responseByte = append(responseByte, rs.Body...)

	return responseByte, nil
}

func NewResponse(code int) *Response {
	return &Response{
		StatusLine: statusline.SetStatuscode(code),
	}
}
