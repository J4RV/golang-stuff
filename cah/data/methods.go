package data

func LoginCorrect(username, pass string) bool {
	u, ok := users[username]
	if !ok {
		return false
	}
	return correctPass(pass, u.Password)
}
