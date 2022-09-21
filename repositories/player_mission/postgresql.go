package player_mission

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

func (repo *postgreRepository) Create(ctx context.Context, m PlayerMissionModel) (lastID int, err error) {
	sql := ` INSERT INTO player_missions (player_id, mission_id, status, created_by, updated_by)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = repo.db.QueryRow(sql, m.PlayerId, m.MissionId, m.Status, m.CreatedBy, m.UpdatedBy).Scan(&lastID)

	return
}

func (repo *postgreRepository) UpdateStatus(ctx context.Context, id int, status string) (err error) {
	sql := `UPDATE player_missions SET status = $1 WHERE id =  $2`

	tx := utils.GetCtxDBTx(ctx)
	if tx != nil {
		utils.LogInfo(ctx, utils.FnTrace(), "there is db transactions")
		_, err = tx.ExecContext(ctx, sql, status, id)
	} else {
		_, err = repo.db.ExecContext(ctx, sql, status, id)
	}

	return
}

func (repo *postgreRepository) Delete(ctx context.Context, id int) (err error) {
	sql := `DELETE player_missions WHERE id =  $1`
	_, err = repo.db.ExecContext(ctx, sql, id)

	return
}

func (repo *postgreRepository) FindAllByPlayerID(ctx context.Context, playerId int) (res []PlayerMissionModel, err error) {
	sql := ` SELECT pm.id, pm.player_id, pm.mission_id, pm.status, pm.created_at, pm.created_by, pm.updated_at, pm.updated_by,
	m.title, m.description, m.gold_bounty
	FROM player_missions pm
	JOIN missions m ON m.id = pm.mission_id
	WHERE player_id =  $1`
	rows, err := repo.db.QueryContext(ctx, sql, playerId)
	if err != nil {
		return
	}
	defer rows.Close()
	var temp PlayerMissionModel
	for rows.Next() {
		err = rows.Scan(&temp.Id, &temp.PlayerId, &temp.MissionId, &temp.Status, &temp.CreatedAt, &temp.CreatedBy, &temp.UpdatedAt, &temp.UpdatedBy,
			&temp.MissionTitle, &temp.MissionDescription, &temp.MissionGoldBounty)
		if err != nil {
			return
		}
		res = append(res, temp)
	}

	return
}

func (repo *postgreRepository) FindOneByPlayerIDAndMissionID(ctx context.Context, playerId, missionId int) (res PlayerMissionModel, err error) {
	sql := ` SELECT pm.id, pm.player_id, pm.mission_id, pm.status, pm.created_at, pm.created_by, pm.updated_at, pm.updated_by,
	m.title, m.description, m.gold_bounty
	FROM player_missions pm
	JOIN missions m ON m.id = pm.mission_id
	WHERE pm.player_id =  $1 AND m.id = $2 LIMIT 1`

	err = repo.db.QueryRowContext(ctx, sql, playerId, missionId).Scan(&res.Id, &res.PlayerId, &res.MissionId, &res.Status, &res.CreatedAt, &res.CreatedBy, &res.UpdatedAt, &res.UpdatedBy,
		&res.MissionTitle, &res.MissionDescription, &res.MissionGoldBounty)

	return
}
