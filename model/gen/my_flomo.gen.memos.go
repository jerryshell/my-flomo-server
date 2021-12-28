package gen

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MemosMgr struct {
	*_BaseMgr
}

// MemosMgr open func
func MemosMgr(db *gorm.DB) *_MemosMgr {
	if db == nil {
		panic(fmt.Errorf("MemosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MemosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("memos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MemosMgr) GetTableName() string {
	return "memos"
}

// Reset 重置gorm会话
func (obj *_MemosMgr) Reset() *_MemosMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MemosMgr) Get() (result Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MemosMgr) Gets() (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MemosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Memos{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MemosMgr) WithID(id string) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取
func (obj *_MemosMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取
func (obj *_MemosMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取
func (obj *_MemosMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithContent content获取
func (obj *_MemosMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithUserID user_id获取
func (obj *_MemosMgr) WithUserID(userID string) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// GetByOption 功能选项模式获取
func (obj *_MemosMgr) GetByOption(opts ...Option) (result Memos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MemosMgr) GetByOptions(opts ...Option) (results []*Memos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_MemosMgr) GetFromID(id string) (result Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MemosMgr) GetBatchFromID(ids []string) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_MemosMgr) GetFromCreatedAt(createdAt time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`created_at` = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量查找
func (obj *_MemosMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容
func (obj *_MemosMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量查找
func (obj *_MemosMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容
func (obj *_MemosMgr) GetFromDeletedAt(deletedAt time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}

// GetBatchFromDeletedAt 批量查找
func (obj *_MemosMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_MemosMgr) GetFromContent(content string) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找
func (obj *_MemosMgr) GetBatchFromContent(contents []string) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_MemosMgr) GetFromUserID(userID string) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_MemosMgr) GetBatchFromUserID(userIDs []string) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MemosMgr) FetchByPrimaryKey(id string) (result Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchIndexByIDxMemosDeletedAt  获取多个内容
func (obj *_MemosMgr) FetchIndexByIDxMemosDeletedAt(deletedAt time.Time) (results []*Memos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Memos{}).Where("`deleted_at` = ?", deletedAt).Find(&results).Error

	return
}
