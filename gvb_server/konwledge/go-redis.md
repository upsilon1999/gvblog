# go-redis

github地址

```sh
https://github.com/redis/go-redis
```

中文文档地址

```sh
https://redis.uptrace.dev/zh/
```

## 前置redis知识

主要是关于redis类型

Redis 主要支持以下几种数据类型：

- **string（字符串）:** 基本的数据存储单元，可以存储字符串、整数或者浮点数。
- **hash（哈希）:**一个键值对集合，可以存储多个字段。
- **list（列表）:**一个简单的列表，可以存储一系列的字符串元素。
- **set（集合）:**一个无序集合，可以存储不重复的字符串元素。
- **zset(sorted set：有序集合):** 类似于集合，但是每个元素都有一个分数（score）与之关联。
- **位图（Bitmaps）：**基于字符串类型，可以对每个位进行操作。
- **超日志（HyperLogLogs）：**用于基数统计，可以估算集合中的唯一元素数量。
- **地理空间（Geospatial）：**用于存储地理位置信息。
- **发布/订阅（Pub/Sub）：**一种消息通信模式，允许客户端订阅消息通道，并接收发布到该通道的消息。
- **流（Streams）：**用于消息队列和日志存储，支持消息的持久化和时间排序。
- **模块（Modules）：**Redis 支持动态加载模块，可以扩展 Redis 的功能。

### String（字符串）

string 是 redis 最基本的类型，你可以理解成与 Memcached 一模一样的类型，一个 key 对应一个 value。

string 类型是二进制安全的。意思是 redis 的 string 可以包含任何数据，比如jpg图片或者序列化的对象。

string 类型是 Redis 最基本的数据类型，string 类型的值最大能存储 512MB。

**常用命令**

- `SET key value`：设置键的值。
- `GET key`：获取键的值。
- `INCR key`：将键的值加 1。
- `DECR key`：将键的值减 1。
- `APPEND key value`：将值追加到键的值之后。

**示例**

```sh
redis 127.0.0.1:6379> SET age 19
OK
redis 127.0.0.1:6379> GET age
"19"
```

>**注意：**一个键最大能存储 512MB。

### Hash（哈希）

Redis hash 是一个键值(key=>value)对集合，类似于一个小型的 NoSQL 数据库。

Redis hash 是一个 string 类型的 field 和 value 的映射表，hash 特别适合用于存储对象。

每个哈希最多可以存储` 2^32 - 1` 个键值对。

**常用命令**

- `HSET key field value`：设置哈希表中字段的值。
- `HGET key field`：获取哈希表中字段的值。
- `HGETALL key`：获取哈希表中所有字段和值。
- `HDEL key field`：删除哈希表中的一个或多个字段。

**示例**

```sh
redis 127.0.0.1:6379> HMSET person name "Zhansan" age 19
"OK"
redis 127.0.0.1:6379> HGET person name
"Zhansan"
redis 127.0.0.1:6379> HGET person age
"19"
```

### List（列表）

Redis 列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）。

列表最多可以存储` 2^32 - 1 `个元素。

**常用命令**

- `LPUSH key value`：将值插入到列表头部。
- `RPUSH key value`：将值插入到列表尾部。
- `LPOP key`：移出并获取列表的第一个元素。
- `RPOP key`：移出并获取列表的最后一个元素。
- `LRANGE key start stop`：获取列表在指定范围内的元素。

**示例**

```sh
redis 127.0.0.1:6379> lpush database redis
(integer) 1
# 可以插入多个值 
redis 127.0.0.1:6379> rpush database mysql oracle
(integer) 3
redis 127.0.0.1:6379> lpush database mongodb
(integer) 4
redis 127.0.0.1:6379> lpush database rabbitmq
(integer) 5

#与limit相似，从0到10
redis 127.0.0.1:6379> lrange database 0 10
1) "rabbitmq"
2) "mongodb"
3) "redis"
4) "mysql"
5) "oracle"
```

### Set（集合）

Redis 的 Set 是 string 类型的无序集合。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。

集合中最大的成员数为 `2^32 - 1(4294967295, 每个集合可存储40多亿个成员)`。

**常用命令**

- `SADD key value`：向集合添加一个或多个成员。
- `SREM key value`：移除集合中的一个或多个成员。
- `SMEMBERS key`：返回集合中的所有成员。
- `SISMEMBER key value`：判断值是否是集合的成员。

