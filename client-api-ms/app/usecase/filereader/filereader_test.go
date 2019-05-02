package filereader

import (
	"fmt"
	"github.com/IhorBondartsov/microservices-with-ddd/client-api-ms/app/domain/models"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestParseJSON_ParseJSON(t *testing.T) {
	a := assert.New(t)

	file, err := os.Open(".testdata/ports.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	entryForResult := make(chan models.Place)
	fr := NewFileReader(file, 3, 100, entryForResult)
	err = fr.ReadPlaces()
	a.NoError(err)

	var i int
	for k := range entryForResult{
			fmt.Println("result", k)
			i++
			if i == 1{
				break
			}
	}
}