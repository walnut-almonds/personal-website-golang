package admin

import "github.com/gin-gonic/gin"

type IAuthCtrl interface {
	InsertSession(ctx *gin.Context)
	DeleteSession(ctx *gin.Context)
}

func newAuth(in digIn) IAuthCtrl {
	return &authCtrl{
		in: in,
	}
}

type authCtrl struct {
	in digIn
}

func (ctl *authCtrl) InsertSession(ctx *gin.Context) {

}

func (ctl *authCtrl) DeleteSession(ctx *gin.Context) {

}
