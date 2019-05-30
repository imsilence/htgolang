module web

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190411191339-88737f569e3a
	golang.org/x/net => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190415081028-16da32be82c5
	golang.org/x/text => github.com/golang/text v0.3.0
)

require github.com/astaxie/beego v1.11.1
