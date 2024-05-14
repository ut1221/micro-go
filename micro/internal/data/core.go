package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"time"
)

type coreRepo struct {
	data *Data
	log  *log.Helper
}

func NewCoreRepo(data *Data, logger log.Logger) biz.CoreRepo {
	return &coreRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// RSetData
// Keys():根据正则获取keys
// Type():获取key对应值得类型
// Del():删除缓存项
// Exists():检测缓存项是否存在
// Expire(),ExpireAt():设置有效期
// TTL(),PTTL():获取有效期
// DBSize():查看当前数据库key的数量
// FlushDB():清空当前数据
// FlushAll():清空所有数据库
// Set():设置键缓存
// SetEX():设置并指定过期时间
// SetNX():设置并指定过期时间,仅当key不存在的时候才设置。
// Get():获取键值
// GetRange():字符串截取
// Incr():增加+1
// IncrBy():按指定步长增加
// Decr():减少-1
// DecrBy():按指定步长减少
// Append():追加
// StrLen():获取长度
func (c coreRepo) RSetData(ctx context.Context, key string, value interface{}) error {
	err := c.data.Rdb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c coreRepo) RSetObj(ctx context.Context, key string, value interface{}) error {
	marshal, _ := json.Marshal(value)
	err := c.data.Rdb.Set(ctx, key, marshal, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c coreRepo) RGetData(ctx context.Context, key string) (interface{}, error) {
	val, err := c.data.Rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c coreRepo) RGetObj(ctx context.Context, key string) ([]byte, error) {
	val, err := c.data.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}
func (c coreRepo) RRemoveData(ctx context.Context, key string) error {
	err := c.data.Rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c coreRepo) RExistsData(ctx context.Context, key string) (int64, error) {
	val, err := c.data.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, err
}
