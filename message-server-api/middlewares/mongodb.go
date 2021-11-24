package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

func MongoDB() martini.Handler {
	mInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "message-server-api",
		Timeout:  60 * time.Second,
	}
	session, err := mgo.DialWithInfo(mInfo)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(mInfo.Database))
	}
}
