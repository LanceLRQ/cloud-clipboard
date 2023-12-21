package responses

type AuthLogin struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}

type OTPInfo struct {
	Secret   string `json:"secret"`
	Url      string `json:"url"`
	TestCode string `json:"test_code"`
}
