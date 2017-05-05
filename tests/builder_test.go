package tests

import (
	"fmt"
	"github.com/jcherianucla/godfried/app/helpers"
	"github.com/jcherianucla/godfried/app/models"
)

func main() {
	user := models.User
	b := helpers.Builder
	b.BuildInsertion(user)
}
