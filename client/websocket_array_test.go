package client

import (
	"errors"
	"testing"

	"github.com/lajosdeme/polymarket-go-api/types"
)

// The CLOB delivers the initial book snapshots of a subscription as a JSON
// array with one element per asset id. Each element must be dispatched to the
// book handler as if it had arrived alone.
func TestHandleMessage_ArrayFrameDispatchesEachElement(t *testing.T) {
	w := NewWebSocketClient("", nil)
	var got []string
	w.SetBookMessageHandler(func(e *types.WebSocketBookEvent) { got = append(got, e.AssetID) })
	var errs []error
	w.SetErrorHandler(func(err error) { errs = append(errs, err) })

	frame := []byte(`[{"event_type":"book","asset_id":"a","market":"m","bids":[],"asks":[],"timestamp":"1","hash":"h"},
	                 {"event_type":"book","asset_id":"b","market":"m","bids":[],"asks":[],"timestamp":"1","hash":"h"}]`)
	w.handleMessage(frame)

	if len(errs) != 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("book handler calls = %v, want [a b]", got)
	}
}

func TestHandleMessage_ObjectFrameUnchanged(t *testing.T) {
	w := NewWebSocketClient("", nil)
	var got []string
	w.SetBookMessageHandler(func(e *types.WebSocketBookEvent) { got = append(got, e.AssetID) })
	w.handleMessage([]byte(`{"event_type":"book","asset_id":"a","market":"m","bids":[],"asks":[],"timestamp":"1","hash":"h"}`))
	if len(got) != 1 || got[0] != "a" {
		t.Fatalf("book handler calls = %v, want [a]", got)
	}
}

func TestHandleMessage_MalformedArrayReportsOneError(t *testing.T) {
	w := NewWebSocketClient("", nil)
	var errs []error
	w.SetErrorHandler(func(err error) { errs = append(errs, err) })
	w.handleMessage([]byte(`[{"event_type":"book"`))
	if len(errs) != 1 || errors.Is(errs[0], nil) {
		t.Fatalf("errors = %v, want exactly one", errs)
	}
}

// Close must not race with startPing creating the ticker.
func TestClose_DuringStartPingIsRaceFree(t *testing.T) {
	for i := 0; i < 50; i++ {
		w := NewWebSocketClient("", nil)
		go w.startPing()
		_ = w.Close()
	}
}
