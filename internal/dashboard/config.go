package dashboard

import (
	"path/filepath"
)

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
)

func getTimeFormat() string {
	return TIME_FORMAT
}

func getMacroSheetName() string {
	return MACRO_SHEET_NAME
}

func getMasterPath() string {
	return filepath.FromSlash(BASE_DIR + MASTER_DIR + DASHBOARD_FILE)
}

func getUserPath(user string) string {
	return filepath.FromSlash(BASE_DIR + USERS_DIR + user + "/" + DASHBOARD_FILE)
}