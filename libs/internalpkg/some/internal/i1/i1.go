package i1

import (
	"fmt"

	"github.com/wittwitt/go.dev/libs/internalpkg/some/internal/i2"

	"github.com/wittwitt/go.dev/libs/internalpkg/some/internal"
)

func I1() {
	fmt.Println("internal i2:I1")

	i2.I2()

	internal.Say()
}
