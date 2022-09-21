package mission

import (
	"context"
	"database/sql"
)

type postgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(db *sql.DB) Repository {
	return &postgreRepository{
		db: db,
	}
}

func (repo *postgreRepository) Create(ctx context.Context, m MissionModel) (lastID int, err error) {
	sql := ` INSERT INTO missions (title, description, gold_bounty, deadline_second,  created_by, updated_by)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = repo.db.QueryRow(sql, m.Title, m.Description, m.GoldBounty, m.DeadlineSecond, m.CreatedBy, m.UpdatedBy).Scan(&lastID)

	return
}

func (repo *postgreRepository) FindAllPagination(ctx context.Context, offset, limit int) (res []MissionModel, count int, err error) {
	sql := ` SELECT id, title, description, gold_bounty, deadline_second,  created_at, created_by, updated_at, updated_by 
	FROM missions ORDER BY id DESC
	OFFSET $1
	LIMIT $2`
	rows, err := repo.db.QueryContext(ctx, sql, offset, limit)
	if err != nil {
		return
	}
	defer rows.Close()

	var temp MissionModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.GoldBounty, &temp.DeadlineSecond, &temp.CreatedAt, &temp.CreatedBy, &temp.UpdatedAt, &temp.UpdatedBy)
		if err != nil {
			return
		}
		res = append(res, temp)
	}

	sql = `SELECT COUNT(id) FROM missions`
	err = repo.db.QueryRow(sql).Scan(&count)

	return
}

func (repo *postgreRepository) FindOneByID(ctx context.Context, id int) (res MissionModel, err error) {
	sql := ` SELECT id, title, description, deadline_second, created_at, created_by, updated_at, updated_by 
	FROM missions WHERE id =  $1`
	row := repo.db.QueryRowContext(ctx, sql, id)

	err = row.Scan(&res.Id, &res.Title, &res.Description, &res.DeadlineSecond, &res.CreatedAt, &res.CreatedBy, &res.UpdatedAt, &res.UpdatedBy)
	if err != nil {
		return
	}

	return
}
