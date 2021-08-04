package main

import "fmt"

type User struct {
	nome  string
	email string
	saldo int
}

func validarSaldo(usuario User, valor int) bool {
	return usuario.saldo >= valor
}

func (u *User) transferir(usuario, usuario2 *User, valor int) (User, string) {
	if validarSaldo(*u, valor) {
		u.saldo += valor
		return (*u), "\n Tranferencia completa"
	} else {
		return *u, "Não foi possivel tranferir"
	}
}

func main() {
	usuario1 := User{nome: "Gabriel", email: "g@gmail", saldo: 1000}
	usuario2 := User{nome: "Rafael", email: "r@gmail", saldo: 1000}
	fmt.Println(usuario1.transferir(&usuario1, &usuario2, 200))
}
