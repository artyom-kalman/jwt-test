package main

var cachedDatabase map[string]*AuthDB = make(map[string]*AuthDB, 0)

func DatabaseFabric(path string) *AuthDB {
	if cachedDatabase[path] != nil {
		return cachedDatabase[path]
	}

	cachedDatabase[path] = NewAuthDB(path)
	return cachedDatabase[path]
}
