1、现象：

 在go中gin框架中，需要接收前端参数时，参数必填，我们一般添加binding:"required"`标签，这样前端参数不给时，gin框架会自动校验，给出error。

 gin的参数校验是基于validator的，如果给了required标签，则不能传入零值，比如字符串的不能传入空串，int类型的不能传入0，bool类型的不能传入false。

 有时候我们需要参数必填，而且需要可以传入零值。比如性别sex有0和1表示，0表示女，1表示男，而且需要必填。这个时候，我们可以通过定义int类型的指针解决该问题。同理，其他类型也是定义指针即可。

2、参考例子

（1）不能接受零值的例子

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "log"
)

// Student
// 假设前端需要传入参数name,age,sex
// name表示学生的姓名(必填)
// age表示学生的年龄(必填)
// sex表示学生的性别,0是女,1是男(必填)
type Student struct {
   Name string `json:"name" binding:"required"`
   Age  int    `json:"age" binding:"required"`
   Sex  int    `json:"sex" binding:"required"`
}

func main() {
   engine := gin.Default() //创建一个默认的引擎

   //请求 http://127.0.0.1:8090/student/add
   engine.POST("/student/add", func(context *gin.Context) {
      student := Student{}
      err := context.ShouldBind(&student)
      if err != nil {
         log.Printf("参数错误:%v", err)
         context.JSON(200, gin.H{"success": false, "msg": err.Error()})
         return
      }
      context.JSON(200, gin.H{"success": true, "msg": ""})
      return
   })

   err := engine.Run("0.0.0.0:8090") //启动引擎,端口是8090
   if err != nil {
      panic(fmt.Sprintf("启动引擎失败,失败信息:%s", err.Error()))
   }
}
```

前端传入参数，性别为0，则报error

```sh
Key: 'Student.Sex' Error:Field validation for 'Sex' failed on the 'required' tag
```

(2)可以接受零值的例子

只需要把

```go
Sex  int    json:"sex" binding:"required"
```

改成

```go
Sex  *int    json:"sex" binding:"required"
```

例如

```go
package main

import (
   "fmt"
   "github.com/gin-gonic/gin"
   "log"
)

// Student
// 假设前端需要传入参数name,age,sex
// name表示学生的姓名(必填)
// age表示学生的年龄(必填)
// sex表示学生的性别,0是女,1是男(必填)
type Student struct {
   Name string `json:"name" binding:"required"`
   Age  int    `json:"age" binding:"required"`
   Sex  *int    `json:"sex" binding:"required"`
}

func main() {
   engine := gin.Default() //创建一个默认的引擎

   //请求 http://127.0.0.1:8090/student/add
   engine.POST("/student/add", func(context *gin.Context) {
      student := Student{}
      err := context.ShouldBind(&student)
      if err != nil {
         log.Printf("参数错误:%v", err)
         context.JSON(200, gin.H{"success": false, "msg": err.Error()})
         return
      }
      context.JSON(200, gin.H{"success": true, "msg": ""})
      return
   })

   err := engine.Run("0.0.0.0:8090") //启动引擎,端口是8090
   if err != nil {
      panic(fmt.Sprintf("启动引擎失败,失败信息:%s", err.Error()))
   }
}
```

注意在插入数据库和使用该值时也需使用指针形式，例如

```go
*student.sex
```

