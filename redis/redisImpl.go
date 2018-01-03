package redis

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

// RedisImpl redis implementation of the Redis interface
type RedisImpl struct {
	connection redis.Conn
}

// ErrNil nil return by the redis
var ErrNil = redis.ErrNil

// New create new connection with redis
func (r *RedisImpl) New(socket, password, readdTimeout, writeTimeout string) {

	// define timeouts
	var (
		rtd, wtd time.Duration
		err      error
	)

	// raise error with read_timeout cannot be parsed as duration
	if rtd, err = time.ParseDuration(readdTimeout); err != nil {
		log.Fatalf("Could not convert Redis read_timeout param to duration %s%s", err.Error(), readdTimeout)
	}

	// raise error with write_timeout cannot be parsed as duration
	if wtd, err = time.ParseDuration(writeTimeout); err != nil {
		log.Fatalf("Could not convert Redis write_timeout param to duration %s%s", err.Error(), writeTimeout)
	}

	if password == "" {
		password = ""
	}
	r.connection, err = redis.Dial("tcp", socket, redis.DialReadTimeout(rtd),
		redis.DialWriteTimeout(wtd), redis.DialPassword(password))
	if err != nil {
		log.Fatal(err)
	}
}

// PushTTL push new key in redis with ttl
func (r *RedisImpl) PushTTL(key string, value string, ttl int) {
	r.connection.Do("SET", key, value, ttl)
}

// Expire modifie expire date of a key
func (r *RedisImpl) Expire(key string, ttl int) {
	r.connection.Do("EXPIRE", key, ttl)
}

// GET get key from redis
func (r *RedisImpl) GET(key string) (string, error) {
	return r.GetString(r.connection.Do("GET", key))
}

// KEYS get keys list from redis
func (r *RedisImpl) KEYS(key string) ([]string, error) {
	return r.Interface(r.connection.Do("KEYS", key))
}

// Push add value to redis
func (r *RedisImpl) Push(key string, value string) {
	r.connection.Do("SET", key, value)
}

// LRange create new lit in redis
func (r *RedisImpl) LRange(key string, start, stop int) ([]string, error) {
	return r.Interface(r.connection.Do("LRANGE", key, start, stop))
}

// ZRange create new lit in redis
func (r *RedisImpl) ZRange(key string, start, stop int) ([]string, error) {
	return r.Interface(r.connection.Do("ZRange", key, start, stop))
}

// RPush create new lit in redis
func (r *RedisImpl) RPush(key, value string) {
	r.connection.Send("RPUSH", key, value)
}

// Close close redis connection
func (r *RedisImpl) Close() {
	r.Close()
}

// FlushAll flush the db
func (r *RedisImpl) FlushAll() (int64, error) {
	return r.Int64(r.connection.Do("FLUSHALL"))
}

// Multi send multi to redis
func (r *RedisImpl) Multi() error {
	return r.connection.Send("MULTI")
}

// Exec execute a redis command
func (r *RedisImpl) Exec() (interface{}, error) {
	return r.connection.Do("EXEC")
}
func (r *RedisImpl) String(valueBytes []byte) string {
	return ""
}

// Values is a helper that converts an array command reply to a []interface{}.
// If err is not equal to nil, then Values returns nil, err. Otherwise, Values
// converts the reply as follows:
//
//  Reply type      Result
//  array           reply, nil
//  nil             nil, ErrNil
//  other           nil, error
func (r *RedisImpl) Values(reply interface{}, err error) ([]interface{}, error) {
	return redis.Values(reply, err)
}

// ScanStruct scans alternating names and values from src to a struct. The
// HGETALL and CONFIG GET commands return replies in this format.
//
// ScanStruct uses exported field names to match values in the response. Use
// 'redis' field tag to override the name:
//
//      Field int `redis:"myName"`
//
// Fields with the tag redis:"-" are ignored.
//
// Integer, float, boolean, string and []byte fields are supported. Scan uses the
// standard strconv package to convert bulk string values to numeric and
// boolean types.
//
// If a src element is nil, then the corresponding field is not modified.
func ScanStruct(src []interface{}, dest interface{}) error {
	return redis.ScanStruct(src, dest)
}

// Interface is a helper that converts an array command reply to a []string. If
// err is not equal to nil, then Strings returns nil, err. Nil array items are
// converted to "" in the output slice. Strings returns an error if an array
// item is not a bulk string or nil.
func (r *RedisImpl) Interface(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// GetString Interface is a helper that converts an array command reply to a []string. If
// err is not equal to nil, then Strings returns nil, err. Nil array items are
// converted to "" in the output slice. Strings returns an error if an array
// item is not a bulk string or nil.
func (r *RedisImpl) GetString(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// Int64 is a helper that converts an array command reply to a []string. If
// err is not equal to nil, then Strings returns nil, err. Nil array items are
// converted to "" in the output slice. Strings returns an error if an array
// item is not a bulk string or nil.
func (r *RedisImpl) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}
