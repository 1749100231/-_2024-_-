package dao

import (
	"JH_2024_MJJ/internal/model"
	"context"
)

func (d *Dao) GetUserByUserName(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username=?", phone).First(&user).Error
	return &user, err
}

func (d *Dao) GetUserByID(ctx context.Context, ID int64) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id=?", ID).First(&user).Error
	return &user, err
}
func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Create(user).Error
}

func (d *Dao) GetToken(ctx context.Context, token string) (*model.TokenTable, error) {
	var tokenTable model.TokenTable
	err := d.orm.WithContext(ctx).Where("token=?", token).First(&tokenTable).Error
	return &tokenTable, err
}

func (d *Dao) GetTokenByID(ctx context.Context, id int64) (*model.TokenTable, error) {
	var tokenTable model.TokenTable
	err := d.orm.WithContext(ctx).Where("user_id=?", id).First(&tokenTable).Error
	return &tokenTable, err
}

func (d *Dao) UpdateToken(ctx context.Context, token *model.TokenTable) error {
	return d.orm.WithContext(ctx).Save(token).Error
}
func (d *Dao) CreateToken(ctx context.Context, token *model.TokenTable) error {
	return d.orm.WithContext(ctx).Create(token).Error
}
