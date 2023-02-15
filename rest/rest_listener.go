package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/lantaris/rest-go-sdk/fmtlog"
)

type RestServer struct {
	RestRouter     *mux.Router
	RestServerConf TRestServerConf
}

// *******************************************************************************
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println("ERROR REQUEST URI:" + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// *******************************************************************************
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ERROR REQUEST URI:" + r.RequestURI)
	w.WriteHeader(http.StatusNotFound)
}

// *******************************************************************************
func (RS RestServer) Start() error {
	var (
		err error = nil
	)
	fmtlog.Debugln("Starting REST server: " + RS.RestServerConf.Name)
	RS.RestRouter = mux.NewRouter().StrictSlash(true)
	for _, service := range RS.RestServerConf.Endpoints {
		fmtlog.Debugln("REST add endpoint '" + service.Name + "':[" + service.Endpoint + ":" + service.Type + "]")
		RS.RestRouter.HandleFunc(service.Endpoint, service.Callback).Methods(service.Type).Name(service.Name)
	}

	fmtlog.Infoln("Starting rest server listen on:", RS.RestServerConf.Listen)
	RS.RestRouter.NotFoundHandler = http.HandlerFunc(notFound)
	err = http.ListenAndServe(RS.RestServerConf.Listen, RS.RestRouter)
	if err != nil {
		fmtlog.Errorln("Error starting REST server: [" + RS.RestServerConf.Name + "]:" + err.Error())
	}
	return err
}
