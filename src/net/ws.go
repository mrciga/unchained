package net

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/KenshiTech/unchained/bls"
	"github.com/KenshiTech/unchained/plugins/uniswap"

	"github.com/gorilla/websocket"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/vmihailenco/msgpack/v5"
)

var signers *xsync.MapOf[*websocket.Conn, uniswap.Signer]
var upgrader = websocket.Upgrader{} // use default options
var addr = "0.0.0.0:9123"

func processHello(conn *websocket.Conn, messageType int, payload []byte) error {

	var signer uniswap.Signer
	err := msgpack.Unmarshal(payload, &signer)

	if err != nil {
		err = conn.WriteMessage(messageType, []byte("packet.invalid"))
		if err != nil {
			fmt.Println("write:", err)
			return err
		}
		return nil
	}

	if signer.Name == "" || len(signer.PublicKey) != 48 {
		conn.WriteMessage(messageType, []byte("conf.invalid"))
		return errors.New("conf.invalid")
	}

	signers.Store(conn, signer)
	err = conn.WriteMessage(messageType, []byte("conf.ok"))

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	return nil
}

func processPriceReport(conn *websocket.Conn, messageType int, payload []byte) error {

	signer, ok := signers.Load(conn)

	if !ok {
		conn.WriteMessage(messageType, []byte("hello.missing"))
		return errors.New("hello.missing")
	}

	var report uniswap.PriceReport
	err := msgpack.Unmarshal(payload, &report)

	if err != nil {
		return nil
	}

	toHash, err := msgpack.Marshal(&report.PriceInfo)

	if err != nil {
		return nil
	}

	hash, err := bls.Hash(toHash)

	if err != nil {
		return nil
	}

	signature, err := bls.RecoverSignature(report.Signature)

	if err != nil {
		return nil
	}

	pk, err := bls.RecoverPublicKey(signer.PublicKey)

	if err != nil {
		return nil
	}

	ok, _ = bls.Verify(signature, hash, pk)

	message := []byte("signature.invalid")
	if ok {
		message = []byte("signature.accepted")
		uniswap.RecordSignature(
			signature,
			signer,
			report.PriceInfo.Block,
		)
	}

	err = conn.WriteMessage(messageType, message)

	if err != nil {
		fmt.Println("write:", err)
		return err
	}

	return nil
}

func handleAtRoot(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}

	defer conn.Close()
	defer signers.Delete(conn)

	for {
		messageType, payload, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read:", err)
			break
		}

		switch payload[0] {
		case 0:
			err := processHello(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}

		case 1:
			err := processPriceReport(conn, messageType, payload[1:])

			if err != nil {
				fmt.Println("write:", err)
			}
		default:
			err = conn.WriteMessage(messageType, []byte("Instruction not supported"))
			if err != nil {
				fmt.Println("write:", err)
			}
		}
	}
}

func StartServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", handleAtRoot)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func init() {
	signers = xsync.NewMapOf[*websocket.Conn, uniswap.Signer]()
}
