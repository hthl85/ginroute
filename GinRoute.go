package ginroute

import "github.com/gin-gonic/gin"

// AppRoute struct
type AppRoute struct {
	Group         string
	Protected     bool
	Routes        []Route
	GroupHandlers []gin.HandlerFunc
}

// Route struct
type Route struct {
	Method        string
	Pattern       string
	RouteHandlers []gin.HandlerFunc
}

// NewGinRoute initializes gin route
func NewGinRoute(engine *gin.Engine, appRoute []AppRoute) {

	// Iterate over given app route and attach them to the gin engine
	for _, routes := range appRoute {
		groupRoute := engine.Group(routes.Group, routes.GroupHandlers...)

		for _, route := range routes.Routes {
			switch route.Method {
			case "GET":
				groupRoute.GET(route.Pattern, route.RouteHandlers...)
			case "POST":
				groupRoute.POST(route.Pattern, route.RouteHandlers...)
			case "PUT":
				groupRoute.PUT(route.Pattern, route.RouteHandlers...)
			case "PATCH":
				groupRoute.PATCH(route.Pattern, route.RouteHandlers...)
			case "HEAD":
				groupRoute.HEAD(route.Pattern, route.RouteHandlers...)
			case "OPTIONS":
				groupRoute.OPTIONS(route.Pattern, route.RouteHandlers...)
			case "DELETE":
				groupRoute.DELETE(route.Pattern, route.RouteHandlers...)
			}
		}
	}
}
