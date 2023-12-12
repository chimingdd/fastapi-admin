package entity

type User struct {
	Id        string `bson:"_id,omitempty"`        // ID
	UserId    int    `bson:"user_id,omitempty"`    // 用户ID
	Nickname  string `bson:"nickname,omitempty"`   // 昵称
	Avatar    string `bson:"avatar,omitempty"`     // 头像
	Gender    int    `bson:"gender,omitempty"`     // 性别[0:保密; 1:男; 2:女]
	Mobile    string `bson:"mobile,omitempty"`     // 手机号
	Email     string `bson:"email,omitempty"`      // 邮箱
	VipLevel  int    `bson:"vip_level,omitempty"`  // 会员等级
	Remark    string `bson:"remark,omitempty"`     // 备注
	Status    int    `bson:"status,omitempty"`     // 状态[1:正常; 2:禁用; -1:删除]
	Creator   string `bson:"creator,omitempty"`    // 创建人
	Updater   string `bson:"updater,omitempty"`    // 更新人
	CreatedAt int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `bson:"updated_at,omitempty"` // 更新时间
}

type Account struct {
	Id            string `bson:"_id,omitempty"`        // ID
	Uid           string `bson:"uid,omitempty"`        // 用户主键ID
	UserId        int    `bson:"user_id,omitempty"`    // 用户ID
	Account       string `bson:"account,omitempty"`    // 账号
	Password      string `bson:"password,omitempty"`   // 密码
	Salt          string `bson:"salt,omitempty"`       // 盐
	LastLoginIP   string `bson:"last_login_ip"`        // 最后登录IP
	LastLoginTime int64  `bson:"last_login_time"`      // 最后登录时间
	Remark        string `bson:"remark,omitempty"`     // 备注
	Status        int    `bson:"status,omitempty"`     // 状态[1:正常; 2:禁用; -1:删除]
	Creator       string `bson:"creator,omitempty"`    // 创建人
	Updater       string `bson:"updater,omitempty"`    // 更新人
	CreatedAt     int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt     int64  `bson:"updated_at,omitempty"` // 更新时间
}
