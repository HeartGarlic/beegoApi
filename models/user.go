package models

import (
	"beegoApi/helper"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User 用户表
type User struct {
	ID             int       `orm:"column(id)" json:"id"`
	GroupId        int       `orm:"column(group_id)" json:"group_id"`                 // 组别ID
	Username       string    `orm:"column(username)" json:"username" form:"username"` // 用户名
	Nickname       string    `orm:"column(nickname)" json:"nickname" form:"nickname"` // 昵称
	Password       string    `orm:"column(password)" json:"password" form:"password"` // 密码
	Salt           string    `orm:"column(salt)" json:"salt"`                         // 密码盐
	Email          string    `orm:"column(email)" json:"email" form:"email"`          // 邮箱
	Mobile         string    `orm:"column(mobile)" json:"mobile" form:"mobile"`       // 手机号
	Avatar         string    `orm:"column(avatar)" json:"avatar" form:"avatar"`       // 头像
	Level          int       `orm:"column(level)" json:"level"`                       // 等级
	Gender         int       `orm:"column(gender)" json:"gender"`                     // 性别
	Birthday       time.Time `orm:"column(birthday)" json:"birthday"`                 // 生日
	Bio            string    `orm:"column(bio)" json:"bio"`                           // 格言
	Money          float64   `orm:"column(money)" json:"money"`                       // 余额
	Score          int       `orm:"column(score)" json:"score"`                       // 积分
	Successions    int       `orm:"column(successions)" json:"successions"`           // 连续登录天数
	MaxSuccessions int       `orm:"column(maxsuccessions)" json:"maxsuccessions"`     // 最大连续登录天数
	PrevTime       int       `orm:"column(prevtime)" json:"prevtime"`                 // 上次登录时间
	LoginTime      int64     `orm:"column(logintime)" json:"logintime"`               // 登录时间
	LoginIp        string    `orm:"column(loginip)" json:"loginip"`                   // 登录IP
	LoginFailure   int       `orm:"column(loginfailure)" json:"loginfailure"`         // 失败次数
	JoinIp         string    `orm:"column(joinip)" json:"joinip"`                     // 加入IP
	JoinTime       int64     `orm:"column(jointime)" json:"jointime"`                 // 加入时间
	CreateTime     int64     `orm:"column(createtime)" json:"createtime"`             // 创建时间
	UpdateTime     int64     `orm:"column(updatetime)" json:"updatetime"`             // 更新时间
	Token          string    `orm:"column(token)" json:"token"`                       // Token
	Status         int       `orm:"column(status)" json:"status"`                     // 状态
	Verification   string    `orm:"column(verification)" json:"verification"`         // 验证
}

// AesKey Aes 加密的key
const AesKey  = "#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR"

// NewUser 构造函数
func NewUser() *User {
	return &User{}
}

// TableName 表名
func (u *User) TableName() string {
	return "fa_user"
}

// Register 注册 返回注册的用户id
func (u *User) Register() (id int64, err error) {
	// 校验用户名 手机号 唯一
	if err = u.CheckUsername(); err != nil {
		return 0, err
	}
	if err = u.CheckMobile(); err != nil {
		return 0, err
	}
	// 密码加密
	err = u.PasswordEncrypt()
	if err != nil {
		return 0, err
	}
	// 插入数据
	// 初始化基础的字段
	u.JoinTime = time.Now().Unix()
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	id, err = orm.NewOrm().Insert(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Login 登陆
func (u *User) Login(Username string, Password string) (user *User, err error) {
	// 查询用户是否存在
	o         := orm.NewOrm()
	u.Username = Username
	err        = o.Read(u, "Username")
	if err != nil {
		return nil, errors.New("用户名或密码不正确1")
	}
	// 判断密码是否正确
	err = u.PasswordDecrypt(Password)
	if err != nil {
		return nil, errors.New("用户名或密码不正确2")
	}
	// 需要更新用户的登陆时间和登陆次数
	u.UpdateTime = time.Now().Unix()
	u.LoginTime  = time.Now().Unix()
	_, err = o.Update(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// CheckMobile 校验手机号是否已存在
func (u *User) CheckMobile() error {
	o := orm.NewOrm()
	exist := o.QueryTable(u.TableName()).Filter("mobile", u.Mobile).Exist()
	if exist {
		return errors.New("手机号已存在")
	}
	return nil
}

// CheckUsername 校验用户名
func (u *User) CheckUsername() error {
	o := orm.NewOrm()
	exist := o.QueryTable(u.TableName()).Filter("username", u.Username).Exist()
	if exist {
		return errors.New("用户名已存在")
	}
	return nil
}

// PasswordDecrypt 密码解密
func (u *User) PasswordDecrypt(Plaintext string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(Plaintext))
	if err != nil {
		return err
	}
	return nil
}

// PasswordEncrypt 密码加密
func (u *User) PasswordEncrypt() error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}

// GenerateToken 生成用户token 使用 aes 对称加密
func (u *User) GenerateToken() (string, error) {
	// 生成token数据
	tokenData             := make(map[string]interface{})
	tokenData["id"]        = u.ID
	tokenData["username"]  = u.Username
	tokenData["mobile"]    = u.Mobile
	tokenData["timestamp"] = time.Now().Unix()
	text, _ := json.Marshal(tokenData)                           // 你要加密的数据
	 // 对称秘钥长度必须是16的倍数
	encrypted, err := helper.AresHelper{}.AesEncrypt(text, []byte(AesKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// ParseToken 解析 token
func (u *User) ParseToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("token不能为空")
	}
	by, _ := base64.StdEncoding.DecodeString(token)
	token  = string(by)
	origin, err := helper.AresHelper{}.AesDecrypt([]byte(token), []byte(AesKey))
	if err != nil {
		return "", err
	}
	return string(origin), nil
}