**提示**

`sadd命令`添加一个 string 元素到 key 对应的 set 集合中，成功返回 1，如果元素已经在集合中返回 0。

```sh
sadd key member
```

**示例**

```sh
redis 127.0.0.1:6379> sadd datalist redis
(integer) 1
redis 127.0.0.1:6379> sadd datalist mongodb
(integer) 1
redis 127.0.0.1:6379> sadd datalist rabbitmq
(integer) 1

#rabbitmq 添加了两次，但根据集合内元素的唯一性，第二次插入的元素将被忽略。
redis 127.0.0.1:6379> sadd datalist rabbitmq
(integer) 0
redis 127.0.0.1:6379> smembers runoob

1) "redis"
2) "rabbitmq"
3) "mongodb"
```

### zset(sorted set：有序集合)

> Redis zset 和 set 一样也是string类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

zset的成员是唯一的,但分数(score)却可以重复。

**常用命令**

- `ZADD key score value`：向有序集合添加一个或多个成员，或更新已存在成员的分数。
- `ZRANGE key start stop [WITHSCORES]`：返回指定范围内的成员。
- `ZREM key value`：移除有序集合中的一个或多个成员。
- `ZSCORE key value`：返回有序集合中，成员的分数值。

**提示**

添加元素到集合，元素在集合中存在则更新对应score

```sh
zadd key score member 
```

**举例**

```sh
redis 127.0.0.1:6379> zadd students 0 zhansan
(integer) 1
redis 127.0.0.1:6379> zadd students 0 lisi
(integer) 1
redis 127.0.0.1:6379> zadd students 0 wanwu
(integer) 1
redis 127.0.0.1:6379> zadd students 0 lisi
(integer) 0

# 分数相同，安装字符串比大小，默认从小到大
redis 127.0.0.1:6379> ZRANGEBYSCORE students 0 1000
1) "lisi"
2) "wanwu"
3) "zhansan"

redis 127.0.0.1:6379> zadd students 1 wanwu
# 分数不同，同分段按字符串比大小排，不同分段默认从小到大从小到大
redis 127.0.0.1:6379> ZRANGEBYSCORE students 0 1000
1) "lisi"
2) "zhansan"
3) "wanwu" 
```

### 各类型适用场景

| 类型                 | 简介                                                   | 特性                                                         | 场景                                                         |
| :------------------- | :----------------------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- |
| String(字符串)       | 二进制安全                                             | 可以包含任何数据,比如jpg图片或者序列化的对象,一个键最大能存储512M | ---                                                          |
| Hash(字典)           | 键值对集合,即编程语言中的Map类型                       | 适合存储对象,并且可以像数据库中update一个属性一样只修改某一项属性值(Memcached中需要取出整个字符串反序列化成对象修改完再序列化存回去) | 存储、读取、修改用户属性                                     |
| List(列表)           | 链表(双向链表)                                         | 增删快,提供了操作某一段元素的API                             | 1,最新消息排行等功能(比如朋友圈的时间线) 2,消息队列          |
| Set(集合)            | 哈希表实现,元素不重复                                  | 1、添加、删除,查找的复杂度都是O(1) 2、为集合提供了求交集、并集、差集等操作 | 1、共同好友 2、利用唯一性,统计访问网站的所有独立ip 3、好友推荐时,根据tag求交集,大于某个阈值就可以推荐 |
| Sorted Set(有序集合) | 将Set中的元素增加一个权重参数score,元素按score有序排列 | 数据插入集合时,已经进行天然排序                              | 1、排行榜 2、带权重的消息队列                                |

# 基于V9的一些小结

## 安装客户端

Go语言中使用第三方库`https://github.com/go-redis/redis`连接Redis数据库并进行操作。使用以下命令下载并安装

```sh
go get github.com/go-redis/redis/v9
```



## 连接redis

>说明：Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。 它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。

`redis_con_test.go`

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)


