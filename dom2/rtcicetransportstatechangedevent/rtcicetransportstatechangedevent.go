package rtcicetransportstatechangedevent

import (
	"github.com/matthewmueller/golly/dom2/rtcicetransportstate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// RTCIceTransportStateChangedEvent struct
// js:"RTCIceTransportStateChangedEvent,omit"
type RTCIceTransportStateChangedEvent struct {
	window.Event
}

// State
func (*RTCIceTransportStateChangedEvent) State() (state *rtcicetransportstate.RTCIceTransportState) {
	js.Rewrite("$<.State")
	return state
}
