package dashboard

import (
	"errors"
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func findRow(book *excelize.File, sheet_name string, value string, rows [2]int, cols [2]int) (int, error) {
	if cols[1] > 27 {
		log.Fatal("column only supported upto Z")
	}

	seed := int('A' - 1)
	for col := cols[0]; col < cols[1]; col++ {
		for row  := rows[0]; row < rows[1]; row++ {
			axis := fmt.Sprintf("%s%d", string(seed + col), row)
			val, err := book.GetCellValue(sheet_name, axis)
			if err != nil {
				log.Fatal(err)
			}		
			if val == value {
				return row, nil
			}
		}	
	} 
	return -1, errors.New("no cell found with the value:" + value)
}