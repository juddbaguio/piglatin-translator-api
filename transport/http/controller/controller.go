package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"piglatin-translator-api/model"
	"piglatin-translator-api/transport/http/response"
	"piglatin-translator-api/usecase"
	"strconv"
)

type Wrapper struct {
	piglatinSvc *usecase.Piglatin
}

func NewContainer(piglatinSvc *usecase.Piglatin) *Wrapper {
	return &Wrapper{
		piglatinSvc: piglatinSvc,
	}
}

func (w *Wrapper) HandleTranslate(wr http.ResponseWriter, r *http.Request) {
	var payload model.TranslationRequestDTO = model.TranslationRequestDTO{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response.NewCustomErrResponse(wr, http.StatusBadRequest, err)
		return
	}

	if payload.Input == "" {
		response.NewCustomErrResponse(wr, http.StatusBadRequest, errors.New("input is required"))
		return
	}

	translationRequest, err := w.piglatinSvc.Translate(payload.Input)
	if err != nil {
		response.NewInternalServerError(wr, err)
		return
	}

	response.NewCreated(wr, nil, translationRequest)
}

func (w *Wrapper) HandleGetTranslationRequest(wr http.ResponseWriter, r *http.Request) {
	page := 1
	if r.URL.Query().Has("page") {
		parsedInt, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			response.NewCustomErrResponse(wr, http.StatusBadRequest, err)
			return
		}

		page = parsedInt
	}

	fetchResponse, err := w.piglatinSvc.GetTranslationRequests(page)
	if err != nil {
		response.NewInternalServerError(wr, err)
		return
	}

	if len(fetchResponse.TranslationList) == 0 {
		response.NewCustomErrResponse(wr, http.StatusNotFound, errors.New("no translation requests found"))
		return
	}

	response.NewSuccess(wr, fetchResponse)
}
