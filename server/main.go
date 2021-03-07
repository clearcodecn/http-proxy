package main

import (
	"encoding/binary"
	"flag"
	http_proxy "github.com/clearcodecn/http-proxy"
	"io"
	"log"
	"net"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":9000", "listen address")
}

func main() {
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("server listen and serve at: ", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		go func() {
			// mark host&port
			b := make([]byte, 2)
			_, err := conn.Read(b)
			if err != nil {
				log.Println(err)
				return
			}
			size := binary.BigEndian.Uint16(b)
			b = make([]byte, size)
			_, err = io.ReadFull(conn, b)
			if err != nil {
				log.Println(err)
				return
			}
			defer conn.Close()

			log.Printf("[proxy]  %s <----> %s\n", conn.LocalAddr().String(), string(b))

			dst, err := net.Dial("tcp", string(b))
			if err != nil {
				log.Println(err)
				return
			}
			defer dst.Close()
			
			// start proxy
			cc := http_proxy.New(dst, dst)
			ri := http_proxy.New(conn, conn)
			go io.Copy(cc, ri)
			io.Copy(ri, cc)
		}()
	}
}
