package postobj

import (
	"encoding/json"
	"fmt"
)

type ConditionKey string

const (
	ConditionObjName            ConditionKey = "key"                     // Object名称
	ConditionSuccessRedirect                 = "success_action_redirect" // 上传成功后的跳转URL地址
	ConditionSuccessStatus                   = "success_action_status"   // 上传成功后的HTTP状态码
	ConditionUserMetaFmt                     = "x-oss-meta-%s"           // 用户自定义元数据的前缀
	ConditionContentType                     = "content-type"            // 上传Object的Content-Type
	ConditionContentDisposition              = "content-disposition"     // 上传Object的Content-Disposition
	ConditionContentEncoding                 = "content-encoding"        // 上传Object的Content-Encoding
	ConditionCacheControl                    = "cache-control"           // 上传Object的Cache-Control
	ConditionExpires                         = "expires"                 // 上传Object的过期时间
)

// ConditionKeyF 是用于格式化ConditionKey
func (c ConditionKey) ConditionKeyF(args ...string) ConditionKey {
	return ConditionKey(fmt.Sprint(c, args))
}

// ContentLengthRange 是允许上传的文件最大和最小范围，单位为字节
type ContentLengthRange struct {
	Min uint64
	Max uint64
}

func (c ContentLengthRange) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{"content-length-range", c.Min, c.Max})
}

// Eq 是表单域的值必须精确匹配声明的值
type Eq struct {
	Key ConditionKey
	Val string
}

func (c Eq) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{"eq", "$" + string(c.Key), c.Val})
}

// StartsWith 是表单域的值必须以指定前缀开始
type StartsWith struct {
	Key ConditionKey
	Val string
}

func (c StartsWith) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{"starts-with", "$" + string(c.Key), c.Val})
}

// In 是以字符串列表的形式指定需包含的检查元素
type In struct {
	Key ConditionKey
	Val []string
}

func (c In) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{"in", "$" + string(c.Key), c.Val})
}

// NotIn 是以字符串列表的形式指定需排除的检查元素
type NotIn struct {
	Key ConditionKey
	Val []string
}

func (c NotIn) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{"not-in", "$" + string(c.Key), c.Val})
}
