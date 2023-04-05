package cache

type Contract interface {
	Get(key string, fallback string) string
	Set(key string, value string, seconds int)
	Exist(key string) bool
	GetMarshal(key string, unMarshal interface{}) error
	SetMarshal(key string, canMarshalVal interface{}, seconds int) error
}