func TestRedis(t *testing.T) {
    //Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
    //它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
    var ctx = context.Background()
    
    
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)
}
```

连接的配置项

```go
type Options struct {
    // 连接网络类型，如: tcp、udp、unix等方式
    // 如果为空默认tcp
    Network string

    // redis服务器地址，ip:port格式，比如：192.168.1.100:6379
    // 默认为 :6379
    Addr string

    // ClientName 是对网络连接设置一个名字，使用 "CLIENT LIST" 命令
    // 可以查看redis服务器当前的网络连接列表
    // 如果设置了ClientName，go-redis对每个连接调用 `CLIENT SETNAME ClientName` 命令
    // 查看: https://redis.io/commands/client-setname/
    // 默认为空，不设置客户端名称
    ClientName string

    // 如果你想自定义连接网络的方式，可以自定义 `Dialer` 方法，
    // 如果不指定，将使用默认的方式进行网络连接 `redis.NewDialer`
    Dialer func(ctx context.Context, network, addr string) (net.Conn, error)

    // 建立了新连接时调用此函数
    // 默认为nil
    OnConnect func(ctx context.Context, cn *Conn) error

    // 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
    // ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
    // 默认为空，不进行认证。
    Username string

    // 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
    // 当redis服务器版本在6.0以下时，仅作为密码认证。
    // ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
    // 默认为空，不进行认证。
    Password string

    // 允许动态设置用户名和密码，go-redis在进行网络连接时会获取用户名和密码，
    // 这对一些认证鉴权有时效性的系统来说很有用，比如一些云服务商提供认证信息有效期为12小时。
    // 默认为nil
    CredentialsProvider func() (username string, password string)

    // redis DB 数据库，默认为0
    DB int

    // 命令最大重试次数， 默认为3
    MaxRetries int

    // 每次重试最小间隔时间
    // 默认 8 * time.Millisecond (8毫秒) ，设置-1为禁用
    MinRetryBackoff time.Duration

    // 每次重试最大间隔时间
    // 默认 512 * time.Millisecond (512毫秒) ，设置-1为禁用
    MaxRetryBackoff time.Duration

    // 建立新网络连接时的超时时间
    // 默认5秒
    DialTimeout time.Duration

    // 从网络连接中读取数据超时时间，可能的值：
    //  0 - 默认值，3秒
    // -1 - 无超时，无限期的阻塞
    // -2 - 不进行超时设置，不调用 SetReadDeadline 方法
    ReadTimeout time.Duration

    // 把数据写入网络连接的超时时间，可能的值：
    //  0 - 默认值，3秒
    // -1 - 无超时，无限期的阻塞
    // -2 - 不进行超时设置，不调用 SetWriteDeadline 方法
    WriteTimeout time.Duration

    // 是否使用context.Context的上下文截止时间，
    // 有些情况下，context.Context的超时可能带来问题。
    // 默认不使用
    ContextTimeoutEnabled bool

    // 连接池的类型，有 LIFO 和 FIFO 两种模式，
    // PoolFIFO 为 false 时使用 LIFO 模式，为 true 使用 FIFO 模式。
    // 当一个连接使用完毕时会把连接归还给连接池，连接池会把连接放入队尾，
    // LIFO 模式时，每次取空闲连接会从"队尾"取，就是刚放入队尾的空闲连接，
    // 也就是说 LIFO 每次使用的都是热连接，连接池有机会关闭"队头"的长期空闲连接，
    // 并且从概率上，刚放入的热连接健康状态会更好；
    // 而 FIFO 模式则相反，每次取空闲连接会从"队头"取，相比较于 LIFO 模式，
    // 会使整个连接池的连接使用更加平均，有点类似于负载均衡寻轮模式，会循环的使用
    // 连接池的所有连接，如果你使用 go-redis 当做代理让后端 redis 节点负载更平均的话，
    // FIFO 模式对你很有用。
    // 如果你不确定使用什么模式，请保持默认 PoolFIFO = false
    PoolFIFO bool

    // 连接池最大连接数量，注意：这里不包括 pub/sub，pub/sub 将使用独立的网络连接
    // 默认为 10 * runtime.GOMAXPROCS
    PoolSize int

    // PoolTimeout 代表如果连接池所有连接都在使用中，等待获取连接时间，超时将返回错误
    // 默认是 1秒+ReadTimeout
    PoolTimeout time.Duration

    // 连接池保持的最小空闲连接数，它受到PoolSize的限制
    // 默认为0，不保持
    MinIdleConns int

    // 连接池保持的最大空闲连接数，多余的空闲连接将被关闭
    // 默认为0，不限制
    MaxIdleConns int

    // ConnMaxIdleTime 是最大空闲时间，超过这个时间将被关闭。
    // 如果 ConnMaxIdleTime <= 0，则连接不会因为空闲而被关闭。
    // 默认值是30分钟，-1禁用
    ConnMaxIdleTime time.Duration

    // ConnMaxLifetime 是一个连接的生存时间，
    // 和 ConnMaxIdleTime 不同，ConnMaxLifetime 表示连接最大的存活时间
    // 如果 ConnMaxLifetime <= 0，则连接不会有使用时间限制
    // 默认值为0，代表连接没有时间限制
    ConnMaxLifetime time.Duration

    // 如果你的redis服务器需要TLS访问，可以在这里配置TLS证书等信息
    // 如果配置了证书信息，go-redis将使用TLS发起连接，
    // 如果你自定义了 `Dialer` 方法，你需要自己实现网络连接
    TLSConfig *tls.Config

    // 限流器的配置，参照 `Limiter` 接口
    Limiter Limiter

    // 设置启用在副本节点只读查询，默认为false不启用
    // 参照：https://redis.io/commands/readonly
    readOnly bool
}
```

## 拓展

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
	Addr:	  "localhost:6379",
	Password: "", // 没有密码，默认值
	DB:		  0,  // 默认DB 0
})
```

