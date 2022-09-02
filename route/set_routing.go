func (app *App) setRouters() {
    for _, route := range routes.DefinedRoutes {
        var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            log.Printf(
                "%s\t%s\t%s\t%s",
                r.Method,
                r.RequestURI,
                route.Name,
                time.Since(start),
            )
            
            route.RouteHandle(app.DB , w, r)
        }
        
        app.Router.HandleFunc(route.Pattern, handler).
            Name(route.Name).
            Methods(route.Method)
    }
}
