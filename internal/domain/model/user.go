package model

import "time"

type User struct {
	ID             string                  `json:"id" bson:"_id,omitempty"`
	FirstName      string                  `json:"firstName" bson:"firstName"`
	LastName       string                  `json:"lastName" bson:"lastName"`
	JobTitle       string                  `json:"jobTitle" bson:"jobTitle"`
	Email          string                  `json:"email" bson:"email"`
	Password       string                  `json:"password" bson:"password"`
	Verification   UserVerificationPayload `json:"verification" bson:"verification"`
	RegisteredTime time.Time               `json:"registeredTime" bson:"registeredTime"`
	LastVisitTime  time.Time               `json:"lastVisitTime" bson:"lastVisitTime"`
	MattermostData MattermostData          `json:"mattermostData" bson:"mattermostData"`
	Blocked        bool                    `json:"blocked" bson:"blocked"`
	Settings       UserSettings            `json:"settings" bson:"settings"`
}

type MattermostData struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserInfoShort struct {
	ID        string `json:"id" bson:"id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
}

type UserVerificationPayload struct {
	VerificationCode            string    `json:"varificationCode" bson:"verificationCode"`
	VerificationCodeExpiresTime time.Time `json:"verificationCodeExpiresTime" bson:"verificationCodeExpiresTime"`
	Verified                    bool      `json:"verified" bson:"verified"`
}

type ForgotPasswordPayload struct {
	Token            string    `json:"token" bson:"token"`
	ResultToken      string    `json:"resultToken" bson:"resultToken"`
	TokenExpiresTime time.Time `json:"tokenExpiresTime" bson:"tokenExpiresTime"`
}

type UserSettings struct {
	UserIconURL    string `json:"userIconUrl" bson:"userIconUrl"`
	BannerImageURL string `json:"bannerImageUrl" bson:"bannerImageUrl"`
	TimeZone       string `json:"timeZone" bson:"timeZone"`
	DateFormat     string `json:"dateFormat" bson:"dateFormat"`
	TimeFormat     string `json:"timeFormat" bson:"timeFormat"`
}
type UpdateUserInfoInput struct {
	FirstName *string
	LastName  *string
	JobTitle  *string
	Email     *string
}

type UpdateUserSettingsInput struct {
	UserIconURL    *string
	BannerImageURL *string
	TimeZone       *string
	DateFormat     *string
	TimeFormat     *string
}
