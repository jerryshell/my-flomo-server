package service

import (
	"errors"
	"github.com/jerryshell/my-flomo-server/db"
	"github.com/jerryshell/my-flomo-server/model"
	"github.com/jerryshell/my-flomo-server/util"
	"log"
)

func SecretList() []model.Secret {
	var secretList []model.Secret
	_ = db.DB.Order("created_at desc").Find(&secretList)
	return secretList
}

func SecretSave(secret *model.Secret) error {
	db.DB.Save(secret)
	return nil
}

func SecretCreate(userId string, secretStr string) (*model.Secret, error) {
	id, err := util.NextIDStr()
	if err != nil {
		return nil, err
	}
	secret := &model.Secret{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserId: userId,
		Secret: secretStr,
	}
	log.Println("secret", secret)
	_ = db.DB.Create(secret)

	return secret, nil
}

func SecretUpdate(id string, content string) (*model.Secret, error) {
	secret := model.Secret{}
	_ = db.DB.First(&secret, id)
	if secret.ID == "" {
		return nil, errors.New("找不到 secret，id: " + id)
	}

	secret.Secret = content
	_ = db.DB.Save(&secret)

	return &secret, nil
}

func SecretDelete(id string) {
	secret := model.Secret{}
	_ = db.DB.First(&secret, id)
	_ = db.DB.Delete(&secret)
}

func GetSecretByUserId(userId string) (*model.Secret, error) {
	secret := model.Secret{}
	if secret.Secret == "" {
		return nil, errors.New("该用户没有 secret")
	}
	db.DB.Where("user_id = ?", userId).First(&secret)
	return &secret, nil
}

func GetSecretBySecret(secret string) (*model.Secret, error) {
	userSecret := model.Secret{}
	if userSecret.Secret == "" {
		return nil, errors.New("该 secret 不存在")
	}
	db.DB.Where("secret = ?", secret).First(&secret)
	return &userSecret, nil
}
