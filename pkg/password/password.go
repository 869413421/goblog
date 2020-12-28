package password

import (
	"fmt"
	"goblog/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Hash进行加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Danger(err, "hash password error")
	}

	return string(bytes)
}

//检查密码和hash是否匹配
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

func IsHashed(str string) bool {
	return len(str) == 60

}
