package utils

func ValidateUsername(username string) bool {
	return len(username) >= 5
}

func ValidatePassword(password string) bool {
	return len(password) >= 8
}

func ValidateNickname(nickname string) bool {
	return len(nickname) >= 1
}
