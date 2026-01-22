package approuter

import (
	"unified_platform/internal/handler"

	"github.com/go-chi/chi/v5"
)

type AppRouter struct {
	R chi.Router
}

func (a AppRouter) Get(path string, h handler.AppHandler) {
	a.R.Get(path, handler.Handle(h))
}

func (a AppRouter) Post(path string, h handler.AppHandler) {
	a.R.Post(path, handler.Handle(h))
}

func (a AppRouter) Put(path string, h handler.AppHandler) {
	a.R.Put(path, handler.Handle(h))
}

func (a AppRouter) Patch(path string, h handler.AppHandler) {
	a.R.Patch(path, handler.Handle(h))
}

func (a AppRouter) Delete(path string, h handler.AppHandler) {
	a.R.Delete(path, handler.Handle(h))
}

func (a AppRouter) Route(path string, fn func(AppRouter)) {
	a.R.Route(path, func(r chi.Router) {
		fn(AppRouter{R: r})
	})
}
