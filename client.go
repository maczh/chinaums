package chinaums

import "github.com/sirupsen/logrus"

type client struct {
	BscPay   *bscPay
	WxAppPay *wxAppPay
	QRPay    *qrPay
}

var (
	Client = &client{
		BscPay:   &bscPay{},
		WxAppPay: &wxAppPay{},
		QRPay:    &qrPay{},
	}

	log = logrus.New()
)
