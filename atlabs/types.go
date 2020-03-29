// Package atlabs implements access to atlabs sandbox.
//
// It futher implements airtime, sms, token and voice sandbox APIs
// as exported functions.
package atlabs

// AtlabsCredentials struct required for innovation sandbox access.
type AtlabsCredentials struct {
	SandboxKey string
}
