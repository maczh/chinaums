package chinaums

import "github.com/sirupsen/logrus"

type client struct {
	BscPay   *bscPay
	WxAppPay *wxAppPay
}

var (
	Client = &client{
		BscPay:   &bscPay{},
		WxAppPay: &wxAppPay{},
	}

	log = logrus.New()
)
