/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/7
*/

package middles

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	reqLogin "go-easy-admin/internal/model/request/login"
	"go-easy-admin/internal/model/system"
	"go-easy-admin/pkg/controller/login"
	"go-easy-admin/pkg/global"
	"go-easy-admin/pkg/utils"
)

func InitAuth() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           viper.GetString("jwt.realm"),                              // jwt标识
		Key:             []byte(viper.GetString("jwt.key")),                        // 服务端密钥
		Timeout:         time.Hour * time.Duration(viper.GetInt("jwt.timeout")),    // token过期时间
		MaxRefresh:      time.Hour * time.Duration(viper.GetInt("jwt.maxRefresh")), // token最大刷新时间(RefreshToken过期时间=Timeout+MaxRefresh)
		PayloadFunc:     payloadFunc,                                               // 有效载荷处理
		IdentityHandler: identityHandler,                                           // 解析Claims
		Authenticator:   loginFunc,                                                 // 校验token的正确性, 处理登录逻辑
		Authorizator:    authorizator,                                              // 用户登录校验成功处理
		Unauthorized:    unauthorized,                                              // 用户登录校验失败处理
		LoginResponse:   loginResponse,                                             // 登录成功后的响应
		LogoutResponse:  logoutResponse,                                            // 登出后的响应
		RefreshResponse: refreshResponse,                                           // 刷新token后的响应
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",        // 自动在这几个地方寻找请求中的token
		TokenHeadName:   "Bearer",                                                  // header名称
		TimeFunc:        time.Now,
	})
	return authMiddleware, err
}

// 登录1
func loginFunc(ctx *gin.Context) (interface{}, error) {
	var loginUser reqLogin.ReqLogin
	if err := ctx.ShouldBind(&loginUser); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	path := strings.Split(ctx.Request.RequestURI, "?")[0]
	// 通用登录
	if !strings.Contains(path, "ldap") {
		var aesLogin = loginUser
		utils.TagAes(&aesLogin)
		err, user := login.NewSysLogin(ctx).GeneralLogin(&aesLogin)
		if err == nil {
			if user.(*system.User).Status != 1 {
				return nil, errors.New("该用户已被禁用")
			}
			ctx.Set("id", user.(*system.User).ID)
			return user, nil
		}

	} else {
		// ldap登录
		err, user := login.NewSysLogin(ctx).LdapLogin(&loginUser)
		if err == nil {
			if user.(*system.User).Status != 1 {
				return nil, errors.New("该用户已被禁用")
			}
			ctx.Set("id", user.(*system.User).ID)
			return user, nil

		}
	}
	return nil, errors.New("用户名或密码错误")
}

// 登录2
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*system.User); ok {
		return jwt.MapClaims{
			"id":            v.ID,
			jwt.IdentityKey: v.ID,
			"username":      v.Username,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	var jwtClaim system.User
	userID, _ := claims[jwt.IdentityKey].(float64)
	userNameStr := fmt.Sprintf("%s", claims["username"])
	jwtClaim.ID = uint(userID)
	jwtClaim.Username = userNameStr
	return &jwtClaim
}

func authorizator(data interface{}, ctx *gin.Context) bool {
	if v, ok := data.(*system.User); ok {
		ctx.Set("username", v.Username)
		ctx.Set("id", v.ID)
		return true
	}
	return false
}

func unauthorized(ctx *gin.Context, code int, message string) {
	response := gin.H{
		"code": code,
		"msg":  "failed",
		"data": message,
	}
	ctx.JSON(http.StatusOK, response)
	return
}

// 登录3
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	userID, _ := ctx.Keys["id"]
	global.ReturnContext(ctx).Successful("success", map[string]interface{}{
		"token":   token,
		"id":      userID,
		"expires": expire.Format("2006-01-02 15:04:05"),
	})
	return
}

func logoutResponse(ctx *gin.Context, code int) {
	global.ReturnContext(ctx).Successful("success", "退出成功")
	return
}

func refreshResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	userID, _ := ctx.Keys["id"]
	global.ReturnContext(ctx).Successful("success", map[string]interface{}{
		"token":   token,
		"id":      userID,
		"expires": expire.Format("2006-01-02 15:04:05"),
	})
	return
}
