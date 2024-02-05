package router

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/murasame29/home/internal/components/top"
	"github.com/murasame29/home/internal/router/middleware"
	"github.com/sirupsen/logrus"
)

type router struct {
	l *logrus.Logger

	mux        *http.ServeMux
	middleware middleware.Middleware
}

// New はルーターを生成します。
func New(l *logrus.Logger) http.Handler {
	router := &router{
		l:          l,
		mux:        http.NewServeMux(),
		middleware: middleware.New(l),
	}

	router.index()

	return router.mux
}

func (r *router) index() {
	handler := top.NewHandler()
	r.mux.Handle("/", middleware.BuildChain(handler, r.middleware.Recover, r.middleware.AccessLog, r.middleware.Counter))
	app.Route("/", top.NewComponent())
}
