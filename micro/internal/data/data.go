package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	slog "log"
	"micro/internal/conf"
	"os"
	"time"
)

var ProviderSet = wire.NewSet(NewData,
	NewDB,
	NewRedis,
	NewCoreRepo,
	NewSysUserRepo,
	NewSysRoleRepo,
	NewSysMenuRepo,
	NewSysDeptRepo,
	NewSysDictRepo,
	NewSysConfigRepo,
	NewSysPostRepo,
)

type Data struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{Db: db, Rdb: rdb}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢查询sql阀值
			Colorful:      true,        //禁用彩色打印
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{
			//SingularTable: true, //表名是否加s
		},
	})
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		panic("failed to connect database")
	}
	//err2 := db.AutoMigrate(&User{})
	//if err2 != nil {
	//	panic("AutoMigrate is failed")
	//}
	return db
}

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		Password:     c.Redis.Password,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	// ping Redis客户端，判断连接是否存在
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Error("Redis database connection failure, err : %v", err)
	}
	log.Info("Cache enabled successfully!")
	return rdb
}
