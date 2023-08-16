package db

import (
	"fmt"

	"github.com/tahaontech/mind_notes/types"
)

type Edge struct {
	ID       string
	SourceID string
	TargetID string
	RootID   string
}

func (d *DB) EdgeGetMany(rootID string) ([]*types.EdgeResp, error) {
	if rootID == "" {
		return nil, fmt.Errorf("rootId is empty")
	}

	rows, err := d.Database.Query("SELECT id, sourceId, targetId FROM edge WHERE rootId = ? ;", rootID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := []*types.EdgeResp{}
	for rows.Next() {
		detail := types.EdgeResp{}
		err := rows.Scan(&detail.ID, &detail.SourceID, &detail.TargetID)
		if err != nil {
			return nil, err
		}
		response = append(response, &detail)
	}

	err = rows.Err()
	return response, err
}

func (d *DB) EdgeAdd(req *types.CreateEdgeReq) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO edge(id, sourceId, targetId, rootId) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.ID, req.SourceID, req.TargetID, req.RootID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) EdgeDeleteByNode(targetID string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM edge WHERE targetId = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(targetID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
