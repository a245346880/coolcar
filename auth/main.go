package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"flag"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//定义命令行参数的值
var addr = flag.String("addr", ":8081", "address to listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")
var privateKeyFile = flag.String("private_key_file", "auth/private.key", "private key file")
var wechatAppID = flag.String("wechat_app_id", "<APPID>", "wechat app id")
var wechatAppSecret = flag.String("wechat_app_secret", "<APPSERET>", "wechat app secret")

func main() {
	//是否解析命令行参数完成
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create loggerr:%v", err)
	}
	//返回一个空值上下文
	c := context.Background()
	//构建Mongo认证信息
	credential := &options.Credential{
		AuthSource:  "admin",
		Username:    "root",
		Password:    "123456",
		PasswordSet: true,
	}
	mongoClinet, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI).SetAuth(*credential))

	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}
	pkFile, err := os.Open(*privateKeyFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}
	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}
	err = server.RunGRPCServer(&server.GRPCConfig{
		Name:   "auth",
		Addr:   *addr,
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIDResolver: &wechat.Service{
					AppID:     *wechatAppID,
					AppSecret: *wechatAppSecret,
				},
				Mongo:          dao.NewMongo(mongoClinet.Database("coolcar")),
				TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privKey),
				TokenExpire:    time.Hour,
				Logger:         logger,
			})
		},
	})
	logger.Sugar().Fatal(err)
}
