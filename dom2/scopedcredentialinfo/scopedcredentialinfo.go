package scopedcredentialinfo

import (
	"github.com/matthewmueller/golly/dom2/cryptokey"
	"github.com/matthewmueller/golly/dom2/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

// ScopedCredentialInfo struct
// js:"ScopedCredentialInfo,omit"
type ScopedCredentialInfo struct {
}

// Credential
func (*ScopedCredentialInfo) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.Credential")
	return credential
}

// PublicKey
func (*ScopedCredentialInfo) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.PublicKey")
	return publicKey
}
