package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func main() {
	appID := "2021000121665075"
	privateKey := "MIIEowIBAAKCAQEAjGqOE2aWuBrGOpNeNcY8TVCQW0vVbWXHWn2OHTFOwzDcROs/iyBoivh9RJCOhUsvkMCXC/Jtx1gGQSMy4ObM6Mizj1MKczIwDKj+dAH1kh8Kt/mUViGjy6PbApQy0lpy0/4USxYx7AHE6Ldu+PNnKT33HsPiJ/eNbvEbsV+cY5ARlPGSVcIBCuJ9kGk6i3qntW/MyyO89HnmDhDY1usatwHBYgNQwTafkanpHfw/SNofv0DZ0KRSQtpv/ndQSRcsyoSBshJdCdUO0HHX1wJTca3AOXUCjOq/vSzX1YAvqoUWJaSj3OSREGi3uWBr8y4f/Ili0YWuw9VyPaKi6MM1IwIDAQABAoIBAFYm4nkAzqSyuMEjvf/cWw9MyOVbB94RPllFA7bhk+OqzNr0Q46HGd16TNGOqAMceFil8YcA/V37ogrBH+xhb4O78H+VgangNx7taQGgWh2HMjpixSJ7jVXaePuCFhR+LknncGgJfCPDih8GvO292aOyQzb7FH1VI8/X2xffA8MJqG2j8Mav9k4Q0dTCclsfRSB1rvG60UgOCKwww4H8DlXpqahmYcdXvOXdoVl+h1qKxeetO71z3NIXv8JAV37Z1gq4W/GVTOi868pMqqSN8GwnwOIm44bV2afibH0Ah+aW9vgmG1J0BmAoutv+SORY6mQB/QPoDCtIkwoisDKmqhECgYEA8UXxGtIPjF5qJhwBCcita90xtWdsCSJ09ble3v62MmG3qKnvD55e+cU8zIJUIKnUTvIcf4bFlUWkJjNRCKaJxRQwigD82ZvWt9hZrRtg2EFnn/UNbGUuUwQMK9yJKcgczLs5kgw3tvSUIpFl6IEqR9SZ/CMK/v8HU/ghXgSZ5D8CgYEAlPyneOK4S0CqAGadf0A2f0S5TRbFzeeyCpXNrCejizqC+nBxMUw9rUaCIF6w8DO4Os1M34/kZxqyy0BiDQRFVG5BfM3xMgA8HhuasD63Abn4Fzbr8x8oaVM5RRRfd59coELtBwoRUvH07lwqGrYP7GVU/g3OOnMkiTz5uPvNJh0CgYEA1FVFVuE/Jb2BlYBXuZCObrr/oDIgdfmJDSfSQlXPao9s3laPOSObWh0m4KRh/Uz4t3GDewfUowXi7GuFCSPnYzXdcdZkr+3iDXGtmhbaJ/eHhtJWueph1lVdkMxJOigOkG7Ev69Y6P5dz/49vVzPJOAJlEgSyU8P1+orE17hQGkCgYBbOxdGI/bcRd2T30JEQMreRfP+K5q2ilFwv+iwqZkw0YSQ3LfDiz8dUtauAa3xeIown7SaFV+0bxna8jLYa6wUlmtNhEJU7uXhPeAMWmrwCLPNa1kyh/rT/B1OmroW4yPyEulX7SdMllL4fsFdl+zKQga4hzWofwd5bjgwuSgXVQKBgFTarKM/s4CanwgdiAbXeF0sfJ83KchoUSKhSK/5FLiVxojWRCgCMEylS18t4M+Aa2jI2i9q3Zj8Kw6VAcccQW7h737Rwg/5nKU7dSp4YjcNx5hnyvJeK5SVf6xESP2DTe4+TnpyUwNVtIPzHfHP0/G16dlx5RSJIWqpqjqfIeRc"
	aliPublickey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsJaQ3NjNAgpVMmUozGiwpb7/q9QjYl033mIHh7nqxuxHFHAQ+nTx7nz7k0vlTPuARmcTog3bZtivXJEEC8SOE46ejjnNiFBD1sCYFRe1LPuGb3RtPggkaPAJoIQQdREUZIdV/8//P/RIWoJjQTj7JIMlSv8UNQTLZtiEC1HQASTMM9w+6LUduTbd+AkpWubmXqj6JuoK6aiDSyl+OP9405MC9YMSeSCUbLv2fWes6vob+00492j9UthAP4/Mx0yUaaRS+pSV5tastfQTVrlqSKCnwxHcavJ6/7kfw1usYyGpO1l6u8EoqbcuEUII+BCPk2iOdkrMqQtI+OgcLmxIIwIDAQAB"
	client, err := alipay.New(appID, privateKey, true)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(aliPublickey)
	if err != nil {
		panic(err)
	}
	p := alipay.TradePagePay{}
	p.NotifyURL = "http://120.25.255.207:33389/"
	p.ReturnURL = "http://127.0.0.1:8089/"
	p.Subject = "生鲜-订单支付"
	p.OutTradeNo = "pluto"
	p.TotalAmount = "1.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	pay, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(pay.String())
}
