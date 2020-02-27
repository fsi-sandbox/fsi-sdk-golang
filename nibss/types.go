// Package nibss implements access to nibss sandbox.
//
// It futher implements bvnr, fingerprint and BVNPlaceholder sandbox APIs
// as exported functions.
package nibss

// NibssCredentials struct required for innovation sandbox access.
type NibssCredentials struct {
	SandboxKey       string
	OrganisationCode string
}

// ResetCredentials struct from Reset API and required nibss sandbox interactions.
type ResetCredentials struct {
	AESKey   string `json:"Aes_key"`
	Code     string `json:"Code"`
	Email    string `json:"Email"`
	IVKey    string `json:"Ivkey"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
