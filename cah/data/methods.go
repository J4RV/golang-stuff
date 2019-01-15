package data

func LoginCorrect(username, pass string) bool {
	u, err := getUserByName(username)
	if err != nil {
		return false
	}
	return correctPass(pass, u.Password)
}
