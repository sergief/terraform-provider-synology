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

		got := Info(c.address, c.query)

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

	apiInfo := Info(srv.URL, "all")

	for _, c := range []struct {
		address string
		account string
		passwd  string
		session string
		format  string
	}{
		{srv.URL, "testuser", "password1234", "TestAuth", "cookie"},
	} {

		got := Login(apiInfo, c.address, c.account, c.passwd, c.session, c.format)

		if got.Sid != "my_sid_token" {
			t.Errorf("Login(%q, %q, %q, %q, %q) == %q", c.address, c.account, c.passwd, c.session, c.format, got)
		}

	}

}

func TestLogout(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	apiInfo := Info(srv.URL, "all")

	for _, c := range []struct {
		address string
		session string
	}{
		{srv.URL, "TestAuth"},
	} {

		got := Logout(apiInfo, c.address, c.session)

		if !got {
			t.Errorf("Logout(%q, %q) ==", c.address, c.session)
		}

	}

}
