package middleware

import (
	"github.com/moeinht/http-protocol/handler"
)

type Middleware func(handler.IHandler) handler.IHandler
