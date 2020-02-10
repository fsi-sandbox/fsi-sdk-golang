package nibss

type NibssCredentials struct {
	SandboxKey       string
	OrganisationCode string
}

type ResetCredentials struct {
	AESKey   string `json:"Aes_key"`
	Code     string `json:"Code"`
	Email    string `json:"Email"`
	IVKey    string `json:"Ivkey"`
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
