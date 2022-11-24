package concurrent

import (
	"github.com/go-leo/leo/global"
)

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				global.Logger().Errorf("%s %+v", "[Panic]", err)
			}
		}()

		f()
	}()
}
