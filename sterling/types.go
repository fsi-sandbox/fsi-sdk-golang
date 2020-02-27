// Package sterling implements access to sterling sandbox.
//
// It futher implements account and Transfer sandbox APIs
// as exported functions.
package sterling

// SterlingCredentials struct required for innovation sandbox access.
type SterlingCredentials struct {
	SandboxKey      string
	SubscriptionKey string
	Appid           string
	Ipval           string
}
