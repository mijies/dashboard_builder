package dashboard

const (
	TIME_FORMAT 	= "20060102030405"

	// files and directories and labels
	BASE_DIR    	= "samples/"
	MASTER_DIR		= "master/"
	USERS_DIR		= "users/"
	DASHBOARD_FILE	= "dashboard.xlsm"

	// command sheet format
	MACRO_SHEET_NAME = "ttl_macro"
	// MACRO_TMP_SHEET_NAME = "ttl_macro_tmp"

	COMMANDS_LABEL	= "[[COMMAND]]"
	SNIPPETS_LABEL	= "[[CONFIG]]"

	COMMANDS_ROW	= 2
	SNIPPETS_ROW	= 2 // offset from the last command chain
	
	
	COMMANDS_STYLE_INDEX = ""
	COMMANDS_STYLE_NAME  = `{"font":{"bold":true}}`
	COMMANDS_STYLE_CHAIN = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
							 "fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":5}}`
	COMMANDS_STYLE_ARGS  = ""

	STYLE_TITLE = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
					"fill":{"type":"gradient","color":["#FFFFFF","#FFE6E6"],"shading":5}}`


)

var COMMANDS_HEADER_STYLE = [4]string{
	`{"font":{"bold":true}}`,
	`{"font":{"bold":true}}`,
	COMMANDS_STYLE_CHAIN,
	`{"font":{"bold":true}}`,
}

var COMMANDS_STYLE = [4]string{
	COMMANDS_STYLE_INDEX,
	COMMANDS_STYLE_NAME,
	COMMANDS_STYLE_CHAIN,
	COMMANDS_STYLE_ARGS,
}

var SNIPPETS_BODY_STYLE  = [4]string{"","","",""}
var SNIPPETS_BLANK_STYLE = SNIPPETS_BODY_STYLE

