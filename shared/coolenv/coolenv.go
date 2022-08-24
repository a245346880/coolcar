package coolenv

import "google.golang.org/protobuf/runtime/protoimpl"

const (
	// 验证生成的代码是否为最新
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	//验证运行的代码是否为最新
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
