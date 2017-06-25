package main

import (
	"fmt"
	"github.com/andreyvit/telegramapi"
	"time"
)

func main() {
	o := telegramapi.Options{
		SeedAddr: telegramapi.Addr{IP: "149.154.167.40", Port: 443},
		PublicKey: `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAwVACPi9w23mF3tBkdZz+zwrzKOaaQdr01vAbU4E1pvkfj4sqDsm6
lyDONS789sVoD/xCS9Y0hkkC3gtL1tSfTlgCMOOul9lcixlEKzwKENj1Yz/s7daS
an9tqw3bfUV/nqgbhGX81v/+7RFAEd+RwFnK7a+XYl9sluzHRyVVaTTveB2GazTw
Efzk2DWgkBluml8OREmvfraX3bkHZJTKX4EQSjBbbdJ2ZXIsRrYOXfaA+xayEGB+
8hdlLmAjbCVfaigxX0CDqWeR1yFL9kwd9P0NsZRPsmoqVwMbMu7mStFai6aIhc3n
Slv8kg9qv1m6XHVQY3PnEw+QQtqSIXklHwIDAQAB
-----END RSA PUBLIC KEY-----`,
		APIHash: "e88ec58aa1ce01f5630e194e9571d751",
		APIID:   123259,
		Verbose: 1000,
	}
	con := telegramapi.New(o, &telegramapi.State{}, &h{})
	con.StartLogin("+989015132328")

	time.Sleep(time.Hour)
}

type h struct {
}

func (s h) HandleConnectionReady() {

	fmt.Println("HandleConnectionReady")
}

func (s h) HandleStateChanged(newState *telegramapi.State) {
	fmt.Println("HandleConnectionReady")

}
