package chinaums

import "github.com/sirupsen/logrus"

type client struct {
	BscPay *bscPay
}

var (
	Client = &client{
		BscPay: &bscPay{},
	}

	log = logrus.New()
)
