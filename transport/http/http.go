package http

import (
	"net/http"
	"piglatin-translator-api/transport/http/controller"
	"piglatin-translator-api/usecase"
)

func NewServer(piglatinService *usecase.Piglatin) *http.Server {
	mux := http.NewServeMux()
	wrapper := controller.NewContainer(piglatinService)
	setupRoutes(mux, wrapper)
	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}
	return server
}

func setupRoutes(mux *http.ServeMux, wrapper *controller.Wrapper) {
	mux.HandleFunc("/translate", EnsureHandlerMethod(http.MethodPost, wrapper.HandleTranslate))
	mux.HandleFunc("/", EnsureHandlerMethod(http.MethodGet, wrapper.HandleGetTranslationRequest))
}
