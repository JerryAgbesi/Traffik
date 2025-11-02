package server

import (
	"net/url"
	"time"
)

type Server struct{
	URL *url.URL `json:"url"`
	ActiveConns int32 `json:"active_conns"`
	ReponseTime time.Duration `json:"response_time"`
}
