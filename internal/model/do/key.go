package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	KEY_COLLECTION = "key"
)

type Key struct {
	gmeta.Meta          `collection:"key" bson:"-"`
	UserId              int      `bson:"user_id,omitempty"`            // 用户ID
	AppId               int      `bson:"app_id,omitempty"`             // 应用ID
	Corp                string   `bson:"corp,omitempty"`               // 公司
	Key                 string   `bson:"key,omitempty"`                // 密钥
	Type                int      `bson:"type,omitempty"`               // 密钥类型[1:应用, 2:模型]
	Weight              int      `bson:"weight"`                       // 权重
	Models              []string `bson:"models"`                       // 模型
	ModelAgents         []string `bson:"model_agents"`                 // 模型代理
	IsAgentsOnly        bool     `bson:"is_agents_only"`               // 是否代理专用
	IsNeverDisable      bool     `bson:"is_never_disable"`             // 是否永不禁用
	IsLimitQuota        bool     `bson:"is_limit_quota"`               // 是否限制额度
	Quota               int      `bson:"quota"`                        // 剩余额度
	UsedQuota           int      `bson:"used_quota,omitempty"`         // 已用额度
	QuotaExpiresRule    int      `bson:"quota_expires_rule,omitempty"` // 额度过期规则[1:固定, 2:时长]
	QuotaExpiresAt      int64    `bson:"quota_expires_at"`             // 额度过期时间
	QuotaExpiresMinutes int64    `bson:"quota_expires_minutes"`        // 额度过期分钟数
	IpWhitelist         []string `bson:"ip_whitelist"`                 // IP白名单
	IpBlacklist         []string `bson:"ip_blacklist"`                 // IP黑名单
	Remark              string   `bson:"remark"`                       // 备注
	Status              int      `bson:"status,omitempty"`             // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled      bool     `bson:"is_auto_disabled"`             // 是否自动禁用
	AutoDisabledReason  string   `bson:"auto_disabled_reason"`         // 自动禁用原因
	Creator             string   `bson:"creator,omitempty"`            // 创建人
	Updater             string   `bson:"updater,omitempty"`            // 更新人
	CreatedAt           int64    `bson:"created_at,omitempty"`         // 创建时间
	UpdatedAt           int64    `bson:"updated_at,omitempty"`         // 更新时间
}
