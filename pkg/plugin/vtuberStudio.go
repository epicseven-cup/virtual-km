package plugin

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/url"
	"time"
	"virtual-km/internal/message"

	"github.com/gorilla/websocket"
)

type VtubeStudio struct {
	name          string
	developer     string
	icon          string
	websocketAddr *url.URL
	apiVersion    string
	outgoing      chan message.Message

	slogger *slog.Logger
	status  chan bool
}

func NewVtubeStudio(name, developer string, icon []byte, websocketAddr *url.URL, apiVersion string) (*VtubeStudio, error) {

	logger := slog.Default()
	timeContext, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()
	conn, respond, err := websocket.DefaultDialer.DialContext(timeContext, websocketAddr.String(), nil)

	if err != nil {
		return nil, err
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	logger.Info(fmt.Sprintf("connected to vtube studio: %s", conn.RemoteAddr()))
	logger.Info(fmt.Sprintf("respond header: %s", respond.Header))

	encodedIcon := base64.StdEncoding.EncodeToString(icon)

	// Respond getting from vtube studio
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				logger.Error(err.Error())
				return
			}
			logger.Info(fmt.Sprintf("vtube studio: %s", message))
		}
	}()

	status := make(chan bool)
	outgoing := make(chan message.Message)

	go func() {
		for {
			select {
			case <-status:
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

	return &VtubeStudio{
		name:          name,
		developer:     developer,
		icon:          encodedIcon,
		websocketAddr: websocketAddr,
		apiVersion:    apiVersion,
		outgoing:      outgoing,
		slogger:       logger,
		status:        status,
	}, nil
}
