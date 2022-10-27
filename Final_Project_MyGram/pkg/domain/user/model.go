package user

import (
	"encoding/json"
	"time"

	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/helpers"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/comment"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/gormmodel"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/photo"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/socialmedia"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type customTime datatypes.Date

var _ json.Unmarshaler = &customTime{}

func (mt *customTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = customTime(t)
	return nil
}

type User struct {
	gormmodel.GormModel
	Username     string                    `gorm:"not null;uniqueIndex" json:"username" valid:"required~username is required"`
	Email        string                    `gorm:"not null;uniqueIndex" json:"email" valid:"required~email is required,email~invalid email format"`
	Password     string                    `gorm:"not null" json:"password" valid:"required~password is required,minstringlength(6)~password has to have a minimum length of 6 characters"`
	DOB          customTime                `gorm:"not null;type:date" json:"dob" valid:"required~date of birth is required"`
	Photos       []photo.Photo             `json:"photos"`
	Comments     []comment.Comment         `json:"comments"`
	SocialMedias []socialmedia.SocialMedia `json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	return
}
