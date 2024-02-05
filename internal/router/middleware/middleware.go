package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type middleware struct {
	l *logrus.Logger
}

// Middleware はミドルウェアを表すインターフェースです。
type Middleware interface {
	Recover(next http.Handler) http.Handler
	Counter(next http.Handler) http.Handler
}

// New はミドルウェアを生成します。
func New(l *logrus.Logger) Middleware {
	return &middleware{
		l: l,
	}
}

// BuildChain はミドルウェアを連結します。
func BuildChain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	if len(m) == 0 {
		return h
	}
	return m[0](BuildChain(h, m[1:cap(m)]...))
}
