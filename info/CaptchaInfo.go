package info

type CaptchaInfo struct {
	ID          *string `json:"id"`
	IsAudio     *bool   `json:"isAudio"`
	VerifyValue *string `json:"verifyValue"`
	Base64Data  *string `json:"base64data"`
}
