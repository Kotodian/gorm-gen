// Code generated by "gorm-gen"; DO NOT EDIT.

package dao

import (
	"context"
	"github.com/Kotodian/gorm-gen/examples/models/common"
	"gorm.io/gorm"
)

// CreateUser create object
func CreateUser(ctx context.Context, db *gorm.DB, user *common.User) error {
	return db.WithContext(ctx).Create(user).Error
}

// ListUser list objects
func UserList(ctx context.Context, db *gorm.DB) ([]*common.User, error) {
	users := make([]*common.User, 0)
	err := db.WithContext(ctx).Where("deleted_at=0").Find(&users).Error
	return users, err
}

// QueryUser query objects by condition
func UserQuery(ctx context.Context, db *gorm.DB, args map[string]interface{}) ([]*common.User, error) {
	users := make([]*common.User, 0)
	err := db.WithContext(ctx).Where(args).Where("deleted_at=0").Find(&users).Error
	return users, err
}

// UpdateUser update object
func UpdateUser(ctx context.Context, db *gorm.DB, args map[string]interface{}) error {
	return db.WithContext(ctx).Model(&common.User{}).Updates(args).Error
}

// DeleteUser delete object
func DeleteUser(ctx context.Context, db *gorm.DB, user common.User) error {
	return db.WithContext(ctx).Delete(&user).Error
}