同时也支持另外一种常见的连接字符串:

```go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

rdb := redis.NewClient(opt)
```



## 基本指令

### Keys():根据正则获取keys

**语法**

```sh
keys切片, err = redis实例.Keys(ctx, 正则表达式).Result()
```

**示例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)


func TestBaseCmd(t *testing.T) {
    //Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})


	//*表示获取所有的key
	//得到的时一个切片
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}
```

### Type():获取key对应值得类型

`Type()`方法用户获取一个key对应值的类型

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)


func TestBaseCmd(t *testing.T) {
    //Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//`Type()`方法用户获取一个key对应值的类型
	//语法 redis实例.Type(ctx,键名)
	vType, err := rdb.Type(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vType) //string
}
```

### Del():删除缓存项

语法

```go
删除缓存项数,错误值 := redis实例.Del(ctx,键名1,键名2,...).Result()
```

示例

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)




func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//Del()删除缓存项，
	//语法 redis实例.Del(ctx,键名1,键名2,...)
	//返回值 删除缓存的项数,错误
	n, err := rdb.Del(ctx, "key1","age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功删除了 %v 个\n", n)
}
```

### Exists():检测缓存项是否存在

`Exists()`方法用于检测某个key是否存在

>注：Exists()方法可以传入多个key,返回的第一个结果表示存在的key的数量,不过工作中我们一般不同时判断多个key是否存在，一般就判断一个key,所以判断是否大于0即可，如果判断判断传入的多个key是否都存在，则返回的结果的值要和传入的key的数量相等

语法

```go
存在key的数量,错误值 := redis实例.Exists(ctx,键名1,键名2,...).Result()
```

示例

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)




func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//Exists():检测缓存项是否存在
	n, err := rdb.Exists(ctx, "age").Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
```

### Expire(),ExpireAt():设置有效期

需要在设置好了缓存项后，在设置有效期

`Expire()`方法是设置某个时间段(time.Duration)后过期，`ExpireAt()`方法是在某个时间点(time.Time)过期失效

**语法**

```go
是否设置成功,错误值 := redis实例.Expire(ctx,键名,过期时间段).Result()
是否设置成功,错误值 := redis实例.ExpireAt(ctx,键名,过期时间点).Result()

区别:
Expire()  相对时间，例如2min后过期，那么就是相对设定的时间2min中后过期，例如设定时间为14:00,那么过期时间为14:02
ExpireAt() 绝对时间，设定一个具体的时间点，例如明天14:02,那么就是明天14:02过期
```

**举例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)




func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//设置有效时间段
	//在某个时间段后过期，例如 2分钟后过期
	//语法 redis实例.Expire(ctx,键名,有效期).Result()
	res, err := rdb.Expire(ctx, "name", time.Minute * 2).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
	
	//设置有效时间点
	//在某个时间点过期，例如 现在过期
	//语法 redis实例.Expire(ctx,键名,过期时间点).Result()
	res, err = rdb.ExpireAt(ctx, "age", time.Now()).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
}
```

### TTL(),PTTL():获取有效期

都可以获取某个键的剩余有效期

**语法**

```go
剩余有效期,err = redis实例.TTL(ctx,键名).Result()
剩余有效期,err = redis实例.PTTL(ctx,键名).Result()

