package git

import (
	lib "gopkg.in/libgit2/git2go.v27"
)

type CredType uint8

const (
	CredTypeUserpassPlaintext CredType = iota
    CredTypeSshKey
    CredTypeSshAgent
)

type Credential interface {
	Type() CredType
}

// OptionsWithCreds provides an interface to get fetch callbacks
type OptionsWithCreds interface {
	authCallbackFunc(string, string, lib.CredType) (lib.ErrorCode, *lib.Cred)
	certCheckCallbackFunc(*lib.Certificate, bool, string) lib.ErrorCode
	creds() Credential
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

func defaultRemoteCallbacks(opts OptionsWithCreds) lib.RemoteCallbacks {
	rcb := lib.RemoteCallbacks{}
	rcb.CredentialsCallback = opts.authCallbackFunc
	rcb.CertificateCheckCallback = opts.certCheckCallbackFunc
	return rcb
}

func defaultAuthCallback(opts OptionsWithCreds, url string, uname string, credType lib.CredType) (lib.ErrorCode, *lib.Cred) {
	if opts.creds() == nil {
		return lib.ErrAuth, nil
	}
	cr := opts.creds()

	switch credType {
	case lib.CredTypeUserpassPlaintext:
		switch cr.(type) {
		case *CredentialsAsPlainText:
			credentials := cr.(*CredentialsAsPlainText)
			errCode, cred := lib.NewCredUserpassPlaintext(credentials.UserName, credentials.Password)
			return lib.ErrorCode(errCode), &cred
		default:
			return lib.ErrAuth, nil
		}
	case lib.CredTypeSshKey:
		switch cr.(type) {
		case *CredentialsAsSshKey:
			credentials := cr.(*CredentialsAsSshKey)
			errCode, cred := lib.NewCredSshKey(credentials.UserName, credentials.PublicKeyPath, credentials.PrivateKeyPath, credentials.Passphrase)
			return lib.ErrorCode(errCode), &cred
		default:
			return lib.ErrAuth, nil
		}
	case lib.CredTypeSshCustom, lib.CredTypeDefault, 70:
		switch cr.(type) {
		case *CredentialsAsSshAgent:
			credentials := cr.(*CredentialsAsSshAgent)
			errCode, cred := lib.NewCredSshKeyFromAgent(credentials.UserName)
			return lib.ErrorCode(errCode), &cred
		default:
			return lib.ErrAuth, nil
		}
	default:
		return lib.ErrAuth, nil
	}
}

func defaultCertCheckCallback (opts OptionsWithCreds, cert *lib.Certificate, valid bool, hostname string) lib.ErrorCode {
	// TODO: look for certificate check
	return 0
}