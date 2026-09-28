package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) TimeoutMiddleware(max_duration time.Duration, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		ctx, cancel := context.WithTimeout(r.Context(), max_duration)
		defer cancel()

		r = r.WithContext(ctx)

		next(w, r, p)
	}
}
