package main

var cachedDatabase map[string]*UsersDatabase = make(map[string]*UsersDatabase, 0)

func DatabaseFabric(path string) *UsersDatabase {
	if cachedDatabase[path] != nil {
		return cachedDatabase[path]
	}

	cachedDatabase[path] = NewUserDatabase(path)
	return cachedDatabase[path]
}
