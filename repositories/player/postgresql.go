package player

import (
	"context"
	"database/sql"
	"idnmedia/utils"
)

type postgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) Repository {
	return &postgreRepository{
		db: db,
	}
}

func (repo *postgreRepository) FindOneByEmail(ctx context.Context, email string) (row PlayerModel, err error) {

	sql := `SELECT id, name, email, password, gold_amount, created_at, created_by, updated_at, updated_by
	FROM players WHERE email = $1`

	err = repo.db.QueryRowContext(ctx, sql, email).Scan(&row.Id, &row.Name, &row.Email, &row.Password, &row.GoldAmount,
		&row.CreatedAt, &row.CreatedBy, &row.UpdatedAt, &row.UpdatedBy)

	return
}

func (repo *postgreRepository) AddGoldAmountByPlayerId(ctx context.Context, id int, goldAmount float64) (currentGold float64, err error) {

	sql := `UPDATE players set gold_amount = gold_amount + $1 WHERE id = $2 RETURNING gold_amount`

	tx := utils.GetCtxDBTx(ctx)
	if tx != nil {
		utils.LogInfo(ctx, utils.FnTrace(), "there is db transactions")
		err = tx.QueryRowContext(ctx, sql, goldAmount, id).Scan(&currentGold)
	} else {
		err = repo.db.QueryRowContext(ctx, sql, goldAmount, id).Scan(&currentGold)
	}

	return
}
