package whitelistrepository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/keremdokumaci/goql/constants"
	"github.com/keremdokumaci/goql/internal/models"
	"github.com/keremdokumaci/goql/internal/repository"
	repositoryutils "github.com/keremdokumaci/goql/internal/repository/utils"
)

type WhitelistRepository interface {
	GetWhitelistByQueryName(queryName string) (*models.Whitelist, error)
}

type whitelistRepository struct {
	repository.Repository[models.Whitelist]
	db *sqlx.DB
}

func (wr whitelistRepository) GetWhitelistByQueryName(queryName string) (*models.Whitelist, error) {
	query := fmt.Sprintf(`
		select w.* from "goql"."whitelists" w
		join "goql"."queries" q on w."query_id" = q."id"
		where q."name"='%s'
	`, queryName)

	model := &models.Whitelist{}

	err := wr.db.QueryRowx(query).StructScan(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func New(dbName constants.DB, db *sql.DB) (WhitelistRepository, error) {
	r, err := repository.NewRepository[models.Whitelist](dbName, db)
	if err != nil {
		return nil, err
	}

	return &whitelistRepository{
		Repository: r,
		db:         sqlx.NewDb(db, repositoryutils.GetDriverNameByDBName(dbName)),
	}, nil
}
