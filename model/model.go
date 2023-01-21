package model

type TranslationRequestDTO struct {
	Input string `json:"input"`
}

type TranslationRequest struct {
	Input       string `json:"input"`
	Translation string `json:"translation"`
}

type TranslationRequestsSummary struct {
	Page            int                  `json:"page"`
	TotalPages      int                  `json:"totalPages"`
	TranslationList []TranslationRequest `json:"translations"`
}
