package gen

import (
	"time"
)

// Memos [...]
type Memos struct {
	ID        string    `gorm:"primaryKey;column:id;type:varchar(191);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3)"`
	DeletedAt time.Time `gorm:"index:idx_memos_deleted_at;column:deleted_at;type:datetime(3)"`
	Content   string    `gorm:"column:content;type:longtext"`
	UserID    string    `gorm:"column:user_id;type:longtext"`
}

// MemosColumns get sql column name.获取数据库列名
var MemosColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Content   string
	UserID    string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Content:   "content",
	UserID:    "user_id",
}

// PluginSecret [...]
type PluginSecret struct {
	ID         string `gorm:"primaryKey;column:id;type:varchar(255);not null"`
	UserID     string `gorm:"column:user_id;type:varchar(255);not null"`
	UserSecret string `gorm:"column:user_secret;type:varchar(255);not null"`
}

// PluginSecretColumns get sql column name.获取数据库列名
var PluginSecretColumns = struct {
	ID         string
	UserID     string
	UserSecret string
}{
	ID:         "id",
	UserID:     "user_id",
	UserSecret: "user_secret",
}

// User [...]
type User struct {
	ID       string `gorm:"primaryKey;column:id;type:varchar(255);not null"`
	Username string `gorm:"column:username;type:varchar(255)"`
	Email    string `gorm:"column:email;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID       string
	Username string
	Email    string
	Password string
}{
	ID:       "id",
	Username: "username",
	Email:    "email",
	Password: "password",
}
