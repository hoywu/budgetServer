package utils

func ValidateUsername(username string) bool {
	return len(username) >= 5
}

func ValidatePassword(password string) bool {
	return len(password) >= 8
}
