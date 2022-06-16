package handler

import "github.com/gin-gonic/gin"

func Route(route *gin.Engine) {
	// go 开发工具
	v1ToolsGroup := route.Group("/api/v1/tools")
	{
		// sql 转 gorm
		v1ToolsGroup.POST("/sql_to_gorm", SqlToGorm)
		// sql 转 go-zero
		v1ToolsGroup.POST("/sql_to_go_zero", SqlToGoZero)
		// sql 转 ent
		v1ToolsGroup.POST("/sql_to_ent", SqlToEnt)
		// yaml 转 go struct
		v1ToolsGroup.POST("/yaml_to_go", YamlToGo)
		// xml 转 json
		v1ToolsGroup.POST("/xml_to_json", XmlToJson)
		// sql 转 es
		v1ToolsGroup.POST("/sql_to_es", SqlToEs)
	}
}
