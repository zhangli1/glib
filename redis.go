package lib

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Host   string
	Port   string
	Passwd string
	DbNum  int
	Conn   redis.Conn
}

func NewRedis(host string, port string, passwd string, db_num int) *Redis {
	c := conn(host, port, passwd, db_num)
	return &Redis{Host: host, Port: port, Passwd: passwd, DbNum: db_num, Conn: c}
}

func conn(host string, port string, passwd string, db_num int) redis.Conn {
	c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err)
	}
	err = c.Send("auth", passwd)
	if err != nil {
		fmt.Println(err)
	}
	err = c.Send("select", db_num)
	if err != nil {
		fmt.Println(err)
	}

	return c
}

func (self *Redis) Set(key string, value string) error {
	if self.Conn == nil {
		self.Conn = conn(self.Host, self.Port, self.Passwd, self.DbNum)
	}
	_, err := self.Conn.Do("SET", key, value)
	return err
}

func (self *Redis) GET(key string) (interface{}, error) {
	if self.Conn == nil {
		self.Conn = conn(self.Host, self.Port, self.Passwd, self.DbNum)
	}
	re, err := self.Conn.Do("GET", key)
	return re, err
}

func (self *Redis) Close() {
	if self.Conn != nil {
		self.Conn.Close()
	}
}
