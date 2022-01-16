package service

import (
	"testing"
)

func TestUserListByEmailIsNotNull(t *testing.T) {
	userListByEmailIsNotNull := UserListByEmailIsNotNull()
	t.Logf("len(userListByEmailIsNotNull): %v", len(userListByEmailIsNotNull))
	if userListByEmailIsNotNull == nil {
		t.Error("UserListByEmailIsNotNull() is nil")
	}
}
