package verificationcode

import "context"

const ClientTypeWeb = ClientType(0)

type ClientType int

func (t ClientType) CreateClient(addr string, ctx context.Context) *Client {
	return CreateClient(t, addr, ctx)
}

type Client struct {
	Type    ClientType
	Addr    string
	Meta    map[string]string
	Context context.Context
}

func (c *Client) CreateContext(store Store, opt *Options) *Context {
	return CreateContext(store, c, opt)
}
func (c *Client) GetMetaValue(key string) string {
	if c.Meta == nil {
		return ""
	}
	return c.Meta[key]
}
func (c *Client) MergeMetaValue(key string, value string) *Client {
	if c.Meta == nil {
		c.Meta = map[string]string{}
	}
	c.Meta[key] = value
	return c
}
func (c *Client) MergeMeta(m map[string]string) *Client {
	c.Meta = m
	return c
}
func CreateClient(clienttype ClientType, addr string, ctx context.Context) *Client {
	return &Client{
		Type:    clienttype,
		Addr:    addr,
		Context: ctx,
	}
}
