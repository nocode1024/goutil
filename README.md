# Nocode GoLang Utils

## Quick start

### ncjson
> 简单地从JSON中获取数据，path参考：[GJSON](https://github.com/tidwall/gjson)

#### 导入
```go
import "github.com/nocode1024/goutil/ncjson"
```

#### 使用示例
```go
title1 := ncjson.GetStrB(jsonData, "Title.0")
title2 := ncjson.GetStrB(string(jsonData), "Title.1")
```