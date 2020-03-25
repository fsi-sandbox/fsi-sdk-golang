// Package union implements access to union sandbox.
//
// It futher implements token and customer sandbox APIs
// as exported functions.
package union

// UnionCredentials struct required for innovation sandbox access.
type UnionCredentials struct {
	SandboxKey string
}
