package msg

// Publish 메세지 포멧
// "event": "user.created"

type Event string

const errKey = "only:error:http:mathod:pk:userid_234"
const errData = "{message,stack}"

type EventMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Ts    int64       `json:"ts"`
}
