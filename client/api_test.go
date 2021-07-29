package client

import (
	"testing"
)

func TestInfo(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	for _, c := range []struct {
		address string
		query   string
	}{
		{srv.URL, "all"},
	} {

		got, err := Info(c.address, c.query)

		if err != nil {
			t.Error(err)
		}

		if got == nil {
			t.Errorf("Info(%q, %q) == %q", c.address, c.query, got)
		}

		if len(got) != 713 {
			t.Error("Not all entries retrieved")
		}
	}

}

func TestLogin(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	apiInfo, err := Info(srv.URL, "all")
	if err != nil {
		t.Error(err)
	}
	for _, c := range []struct {
		address string
		account string
		passwd  string
		session string
		format  string
	}{
		{srv.URL, "testuser", "password1234", "TestAuth", "cookie"},
	} {

		got, errLogin := Login(apiInfo, c.address, c.account, c.passwd, c.session, c.format)
		if errLogin != nil {
			t.Error(errLogin)
		}
		if got.Sid != "my_sid_token" {
			t.Errorf("Login(%q, %q, %q, %q, %q) == %q", c.address, c.account, c.passwd, c.session, c.format, got)
		}

	}

}

func TestLogout(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	apiInfo, err := Info(srv.URL, "all")
	if err != nil {
		t.Error(err)
	}

	for _, c := range []struct {
		address string
		session string
	}{
		{srv.URL, "TestAuth"},
	} {

		got, errLogout := Logout(apiInfo, c.address, c.session)
		if errLogout != nil {
			t.Error(errLogout)
		}
		if !got {
			t.Errorf("Logout(%q, %q) ==", c.address, c.session)
		}

	}

}
