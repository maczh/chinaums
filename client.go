package chinaums

import "github.com/sirupsen/logrus"

type client struct {
	BscPay   *bscPay
	WxAppPay *wxAppPay
	QRPay    *qrPay
	Bum      *businessUnifyMulti
}

var (
	Client = &client{
		BscPay:   &bscPay{},
		WxAppPay: &wxAppPay{},
		QRPay:    &qrPay{},
		Bum:      &businessUnifyMulti{},
	}

	log = logrus.New()
)
