// 响应错误码及其对应值
package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 // 系统错误
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
	}
)
