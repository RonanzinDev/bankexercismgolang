# Exercício de transferência bancária

Primeiro crio uma struct de User e defino seus campos

```golang
type User struct {
    nome  string
    email string
    saldo int
}
```

Segundamente eu faço uma função para verificar se o saldo é suficiente para completar a transação.

```golang
func validarSaldo(usuario User, valor int) bool {
   return usuario.saldo >= valor
}
```

Após isso criamos a função para a transferência:

```golang
func (u *User) transferir(usuario, usuario2 *User, valor int) (User, string) {
    if validarSaldo(*u, valor) {
        u.saldo += valor
        return (*u), "\n Transferencia completa"
    } else {
        return *u, "Não foi possível transferir"
    }
}
```

E assim chamando a função `main()`

```golang
func main() {
    usuario1 := User{nome: "Gabriel", email: "g@gmail", saldo: 1000}
    usuario2 := User{nome: "Rafael", email: "r@gmail", saldo: 1000}
    fmt.Println(usuario1.transferir(&usuario1, &usuario2, 200))
}
```

E o retorno será :

```go
{Gabriel g@gmail 1200}
 Transferencia completa
```
