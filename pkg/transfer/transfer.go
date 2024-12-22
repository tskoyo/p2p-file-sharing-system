package transfer

import (
	"io"
	"log"
	"os"
	"p2p-file-sharing-system/pkg/peer"
)

func SendFile(client *peer.Client, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(client.Conn, file)
	if err != nil {
		return err
	}

	log.Println("[CLIENT]: File sent successfully!")
	return nil
}
