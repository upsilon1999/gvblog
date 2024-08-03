package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

const prefix = "logout_"

// Logout 针对注销的操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(core.RedisCtx,prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) (bool,error) {
	// keys := global.Redis.Keys(core.RedisCtx,prefix + "*").Val()
	keys,err:=global.Redis.Keys(core.RedisCtx,prefix + "*").Result()
	if err != nil {
        return false,err
	}
	if utils.InList(prefix+token, keys) {
		return true,nil
	}
	return false,nil
}
