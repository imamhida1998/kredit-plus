package model

type RedisStoreRequest struct {
	KeyValue string `json:"keyValue"`
	Value    string `json:"value"`
	Lifetime int    `json:"lifetime"`
}

type RedisValueEntity struct {
	Value string `json:"value"`
}
