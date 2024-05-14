package constants

const (
	// CachePrefix 应用缓存数据前缀
	CachePrefix = "APP:"

	CacheModelMem   = "memory:"
	CacheModelRedis = "redis:"
	CacheModelDist  = "dist:"
	CacheCaptcha    = CachePrefix + "captcha:"
	CacheMenu       = CachePrefix + "menu:"
	// CacheSysDict 字典缓存菜单KEY
	CacheSysDict = CachePrefix + "sysDict:"
	// CacheSysDictTag 字典缓存标签
	CacheSysDictTag = CachePrefix + "sysDictTag:"
	// CacheSysConfigTag 系统参数配置
	CacheSysConfigTag = CachePrefix + "sysConfigTag:"
)
