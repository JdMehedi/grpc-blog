package main

import (
	"blog/cms/handler"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	bgv "blog/gunk/v1/post"
)

func main(){
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("cms/env/config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("error loading configuration: %v", err)
	}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	 store := sessions.NewCookieStore([]byte(config.GetString("session.secret")))
	
	conn, err:=grpc.Dial(
		fmt.Sprintf("%s:%s",config.GetString("blog.host"),config.GetString("blog.port")),
		grpc.WithInsecure(),
	)
	if err!=nil{
		log.Fatal(err)
	}
	tc:=bgv.NewPostServiceClient(conn)
	r:=handler.New( decoder, store,tc)
	host,port:=config.GetString("server.host"),config.GetString("server.port")
	
	log.Printf("Server is starting on: http://%s:%s\n",host,port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s",host,port),r); err !=nil{
		log.Fatal(err)
	}
}