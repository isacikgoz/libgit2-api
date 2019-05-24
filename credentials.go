package git

type CredType uint8

const (
	CredTypeUserpassPlaintext CredType = iota
    CredTypeSshKey
    CredTypeSshAgent
)

type Credential interface {
	Type() CredType
}

type CredentialsAsPlainText struct {
	UserName string
	Password string
}

func (c *CredentialsAsPlainText) Type() CredType {
	return CredTypeUserpassPlaintext
}

type CredentialsAsSshKey struct {
	UserName string
	PublicKeyPath string
	PrivateKeyPath string
	Passphrase string
}

func (c *CredentialsAsSshKey) Type() CredType {
	return CredTypeSshKey
}

type CredentialsAsSshAgent struct {
	UserName string
}

func (c *CredentialsAsSshAgent) Type() CredType {
	return CredTypeSshAgent
}
