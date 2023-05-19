package bilibili

import (
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/go-resty/resty/v2"
	"github.com/skip2/go-qrcode"
	"time"
)

const (
	GenerateQRCodeurl = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	PollQRCodeurl     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
)

type QRCode struct {
	Url       string `json:"url"`
	QRCodeKey string `json:"qrcode_key"`
}

func (c *QRCode) Encode() ([]byte, error) {
	return qrcode.Encode(c.Url, qrcode.Medium, 256)
}

func (c *QRCode) Print() {
	qrcodeTerminal.New2(
		qrcodeTerminal.ConsoleColors.BrightBlack,
		qrcodeTerminal.ConsoleColors.BrightWhite,
		qrcodeTerminal.QRCodeRecoveryLevels.Low,
	).Get(c.Url).Print()
}

type GenerateQRCodeResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    QRCode `json:"data"`
}

type PollQRCode struct {
	Url          string `json:"url"`
	RefreshToken string `json:"refresh_token"`
	Timestamp    int    `json:"timestamp"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

type PollQRCodeResult struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	TTL     int        `json:"ttl"`
	Data    PollQRCode `json:"data"`
}

func GenerateQRCode() (*QRCode, error) {
	var result GenerateQRCodeResult
	_, err := NewRequest().SetResult(&result).Get(GenerateQRCodeurl)

	return &result.Data, err
}

func LoginByQRCode() (*resty.Client, error) {
	qrCode, err := GenerateQRCode()
	if err != nil {
		return nil, err
	}
	qrCode.Print()

	for {
		<-time.After(2 * time.Second)
		var result PollQRCodeResult
		resp, err := NewRequest().
			SetResult(&result).
			SetQueryParam("qrcode_key", qrCode.QRCodeKey).
			Get(PollQRCodeurl)
		if err != nil {
			return nil, err
		}

		switch result.Data.Code {
		case 0:
			return NewClient().SetCookies(resp.Cookies()), nil
		case 86101: // 未扫码
			continue
		case 86038: // 二维码已失效
			if qrCode, err = GenerateQRCode(); err != nil {
				return nil, err
			}
			qrCode.Print()
		case 86090: // 二维码已扫码未确认
			continue
		}
	}
}
