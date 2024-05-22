package Controller

import (
	_ "github.com/mattn/go-sqlite3"
)

type APIResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func isInSlice(slice []any, val any) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
