package SDK

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) InitRouter() {
	s.Router.Use(clientIPMiddleware())
	s.Router.Any("/", s.HandleDefault)
	s.Router.Any("/index.html", s.HandleDefault)
	s.Router.Any("/sdk/dataUpload", s.SdkDataUploadHandler)
	s.Router.POST("/apm/dataUpload", s.apmdataUpload)
	s.Router.POST("/common/h5log/log/batch", s.commonh5log)

	// 调度
	s.Router.GET("/query_dispatch", s.QueryDispatchHandler)
	s.Router.GET("/query_gateway", s.QueryGatewayHandler)
	s.Router.GET("/query_gateway_capture_cn", s.QueryGatewayHandlerCaptureCn)
	s.Router.GET("/query_gateway_capture_os", s.QueryGatewayHandlerCaptureOs)
	// 登录
	s.Router.POST("/account/risky/api/check", s.RiskyApiCheckHandler)
	Global := s.Router.Group("/:game_biz")
	{
		Global.GET("/combo/granter/api/getConfig", s.ComboGranterApiGetConfigHandler) // 获取服务器配置
		Global.POST("/mdk/shield/api/login", s.LoginRequestHandler)                   // 账号登录
		Global.POST("/mdk/shield/api/verify", s.VerifyRequestHandler)                 // token登录
		Global.POST("/combo/granter/login/v2/login", s.V2LoginRequestHandler)         // 获取combo token
		Global.GET("/mdk/agreement/api/getAgreementInfos", s.GetAgreementInfos)
	}

	// API
	s.Router.Any("/api", s.Api) // Api

	// 杂
	s.Router.POST("/data_abtest_api/config/experiment/list", s.GetExperimentListHandler)
}

func (s *Server) HandleDefault(c *gin.Context) {
	c.String(200, "hkrpg-go")
}
