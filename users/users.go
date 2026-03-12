package users

import (
	"fmt"
	"time"

	"github.com/freneklopez/godesde0/modelos"
)

func AltaUser() {
	u := new(modelos.User)
	u.AddUser(30, "Eric", time.Now(), true)
	fmt.Println(u)

}
