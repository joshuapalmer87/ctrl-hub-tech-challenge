package httpserver_test

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/httpserver"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const serverAddr = "http://localhost:8090" // TODO - construct from config when added

func TestServerPing(t *testing.T) {
	setup(t)
	url := serverAddr + "/ping"
	resp, err := http.Get(url)
	if err == nil {
		require.NoError(t, err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Spins up the server as a seperate process and checks that it's running, and sets up shutdown smoothly
func setup(t *testing.T) {
	server := httpserver.NewHttpServer()

	// Serve blocks (http.ListenAndServe), so run it in the background.
	go server.Serve()
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			t.Errorf("server shutdown: %v", err)
		}
	})

	url := serverAddr + "/ping"

	// The server boots asynchronously, so retry until it accepts
	// connections, giving up after a short timeout.
	// TODO - is there a nicer way to do this?
	var resp *http.Response
	var err error
	for attempt := 0; attempt < 50; attempt++ {
		resp, err = http.Get(url)
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("could not reach %s: %v", url, err)
	}
	defer resp.Body.Close()
}
