package aliyun



// 阿里云访问秘钥
type AliYunAccessKey struct {
	AccessKeyID string `yaml:"AccessKeyID,omitempty"`
	AccessKeySecret string `yaml:"AccessKeySecret,omitempty"`
}

type AliYunOssConf struct {
	HostOuter string `yaml:"HostOuter,omitempty"`   // 外网访问地址
	HostIn string `yaml:"HostIn,omitempty"`   // 内网访问地址
	CallbackUrl string `yaml:"CallbackUrl,omitempty"`   // oss 回调
	ExpireTime int64 `yaml:"ExpireTime,omitempty"`
	UploadDir string `yaml:"UploadDir,omitempty"`
}





