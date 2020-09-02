package handler

import (
	"log"
)

func (u *User) ActionLogin() {
	log.Println("hello login", u.ctx.IP)
}
func (u *User) ActionIndex() {
	log.Println("hello login", u.ctx.Host)
}

func (u *User) ActionTest() {
	log.Println("hello test", u.ctx.IP)
}

func (u *User) ActionUserAdd() {
	log.Println("hello user add", u.ctx.IP)
}

func (u *User) ActionUserDelete() {
	log.Println("delete user ok", u.ctx.IP)
}
