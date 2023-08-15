package types

// node responses
type NodeResp struct {
	ID        string  `json:"id"`
	Label     string  `json:"label"`
	Root      bool    `json:"root"`
	PositionX float32 `json:"positionX"`
	PositionY float32 `json:"positionY"`
}

type RootNodesResp struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// edge responses
type EdgeResp struct {
	ID       string `json:"id"`
	SourceID string `json:"sourceId"`
	TargetID string `json:"targetId"`
}

// mind map response
type MindMapResp struct {
	Category string     `json:"category"` // rootID->label
	Nodes    []NodeResp `json:"nodes"`
	Edges    []EdgeResp `json:"edges"`
}

// document responses

type DocumentResp struct {
	ID     string `json:"id"`
	NodeID string `json:"nodeId"`
	Data   string `json:"data"`
}

// error
type ErrorResp struct {
	Error string `json:"error"`
}

// ok response
type OkResp struct {
	Msg string `json:"msg"`
}
