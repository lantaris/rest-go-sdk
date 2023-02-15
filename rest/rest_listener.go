package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/lantaris/rest-go-sdk/logger"
)

type RestServer struct {
	RestRouter     *mux.Router
	RestServerConf TRestServerConf
}

// *******************************************************************************
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// *******************************************************************************
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

// *******************************************************************************
func (RS RestServer) Start() error {
	var (
		err error = nil
	)
	logger.Debugln("Starting REST server: " + RS.RestServerConf.Name)
	RS.RestRouter = mux.NewRouter().StrictSlash(true)
	for _, service := range RS.RestServerConf.Endpoints {
		logger.Debugln("REST add endpoint '" + service.Name + "':[" + service.Endpoint + ":" + service.Type + "]")
		RS.RestRouter.HandleFunc(service.Endpoint, service.Callback).Methods(service.Type).Name(service.Name)
	}

	logger.Infoln("Starting rest server listen on:", RS.RestServerConf.Listen)
	RS.RestRouter.NotFoundHandler = http.HandlerFunc(notFound)
	err = http.ListenAndServe(RS.RestServerConf.Listen, RS.RestRouter)
	if err != nil {
		logger.Errorln("Error starting REST server: [" + RS.RestServerConf.Name + "]:" + err.Error())
	}
	return err
}
