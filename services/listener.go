package services

import (
	"database/sql"
	"encoding/json"
	"github.com/metaleaf-io/gator/conf"
	"github.com/metaleaf-io/log"
	"net"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	k   = 1024
)

func StartListener(appConfig *conf.AppConfig, db *sql.DB) {
	port := ":" + strconv.FormatInt(int64(appConfig.Gator.Port), 10)
	log.Info("Starting UDP log listener", log.Int16("port", appConfig.Gator.Port))

	addr := resolve(port)
	sock := listen(addr)
	go receive(sock, db)
}

func resolve(address string) *net.UDPAddr {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic(err)
	}
	return addr
}

func listen(addr *net.UDPAddr) *net.UDPConn {
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	return sock
}

func receive(sock *net.UDPConn, db *sql.DB) {
	// Buffer set to the maximum size of a UDP datagram
	buf := make([]byte, 64*k)

	for {
		// Receive the datagram (this is a blocking call).
		l, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		// Decode the datagram.
		var j map[string]interface{}
		json.Unmarshal(buf[0:l], &j)

		// Re-encode the structured data fields to a string that can be saved in the database.
		fields, err := json.Marshal(j["fields"])
		if err != nil {
			log.Error("Could not encode structured fields.")
			fields = []byte("{}")
		}

		// Execute the database insert in a goroutine so that it doesn't slow down the server.
		go func() {
			// Save the message in the database.
			if _, err := db.Exec("call gator_save($1, $2, $3, $4, $5)", j["time"], j["name"], j["level"], fields, j["message"]); err != nil {
				log.Error(err.Error())
			}
		}()
	}
}
