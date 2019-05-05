package web

import (
	"net/http"

	"github.com/blend/go-sdk/logger"
	"github.com/blend/go-sdk/webutil"
)

// CtxOption is an option for a context.
type CtxOption func(*Ctx)

// OptCtxID sets the context request id.
func OptCtxID(id string) CtxOption {
	return func(c *Ctx) { c.ID = id }
}

// OptCtxApp sets the context app.
func OptCtxApp(a *App) CtxOption {
	return func(c *Ctx) { c.App = a }
}

// OptCtxAuth sets the context auth manager.
func OptCtxAuth(a AuthManager) CtxOption {
	return func(c *Ctx) { c.Auth = a }
}

// OptCtxViews sets the context view cache.
func OptCtxViews(v *ViewCache) CtxOption {
	return func(c *Ctx) { c.Views = v }
}

// OptCtxState sets the context state.
func OptCtxState(s State) CtxOption {
	return func(c *Ctx) { c.State = s }
}

// OptCtxRoute sets the context route.
func OptCtxRoute(r *Route) CtxOption {
	return func(c *Ctx) { c.Route = r }
}

// OptCtxRouteParams sets the context route params.
func OptCtxRouteParams(r RouteParameters) CtxOption {
	return func(c *Ctx) { c.RouteParams = r }
}

// OptCtxLog sets the context logger.
func OptCtxLog(log logger.Log) CtxOption {
	return func(c *Ctx) { c.Log = log }
}

// OptCtxTracer sets the context tracer.
func OptCtxTracer(tracer Tracer) CtxOption {
	return func(c *Ctx) { c.Tracer = tracer }
}

// OptCtxDefaultProvider sets the context default result provider.
func OptCtxDefaultProvider(rp ResultProvider) CtxOption {
	return func(c *Ctx) { c.DefaultProvider = rp }
}

// OptCtxRouteParamValue sets the context default result provider.
func OptCtxRouteParamValue(key, value string) CtxOption {
	return func(c *Ctx) {
		if c.RouteParams == nil {
			c.RouteParams = make(RouteParameters)
		}
		c.RouteParams[key] = value
	}
}

// CtxRequestOption is a ctx option that wraps a request option.
func CtxRequestOption(opt func(*http.Request) error) CtxOption {
	return func(c *Ctx) {
		opt(c.Request)
	}
}

// OptCtxQueryValue sets a query value on a context.
func OptCtxQueryValue(key, value string) CtxOption {
	return CtxRequestOption(webutil.OptQueryValue(key, value))
}

// OptCtxHeaderValue sets a header value on a context.
func OptCtxHeaderValue(key, value string) CtxOption {
	return CtxRequestOption(webutil.OptHeaderValue(key, value))
}

// OptCtxPostFormValue sets a form value on a context.
func OptCtxPostFormValue(key, value string) CtxOption {
	return CtxRequestOption(webutil.OptPostFormValue(key, value))
}

// OptCtxCookieValue sets a cookie value on a context.
func OptCtxCookieValue(key, value string) CtxOption {
	return CtxRequestOption(webutil.OptCookieValue(key, value))
}

// OptCtxBodyBytes sets a post body on a context.
func OptCtxBodyBytes(body []byte) CtxOption {
	return CtxRequestOption(webutil.OptBodyBytes(body))
}

// OptCtxPostedFiles sets posted files on a context.
func OptCtxPostedFiles(files ...webutil.PostedFile) CtxOption {
	return CtxRequestOption(webutil.OptPostedFiles(files...))
}
