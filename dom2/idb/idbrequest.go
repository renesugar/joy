package idb

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/idbrequestreadystate"
)

// js:"IDBRequest,omit"
type IDBRequest interface {
	EventTarget

	// Error
	Error() (err *domerror.DOMError)

	// Onerror
	Onerror() (onerror func(Event))

	// Onerror
	SetOnerror(onerror func(Event))

	// Onsuccess
	Onsuccess() (onsuccess func(Event))

	// Onsuccess
	SetOnsuccess(onsuccess func(Event))

	// ReadyState
	ReadyState() (readyState *idbrequestreadystate.IDBRequestReadyState)

	// Result
	Result() (result interface{})

	// Source
	Source() (source interface{})

	// Transaction
	Transaction() (transaction *IDBTransaction)
}
