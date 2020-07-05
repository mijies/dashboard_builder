package config

import (
	"path/filepath"
	"github.com/mijies/dashboard_generator/pkg/utils"
)

const (
	TIME_FORMAT 	= "20060102030405"

	// files and directories
	BASE_PATH    	= "samples/"
	COMMANDS_DIR 	= "commands/"
	COMMANDS_FILE	= "commands.json"
	TTL_CODES_DIR	= "ttl_codes/"

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


type Config interface {
	GetNewSheetName()	string
	GetCommandsDir()	string
	GetCommandsFile()	string
}

type config struct {
}

func NewConfig() Config {
	return Config(&config{})
}

func(i *config) GetNewSheetName() string {
	return MACRO_SHEET_NAME_PREFIX + utils.GetTimeStr(TIME_FORMAT)
}

func(i *config) GetCommandsDir() string {
	return filepath.FromSlash(BASE_PATH + COMMANDS_DIR)
}

func(i *config) GetCommandsFile() string {
	return COMMANDS_FILE
}
