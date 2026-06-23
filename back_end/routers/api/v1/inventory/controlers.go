package inventory

import (
	"alto_server/common/pkg/e"
	. "alto_server/common/utils"

	"github.com/gin-gonic/gin"
)

func todoInventoryHandler(c *gin.Context) {
	RES(c, e.SUCCESS, gin.H{"status": e.SUCCESS, "message": "Success", "data": "2323232"})
}
