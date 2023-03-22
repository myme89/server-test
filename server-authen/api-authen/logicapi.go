package api

import (
	"context"
	"fmt"
	"server-test/server-authen/pb_authen"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
)

func GetLoginToken(username string) (string, error) {
	//generate JWT
	var jwtString string
	// mySigningKey, err := helpers.ConfigGet("jwt", "signkey")
	// if err != nil {
	// 	log.Error("cannot get sign key for JWT , set to default")
	// 	mySigningKey = "weriwoxcr342f234"
	// }

	mySigningKey := "weriwoxcr342f234"
	var mapClaim jwt.StandardClaims
	mapClaim.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()
	mapClaim.Issuer = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaim)
	jwtString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func (serverAuthen *ServerAuthen) SignUp(ctx context.Context, res *pb_authen.UserResquest) (*pb_authen.UserRespone, error) {

	infoSignUp := res.GetUserinfo()

	fmt.Println("username: ", infoSignUp.Username)
	fmt.Println("password: ", infoSignUp.Password)
	// fmt.Println("lastname: ", infoSignUp.Lastname)
	// fmt.Println("firstname: ", infoSignUp.Firstname)

	resp, err := serverAuthen.clientDatabase.SignUpAcc(ctx, infoSignUp.Username, infoSignUp.Password, infoSignUp.Firstname, infoSignUp.Lastname)

	if err != nil {
		fmt.Println(err)
	}

	noti := &pb_authen.UserRespone{Noti: resp}

	return noti, nil
}

func (serverAuthen *ServerAuthen) SignIn(ctx context.Context, res *pb_authen.SignInResquest) (*pb_authen.SignInRespone, error) {

	infoSignUp := res.GetUserinfo()

	resp, err := serverAuthen.clientDatabase.LoginAccClient(ctx, infoSignUp.Username, infoSignUp.Password)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "User or password is incorrect")
	}

	token, err := GetLoginToken(resp.Userinfo.Id)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Create Token Failed")
	}

	userInfo := &pb_authen.UserInfo{
		Username:  resp.Userinfo.Username,
		Lastname:  resp.Userinfo.Lastname,
		Firstname: resp.Userinfo.Firstname,
	}

	noti := &pb_authen.SignInRespone{
		Userinfo: userInfo,
		Token:    token,
	}

	return noti, nil
}

func (serverAuthen *ServerAuthen) AuthenToken(ctx context.Context, res *pb_authen.AuthenTokenResquest) (*pb_authen.AuthenTokenRespone, error) {

	token := res.GetToken()

	tokenParse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		signkey := "weriwoxcr342f234"
		// if err != nil {
		// 	log.Error("get signkey in JWT authen error ", err)
		// 	return nil, status.Errorf(codes.Unimplemented, "Cannot parse token claim to map claim")

		// }
		return []byte(signkey), nil
	})
	var idUserName string
	if err == nil && tokenParse.Valid {
		userInfo, ok := tokenParse.Claims.(jwt.MapClaims)
		if !ok {
			log.Error("Cannot parse token claim to map claim ")
			return nil, status.Errorf(codes.Unauthenticated, "Cannot parse token claim to map claim")
		}
		idUserName, ok = userInfo["iss"].(string)
		if !ok {
			log.Error("Cannot parse contact claim ")
			return nil, status.Errorf(codes.Unauthenticated, "Cannot parse contact claim")

		}
		log.Info(" Get request from user :", idUserName)
	}

	return &pb_authen.AuthenTokenRespone{Iduser: idUserName}, nil
}