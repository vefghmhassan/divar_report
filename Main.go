package main

import (
	"com.snap.co/sample/models"
	"com.snap.co/sample/repository"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
)

var sizedata  =0
var widgetList models.Divar
var page int
func main() {
page=1
	//widgetList := repository.GetDivarByCategory(page)

AddSize()

	fmt.Println()
}
func getReadExle() {
	var items []Price
	f, err := excelize.OpenFile("/home/nobaar/Desktop/go_read_exle/truck.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("Sheet 1")
	if err != nil {
		// fmt.Println(err)
		return
	}

	for _, row := range rows {
		for i := 0; i < len(row); i++ {

			var msg = new(Price)
			msg.priceNobaar = row[0]
			msg.priceSnapp = row[1]
			msg.id = row[2]
			s := Price{msg.priceNobaar, msg.priceSnapp, msg.id}
			items = append(items, s)

		}
	}
}
func AddSize()  {
	widgetList := repository.GetDivarByCategory(page)
	size := len(widgetList.WidgetList)
	if size !=0{
		page++
		sizedata = sizedata+size
		fmt.Println(sizedata)


		widgetList=nil
		AddSize()
	} else {
		fmt.Printf(string(sizedata))
	}

}
