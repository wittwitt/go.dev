package some

import (
	"github.com/wittwitt/go.dev/libs/internalpkg/some/internal/i2" // invalid use of internal package
)

func Some() {
	i2.I2() // 不能引用internal，internal只能内部用
}
