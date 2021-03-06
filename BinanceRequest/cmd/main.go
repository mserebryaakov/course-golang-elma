package main

import (
	"github.com/mserebryaakov/course-golang/BinanceRequest/pkg"
	"github.com/mserebryaakov/course-golang/BinanceRequest/postgresql"
)

func main() {
	st := postgresql.New()
	st.Open()
	defer st.Close()
	data, _ := pkg.RunGet()
	for _, value := range *data {
		st.InsertBinanceData(value.Symbol, value.Price)
	}
}
