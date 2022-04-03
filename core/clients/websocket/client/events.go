package client

import "github.com/kyaxcorp/go-core/core/helpers/function"

func (c *Client) OnConnect(name string, callback OnConnect) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onConnect.Set(name, callback)
	return true
}

func (c *Client) OnConnectRemove(name string) {
	c.onConnect.Del(name)
}

func (c *Client) OnReceive(name string, callback OnReceive) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReceive.Set(name, callback)
	return true
}

// OnMessage analog to OnReceive
func (c *Client) OnMessage(name string, callback OnReceive) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReceive.Set(name, callback)
	return true
}

func (c *Client) OnReceiveRemove(name string) {
	c.onReceive.Del(name)
}

// OnMessageRemove analog to OnReceiveRemove
func (c *Client) OnMessageRemove(name string) {
	c.onReceive.Del(name)
}

func (c *Client) OnText(name string, callback OnText) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onText.Set(name, callback)
	return true
}

func (c *Client) OnTextRemove(name string) {
	c.onText.Del(name)
}

func (c *Client) OnBinary(name string, callback OnBinary) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onBinary.Set(name, callback)
	return true
}

func (c *Client) OnBinaryRemove(name string) {
	c.onBinary.Del(name)
}

func (c *Client) OnSend(name string, callback OnSend) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onSend.Set(name, callback)
	return true
}

func (c *Client) OnSendRemove(name string) {
	c.onSend.Del(name)
}

func (c *Client) OnSendError(name string, callback OnSendError) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onSendError.Set(name, callback)
	return true
}

func (c *Client) OnSendErrorRemove(name string) {
	c.onSendError.Del(name)
}

func (c *Client) OnReadError(name string, callback OnReadError) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReadError.Set(name, callback)
	return true
}

func (c *Client) OnReadErrorRemove(name string) {
	c.onReadError.Del(name)
}

func (c *Client) OnError(name string, callback OnError) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onError.Set(name, callback)
	return true
}

func (c *Client) OnErrorRemove(name string) {
	c.onError.Del(name)
}

func (c *Client) OnBeforeDisconnect(name string, callback OnBeforeDisconnect) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onBeforeDisconnect.Set(name, callback)
	return true
}

func (c *Client) OnBeforeDisconnectRemove(name string) {
	c.onBeforeDisconnect.Del(name)
}

func (c *Client) OnDisconnect(name string, callback OnDisconnect) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onDisconnect.Set(name, callback)
	return true
}

func (c *Client) OnDisconnectRemove(name string) {
	c.onDisconnect.Del(name)
}

func (c *Client) OnLinkDisconnect(name string, callback OnLinkDisconnect) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onLinkDisconnect.Set(name, callback)
	return true
}

func (c *Client) OnLinkDisconnectRemove(name string) {
	c.onLinkDisconnect.Del(name)
}

func (c *Client) OnReconnected(name string, callback OnReconnected) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReconnected.Set(name, callback)
	return true
}

func (c *Client) OnReconnectedRemove(name string) {
	c.onReconnected.Del(name)
}

func (c *Client) OnBeforeConnectToServer(name string, callback OnBeforeConnectToServer) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onBeforeConnectToServer.Set(name, callback)
	return true
}

func (c *Client) OnBeforeConnectToServerRemove(name string) {
	c.onBeforeConnectToServer.Del(name)
}

func (c *Client) OnConnectError(name string, callback OnConnectError) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onConnectError.Set(name, callback)
	return true
}

func (c *Client) OnConnectErrorRemove(name string) {
	c.onConnectError.Del(name)
}

func (c *Client) OnConnectFailed(name string, callback OnConnectFailed) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onConnectFailed.Set(name, callback)
	return true
}

func (c *Client) OnConnectFailedRemove(name string) {
	c.onConnectFailed.Del(name)
}

func (c *Client) OnConnectFailedAll(name string, callback OnConnectFailedAll) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onConnectFailedAll.Set(name, callback)
	return true
}

func (c *Client) HasOnConnectFailedAll(name string) bool {
	return c.onConnectFailedAll.Has(name)
}

func (c *Client) OnConnectFailedAllRemove(name string) {
	c.onConnectFailedAll.Del(name)
}

func (c *Client) OnConnectSuccess(name string, callback OnConnectSuccess) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onConnectSuccess.Set(name, callback)
	return true
}

func (c *Client) HasOnConnectSuccess(name string) bool {
	return c.onConnectSuccess.Has(name)
}

func (c *Client) RemoveOnConnectSuccess(name string) {
	c.onConnectSuccess.Del(name)
}

func (c *Client) OnConnectSuccessRemove(name string) {
	c.onConnectSuccess.Del(name)
}

func (c *Client) OnTerminate(name string, callback OnTerminate) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onTerminate.Set(name, callback)
	return true
}

func (c *Client) OnTerminateRemove(name string) {
	c.onTerminate.Del(name)
}

func (c *Client) OnStopConnectingFinish(name string, callback OnStopConnectingFinish) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onStopConnectingFinish.Set(name, callback)
	return true
}

func (c *Client) OnStopConnectingFinishRemove(name string) {
	c.onStopConnectingFinish.Del(name)
}

func (c *Client) OnReconnecting(name string, callback OnReconnecting) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReconnecting.Set(name, callback)
	return true
}

func (c *Client) OnReconnectingRemove(name string) {
	c.onReconnecting.Del(name)
}

func (c *Client) OnReconnectFailed(name string, callback OnReconnectFailed) bool {
	if !function.IsCallable(callback) || name == "" {
		return false
	}
	c.onReconnectFailed.Set(name, callback)
	return true
}

func (c *Client) OnReconnectFailedRemove(name string) {
	c.onReconnectFailed.Del(name)
}
