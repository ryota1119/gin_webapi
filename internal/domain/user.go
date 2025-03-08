package domain

import "strconv"

// UserID ユーザーID
type UserID uint

// User ユーザー情報
type User struct {
	ID       UserID
	Name     string
	Email    string
	Password string
}

// String はUserIDをstring型にキャストする
func (uid UserID) String() string {
	return strconv.FormatUint(uint64(uid), 10)
}
