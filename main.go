package postobj

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
)

type PostObject struct {
	accessKeyId     string
	accessKeySecret string
	policy          *Policy
}

// UploadParams 直传OSS的参数
func (helper *PostObject) UploadParams() (res UploadParams) {
	res = UploadParams{
		AccessKeyId: helper.accessKeyId,
		Signature:   helper.signature(),
	}
	if helper.policy != nil {
		res.Policy = helper.policyBase64()
	}
	return
}

// With 用于设置Policy
func (helper *PostObject) With(p Policy) *PostObject {
	return &PostObject{
		accessKeyId:     helper.accessKeyId,
		accessKeySecret: helper.accessKeySecret,
		policy:          &p,
	}
}

func (helper *PostObject) policyBase64() string {
	if helper.policy == nil {
		return ""
	}
	policyTxt, err := json.Marshal(helper.policy)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(policyTxt)
}

func (helper *PostObject) signature() string {
	mac := hmac.New(sha1.New, []byte(helper.accessKeySecret))
	mac.Write([]byte(helper.policyBase64()))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func New(accessKeyId, accessKeySecret string) *PostObject {
	return &PostObject{
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
	}
}

type UploadParams struct {
	AccessKeyId string `json:"accessKeyId"`
	Policy      string `json:"policy"`
	Signature   string `json:"signature"`
}
