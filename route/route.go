package route

import (
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/controllers"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func RouteApp(appContext *controllers.AppContext) (r *router) {
	appC := appContext
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	router := NewRouter()
	router.NotFound = http.HandlerFunc(controllers.MyNotFound)
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	// router.Get("/admin", commonHandlers.Append(appC.AuthHandler).ThenFunc(appC.AdminHandler))
	router.Get("/dashboard", commonHandlers.ThenFunc(appC.Dashboard))
	//User resources
	router.Get("/users", commonHandlers.ThenFunc(appC.UserIndex))
	router.POST("/users", appC.UserStore)
	router.POST("/users/:id", appC.UserUpdate)
	router.Get("/new/user", commonHandlers.ThenFunc(controllers.UserCreate))
	router.Get("/users/:id/edit", commonHandlers.ThenFunc(appC.UserEdit))

	//Mountain resources
	router.Get("/mountains", commonHandlers.ThenFunc(appC.MountainIndex))
	router.POST("/mountains", appC.MountainStore)
	router.POST("/tracks", appC.TrackStore)
	router.POST("/mountains/:id", appC.MountainUpdate)
	router.Get("/new/mountain", commonHandlers.ThenFunc(controllers.MountainCreate))
	router.Get("/mountains/:id/edit", commonHandlers.ThenFunc(appC.MountainEdit))
	router.Get("/mountains/:id/tracks", commonHandlers.ThenFunc(appC.TrackIndex))
	router.Get("/mountains/:id/tracks/create", commonHandlers.ThenFunc(appC.TrackCreate))
	router.Get("/", commonHandlers.ThenFunc(appC.IndexHandler))
	router.Get("/teas/:query", commonHandlers.ThenFunc(appC.TeaHandler))
	return router
}
