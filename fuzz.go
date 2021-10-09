// +build fuzz

package proxy

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/plugin/pkg/fuzz"
<<<<<<< HEAD
=======

	"github.com/coredns/caddy"
>>>>>>> e865259b26bf612690d989e932913f1b5c2b34fe
)

// Fuzz fuzzes proxy.
func Fuzz(data []byte) int {
	c := caddy.NewTestController("dns", "proxy . 8.8.8.8:53")
	up, err := NewStaticUpstreams(&c.Dispenser)
	if err != nil {
		return 0
	}
	p := &Proxy{Upstreams: &up}

	return fuzz.Do(p, data)
}
