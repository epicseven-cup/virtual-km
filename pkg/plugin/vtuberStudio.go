package plugin

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"net/url"
	"virtual-km/internal/message"

	"github.com/gorilla/websocket"
)

type VtubeStudio struct {
	Name          string
	Developer     string
	Icon          string
	WebsocketAddr *url.URL
	ApiVersion    string
	outgoing      chan message.Message

	slogger *slog.Logger
	status  chan bool
}

func NewVtubeStudio(name, developer string, icon []byte, websocketAddr *url.URL, apiVersion string) (*VtubeStudio, error) {

	logger := slog.Default()
	//timeContext, cancel := context.WithTimeout(context.Background(), time.Second*50)
	//defer cancel()
	//conn, respond, err := websocket.DefaultDialer.DialContext(timeContext, websocketAddr.String(), nil)
	conn, respond, err := websocket.DefaultDialer.Dial(websocketAddr.String(), nil)

	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("connected to vtube studio: %s", conn.RemoteAddr()))
	logger.Info(fmt.Sprintf("respond header: %s", respond.Header))

	encodedIcon := base64.StdEncoding.EncodeToString(icon)

	// Respond getting from vtube studio
	go func() {
		for {
			msg := message.Message{}
			err := conn.ReadJSON(&msg)
			if err != nil {
				logger.Error(err.Error())
			}
			logger.Info(fmt.Sprintf("msg: %+v", msg))
		}
	}()

	status := make(chan bool)
	outgoing := make(chan message.Message)

	logger.Info("Starting websocket outgoing go routine")

	go func() {
		for {
			select {
			case <-status:
				err := conn.Close()
				if err != nil {
					logger.Error(err.Error())
					return
				}
				return

			case msg := <-outgoing:
				err := conn.WriteJSON(msg)
				if err != nil {
					logger.Error(err.Error())
					return
				}
			}
		}
	}()

	logger.Info(fmt.Sprintf("vtube studio: %s, starting", name))

	return &VtubeStudio{
		Name:          name,
		Developer:     developer,
		Icon:          encodedIcon,
		WebsocketAddr: websocketAddr,
		ApiVersion:    apiVersion,
		outgoing:      outgoing,
		slogger:       logger,
		status:        status,
	}, nil
}

func (v *VtubeStudio) Start() {
	msg := message.Message{
		ApiName:     v.Name,
		ApiVersion:  v.ApiVersion,
		RequestID:   uuid.New().String(),
		MessageType: "APIStateRequest",
	}
	v.outgoing <- msg
}

func (v *VtubeStudio) SendMessage(messageType string, data message.Data) message.Message {
	msg := message.Message{
		ApiName:     v.Name,
		ApiVersion:  v.ApiVersion,
		RequestID:   uuid.New().String(),
		MessageType: messageType,
		Data:        data,
	}
	v.outgoing <- msg
	return msg
}
