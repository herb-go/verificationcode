package verificationcode

type Options struct {
	User  string
	Scene string
	Reset bool
}

func (o *Options) MergeUser(user string) *Options {
	o.User = user
	return o
}
func (o *Options) MergeScene(scene string) *Options {
	o.Scene = scene
	return o
}
func (o *Options) MergeReset(reset bool) *Options {
	o.Reset = reset
	return o
}

func NewOptions() *Options {
	return &Options{}
}
