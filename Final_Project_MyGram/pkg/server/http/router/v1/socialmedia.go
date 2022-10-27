package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/config/gin"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/socialmedia"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/middleware"
	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/server/http/router"
)

type SocialMediaRouterImpl struct {
	ginEngine          engine.HttpServer
	routerGroup        *gin.RouterGroup
	socialMediaHandler socialmedia.SocialMediaHandler
	authMiddleware     middleware.AuthMiddleware
}

func (p *SocialMediaRouterImpl) get() {
	p.routerGroup.GET("", p.authMiddleware.CheckJWTAuth, p.socialMediaHandler.GetSocialMediasHdl)
}

func (p *SocialMediaRouterImpl) post() {
	p.routerGroup.POST("", p.authMiddleware.CheckJWTAuth, p.socialMediaHandler.CreateSocialMediaHdl)
}

func (p *SocialMediaRouterImpl) put() {
	p.routerGroup.PUT("/:socialMediaId", p.authMiddleware.CheckJWTAuth, p.socialMediaHandler.UpdateSocialMediaHdl)
}

func (p *SocialMediaRouterImpl) delete() {
	p.routerGroup.DELETE("/:socialMediaId", p.authMiddleware.CheckJWTAuth, p.socialMediaHandler.DeleteSocialMediaHdl)
}

func (p *SocialMediaRouterImpl) Routers() {
	p.get()
	p.post()
	p.put()
	p.delete()
}

func NewSocialMediaRouter(ginEngine engine.HttpServer, socialMediaHandler socialmedia.SocialMediaHandler, authMiddleware middleware.AuthMiddleware) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/mygram/v1/socialmedias")
	return &SocialMediaRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, socialMediaHandler: socialMediaHandler, authMiddleware: authMiddleware}
}