区别:
TTL获取的是秒，PTTL获取的是毫秒
```

**举例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//获取剩余有效期,单位:秒(s)
	//语法 redis实例.TTL(ctx,键名)
	ttl, err := rdb.TTL(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ttl)

	//获取剩余有效期,单位:毫秒(ms)
	//语法 redis实例.PTTL(ctx,键名)
	pttl, err := rdb.PTTL(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pttl)
}
```

**对于值的解释**

```sh
[-2ns]
键已过期，或者不存在

[-1ns]
永久有效，即我们没有设置过期时间

[2m0s]
剩余有效时间为2m0s
所以获取的不是时间戳，而是redis.DurationCmd
```

通过查看源码来获得值

```go
type DurationCmd struct {
	baseCmd

	val       time.Duration
	precision time.Duration
}

var _ Cmder = (*DurationCmd)(nil)

func NewDurationCmd(ctx context.Context, precision time.Duration, args ...interface{}) *DurationCmd {
	return &DurationCmd{
		baseCmd: baseCmd{
			ctx:  ctx,
			args: args,
		},
		precision: precision,
	}
}

func (cmd *DurationCmd) SetVal(val time.Duration) {
	cmd.val = val
}

func (cmd *DurationCmd) Val() time.Duration {
	return cmd.val
}

func (cmd *DurationCmd) Result() (time.Duration, error) {
	return cmd.val, cmd.err
}

func (cmd *DurationCmd) String() string {
	return cmdString(cmd, cmd.val)
}

func (cmd *DurationCmd) readReply(rd *proto.Reader) error {
	n, err := rd.ReadInt()
	if err != nil {
		return err
	}
	switch n {
	// -2 if the key does not exist
	// -1 if the key exists but has no associated expire
	case -2, -1:
		cmd.val = time.Duration(n)
	default:
		cmd.val = time.Duration(n) * cmd.precision
	}
	return nil
}
```

所以我们可以通过以下方法获得值

```go
//此时得到的ttl是redis.Duration
ttl, err := rdb.TTL(ctx, "name")

//此时得到的ttl是time.Duration
ttl, err := rdb.TTL(ctx, "name").Result()
```

我们想转为数字只需要考虑time.Duration的处理

**过期后的结果**

在redis中如果一个键过期了，那么他的值将变成`<nil>`，实际上他已经从内存中移除了，所以Keys()获取不到他



### DBSize():查看当前数据库key的数量

语法

```sh
key的数量,err = redis实例.DBSize(ctx).Result()
```

示例

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)




func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//查看当前数据库key的数量
	//语法: redis实例.DBSize(ctx)
	num, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("数据库有 %v 个缓存项\n", num)
}
```

### FlushDB():清空当前数据

语法

```go
//清空当前数据库，例如连接的是索引为0的数据库，那么清空的就是0号数据库
是否清除成功,err = redis实例.FlushDB(ctx).Result()
```

示例

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)


func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	 //清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
	 //语法 redis实例.FlushDB(ctx)
	 res, err := rdb.FlushDB(ctx).Result()
	 if err != nil {
		 panic(err)
	 }
	 fmt.Println(res)//OK
}
```

### FlushAll():清空所有数据库

语法

```go
是否清除成功,err = redis实例.FlushAll(ctx).Result()
```

示例

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	//清空该连接上所有数据库
	// 语法 redis实例.FlushAll(ctx)
	res, err := rdb.FlushAll(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)//OK
}
```

## 字符串(string)类型

### Set():设置

仅仅支持字符串(包含数字)操作，不支持内置数据编码功能。

如果需要存储Go的非字符串类型，需要提前手动序列化，获取时再反序列化。

**语法**

```go
//如果过期时间段设置为0代表永不过期
错误值 = redis实例.Set(ctx,键名,键值,过期时间段).Err()
```

**示例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisString(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)

	//Set设置
	//语法 redis实例.Set(ctx,键名,键值,过期时间段)
	//Set方法的最后一个参数表示过期时间，0表示永不过期
	err = rdb.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	//key2将会在两分钟后过期失效
	err = rdb.Set(ctx, "key2", "value2", time.Minute * 2).Err()
	if err != nil {
		panic(err)
	}

}
```

