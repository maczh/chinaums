package chinaums

type sdkConfig struct {
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	SourceId   string `json:"sourceId"`
	EncryptKey string `json:"encryptKey"`
	ApiHost    string `json:"apiHost"`
	MerchantId string `json:"merchantId"`
	TerminalId string `json:"terminalId"`
}

var config = map[string]sdkConfig{
	"test": sdkConfig{
		AppId:      "10037e6f803bc8ab01805fd07db4000d",
		AppKey:     "18a3a8d41b9f44d9acdf39c360f7a971",
		SourceId:   "103A",
		EncryptKey: "GAhPWQ8D4hXanneneaydaHYHiwn64p7y3A46Jpss87aKWsy5",
		ApiHost:    "https://test-api-open.chinaums.com",
		MerchantId: "898201612345678",
		TerminalId: "88880001",
	},
	"prod": {
		AppId:      "8a81c1dc96cf8e4c01996126a99b12b0",
		AppKey:     "78ed50b460494c889fff3c2def6b5268",
		SourceId:   "3GET",
		EncryptKey: "DrynX3P3GZQjwbENGYpFemQK6CMxxyx4kcSCp4G4CkzEdG58",
		ApiHost:    "https://api-mop.chinaums.com",
		MerchantId: "",
		TerminalId: "",
	},
}

const (
	env    = "test"
	URI_BC = "/v4/poslink/transaction/pay"
)
