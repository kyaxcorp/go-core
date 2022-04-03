package connection

import "github.com/gin-gonic/gin"

func GetConnectionDetailsFromCtx(c *gin.Context) *ConnDetails {
	var _connDetails *ConnDetails

	connDetails, connIfExists := c.Get(HttpContextConnDetailsKey)
	if connIfExists && connDetails != nil {
		// Set the data into the Client NonPtrObj
		_connDetails = connDetails.(*ConnDetails)
	} else {
		_connDetails = &ConnDetails{}
	}
	return _connDetails
}
