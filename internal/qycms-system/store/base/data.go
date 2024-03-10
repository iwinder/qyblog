package base

import "github.com/iwinder/qyblog/internal/pkg/options"

type Data struct {
	// TODO wrapped database client
	Db    *MySqlStore
	Redis *RedisStore
}

func NewData(opt *options.DataOptions) (*Data, error) {
	var err error
	var mysqlStore *MySqlStore
	var redisStore *RedisStore
	mysqlStore, err = GetMySQLFactoryOr(opt.MySQLOptions)
	if err != nil {
		return nil, err
	}
	redisStore, err = GetRedisClientOr(opt.RedisOptions)
	if err != nil {
		return nil, err
	}
	return &Data{Db: mysqlStore, Redis: redisStore}, nil

}
func (d *Data) Close() {
	d.Redis.Close()
}
