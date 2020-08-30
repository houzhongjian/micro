package handler

import (
	"log"
)

func (u *User) ActionLogin() {
	log.Println("hello login", u.Ctx.IP)
}
func (u *User) ActionIndex() {
	log.Println("hello login", u.Ctx.Host)
}

func (u *User) ActionTest() {
	log.Println("hello test", u.Ctx.IP)
}

func (u *User) ActionUserAdd() {
	log.Println("hello user add", u.Ctx.IP)
}

func (u *User) ActionUserDelete() {
	log.Println("delete user ok", u.Ctx.IP)
}
