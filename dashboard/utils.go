package dashboard

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func getMasterPath() string {
	return filepath.FromSlash(BASE_DIR + MASTER_DIR + DASHBOARD_FILE)
}

func getUserPath(user string) string {
	return filepath.FromSlash(BASE_DIR + USERS_DIR + user + "/" + DASHBOARD_FILE)
}

func findRow(book *excelize.File, sheet_name string, value string, rows [2]int, cols [2]int) (int, error) {
	if cols[1] > 27 {
		log.Fatal("column only supported upto Z")
	}

	COL_SEED := int('A' - 1)
	for col := cols[0]; col < cols[1]; col++ {
		for row  := rows[0]; row < rows[1]; row++ {
			axis := fmt.Sprintf("%s%d", string(COL_SEED + col), row)
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