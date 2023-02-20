package main

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64  `json,bun:",pk"`
	Name string `json,bun:"name,notnull"`
}
