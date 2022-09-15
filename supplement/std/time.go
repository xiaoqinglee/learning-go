package std

import (
	"github.com/k0kubun/pp/v3"
	"time"
)

func LocalizedTime() {

	//"2022-04-02T23:00:00Z"
	//"2022-04-03T07:00:00+08:00"
	//"=========================================="
	//2022 4 3
	//"2022-04-03T00:00:00+08:00"
	//"2022-04-04T00:00:00+08:00"
	//"=========================================="
	//2022 4 2

	z := time.UTC
	beijing, _ := time.LoadLocation("Asia/Shanghai")

	t, _ := time.Parse(time.RFC3339, "2022-04-02T23:00:00Z")
	pp.Println(t.In(z).Format(time.RFC3339))
	pp.Println(t.In(beijing).Format(time.RFC3339))

	pp.Println("==========================================")
	localized := t.In(beijing)
	year, month, day := localized.Date()
	pp.Println(year, month, day)

	start := time.Date(year, month, day, 0, 0, 0, 0, beijing)
	end := start.Add(time.Hour * time.Duration(24))

	pp.Println(start.In(beijing).Format(time.RFC3339))
	pp.Println(end.In(beijing).Format(time.RFC3339))

	pp.Println("==========================================")
	utc := t.In(time.UTC)
	year, month, day = utc.Date()
	pp.Println(year, month, day)
}
