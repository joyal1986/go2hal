package chef

import (
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/weAutomateEverything/go2hal/gokit"
	"github.com/weAutomateEverything/go2hal/machineLearning"
	"net/http"
	"encoding/json"
	"context"
)

//MakeHandler returns a restful http handler for the chef delivery service
//the Machine Learning service can be set to nil if you do not wish to log the http requests
func MakeHandler(service Service, logger kitlog.Logger, ml machineLearning.Service) http.Handler {
	opts := gokit.GetServerOpts(logger, ml)
	chefDeliveryEndpoint := kithttp.NewServer(makeChefDeliveryAlertEndpoint(service), gokit.DecodeString, gokit.EncodeResponse, opts...)
	findNodesEndpoint := kithttp.NewServer(makeFindNodesHandler(service), decodeFindNodesRequest, gokit.EncodeResponse, opts...)
	keyboardRecipeHandler := kithttp.NewServer(makeSendKeyboardRecipeHandler(service), decodeSendKeyboardRequest, gokit.EncodeResponse, opts...)
	keyboardEnvironmentHandler := kithttp.NewServer(makeSendKeyboardEnvironmentHandler(service), decodeSendKeyboardRequest, gokit.EncodeResponse, opts...)
	keyboardNodeHandler := kithttp.NewServer(makeSendKeyboardNodesHandler(service), decodeSendKeyboardNodeRequest, gokit.EncodeResponse, opts...)
	r := mux.NewRouter()

	r.Handle("/delivery", chefDeliveryEndpoint).Methods("POST")
	r.Handle("/chef/nodes", findNodesEndpoint).Methods("POST")
	r.Handle("/chef/keyboard/recipes", keyboardRecipeHandler).Methods("POST")
	r.Handle("/chef/keyboard/environments", keyboardEnvironmentHandler).Methods("POST")
	r.Handle("/chef/keyboard/nodes", keyboardNodeHandler).Methods("POST")

	return r

}
func decodeFindNodesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	v := FindNodesRequest{}
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}
func decodeSendKeyboardRequest(_ context.Context, r *http.Request) (interface{}, error) {
	v := SendKeyboardRequest{}
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}
func decodeSendKeyboardNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	v := SendKeyboardNodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&v)
	return v, err
}