### SetNX():设置并指定过期时间

设置键的同时，设置过期时间

**语法**

```sh
错误值 = redis实例.SetNX(ctx,键名,键值,过期时间段).Err()
```

**区别**

之前版本还有一个`SetEX`,功能设置键的同时设置过期时间，必传过期时间，由于被`Set`功能囊括，所以废弃了

>注：SetNX()与Set()的区别是，SexNX()仅当key不存在的时候才设置，如果key已经存在则不做任何操作，而Set()方法不管该key是否已经存在缓存中直接覆盖

介绍了`SetNX()`与`Set()`的区别后，还有一点需要注意的时候，如果我们想知道我们调用SetNX()是否设置成功了，可以接着调用Result()方法，返回的第一个值表示是否设置成功了，如果返回false,说明缓存Key已经存在，此次操作虽然没有错误，但是是没有起任何效果的。如果返回true，表示在此之前key是不存在缓存中的，操作是成功的

**示例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)



func TestRedisString(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)

	//SetNX设置
	//设置键的同时设置过期时间
	//语法 redis实例.SetEX(ctx,键名,键值,过期时间段)
	err = rdb.SetNX(ctx, "key3", "value", time.Hour * 2).Err()
	if err != nil {
		panic(err)
	}

}
```

### Get():获取

如果要获取的key在缓存中并不存在，`Get()`方法将会返回`redis.Nil`

**语法**

```go
值,错误 = redis实例.Get(ctx,键名).Result()
```

**举例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)



func TestRedisString(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)

	//Get获取
	//语法 redis实例.Get(ctx,键名)
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key1: %v\n", val)
	
	val2, err := rdb.Get(ctx, "key-not-exist").Result()
	if err == redis.Nil {
		fmt.Println("key不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("值为: %v\n", val2)
	}
}
```

### GetRange():字符串截取

`GetRange()`方法可以用来截取字符串的部分内容,第二个参数是下标索引的开始位置，第三个参数是下标索引的结束位置(不是要截取的长度) 

>注：即使key不存在，调用GetRange()也不会报错，只是返回的截取结果是空"",可以使用`fmt.Printf("%q\n", val)`来打印测试

**语法**

```sh
值,错误 = redis实例.GetRange(ctx,键名,开始下标,结束下标).Result()
```

**举例**

```go
package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)



func TestRedisString(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)

	//GetRange获取
	//语法 redis实例.GetRange(ctx,键名,开始下标,结束下标)
	val, err := rdb.GetRange(ctx, "key1",1,3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key1: %v\n", val)
}
```

### Incr():增加+1

`Incr()`、`IncrBy()`都是操作数字，对数字进行增加的操作，incr是执行`原子`加1操作，incrBy是增加指定的数

所谓原子操作是指不会被线程调度机制打断的操作：这种操作一旦开始，就一直运行到结束，中间不会有任何context witch(切换到另一个线程).

(1)在单线程中，能够在单条指令中完成的操作都可以认为是“原子操作”，因为中断只能发生于指令之间。

(2)在多线程中，不能被其它进程(线程)打算的操作叫原子操作。

Redis单命令的原子性主要得益于Redis的单线程

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.Incr(ctx, "number").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key当前的值为: %v\n", val)
}
```

### IncrBy():按指定步长增加

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.IncrBy(ctx, "number", 12).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key当前的值为: %v\n", val)
}
```

### Decr():减少-1

`Decr()`和`DecrBy()`方法是对数字进行减的操作，和Incr正好相反

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.Decr(ctx, "number").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key当前的值为: %v\n", val)
}
```

### DecrBy():按指定步长减少

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.DecrBy(ctx, "number", 12).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key当前的值为: %v\n", val)
}
```

### Append():追加

`Append()`表示往字符串后面追加元素，返回值是字符串的总长度

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "hello", 0).Err()
	if err != nil {
		panic(err)
	}
	length, err := rdb.Append(ctx, "key", " world!").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("当前缓存key的长度为: %v\n", length) //12
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("当前缓存key的值为: %v\n", val) //hello world!
}
```

### StrLen():获取长度

`StrLen()`方法可以获取字符串的长度

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "hello world!", 0).Err()
	if err != nil {
		panic(err)
	}
	length, err := rdb.StrLen(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("当前缓存key的长度为: %v\n", length) //12
}
```

