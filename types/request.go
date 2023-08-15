package types

// Node Requests

type CreateNodeReq struct {
	ID        string  `json:"id"`
	Label     string  `json:"label"`
	Root      bool    `json:"root"`
	PositionX float32 `json:"positionX"`
	PositionY float32 `json:"positionY"`
	RootID    string  `json:"rootId"`
}

type UpdateNodeLabelReq struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type UpdateNodePosReq struct {
	ID        string  `json:"id"`
	PositionX float32 `json:"positionX"`
	PositionY float32 `json:"positionY"`
}

// Edge Requests

type CreateEdgeReq struct {
	ID       string `json:"id"`
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
	RootID   string `json:"rootId"`
}

// Documents Requests

type CreateDocumentReq struct {
	ID     string `json:"id"`
	NodeID string `json:"nodeId"`
	Data   string `json:"data"`
}

type UpdateDocumentReq struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}
