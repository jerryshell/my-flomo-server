package service

import (
	"testing"
)

func TestUserListByEmailIsNotNull(t *testing.T) {
	userListByEmailIsNotNull, err := UserListByEmailIsNotNull()
	if err != nil {
		t.Error(err)
	}
	t.Logf("len(userListByEmailIsNotNull): %v", len(userListByEmailIsNotNull))
	if userListByEmailIsNotNull == nil {
		t.Error("UserListByEmailIsNotNull() is nil")
	}
}
