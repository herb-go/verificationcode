package verificationcode

type Context struct {
	Store  Store
	Client *Client
}

func CreateContext(user string, store Store, client *Client) *Context {
	return &Context{
		Store:  store,
		Client: client,
	}
}