如上所述都是常用的字符串操作，此外，字符串(string)类型还有`MGet()`、`Mset()`、`MSetNX()`等同时操作多个key的方法，

## 列表(list)类型

### LPush():将元素压入链表

可以使用`LPush()`方法将数据从左侧压入链表

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//返回值是当前列表元素的数量
	n, err := rdb.LPush(ctx, "list", 1, 2, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
```

也可以从右侧压如链表，对应的方法是`RPush()`

### LInsert():在某个位置插入新元素

位置的判断，是根据相对的参考元素判断

示例：

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//在名为key的缓存项值为100的元素前面插入一个值，值为123
	err := rdb.LInsert(ctx, "key", "before", "100", 123).Err()
	if err != nil {
		panic(err)
	}
}
```

> 注：即使key列表里有多个值为100的元素，也只会在第一个值为100的元素前面插入123，并不会在所有值为100的前面插入123,客户端还提供了从前面插入和从后面插入的LInsertBefore()和LInsertAfer()方法

### LSet():设置某个元素的值

示例：

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//下标是从0开始的
	err := rdb.LSet(ctx, "list", 1, 100).Err()
	if err != nil {
		panic(err)
	}
}
```

### LLen():获取链表元素个数

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	length, err := rdb.LLen(ctx, "list").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("当前链表的长度为: %v\n", length)
}
```

### LIndex():获取链表下标对应的元素

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.LIndex(ctx, "list", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("下标为0的值为: %v\n", val)
}
```

### LRange():获取某个选定范围的元素集

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	vals, err := rdb.LRange(ctx, "list", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("从下标0到下标2的值: %v\n", vals)
}
```

### 从链表左侧弹出数据

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.LPop(ctx, "list").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("移除的元素为: %v\n", val)
}
```

与之相对的，从右侧弹出数据的方法为`RPop()`

### LRem():根据值移除元素

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	n, err := rdb.LRem(ctx, "list", 2, "100").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("移除了: %v 个\n", n)
}
```

## 集合(set)类型

Redis set对外提供的功能与list类似是一个列表的功能，特殊之处在于set是可以自动排重的，当你需要存储一个列表数据，又不希望出现重复数据，set是一个很好的选择，并且set提供了判断某个成员是否在一个set集合内的接口，这是也是list所不能提供了。

Redis的Set是string类型的无序集合。它底层其实是一个value为null的hash表，所以添加、删除、查找的复杂度都是O(1)。

集合数据的特征：

1. 元素不能重复，保持唯一性
2. 元素无序，不能使用索引(下标)操作

### SAdd():添加元素

示例：

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.SAdd(ctx, "team", "kobe", "jordan")
	rdb.SAdd(ctx, "team", "curry")
	rdb.SAdd(ctx, "team", "kobe") //由于kobe已经被添加到team集合中，所以重复添加时无效的
}
```

### SPop():随机获取一个元素

无序性，是随机的

`SPop()`方法是从集合中随机取出元素的，如果想一次获取多个元素，可以使用`SPopN()`方法

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	val, err := rdb.SPop(ctx, "team").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
```

### SRem():删除集合里指定的值

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	n, err := rdb.SRem(ctx, "team", "kobe", "v2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
```

### SSMembers():获取所有成员

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	vals, err := rdb.SMembers(ctx, "team").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)
}
```

### SIsMember():判断元素是否在集合中

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	exists, err := rdb.SIsMember(ctx, "team", "jordan").Result()
	if err != nil {
		panic(err)
	}
	if exists {
		fmt.Println("存在集合中")
	} else {
		fmt.Println("不存在集合中")
	}
}
```

### SCard():获取集合元素个数

获取集合中元素个数

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	total, err := rdb.SCard(ctx, "team").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("集合总共有 %v 个元素\n", total)
}
```

### SUnion():并集,SDiff():差集,SInter():交集

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.SAdd(ctx, "setA", "a", "b", "c", "d")
	rdb.SAdd(ctx, "setB", "a", "d", "e", "f")

	//并集
	union, err := rdb.SUnion(ctx, "setA", "setB").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(union)

	//差集
	diff, err := rdb.SDiff(ctx, "setA", "setB").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(diff)

	//交集
	inter, err := rdb.SInter(ctx, "setA", "setB").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(inter)
}
```

## 有序集合(zset)类型

Redis 有序集合和集合一样也是string类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。

有序集合的成员是唯一的,但`分数(score)`却可以重复。

集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是O(1)。 集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。

### ZAdd():添加元素

添加6个元素1~6,分值都是0

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 1,
	})
	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 2,
	})
	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 3,
	})
	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 4,
	})
	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 5,
	})
	rdb.ZAdd(ctx, "zSet", &redis.Z{
		Score: 0,
		Member: 6,
	})
}
```

