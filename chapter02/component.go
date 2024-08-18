package chapter02

const DefaultUpdateOrder = 100

type Component interface {
	Update(deltaTime float32)
	GetOwner() Actor
	GetUpdateOrder() int
	Destroy()
}

type component struct {
	owner       Actor
	updateOrder int
}

func NewComponent(owner Actor, updateOrder int) Component {
	c := &component{
		owner:       owner,
		updateOrder: updateOrder,
	}

	return c
}

func (c *component) Update(deltaTime float32) {
}

func (c *component) GetOwner() Actor {
	return c.owner
}

func (c *component) GetUpdateOrder() int {
	return c.updateOrder
}

func (c *component) Destroy() {
	c.owner.RemoveComponent(c)
}
