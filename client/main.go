package main

import (
	"encoding/binary"
	"flag"
	http_proxy "github.com/clearcodecn/http-proxy"
	"io"
	"log"
	"net"
	"net/http"
)

var (
	serverAddr string
	localAddr  string
)

func init() {
	flag.StringVar(&serverAddr, "s", "", "server address")
	flag.StringVar(&localAddr, "l", ":9001", "local address")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	server := new(s)
	log.Fatal(http.ListenAndServe(localAddr, server))
}

var tunnelEstablishedResponseLine = []byte("HTTP/1.1 200 Connection established\r\n\r\n")

type s struct{}

func (s *s) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpConn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		return
	}
	defer httpConn.Close()

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return
	}
	defer conn.Close()

	if r.Method == http.MethodConnect {
		_, err = httpConn.Write(tunnelEstablishedResponseLine)
		if err != nil {
			log.Println("failed to serve https", err)
			conn.Close()
			return
		}
	}
	host := r.URL.Host
	hostSize := []byte(host)
	b := make([]byte, 2+len(hostSize))
	binary.BigEndian.PutUint16(b[:2], uint16(len(hostSize)))
	copy(b[2:], hostSize)
	_, err = conn.Write(b)
	if err != nil {
		log.Println(err)
		return
	}

	dst := http_proxy.New(conn, conn)
	ori := http_proxy.New(httpConn, httpConn)

	go io.Copy(ori, dst)
	io.Copy(dst, ori)
}