### ZIncrBy():增加元素分值

分值可以为负数，表示递减

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math/rand"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "1")
	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "2")
	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "3")
	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "4")
	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "5")
	rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "6")
}
```

### ZRange()、ZRevRange():获取根据score排序后的数据段

根据分值排序后的，升序和降序的列表获取

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//获取排行榜
	//获取分值(点击率)前三的文章ID列表
	res, err := rdb.ZRevRange(ctx, "zSet", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### ZRangeByScore()、ZRevRangeByScore():获取score过滤后排序的数据段

根据分值过滤之后的列表

需要提供分值区间

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	res, err := rdb.ZRangeByScore(ctx, "zSet", &redis.ZRangeBy{
		Min:    "40",
		Max:    "85",
	}).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### ZCard():获取元素个数

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	count, err := rdb.ZCard(ctx, "zSet").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
```

### ZCount():获取区间内元素个数

获取分值在[40, 85]的元素的数量

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	n, err := rdb.ZCount(ctx, "zSet", "40", "85").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
```

### ZScore():获取元素的score

获取元素分值

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	score, err := rdb.ZScore(ctx, "zSet", "5").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(score)
}
```

### ZRank()、ZRevRank():获取某个元素在集合中的排名

`ZRank()`方法是返回元素在集合中的升序排名情况，从0开始。`ZRevRank()`正好相反

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	res, err := rdb.ZRevRank(ctx, "zSet", "2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### ZRem():删除元素

`ZRem()`方法支持通过元素的值来删除元素

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

    //通过元素值来删除元素
	res, err := rdb.ZRem(ctx, "zSet", "2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### ZRemRangeByRank():根据排名来删除

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//按照升序排序删除第一个和第二个元素
	res, err := rdb.ZRemRangeByRank(ctx, "zSet",  0, 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### ZRemRangeByScore():根据分值区间来删除

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	//删除score在[40, 70]之间的元素
	res, err := rdb.ZRemRangeByScore(ctx, "zSet", "40", "70").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

## 哈希(hash)类型

Redis hash 是一个 string 类型的 field（字段） 和 value（值） 的映射表，hash 特别适合用于存储对象。

Redis 中每个 hash 可以存储 232 - 1 键值对（40多亿）。

当前服务器一般都是将用户登录信息保存到Redis中，这里存储用户登录信息就比较适合用hash表。hash表比string更合适，如果我们选择使用string类型来存储用户的信息的时候，我们每次存储的时候就得先序列化(json_encode()、serialize())成字符串才能存储redis,

从redis拿到用户信息后又得反序列化(UnMarshal()、Marshal())成数组或对象，这样开销比较大。如果使用hash的话我们通过key(用户ID)+field(属性标签)就可以操作对应属性数据了，既不需要重复存储数据，也不会带来序列化和并发修改控制的问题。

### HSet():设置

`HSet()`方法支持如下格式

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.HSet(ctx, "user", "key1", "value1", "key2", "value2")
	rdb.HSet(ctx, "user", []string{"key3", "value3", "key4", "value4"})
	rdb.HSet(ctx, "user", map[string]interface{}{"key5": "value5", "key6": "value6"})
}
```

### HMset():批量设置

示例：

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	rdb.Del(ctx, "user")
	rdb.HMSet(ctx, "user", map[string]interface{}{"name":"kevin", "age": 27, "address":"北京"})
}
```

### HGet():获取某个元素

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	address, err := rdb.HGet(ctx, "user", "address").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
```

### HGetAll():获取全部元素

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	user, err := rdb.HGetAll(ctx, "user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
```

### HDel():删除某个元素

`HDel()`支持一次删除多个元素

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	res, err := rdb.HDel(ctx, "user", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### HExists():判断元素是否存在

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	res, err := rdb.HExists(ctx, "user", "address").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

### HLen():获取长度

示例：

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})

	res, err := rdb.HLen(ctx, "user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```



