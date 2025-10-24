package basslink

import "github.com/vannleonheart/goutil"

type RecaptchaConfig struct {
	SiteKey   string `json:"site_key"`
	SecretKey string `json:"secret_key"`
	VerifyUrl string `json:"verify_url"`
}

type RecaptchaClient struct {
	Config *RecaptchaConfig
}

type RecaptchaVerifyRequest struct {
	Secret   string  `json:"secret"`
	Response string  `json:"response"`
	RemoteIp *string `json:"remoteip"`
}

type RecaptchaVerifyResponse struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func NewRecaptchaClient(config *RecaptchaConfig) *RecaptchaClient {
	return &RecaptchaClient{
		Config: config,
	}
}

func (c *RecaptchaClient) Verify(token string, ipAddress *string) (bool, error) {
	data := RecaptchaVerifyRequest{
		Secret:   c.Config.SecretKey,
		Response: token,
		RemoteIp: ipAddress,
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	var response RecaptchaVerifyResponse
	if _, err := goutil.SendHttpPost(c.Config.VerifyUrl, &data, &headers, &response, nil); err != nil {
		return false, err
	}

	return response.Success, nil
}
