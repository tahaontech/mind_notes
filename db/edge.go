package db

type Edge struct {
	ID       string
	SourceID string
	TargetID string
}

func (d *DB) AddEdge() error {
	return nil
}
