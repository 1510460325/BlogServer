package info

import "blog/dal/domain"

type CreateMessageInfo struct {
	domain.BlogMessageDomain
	CaptchaID   *string `json:"captchaID"`
	VerifyValue *string `json:"verifyValue"`
}
