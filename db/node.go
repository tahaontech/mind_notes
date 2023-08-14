package db

type Node struct {
	ID        string
	Label     string
	Root      bool
	PositionX float32
	PositionY float32
}

func (d *DB) AddNode() error {
	return nil
}
