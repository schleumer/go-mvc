package gomvc

type Controller struct {
	App     App
	Wrapper Wrapper
}

func (c Controller) Respond(str string) {
	c.Wrapper.Write(str)
}

