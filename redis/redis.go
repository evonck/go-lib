package redis

type Redis interface {
	New(socketConf, password, readdTimeout, writeTimeout string)
	LRange(key string, start, stop int) ([]string, error)
	ZRange(key string, start, stop int) ([]string, error)
	RPush(key, value string)
	Close()
	FlushAll() (int64, error)
	Multi() error
	Exec() (interface{}, error)
	String(valueBytes []byte) string
	PushTTL(key string, value string, ttl int)
	Expire(key string, ttl int)
	GET(key string) (string, error)
	Push(key string, value string)
	KEYS(key string) ([]string, error)
}
