package repo

import (
	"database/sql"
	"piglatin-translator-api/model"
)

type PiglatinRepo struct {
	db *sql.DB
}

func NewPiglatinRepo(db *sql.DB) *PiglatinRepo {
	return &PiglatinRepo{
		db: db,
	}
}

func (p *PiglatinRepo) SaveTranslationRequest(input string, translated string) error {
	db := p.db
	_, err := db.Exec("INSERT INTO piglatin_requests (request,translation) VALUES ($1,$2);", input, translated)
	return err
}

func (p *PiglatinRepo) GetTranslationRequests(page int) (*model.TranslationRequestsSummary, error) {
	db := p.db
	limit := 20
	rows, err := db.Query(`SELECT id, request, translation FROM piglatin_requests ORDER BY id desc LIMIT $1 OFFSET $2;`, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translationRequests []model.TranslationRequest
	for rows.Next() {
		var row model.TranslationRequest
		if err := rows.Scan(&ignoreColumn{}, &row.Input, &row.Translation); err != nil {
			return nil, err
		}
		translationRequests = append(translationRequests, row)
	}

	var totalPages int
	err = db.QueryRow(`SELECT ceiling(COUNT(*) * 1.0 / $1) total_pages FROM piglatin_requests`, limit).Scan(&totalPages)
	if err != nil {
		return nil, err
	}

	return &model.TranslationRequestsSummary{
		Page:            page,
		TotalPages:      totalPages,
		TranslationList: translationRequests,
	}, nil
}

func (p *PiglatinRepo) FindOneTranslationRequest(input string) (*model.TranslationRequest, error) {
	db := p.db
	var translationRequest *model.TranslationRequest = &model.TranslationRequest{}
	err := db.QueryRow("SELECT request, translation FROM piglatin_requests WHERE request = $1", input).Scan(&translationRequest.Input, &translationRequest.Translation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return translationRequest, nil
}
