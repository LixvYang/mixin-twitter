package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lixvyang/mixin-twitter/internal/model"
	"github.com/lixvyang/mixin-twitter/internal/utils/errmsg"
)

func MixinOauth(c *gin.Context) {
	var code = c.Query("code")
	// access_token, _, err := mixin.AuthorizeToken(c, viper.GetString("mixin.ClientId"), viper.GetString("mixin.AppSecret"), code, "")
	access_token, _, err := mixin.AuthorizeToken(c, "30aad5a5-e5f3-4824-9409-c2ff4152724e", "dfbec1086bc839838e1921634aa062f87a624458918cb6929d7831fe676cfed6", code, "")
	if err != nil {
		log.Printf("AuthorizeToken: %v", err)
		// if viper.GetString("server.AppMode") == "release" {
		c.Redirect(http.StatusPermanentRedirect, "http://43.143.154.162")
		// } else {
		// 	c.Redirect(http.StatusPermanentRedirect, "http://localhost:8080")
		// }
		return
	}

	userinfo, err := GetUserInfo(access_token)
	if err != nil {
		// if viper.GetString("server.AppMode") == "release" {
		c.Redirect(http.StatusPermanentRedirect, "http://43.143.154.162")
		// } else {
		// 	c.Redirect(http.StatusPermanentRedirect, "http://localhost:8080")
		// }
		return
	}

	user := model.User{
		Uid:            userinfo.UserID,
		AvatarUrl:      userinfo.AvatarURL,
		FullName:       userinfo.FullName,
		SessionId:      userinfo.SessionID,
		IdentityNumber: userinfo.IdentityNumber,
	}

	session := sessions.Default(c)
	if checked := model.CheckUser(userinfo.UserID); checked != errmsg.SUCCSE {
		// 如果用户不存在
		if coded := model.CreateUser(&user); coded != errmsg.SUCCSE {
			fmt.Println("Get userInfo fail!!!")
		}
	} else {
		if coded := model.UpdateUser(userinfo.UserID, &user); coded != errmsg.SUCCSE {
			log.Println("Update userInfo fail!!!")
		}
		session.Clear()
	}
	sessionToken, _ := uuid.NewV4()
	session.Set("userId", user.Uid)
	session.Set("token", sessionToken.String())
	err = session.Save()
	if err != nil {
		log.Fatal(err.Error())
	}
	// if viper.GetString("server.AppMode") == "release" {
	c.Redirect(http.StatusPermanentRedirect, "http://43.143.154.162")
	// } else {
	// 	c.Redirect(http.StatusPermanentRedirect, "http://localhost:8080")
	// }
}

type UserInfo struct {
	Data Data `json:"data"`
}

type Data struct {
	UserID         string `json:"user_id"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	Biography      string `json:"biography"`
	AvatarURL      string `json:"avatar_url"`
	SessionID      string `json:"session_id"`
	PinToken       string `json:"pin_token"`
	PinTokenBase64 string `json:"pin_token_base64"`
	Phone          string `json:"phone"`
}

func GetUserInfo(access_token string) (Data, error) {
	// 形成请求
	var userInfoUrl = "https://api.mixin.one/me" // mixin
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return Data{}, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return Data{}, err
	}
	defer res.Body.Close()
	// 将响应的数据写入 userInfo 中，并返回
	var userInfo UserInfo
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return Data{}, err
	}
	return userInfo.Data, nil
}
