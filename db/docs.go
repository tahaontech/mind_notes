package db

import "github.com/tahaontech/mind_notes/types"

type Document struct {
	ID     string
	NodeID string
	Data   string
}

func (d *DB) DocumentGetOne(nodeID string) (*types.DocumentResp, error) {
	rows, err := d.Database.Query("SELECT id, nodeId, data FROM document WHERE nodeId = ?;", nodeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		detail := types.DocumentResp{}
		err := rows.Scan(&detail.ID, &detail.NodeID, &detail.Data)
		if err != nil {
			return nil, err
		} else {
			return &detail, nil
		}
	}

	err = rows.Err()
	return nil, err
}

func (d *DB) DocumentAdd(req *types.CreateDocumentReq) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO document(id, nodeId, data) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.ID, req.NodeID, req.Data)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) DocumentUpdate(req *types.UpdateDocumentReq) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE document SET data = ? WHERE nodeId = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Data, req.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) DocumentDelete(nodeID string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM document WHERE nodeId = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(nodeID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
