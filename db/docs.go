package db

type Document struct {
	ID     string
	NodeId string
	Data   string
}

func (d *DB) AddDocument() error {
	return nil
}
