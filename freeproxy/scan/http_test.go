package main

import (
	"encoding/base64"
	"testing"
)

func TestBase64(t *testing.T) {
	//sn := base64.StdEncoding.EncodeToString([]byte("ip=1.1.1.1&port=999&type=socks5"))
	sn := base64.StdEncoding.EncodeToString([]byte("1.1.1.1&999&socks5"))
	t.Log(sn)

	sc, _ := base64.StdEncoding.DecodeString(sn)
	t.Log(string(sc))

}

func TestB64(t *testing.T) {
	b64 := base64.NewEncoding("123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912")
	sn := b64.EncodeToString([]byte("1.1.1.1&999&socks5"))

	t.Log(sn)

	sc, err := b64.DecodeString(sn)
	t.Log(string(sc), err)

	sc, err = base64.StdEncoding.DecodeString(sn)
	t.Log(string(sc), err)

}
