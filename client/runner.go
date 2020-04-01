package client

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v6"
	pg "github.com/tradingAI/go/db/postgres"
	minio2 "github.com/tradingAI/go/s3/minio"
)

type Client struct {
	Conf  Conf
	DB    *gorm.DB
	Minio *minio.Client
	ID string
}

func New(conf Conf) (c *Client, err error) {
	// make server
	c = &Client{
		Conf: conf,
	}

	// Init db
	c.DB, err = pg.NewPostgreSQL(conf.DB)
	if err != nil {
		glog.Error(err)
		return
	}

	c.Minio, err = minio2.NewMinioClient(c.Conf.Minio)
	if err != nil {
		glog.Error(err)
		return
	}

	// TODO: use uuid
	c.ID = "test"

	return
}

func (c *Client) StartOrDie() (err error) {
	glog.Info("Starting runner")
	d := time.Duration(time.Second * c.Conf.HeartbeatSeconds)
	t := time.NewTicker(d)
    defer t.Stop()

    for {
            <- t.C
			c.hearbeat()
			c.listen()
    }
	return
}

func (c *Client) Free() {
	if err := c.DB.Close(); err != nil {
		glog.Warning(err)
	}
	return
}

func (c *Client) heartbeat() (err error) {
	glog.Infof("runner[%s] heartbeat", c.ID)
	// TODO: collect machine info call rpc hearbeat
	return
}

func (c *Client) listen() (err error) {
	glog.Infof("runner[%s] listen job from redis", c.ID)
	// TODO: listen redis status and excute actions
	return
}
