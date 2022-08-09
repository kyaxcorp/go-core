package connection

import "github.com/gookit/color"

func (c *ConnDetails) Process() {
	c.generateDetails()
	// Debug the connection

	c.Logger.Logger.Info().
		Str("host", c.Host).
		Str("client_ip", c.ClientIPAddress).
		Int("client_port", c.ClientPort).
		Str("user_agent", c.UserAgent).
		Str("remote_addr", c.RemoteAddr).
		Str("request_path", c.RequestPath).
		Str("referer", c.Referer).
		Msg(color.Style{color.LightGreen, color.OpBold}.Render("new connection"))

}
