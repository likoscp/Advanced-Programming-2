package models

import "testing"

func TestIsValid(t *testing.T) {
	invalidUsers := []*User{
		{Username: "hello1", Password: "123321", Email: "sabd"},
		{Username: "", Password: "", Email: ""},
		{Username: "", Password: "123321", Email: "sabdp1@gmai.cos"},
		{Username: "barcek221", Password: "123321", Email: ""},
		{Username: "barcek2271", Password: "", Email: "sabdp1@gmai.cos"},
		{Username: "barcek2281", Password: "121", Email: "sabdp1@gmai.cos"},
		{Username: "bar", Password: "123321", Email: "sabdp1@gmai.cos"},

	}
	for _, user := range invalidUsers {
		err := user.IsValid()
		if err == nil {
			t.Fatalf("user should be invalid, %v", err)
		}
	}

	validUsers := []*User{
		{Username: "barcke282", Password: "awdawda", Email: "sabdp123@asdf.casod"},
		{Username: "zxcvvczx", Password: "wadawdwad", Email: "sabdp123@asdf.casod"},
		{Username: "123ssddsa", Password: "awdawda", Email: "sabdp123@asdf.casod"},
		{Username: "bazxcvrcke282", Password: "awdddawda", Email: "sabdp123@asdf.casod"},
	}

	for _, user := range validUsers {
		err := user.IsValid()
		if err != nil {
			t.Fatalf("user should be valid, %v", err)
		}
	}

}

func TestHashPassword(t *testing.T) {
	oldPassword := "password"
	u := &User{
		Email: "sabdp123@asdda.com",
		Password: oldPassword,
		Username: "adsadaa",
	}

	if err := u.IsValid(); err != nil {
		t.Fatalf("user should be valid, error: %v", err)
	}

	err := u.HashPassword()

	if err != nil {
		t.Fatalf("user should be ok, error: %v", err)
	}

	if u.Password == oldPassword{
		t.Fatalf("user's password should be different, old: %v, new: %v", oldPassword, u.Password)
	}

}

func TestComparePassword(t *testing.T) {
	u1 := &User{
		Email: "asdas@asd.com",
		Password: "password",
		Username: "123321",
	}
	u2 := &User{
		Email: "asdas@asd.com",
		Password: "password",
		Username: "123321",
	}

	if err := u1.IsValid(); err != nil {
		t.Fatalf("error: %v", err)
	}

	if err := u2.IsValid(); err != nil {
		t.Fatalf("error: %v", err)
	}

	u1.HashPassword()
	u2.HashPassword()

	if u1.HashPassword() != u2.HashPassword() {
		t.Fatalf("users password should be same")
	}
}