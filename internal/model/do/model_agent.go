package do

import "github.com/gogf/gf/v2/util/gmeta"

type ModelAgent struct {
	gmeta.Meta           `collection:"model_agent" bson:"-"`
	ProviderId           string   `bson:"provider_id,omitempty"`   // 提供商ID
	Name                 string   `bson:"name,omitempty"`          // 模型代理名称
	BaseUrl              string   `bson:"base_url,omitempty"`      // 模型代理地址
	Path                 string   `bson:"path"`                    // 模型代理地址路径
	Weight               int      `bson:"weight"`                  // 权重
	Models               []string `bson:"models"`                  // 绑定模型
	IsEnableModelReplace bool     `bson:"is_enable_model_replace"` // 是否启用模型替换
	ReplaceModels        []string `bson:"replace_models"`          // 替换模型
	TargetModels         []string `bson:"target_models"`           // 目标模型
	IsNeverDisable       bool     `bson:"is_never_disable"`        // 是否永不禁用
	LbStrategy           int      `bson:"lb_strategy,omitempty"`   // 密钥负载均衡策略[1:轮询, 2:权重]
	Remark               string   `bson:"remark"`                  // 备注
	Status               int      `bson:"status,omitempty"`        // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled       bool     `bson:"is_auto_disabled"`        // 是否自动禁用
	AutoDisabledReason   string   `bson:"auto_disabled_reason"`    // 自动禁用原因
	Creator              string   `bson:"creator,omitempty"`       // 创建人
	Updater              string   `bson:"updater,omitempty"`       // 更新人
	CreatedAt            int64    `bson:"created_at,omitempty"`    // 创建时间
	UpdatedAt            int64    `bson:"updated_at,omitempty"`    // 更新时间
}
