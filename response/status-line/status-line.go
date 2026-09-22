package statusline

type ResponseLine struct {
	Version string
	Status  int
	Reason  string
}

const (
	Ok                  = "OK"
	Created             = "Created"
	NotFound            = "Not Found"
	InternalServerError = "Internal Server Error"
	NoResponse          = "NoResponse"
)

func SetStatuscode(code int) ResponseLine {

	responseLine := ResponseLine{
		Version: "HTTP/1.1",
	}

	switch code {
	case 200:
		responseLine.Reason = Ok
		responseLine.Status = 200

	case 201:
		responseLine.Reason = Created
		responseLine.Status = 201

	case 404:
		responseLine.Reason = NotFound
		responseLine.Status = 404

	case 500:
		responseLine.Reason = InternalServerError
		responseLine.Status = 500
	default:
		responseLine.Reason = NoResponse
		responseLine.Status = 204
	}

	return responseLine
}
