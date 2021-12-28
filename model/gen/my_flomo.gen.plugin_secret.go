package gen

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PluginSecretMgr struct {
	*_BaseMgr
}

// PluginSecretMgr open func
func PluginSecretMgr(db *gorm.DB) *_PluginSecretMgr {
	if db == nil {
		panic(fmt.Errorf("PluginSecretMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PluginSecretMgr{_BaseMgr: &_BaseMgr{DB: db.Table("plugin_secret"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PluginSecretMgr) GetTableName() string {
	return "plugin_secret"
}

// Reset 重置gorm会话
func (obj *_PluginSecretMgr) Reset() *_PluginSecretMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PluginSecretMgr) Get() (result PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PluginSecretMgr) Gets() (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PluginSecretMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PluginSecretMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取
func (obj *_PluginSecretMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithUserSecret user_secret获取
func (obj *_PluginSecretMgr) WithUserSecret(userSecret string) Option {
	return optionFunc(func(o *options) { o.query["user_secret"] = userSecret })
}

// GetByOption 功能选项模式获取
func (obj *_PluginSecretMgr) GetByOption(opts ...Option) (result PluginSecret, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PluginSecretMgr) GetByOptions(opts ...Option) (results []*PluginSecret, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_PluginSecretMgr) GetFromID(id string) (result PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PluginSecretMgr) GetBatchFromID(ids []string) (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_PluginSecretMgr) GetFromUserID(userID string) (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_PluginSecretMgr) GetBatchFromUserID(userIDs []string) (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromUserSecret 通过user_secret获取内容
func (obj *_PluginSecretMgr) GetFromUserSecret(userSecret string) (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`user_secret` = ?", userSecret).Find(&results).Error

	return
}

// GetBatchFromUserSecret 批量查找
func (obj *_PluginSecretMgr) GetBatchFromUserSecret(userSecrets []string) (results []*PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`user_secret` IN (?)", userSecrets).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PluginSecretMgr) FetchByPrimaryKey(id string) (result PluginSecret, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PluginSecret{}).Where("`id` = ?", id).Find(&result).Error

	return
}
