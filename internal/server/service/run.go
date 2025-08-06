package service

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/gin-gonic/gin"

	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Exam-Center
// @version		1.0
// @description	Exam-Center API.
// @termsOfService	https://descart.com
// @contact.name	Alex
// @contact.url	https://descart.com
// @contact.email	alexli@descart.com
// @host			localhost:46689
// @BasePath		/v1/api
func (s *Services) Doc(ctx context.Context, wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", s.cfg.AddressDoc())
	if err != nil {
		panic(err)
	}
	s.logger.Infof("Listen doc service on %s", s.cfg.AddressDoc())

	publicEngine := gin.Default()
	{
		v1Doc := publicEngine.Group("/v1/docs")
		v1Doc.GET("/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	}
	publicService := &http.Server{
		Handler: publicEngine.Handler(),
	}

	goo.Goo(func() {
		wg.Add(1)
		defer wg.Done()
		err = publicService.ServeTLS(listener, s.cfg.Cert, s.cfg.Key)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}, func(err error) { panic(err) })

	goo.Go(func() {
		wg.Add(1)
		defer wg.Done()
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
		defer cancel()
		_ = publicService.Shutdown(ctx)
	})
}

func (s *Services) Serve(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	listener, err := net.Listen("tcp", s.cfg.Address())

	s.logger.Printf("url is", s.cfg.Address())
	if err != nil {
		panic(err)
	}
	s.logger.Infof("Listen service on %s", s.cfg.Address())
	service := s.Http

	engine := gin.Default()
	V1API := engine.Group("v1/api")

	s.logger.Printf("service.RegisterRouter runnint")
	service.RegisterRouter(V1API)

	apiService := &http.Server{
		Handler: engine.Handler(),
	}
	s.logger.Printf("Serve goo")

	wg.Add(1)
	goo.Goo(func() {
		defer wg.Done()
		// todo: add ssl & sslkey
		s.logger.Printf("s.cfg.Cert is", s.cfg.Cert)
		err = apiService.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}, func(err error) { panic(err) })

	wg.Add(1)
	goo.Go(func() {
		defer wg.Done()
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
		defer cancel()
		_ = apiService.Shutdown(ctx)
	})

}
