package inventory

import (
	"path/filepath"
)

const (
	TIME_FORMAT 	= "20060102030405"

	// files and directories
	BASE_PATH    	= "~/dashboard_generator_test/"
	COMMANDS_DIR 	= "commands/"
	COMMANDS_FILE	= "commands.json"
	CODES_DIR 		= "codes/"

	// command sheet format
	MACRO_SHEET_NAME_PREFIX = "ttl_macro_"
	DESCRIPTION_COLUMN = 1

	COMMAND_LABEL	= "[[COMMANDS]]"
	COMMAND_COLUMN	= 2
	COMMAND_ROW		= 2

	ARGS_LABEL		= "[[ARGS]]"
	ARGS_COLUMN		= 3

	CODES_LABEL		= "[[CODES]]"
)


type Inventory interface {
	GetCommandsFile() string
}

type inventory struct {
}

func NewInventory() Inventory {
	return Inventory(&inventory{})
}

func(i *inventory) GetCommandsFile() string {
	return filepath.FromSlash(BASE_PATH + COMMANDS_DIR + COMMANDS_FILE)
}
