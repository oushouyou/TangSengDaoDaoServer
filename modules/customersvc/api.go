package customersvc

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/config"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/wkhttp"
	"github.com/golang-jwt/jwt"
)

type CustomersvcGroup struct {
	service IService
}

func New(ctx *config.Context) *CustomersvcGroup {
	return &CustomersvcGroup{
		service: NewService(ctx),
	}
}

func (c *CustomersvcGroup) Route(r *wkhttp.WKHttp) {
	r.POST("/v1/awakenTheGroup", c.AwakenTheGroup)
}

type custormerReq struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (c *CustomersvcGroup) AwakenTheGroup(r *wkhttp.Context) {
	// 基于JWT 获取用户名
	jwtToken := r.Request.Header.Get("jwtToken")
	// 解析token
	publicKeyBytes, err := os.ReadFile("tsdddata/rsa-public-key.pem")
	if err != nil {
		r.ResponseError(errors.New("获取rsa公钥文件出错"))
		return
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		r.ResponseError(errors.New("解释rsa公钥文件失败"))
		return
	}
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 确保算法是期望的
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		r.ResponseError(errors.New("解释jwtToken出错"))
		return
	}
	mc := token.Claims.(jwt.MapClaims)
	username := mc["username"].(string)
	email := mc["email"].(string)

	if username == "" || email == "" {
		r.ResponseError(errors.New("用户名和email不能为空"))
		return
	}
	customer := Customer{
		Username: username,
		Phone:    "",
		Email:    email,
	}
	resp, err := c.service.AwakenTheGroup(customer)
	if err != nil {
		r.ResponseError(err)
		return
	}

	// 基于JWT 刷新登陆 Token
	r.JSON(http.StatusOK, &CustomersvcGroupResp{
		GroupId: resp,
	})
}

type CustomersvcGroupResp struct {
	GroupId string `json:"group_id,omitempty"`
}
