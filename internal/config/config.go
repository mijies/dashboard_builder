package config

import (
	"path/filepath"
	// "github.com/mijies/dashboard_generator/pkg/utils"
)

const (
	TIME_FORMAT 	= "20060102030405"

	// files and directories and labels
	BASE_PATH    	= "samples/"
	COMMANDS_DIR 	= "commands/"
	COMMANDS_FILE	= "commands.json"
	COMMANDS_LABEL	= "[[COMMANDS]]"
	TTL_CODES_DIR	= "ttl_codes/"
	TTL_CODES_LABEL	= "[[SNIPPETS]]"

	// command sheet format
	MACRO_SHEET_NAME   	 = "ttl_macro"
	MACRO_TMP_SHEET_NAME = "ttl_macro_tmp"
)


type Config interface {
	GetTimeFormat()			string
	GetMacroSheetName()		string
	GetMacroTmpSheetName()	string
	GetCommandsDir()		string
	GetCommandsFile()		string
	GetCommandsLabel()		string
	GetTTLCodesDir()		string
	GetTTLCodesLabel()		string
}

type config struct {
}

func NewConfig() Config {
	return Config(&config{})
}

func(i *config) GetTimeFormat() string {
	return TIME_FORMAT
}

func(i *config) GetMacroSheetName() string {
	return MACRO_SHEET_NAME
}

func(i *config) GetMacroTmpSheetName() string {
	return MACRO_TMP_SHEET_NAME
}

// commands
func(i *config) GetCommandsDir() string {
	return filepath.FromSlash(BASE_PATH + COMMANDS_DIR)
}

func(i *config) GetCommandsFile() string {
	return COMMANDS_FILE
}

func(i *config) GetCommandsLabel() string {
	return COMMANDS_LABEL
}

// ttl_codes
func(i *config) GetTTLCodesDir() string {
	return filepath.FromSlash(BASE_PATH + TTL_CODES_DIR)
}

func(i *config) GetTTLCodesLabel() string {
	return TTL_CODES_LABEL
}