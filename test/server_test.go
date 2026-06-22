package httpserver

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/httpserver"
	"net/http"
	"testing"
	"time"
)

// TODO - consider if we can share a single server instance for testing or if we need isolation
func TestServerPing(t *testing.T) {
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

	const url = "http://localhost:8090/ping" // TODO - construct from config when added

	// The server boots asynchronously, so retry until it accepts
	// connections, giving up after a short timeout.
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

	if resp.StatusCode != http.StatusOK {
		t.Errorf("GET %s: expected status %d OK, got %d", url, http.StatusOK, resp.StatusCode)
	}
}
