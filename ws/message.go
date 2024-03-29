package ws

import (
	"encoding/json"

	"github.com/gofiber/contrib/websocket"
)

type ActionType string
type ServerActionType string

const (
	CommunityChatMessage ActionType = "community_chat_message"
	Notification         ActionType = "notification"
)

const (
	Logout   ServerActionType = "logout"
	NoAction ServerActionType = "no_action"
)

type MessageFormat struct {
	Token            string      `json:"token"`
	Data             interface{} `json:"data"`
	ActionType       ActionType  `json:"actionType"`
	ClientConnection *websocket.Conn
}

type ServerToClientMessageFormat struct {
	Data       interface{}      `json:"data"`
	ActionType ServerActionType `json:"actionType"`
}

func SendServerMessage(wsConn *websocket.Conn, data string, actionType ServerActionType) {
	message := ServerToClientMessageFormat{
		Data:       data,
		ActionType: actionType,
	}
	serverMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

	err = wsConn.WriteMessage(websocket.TextMessage, serverMsg)
	if err != nil {
		wsConn.Close()
		Unregister <- wsConn
		return
	}
}
