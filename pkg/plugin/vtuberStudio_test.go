package plugin

import (
	"net/url"
	"os"
	"testing"
)

func TestNewVtubeStudio(t *testing.T) {
	file, err := os.ReadFile("../asset/icon.jpg")
	if err != nil {
		t.Fatal(err)
		return
	}

	websocketAddr := url.URL{Scheme: "ws", Host: "localhost:8001", Path: "/"}

	v, err := NewVtubeStudio("VTubeStudioPublicAPI", "tomato", file, &websocketAddr, "1.0")
	if err != nil {
		t.Fatal(err)
		return
	}

	v.Start()
	for {
	}
}
