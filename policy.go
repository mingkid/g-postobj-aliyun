package postobj

import (
	"encoding/json"
	"time"
)

// Policy 是Post请求的Policy表单域，用于验证请求的合法性。声明了Post请求必须满足的条件。强烈建议使用该域来限制Post请求
type Policy struct {
	Expiration time.Time        `json:"expiration"`
	Conditions []json.Marshaler `json:"conditions"`
}

func (p *Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"expiration": p.Expiration.Format("2006-01-02T15:04:05Z"),
		"conditions": p.Conditions,
	})
}
