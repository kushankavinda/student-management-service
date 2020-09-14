package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/student-management-service/servises"
	"github.com/student-management-service/util"
	viewModel "github.com/student-management-service/viewModels"
	//	"github.com/student-management-service/servises/adaptor"
)

type userControllers struct {
	userIDPattern *regexp.Regexp
}

func (uc userControllers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("ServeHTTP implementation")
	if r.URL.Path == "/api/user/gen-key" {
		switch r.Method {
		case http.MethodPost:
			uc.onboardToDigitalSignature(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userControllers) onboardToDigitalSignature(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequestDigitalSignature(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(util.INVALID_REQUEST))
		return
	}
	log.Println("/api/user/biometric API Receive a Req ", u)
	isSuccess, responseMessage := servises.OnboardToDigitalSignature(u)
	if isSuccess {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseMessage))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(responseMessage))
	}
}

func (uc *userControllers) parseRequestDigitalSignature(r *http.Request) (viewModel.DigitalSignatureOnboard, error) {
	dec := json.NewDecoder(r.Body)
	var u viewModel.DigitalSignatureOnboard
	err := dec.Decode(&u)
	if err != nil {
		return viewModel.DigitalSignatureOnboard{}, err
	}
	return u, nil
}

func implementController() *userControllers {
	log.Println("controller implementation")
	return &userControllers{}
}
