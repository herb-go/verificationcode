package verificationcode

type Context struct {
	User   string
	Store  Store
	Client *Client
}

func CreateContext(user string, store Store, client *Client) *Context {
	return &Context{
		User:   user,
		Store:  store,
		Client: client,
	}
}
