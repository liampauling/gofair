package streaming

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"encoding/json"
)


func ReadFile(directory string) {
	fmt.Printf("Parsing file %v\n", directory)

	file, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		var mc MarketChangeMessage
		err := json.Unmarshal(scanner.Bytes(), &mc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mc)
		//fmt.Println(mc.MarketChanges)

		//for _, market := range mc.MarketChanges {
		//	for _, runner := range market.RunnerChange {
		//		fmt.Println(runner)
		//	}
		//}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
