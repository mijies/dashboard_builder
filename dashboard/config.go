package dashboard

const (
	TIME_FORMAT 	= "20060102030405"

	// files and directories and labels
	BASE_DIR    	= "samples/"
	MASTER_PATH		= "master/dashboard.xlsm"
	USERS_DIR		= "users/"

	// command sheet format
	MACRO_SHEET_NAME = "ttl_macro"

	COMMANDS_LABEL	= "[[COMMAND]]"
	SNIPPETS_LABEL	= "[[DEFINITION]]"

	LOGIN_NAME_MATCH = "foobar"

	COMMANDS_ROW	= 2
	SNIPPETS_ROW	= 2 // offset from the last command chain

	NO_COLUMN_WIDTH    = 3
	NAME_COLUMN_WIDTH  = 20
	CHAIN_COLUMN_WIDTH = 70
	ARGS_COLUMN_WIDTH  = 20
	
	ARGS_EXTRA_COLUMN_COUNT = 6
	COL_SEED = int('A')

	COMMANDS_STYLE_HEADER = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
							 "fill":{"type":"gradient","color":["#FFFFFF","#FFE6E6"],"shading":5},
							 "font":{"bold":true}}`
	COMMANDS_STYLE_INDEX = ""
	COMMANDS_STYLE_NAME  = `{"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`
	COMMANDS_STYLE_NAME_BOLD = `{"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1},"font":{"bold":true}}`
	COMMANDS_STYLE_CHAIN = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
							 "fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":5}}`
	COMMANDS_STYLE_ARGS  = ""

	SNIPPETS_STYLE_NAME  = `{"border":[{"type":"left","style":1},{"type":"right","style":1},{"type":"top","style":1},{"type":"bottom","style":1}],
							 "fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":5}}`
	SNIPPETS_STYLE_BODY  = `{"alignment":{"vertical":"center","wrap_text":true}}`
)

var COLUMN_WIDTH_SLICE = [4 + ARGS_EXTRA_COLUMN_COUNT]int{
	NO_COLUMN_WIDTH, NAME_COLUMN_WIDTH, CHAIN_COLUMN_WIDTH, ARGS_COLUMN_WIDTH,
	ARGS_COLUMN_WIDTH, ARGS_COLUMN_WIDTH, ARGS_COLUMN_WIDTH,
	ARGS_COLUMN_WIDTH, ARGS_COLUMN_WIDTH, ARGS_COLUMN_WIDTH,
}

var COMMANDS_STYLE_HEADERS = [4 + ARGS_EXTRA_COLUMN_COUNT]string{
	COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,COMMANDS_STYLE_HEADER,
	COMMANDS_STYLE_HEADER,COMMANDS_STYLE_HEADER,
}

var COMMANDS_STYLE = [4]string{
	COMMANDS_STYLE_INDEX,
	COMMANDS_STYLE_NAME,
	COMMANDS_STYLE_CHAIN,
	COMMANDS_STYLE_ARGS,
}

var SNIPPETS_STYLE_HEADERS = [3]string{"","",COMMANDS_STYLE_HEADER}
var SNIPPETS_STYLE_NAMES   = [3]string{"","",SNIPPETS_STYLE_NAME}
var SNIPPETS_STYLE_BODIES  = [3]string{"","",SNIPPETS_STYLE_BODY}
var SNIPPETS_STYLE_BLANKS  = SNIPPETS_STYLE_BODIES

