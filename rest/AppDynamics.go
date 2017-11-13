package rest

import (
	"net/http"
	"github.com/zamedic/go2hal/database"
	"io/ioutil"
	"github.com/zamedic/go2hal/service"
	"encoding/json"
	"log"
)

type appdynamicsQueueEndpoint struct {
	Name        string
	Application string
	Metricpath  string
}

type appdynamicsEndpoint struct {
	Endpoint string
}

func receiveAppDynamicsAlert(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		service.SendError(err)
		return
	}

	database.SaveAudit("APPDYNAMICS", string(body))
	database.ReceiveAppynamicsMessage()
	service.SendAppdynamicsAlert(string(body))
}

func addAppDynamicsEndpoint(w http.ResponseWriter, r *http.Request) {
	e := appdynamicsEndpoint{}
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = service.AddAppdynamicsEndpoint(e.Endpoint)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}

func addAppdynamicsQueueEndpoint(w http.ResponseWriter, r *http.Request) {
	var appDynamics appdynamicsQueueEndpoint
	err := json.NewDecoder(r.Body).Decode(&appDynamics)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = service.AddAppDynamicsQueue(appDynamics.Name, appDynamics.Application, appDynamics.Metricpath)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
