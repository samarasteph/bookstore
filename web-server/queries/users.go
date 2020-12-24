package queries

var users map[string]string = make(map[string]string)

// AddUser add user with credentials in database (in memory)
func AddUser(name string, pwd string) {
	users[name] = pwd
}

// GetUserCred return user Cred from user name
func GetUserCred(name string) (string, bool) {
	cred, exists := users[name]
	return cred, exists
}
