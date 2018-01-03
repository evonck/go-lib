package redis

type RedisTest struct {
	GetResponse string
}

func (test *RedisTest) SetGetResponse(value string) {
	test.GetResponse = value
}

func (test *RedisTest) New(socket, passwrod, readdTimeout, writeTimeout string) {
	return
}

func (test *RedisTest) LRange(key string, start, stop int) ([]string, error) {
	testLrange := []string{"ok"}
	return testLrange, nil
}

func (test *RedisTest) ZRange(key string, start, stop int) ([]string, error) {
	testLrange := []string{"ok"}
	return testLrange, nil
}

func (test *RedisTest) RPush(key, value string) {
	return
}

func (test *RedisTest) Close() {
	return
}

func (test *RedisTest) FlushAll() (int64, error) {
	return 42, nil
}

func (test *RedisTest) Multi() error {
	return nil
}

func (test *RedisTest) Exec() (interface{}, error) {
	return nil, nil
}

func (test *RedisTest) String(valueBytes []byte) string {
	return ""
}

func (test *RedisTest) PushTTL(key string, value string, ttl int) {
	return
}

func (test *RedisTest) Expire(key string, ttl int) {
	return
}

func (test *RedisTest) GET(key string) (string, error) {
	return test.GetResponse, nil
}

func (test *RedisTest) Push(key string, value string) {
	return
}

// KEYS get keys list from redis
func (test *RedisTest) KEYS(key string) ([]string, error) {
	return nil, nil
}
