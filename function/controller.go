package Controller

import (
	_ "github.com/mattn/go-sqlite3"
)

type APIResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
