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
	NoResponse          = "No Response"
	BadRequest          = "Bad Request"
	NotAllowed          = "Method Not Allowed"
)

func SetStatuscode(code int) ResponseLine {

	responseLine := ResponseLine{
		Version: "HTTP/1.1",
	}

	switch code {
	case 200:
		responseLine.Reason = Ok
		responseLine.Status = code

	case 201:
		responseLine.Reason = Created
		responseLine.Status = code

	case 404:
		responseLine.Reason = NotFound
		responseLine.Status = code
	case 400:
		responseLine.Reason = BadRequest
		responseLine.Status = code

	case 500:
		responseLine.Reason = InternalServerError
		responseLine.Status = code
	case 405:
		responseLine.Reason = NotAllowed
		responseLine.Status = code
	default:
		responseLine.Reason = NoResponse
		responseLine.Status = code
	}

	return responseLine
}
