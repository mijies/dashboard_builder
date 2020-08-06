package dashboard

import (
	"fmt"
	"log"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mijies/dashboard_builder/utils"
)

type dbook interface {
	load()
	parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command)
	build()
	render()
}

type dashboard struct {
	path		string
	book		*excelize.File
	commands	commands
	snippets	snippets
}
type targetBook struct {
	dashboard
}
type masterBook struct {
	dashboard
}
type userBook struct {
	dashboard
}

func(d *dashboard) load() {
	file, err := excelize.OpenFile(d.path)
    if err != nil {
        log.Fatal(err)
	}
	d.book = file
}
func(d *masterBook) load() {
	d.dashboard.load()
	d.commands.load(d.book)
	d.snippets.load(d.book)
	d.book = nil
}
func(d *userBook) load() {
	d.dashboard.load()
	d.commands.load(d.book)
	d.snippets.load(d.book)
	d.book = nil
}

func(d *dashboard) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
}
func(d *targetBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	d._parse_snippets(*<- sch, *<- sch)
	d._parse_commands(*<- cch, *<- cch)
}
func(d *masterBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	sch <- d.snippets.parse().(*[]snippet)
	mtx <- true
	cch <- d.commands.parse().(*[]command)
	mtx <- true
}
func(d *userBook) parse(mtx chan bool, sch chan *[]snippet, cch chan *[]command) {
	<- mtx // wait for masterBook snippets
	sch <- d.snippets.parse().(*[]snippet)
	<- mtx
	cch <- d.commands.parse().(*[]command)
}

func(d *targetBook) _parse_snippets(ms []snippet, us []snippet) {
	for _, master := range ms {
		for _, user := range us {
			for k, _ := range master.snipMap {
				if _, dup := user.snipMap[k]; dup {
					log.Fatal("code name duplication with your custom code: " + k)
				}
			}
		}
	}
	ms = append(ms, us...)
	d.snippets.snippets = append(d.snippets.snippets, ms...)
}

func(d *targetBook) _parse_commands(mc []command, uc []command) {
	var c command
	m, u, offset := 0, 0, 0
	// prioritize user chains indces
	for i := 1; i <= len(mc) + len(uc); i++ {
		if m < len(mc) {
			c = mc[m]
		}
		if u < len(uc) && uc[u].index <= c.index + offset {
			c = uc[u]
			offset++
			u++
		} else {
			m++
		}
		c.index = i
		d.commands.chains = append(d.commands.chains, c)
	}
}

func(d *dashboard) build() {
}
func(d *targetBook) build() {
	// create a new book
	new_path := utils.AddTimestampToFilename(d.path, TIME_FORMAT, "xlsm")
	if err := d.book.SaveAs(new_path); err != nil {
        log.Fatal(err)
	}

	// swap with the new one
	d.path = new_path
	d.load()

	// delete the exsisting macro sheet
	d.book.DeleteSheet(MACRO_SHEET_NAME)

	// copy the template sheet
	d.book.NewSheet(MACRO_SHEET_NAME)

	d.render()
	if err := d.book.Save(); err != nil {
        log.Fatal(err)
	}
}

func(d *dashboard) render() {
	d._render_commands()
	d._render_snippets()
	d._layout_sheet()
}

func(d *dashboard) _render_commands() {
	rowi := COMMANDS_ROW
	cols := []string{"No","Name",COMMANDS_LABEL,"[[ARGS]]","","","","","",""}
	d._render_row(rowi, cols, COMMANDS_STYLE_HEADERS[:])

	max_chain_width := 0
	max_args_count  := 0
	iter := d.commands.into_iter()
	for iter.hasNext() {
		rowi++
		cols, styles := iter.next()
		max_chain_width = utils.MaxInt(max_chain_width, len(cols[2]))
		max_args_count  = utils.MaxInt(max_args_count,  len(cols[3:]))
		d._render_row(rowi, cols, styles)
	}
}

func(d *dashboard) _render_snippets() {
	rowi := COMMANDS_ROW + d.commands.len() + SNIPPETS_ROW
	cols := []string{"","",SNIPPETS_LABEL}
	d._render_row(rowi, cols, SNIPPETS_STYLE_HEADERS[:])

	iter := d.snippets.into_iter()
	for iter.hasNext() {
		rowi++
		cols, styles := iter.next()
		d._render_row(rowi, cols, styles)
	}
}

func(d *dashboard) _render_row(rowi int, cols []string, styles []string) {
	COL_SEED := int('A')
	for i, v := range cols {
		axis := fmt.Sprintf("%s%d", string(COL_SEED + i), rowi)
		d.book.SetCellValue(MACRO_SHEET_NAME, axis, v)
		if len(styles[i]) != 0 {
			d.setRowStyle(MACRO_SHEET_NAME, axis, styles[i])
		}
	}
}

func(d *dashboard) setRowStyle(sheet_name string, axis string, style_str string) {
	style, err := d.book.NewStyle(style_str)
	if err != nil {
		log.Fatal(err)
	}
	err = d.book.SetCellStyle(sheet_name, axis, axis, style)
}

func(d *dashboard) _layout_sheet() {
	COL_SEED := int('A')
	for i, width := range COLUMN_WIDTH_SLICE {
		err := d.book.SetColWidth(MACRO_SHEET_NAME, string(COL_SEED + i), string(COL_SEED + i), float64(width))
		if err != nil {
			log.Fatal(err)
		}
	}
}