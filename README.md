# Simplelog
**一个简单的结构化日志**

## 支持特性
+  **JSON 结构数据**
+  **Level输出**
+  **Hook勾子**
+  **日志文件大小设置，滚动更新**
+  **直接写入或缓冲写入**
## 未支持特性
+  **HTTP更新Level**
## 使用
```html
  get get github.com/tanzy2018/simplelog
```

- 基本使用
```go

package main
import (
	"github.com/tanzy2018/simplelog"
	"github.com/tanzy2018/simplelog/encode"
)

func main() {
	newLog := simplelog.New()
	defer newLog.Sync()
	newLog.Info("profile", encode.String("name", "Tom"), encode.Int("id", 10), encode.Any("flag", false))
}
// 输出:
/*
{"time":"2020-08-11 10:28:18","level":"info","msg":"profile","name":"Tom","id":10,"flag":false}
*/
    
  ```

  - 使用文件输出
```go
package main

import (
	"github.com/tanzy2018/simplelog"
	"github.com/tanzy2018/simplelog/encode"
)

func main() {
	newLog := simplelog.New().WithFileWriter("", "", "output.txt")
	defer newLog.Sync()
	newLog.Info("profile", encode.String("name", "Tom"), encode.Int("id", 10), encode.Any("flag", false))
}

// output 
/*simplelog.txt
{"time":"2020-08-11 11:00:32","level":"info","msg":"profile","name":"Tom","id":10,"flag":false}
*/

```

- 使用勾子 Hook
```go

package main

import (
	"github.com/tanzy2018/simplelog"
	"github.com/tanzy2018/simplelog/encode"
)

func main() {
	newLog := simplelog.New()
	newLog.Hook(func() encode.Meta {
		return encode.Any("map", map[string]interface{}{"fruit": []string{"apple", "peach"}})
	})
	newLog.Hook(func() encode.Meta {
		return encode.Bools("flags", []bool{false, true, false})
	})
	defer newLog.Sync()
	newLog.Info("profile", encode.String("name", "Tom"))
	newLog.Info("profile", encode.String("name", "Jeiry"), encode.Int("say", 2))
}
// output

/*
{"time":"2020-08-11 11:13:19","level":"info","msg":"profile","map":{"fruit":["apple","peach"]},"flags":[false,false,true,false],"name":"Tom"}
{"time":"2020-08-11 11:13:19","level":"info","msg":"profile","map":{"fruit":["apple","peach"]},"flags":[false,false,true,false],"name":"Jeiry","say":2}
*/

```