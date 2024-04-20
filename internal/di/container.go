package di

import "github.com/sarulabs/di"

// Container data
type Container struct {
	ctn di.Container
}

// NewContainer create a new container object
func NewContainer() *Container {
	builder, _ := di.NewBuilder()

	return &Container{
		ctn: builder.Build(),
	}
}

// Resolve container by name
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}
