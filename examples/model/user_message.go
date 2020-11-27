package model

import (
	"time"

	gormx "github.com/jinzhu/gorm"
)

type UserMessage struct {
	Id          int64      `json:"id" gorm:"primary_key;type:INT AUTO_INCREMENT"`
	NotifyUid   string     `json:"notify_uid" gorm:"column:notify_uid;NOT NULL;default:'';comment:'通知者id';type:VARCHAR(255)"`
	Username    string     `json:"username" gorm:"column:username;NOT NULL;default:'';comment:'用户名';type:VARCHAR(255)"`
	Avatar      string     `json:"avatar" gorm:"column:avatar;NOT NULL;default:'';comment:'头像';type:VARCHAR(255)"`
	ReceiveUid  string     `json:"receive_uid" gorm:"column:receive_uid;NOT NULL;default:'';comment:'发送者id';type:VARCHAR(255)"`
	MessageType string     `json:"message_type" gorm:"column:message_type;NOT NULL;default:'';comment:'消息类型';type:VARCHAR(255)"`
	SendUid     string     `json:"send_uid" gorm:"column:send_uid;NOT NULL;default:'';comment:'发送者id';type:VARCHAR(255)"`
	Content     string     `json:"content" gorm:"column:content;NOT NULL;default:'';comment:'消息内容';type:VARCHAR(255)"`
	Status      int64      `json:"status" gorm:"column:status;NOT NULL;default:0;comment:'0未已查看1已经查看';type:TINYINT"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:created_at;NOT NULL;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at;NOT NULL;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:DATETIME"`
	table       string     `json:"-"`
}

var userMessageMap map[int64]*UserMessage = make(map[int64]*UserMessage)

// @desc
// @auth liuguoqiang 2020-11-26
// @param
// @return
func UserMessageModel(index int64) *UserMessage {
	return userMessageMap[index]
}

func (this *UserMessage) TableName() string {
	return this.table
}
func (*UserMessage) Insert(db *gormx.DB, data *UserMessage) error {
	if db == nil {
		db = Orm()
	}
	return db.Create(data).Error
}

func (*UserMessage) GetOne(where string, args ...interface{}) (*UserMessage, error) {
	var obj UserMessage
	return &obj, Orm().Where(where, args...).Take(&obj).Error
}

func (*UserMessage) GetList(where string, args ...interface{}) ([]*UserMessage, error) {
	var list []*UserMessage
	db := Orm()
	return list, db.Where(where, args...).Find(&list).Error
}

func (*UserMessage) GetListPage(pageSize int64, where string, args ...interface{}) ([]*UserMessage, error) {
	var list []*UserMessage
	db := Orm()
	return list, db.Where(where, args...).Limit(pageSize).Find(&list).Error
}

func (*UserMessage) GetCount(where string, args ...interface{}) (int, error) {
	var number int
	err := Orm().Model(&UserMessage{}).Where(where, args...).Count(&number).Error
	return number, err
}

func (*UserMessage) Delete(db *gormx.DB, where string, args ...interface{}) error {
	if db == nil {
		db = Orm()
	}
	return db.Where(where, args...).Delete(&UserMessage{}).Error
}

func (*UserMessage) Update(db *gormx.DB, data map[string]interface{}, where string, args ...interface{}) error {
	if db == nil {
		db = Orm()
	}
	return db.Model(&UserMessage{}).Where(where, args...).Update(data).Error
}
