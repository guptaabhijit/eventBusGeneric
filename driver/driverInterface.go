package driver

type Driver interface {
	Publish(channel string, message interface{})
	Subscribe(channel string, chanOrCallBack interface{})
}
