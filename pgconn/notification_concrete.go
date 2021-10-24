package pgconn

import "github.com/jackc/pgconn"

type ConcreteNotification struct {
	*pgconn.Notification
}

func (c *ConcreteNotification) Channel() string {
	if c == nil || c.Notification == nil {
		return ""
	}
	return c.Notification.Channel
}

func (c *ConcreteNotification) Payload() string {
	if c == nil || c.Notification == nil {
		return ""
	}
	return c.Notification.Payload
}

func (c *ConcreteNotification) PID() uint32 {
	if c == nil || c.Notification == nil {
		return 0
	}
	return c.Notification.PID
}

func (c *ConcreteNotification) SetChannel(s string) Notification {
	if c == nil {
		panic("Cannot set a channel on a nil notification")
	}
	if c.Notification == nil {
		c.Notification = &pgconn.Notification{}
	}
	c.Notification.Channel = s
	return c
}

func (c *ConcreteNotification) SetPayload(s string) Notification {
	if c == nil {
		panic("Cannot set a payload on a nil notification")
	}
	if c.Notification == nil {
		c.Notification = &pgconn.Notification{}
	}
	c.Notification.Payload = s
	return c
}

func (c *ConcreteNotification) SetPID(u uint32) Notification {
	if c == nil {
		panic("Cannot set a PID on a nil notification")
	}
	if c.Notification == nil {
		c.Notification = &pgconn.Notification{}
	}
	c.Notification.PID = u
	return c
}
