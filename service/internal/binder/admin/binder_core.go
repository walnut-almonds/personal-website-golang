package admin

import "go.uber.org/dig"

func provideCore(binder *dig.Container) {
	if err := binder.Provide(auth.NewAuth); err != nil {
		panic(err)
	}
}
