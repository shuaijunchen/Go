// excel_image.go
package main

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
)

func main() {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 20, "C2": 30, "D2": 30, "B3": 50, "C3": 20, "D3": 40, "B4": 60, "C4": 70, "D4": 80}
	xlsx := excelize.CreateFile()
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Sheet1", k, v)
	}

	format := `{"type":"bar3D","series":[{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$2:$D$2"},{"name":"=Sheet1!$A$2","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$3:$D$3"},{"name":"=Sheet1!$A$3","categories":"=Sheet1!$B$1:$D$1","values":"=Sheet1!$B$4:$D$4"}],"title":{"name":"Fruit 3D Line Chart"}}`

	xlsx.AddChart("Sheet1", "E1", format)

	err := xlsx.WriteTo("./Workbook.xlsx")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
