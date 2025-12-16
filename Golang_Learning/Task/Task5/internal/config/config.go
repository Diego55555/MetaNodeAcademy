// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

var G_secret = []byte("KpcL4pmBaYkxx")

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
}
