package gomvc

import (
	"fmt"
)

type Controller struct {
	App     App
	Wrapper Wrapper
}

func (c Controller) Respond(str string) {
	c.Wrapper.Write(str)
}

func (c Controller) SetApp(a App) {
	fmt.Println(c)
	fmt.Println(a)
	//c.App = a
}

func (c Controller) SetWrapper(w Wrapper) {
	//c.Wrapper = w
}
