package key_value_store

type webClient struct {
	url string
}

func NewWebClient(url string) (SimpleKeyValueStore) {
	return &webClient{
		url: url,
	}
}

func (store *webClient) Get(key string) string {

}

func (store *webClient) Set(key string, value string) {

}
