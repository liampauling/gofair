package gofair

import (
	"bufio"
	"encoding/json"
	"gofair/streaming"
	"log"
	"os"
)

func (h *Historical) ParseHistoricalData(directory string, listener streaming.Listener) error {
	file, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		var mc streaming.MarketChangeMessage
		err := json.Unmarshal(scanner.Bytes(), &mc)
		if err != nil {
			log.Fatal(err, t)
		}

		listener.OnData(mc)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("Reading complete")

	// close channel
	close(listener.OutputChannel)

	return nil
}
