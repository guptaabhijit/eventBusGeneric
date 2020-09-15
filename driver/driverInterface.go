package driver

type Driver interface {
	Publish(channel string, message interface{})
	Subscribe(channel string, chanOrCallBack interface{})
}

type Client struct {
	driver Driver
}

func New(driver Driver) *Client {
	client := &Client{driver: driver}
	return client
}

func (c *Client) PublishMessage(channel string, message interface{}) {
	c.driver.Publish(channel, message)
}

func (c *Client) SubscribeMessage(channel string, chanOrCallBack interface{}) {
	c.driver.Subscribe(channel, chanOrCallBack)
}
