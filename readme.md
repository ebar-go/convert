## 说明
提供Golang里各种类型的转换方法

## Install 
```
go get github.com/ebar-go/convert
```


## 字符串
- 转int
```go
convert.String("20210206").ToInt()
```

- 转byte
```go
convert.String("hello").ToByte()
```

- 转float
```go
convert.String("14.43").ToFloat()
```

## 整形
- 转字符串
```go
convert.Number(1).ToString()
```
- 转byte
```go
convert.Number(2).ToByte()
```

## 字节
- 转字符串
```go
convert.Byte([]byte{'h','e','l','l','o'})).ToString()
```

## 数组
- 整形转字符串数组
```go
convert.IntArray([]int{1,2,3}).ToString()
```

- 字符串转整形数组
```go
convert.StringArray([]string{"1","2","3"}).ToInt()
```

## 接口
```go
convert.Interface("20210206").ToInt()
```

- 转byte
```go
convert.Interface("hello").ToString()
```

- 转float
```go
convert.Interface(1.345).ToFloat()
```

## Json
- 转map
```go
Json(`{"name":"convert"}`).ToMap()
```

- 转struct
```go
p := new(struct{
	Name string
})
Json(`{"name":"convert"}`).ToStruct(p)
```

## HttpResponse
- 转string
```go
response, err := Response(resp)
response.ToString()
```

- 转byte
```go
response, err := Response(resp)
response.ToByte()
```