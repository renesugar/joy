package window

import (
	"github.com/matthewmueller/golly/dom2/mediakeysessiontype"
	"github.com/matthewmueller/golly/js"
)

// MediaKeys struct
// js:"MediaKeys,omit"
type MediaKeys struct {
}

// CreateSession
func (*MediaKeys) CreateSession(sessionType *mediakeysessiontype.MediaKeySessionType) (m *MediaKeySession) {
	js.Rewrite("$<.CreateSession($1)", sessionType)
	return m
}

// SetServerCertificate
func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	js.Rewrite("await $<.SetServerCertificate($1)", serverCertificate)
}
