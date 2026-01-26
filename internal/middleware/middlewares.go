package middleware

type Middlewares struct {
	Global *globalMiddlewares
}

func NewMiddlewares(origins []string) *Middlewares {
	return &Middlewares{
		Global: newGlobalMiddlewares(origins),
	}
}
