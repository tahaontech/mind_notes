package db

import (
	"fmt"

	"github.com/tahaontech/mind_notes/types"
)

type Node struct {
	ID        string
	Label     string
	Root      bool
	PositionX float32
	PositionY float32
	RootID    string
}

func (d *DB) NodeGetRoots() ([]*types.RootNodesResp, error) {
	// return all root nodes
	rows, err := d.Database.Query("SELECT id, label FROM node WHERE root = true;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := []*types.RootNodesResp{}
	for rows.Next() {
		detail := types.RootNodesResp{}
		err := rows.Scan(&detail.ID, &detail.Label)
		if err != nil {
			return nil, err
		}
		response = append(response, &detail)
	}

	err = rows.Err()
	return response, err
}

func (d *DB) NodeGetMany(rootID string) ([]*types.NodeResp, error) {
	if rootID == "" {
		return nil, fmt.Errorf("root Id is empty")
	}
	// return all nodes with this rootID
	rows, err := d.Database.Query("SELECT id, label, root, positionX, positionY FROM node WHERE rootId = ? OR id = ?", rootID, rootID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := []*types.NodeResp{}
	for rows.Next() {
		detail := types.NodeResp{}
		err := rows.Scan(&detail.ID, &detail.Label, &detail.Root, &detail.PositionX, &detail.PositionY)
		if err != nil {
			return nil, err
		}
		response = append(response, &detail)
	}

	err = rows.Err()
	return response, err
}

func (d *DB) NodeGetOne(nodeID string) (*types.NodeResp, error) {
	stm, err := d.Database.Prepare("select id, label, root, positionX, positionY  FROM node WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	resp := types.NodeResp{}

	err = stm.QueryRow(nodeID).Scan(&resp.ID, &resp.Label, &resp.Root, &resp.PositionX, &resp.PositionY)

	return &resp, err
}

func (d *DB) NodeAdd(req *types.AddNodeReq) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO node(id, label, root, positionX, positionY, rootId) VALUES (?, ?, ?, ?, ?, ?);`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.ID, req.Label, req.Root, req.PositionX, req.PositionY, req.RootID)
	if err != nil {
		return err
	}

	query2 := `INSERT INTO edge(id, sourceId, targetId, rootId) VALUES (?, ?, ?, ?);`
	stmt2, err := tx.Prepare(query2)
	if err != nil {
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(req.EdgeID, req.SourceID, req.ID, req.RootID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) NodeAddRoot(req *types.CreateNoderootReq) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO node(id, label, root, positionX, positionY, rootId) VALUES (?, ?, true, 0.0, 0.0, '');`

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.ID, req.Label)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) NodeLabelUpdate(nodeID string, label string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE node SET label=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(label, nodeID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) NodePositionUpdate(nodeID string, xpos, ypos float32) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE node SET positionX = ?, positionY = ? WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(xpos, ypos, nodeID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (d *DB) NodeDelete(nodeID string) error {
	tx, err := d.Database.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM node WHERE id = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(nodeID)
	if err != nil {
		return err
	}
	// n, err := res.RowsAffected()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("The statement has affected %d rows\n", n)

	err = tx.Commit()
	return err
}
