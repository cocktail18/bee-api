package cmd

import (
	"context"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/router"
	"gitee.com/stuinfer/bee-api/sys"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	logger.InitLogger()
}

func Run(cfg *config2.AppConfig) {
	config2.AppConfigIns = cfg
	svr := &http.Server{
		Addr:    cfg.App.Listen,
		Handler: router.NewRouter(),
	}
	if cfg.DB.NeedInit {
		model.InitDB()
		sys.InitDB()
	}
	go func() {
		logrus.Infof("服务启动：%v", svr.Addr)
		err := svr.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	timeout := time.Duration(5) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}
