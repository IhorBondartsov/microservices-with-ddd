package main

import (
	"encoding/json"
	"fmt"
	"github.com/IhorBondartsov/microservices-with-ddd/client-api-ms/app/usecase"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello world")

	m := usecase.NewMemoryAnalyser()

	t, f, e := m.GetTotalAndFreeMemory()
	fmt.Printf("%d | %d | %v \n", t, f, e)

	ReadJSON()

}

func ReadJSON() {
	//	const jsonStream = `
	//	[
	//		{"Name": "Ed", "Text": "Knock knock."},
	//		{"Name": "Sam", "Text": "Who's there?"},
	//		{"Name": "Ed", "Text": "Go fmt."},
	//		{"Name": "Sam", "Text": "Go fmt who?"},
	//		{"Name": "Ed", "Text": "Go fmt yourself!"}
	//	]
	//`
	type Place struct {
		Name        string        `json:"name"`
		City        string        `json:"city"`
		Country     string        `json:"country"`
		Alias       []interface{} `json:"alias"`
		Regions     []interface{} `json:"regions"`
		Coordinates []float64     `json:"coordinates"`
		Province    string        `json:"province"`
		Timezone    string        `json:"timezone"`
		Unlocs      []string      `json:"unlocs"`
		Code        string        `json:"code"`
	}
	dec := json.NewDecoder(strings.NewReader(test))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T: %v\n", t, t)

	var m Place
	// decode an array value (Message)
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", m)

	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// decode an array value (Message)
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", m)



	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// decode an array value (Message)
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", m)



	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// decode an array value (Message)
	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", m)

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}

func ReadFile(){
	file, err := os.Open("client-api-ms/test/testdata/ports.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	buf := make([]byte, 10) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Print(buf[:n]) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}
}

var test = `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "ANPHI": {
    "name": "Philipsburg",
    "city": "Philipsburg",
    "country": "Netherlands Antilles",
    "alias": [],
    "regions": [],
    "coordinates": [
      -63.04713709999999,
      18.0295839
    ],
    "province": "Sint Maarten",
    "timezone": "America/Curacao",
    "unlocs": [
      "ANPHI"
    ]
  }
}`
