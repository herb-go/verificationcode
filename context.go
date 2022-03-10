package verificationcode

type Context struct {
	Options *Options
	Store   Store
	Client  *Client
}

func CreateContext(store Store, client *Client, opt *Options) *Context {
	return &Context{
		Options: opt,
		Store:   store,
		Client:  client,
	}
}
