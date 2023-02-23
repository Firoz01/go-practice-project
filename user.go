package main

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `json,bun:",pk,autoincrement"`
	Name      string    `json,bun:"name,notnull"`
	Password  string    `json,bun:"password"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Employee struct {
	bun.BaseModel `bun:"table:employee"`
	ID            int       `json,bun:"id,pk,autoincrement"`
	Name          string    `json,bun:"name,notnull"`
	Salary        int       `json,bun:"salary"`
	CreatedAt     time.Time `json,bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `json,bun:",nullzero,notnull,default:current_timestamp"`
}
