## 阿里云 PostObject SDK

## Usage

```go
# 组装策略
policy := postobj.Policy{
    Expiration: time.Now().Add(time.Hour * 24),
    Conditions: []json.Marshaler{
        postobj.ContentLengthRange{
            Min: 0,
            Max: 1024 * 1024 * 10,
        },
    },
}

# 创建PostObject实例
postObj := postobj.New("accessKeyId", "accessKeySecret").With(policy)
```
更多Policy Conditions参考文件：[conditions.go](conditions.go)。Conditions和对应的匹配方式请参考[Post Policy](https://help.aliyun.com/document_detail/31988.html#section-d5z-1ww-wdb)