package git

import (
	"testing"

	lib "gopkg.in/libgit2/git2go.v27"
)

func TestDefaultAuthCallbackFunc(t *testing.T) {
	opts := &CloneOptions{}
	var tests = []struct {
		inputOpts OptionsWithCreds
		inputURL string
		inoutUsr string
		inputCred lib.CredType
		outErrCode lib.ErrorCode
		outCredential *lib.Cred
	}{
		{opts, "https://github.com/isacikgoz/gia.git", "", lib.CredTypeUserpassPlaintext, lib.ErrAuth, nil},
	}
	for _, test := range tests {
		if errCode, _ := defaultAuthCallback(test.inputOpts, test.inputURL, test.inoutUsr, test.inputCred); errCode != test.outErrCode {
			t.Error("test failed.")
		}
	}	
}

func TestDefaultCertCheckCallback(t *testing.T) {
	opts := &CloneOptions{}
	var tests = []struct {
		inputOpts OptionsWithCreds
		inputCert *lib.Certificate
		inputValid bool
		inputHost string
		outErrCode lib.ErrorCode
	}{
		{opts, nil, false, "", 0},
	}
	for _, test := range tests {
		if errCode := defaultCertCheckCallback(test.inputOpts, test.inputCert, test.inputValid, test.inputHost); errCode != test.outErrCode {
			t.Error("test failed.")
		}
	}
}