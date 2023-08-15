package db

import "github.com/tahaontech/mind_notes/types"

type Document struct {
	ID     string
	NodeID string
	Data   string
}

func (d *DB) DocumentGetOne(nodeID string) (*types.DocumentResp, error) {
	stm, err := d.Database.Prepare("select id, nodeId, data FROM document WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	resp := types.DocumentResp{}

	err = stm.QueryRow(nodeID).Scan(&resp.ID, &resp.NodeID, &resp.Data)

	return &resp, err
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

func (d *DB) DocumentDelete(nodeID string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE document WHERE nodeId = ?;")
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
