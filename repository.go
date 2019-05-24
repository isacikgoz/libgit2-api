package git

import (
	lib "gopkg.in/libgit2/git2go.v27"
)

// Repository is the wrapper and main interface to git repository
type Repository struct {
	essence *lib.Repository
	dir string
}

// CloneOptions are mostly used git clone options from a remote
type CloneOptions struct {
	Bare bool
	Recursive bool
	Depth int
	Credentials Credential
}

// Clone fetches a git repository from a given url
func Clone(path string, url string, opts *CloneOptions) (*Repository, error) {
	options := &lib.CloneOptions{
		Bare: opts.Bare,
	}
	fetchOptions := &lib.FetchOptions{}

	remoteCallbacks := lib.RemoteCallbacks{}
	// authCallbackFunc returns specific data to given auth type
	remoteCallbacks.CredentialsCallback = opts.authCallbackFunc
	// certificateCheckCallback added for user cancelled hostkey check error
	remoteCallbacks.CertificateCheckCallback = opts.certificateCheckCallback
	fetchOptions.RemoteCallbacks = remoteCallbacks

	options.FetchOptions = fetchOptions
	r, err:= lib.Clone(url, path, options)
	if err != nil {
		return nil, err
	}
	repository := &Repository{
		essence: r,
		dir: path,
	}
	return repository, nil
}

func (opts *CloneOptions) authCallbackFunc(url string, uname string, credType lib.CredType) (lib.ErrorCode, *lib.Cred) {
	if opts.Credentials == nil {
		return lib.ErrAuth, nil
	}
	cr := opts.Credentials

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

func (opts *CloneOptions) certificateCheckCallback (cert *lib.Certificate, valid bool, hostname string) lib.ErrorCode {
	// TODO: look for certificate check
	return 0
}