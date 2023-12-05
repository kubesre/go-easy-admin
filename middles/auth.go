package middles

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
	"strconv"
	"time"
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

// 用户登录成功后被调用，并且会接收一个参数 data，表示用户信息
// 用户登录是执行顺序 2
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v.ID,
			"username":      v.UserName,
		}
	}
	// TODO 将用户数据同步进缓存

	return jwt.MapClaims{}
}

// 处理jwt 1
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var jwtClaim models.User
	userID, _ := strconv.Atoi(fmt.Sprintf("%d", claims[jwt.IdentityKey]))
	userNameStr := fmt.Sprintf("%s", claims["username"])
	jwtClaim.ID = uint(userID)
	jwtClaim.UserName = userNameStr
	return &jwtClaim
}

// 用户登录是执行顺序 1
func loginFunc(c *gin.Context) (interface{}, error) {
	var loginUser models.User
	if err := c.ShouldBind(&loginUser); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userName := loginUser.UserName
	password := loginUser.Password
	ok, id := dao.NewUserInterface().ExitUser(userName, password)
	if ok {
		loginUser.ID = id
		return &loginUser, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// 处理jwt 2
func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok {
		c.Set("username", v.UserName)
		return true
	}
	return false
}

// 处理jwt 3
func unauthorized(c *gin.Context, code int, message string) {
	global.ReturnContext(c).Failed("failed", message)
	return
}

// 用户登录是执行顺序 3
func loginResponse(c *gin.Context, code int, token string, expires time.Time) {
	global.ReturnContext(c).Successful("success", map[string]string{
		"token":   token,
		"expires": expires.Format("2006-01-02 15:04:05"),
	})
	return
}
func logoutResponse(c *gin.Context, code int) {
	global.ReturnContext(c).Successful("success", "退出成功")
	return
}

func refreshResponse(c *gin.Context, code int, token string, expires time.Time) {
	global.ReturnContext(c).Successful("success", map[string]string{
		"token":   token,
		"expires": expires.Format("2006-01-02 15:04:05"),
	})
	return
}
