package entity

// Publish 메세지 포멧
// "event": "user.created"

type EventMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Ts    int64       `json:"ts"`
}
