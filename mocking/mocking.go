package mocking

type DataStore interface {
	Set(key string, value string)
}

type Contacts struct {
	store DataStore
}

func (contacts *Contacts) SaveName(name string) {
	contacts.store.Set("name", name)
}
