package filereader

import (
	"encoding/json"
	"fmt"
	"github.com/IhorBondartsov/microservices-with-ddd/client-api-ms/app/domain/models"
	"github.com/IhorBondartsov/microservices-with-ddd/client-api-ms/app/usecase/filereader/jsonparser"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
)

type FileReader interface {
	ReadPlaces() error
}

type fileReader struct {
	buffer       []byte
	file         *os.File
	objectReader jsonparser.ParserJSONByParts
	readFile     bool

	workers        int
	entryForResult chan<- models.Place
}

func NewFileReader(file *os.File, workers, bufferSize int, entryForResult chan<- models.Place) FileReader {
	return &fileReader{
		buffer:         make([]byte, bufferSize),
		file:           file,
		objectReader:   jsonparser.NewJSONObjectReader(),
		entryForResult: entryForResult,
		workers:        workers,
	}
}

func (fr *fileReader) ReadPlaces() error {
	if fr.readFile {
		return errors.New("cant read file at the same time")
	}
	if fr.file == nil {
		return errors.New("have not file")
	}

	// Create workers for unmarshaling and sending result to customer
	results := make(chan []string)
	fr.runPoolWorkers(results)

	fr.readFile = true
	defer func() { fr.readFile = false }()

	for {
		// Read file
		n, err := fr.file.Read(fr.buffer)
		if n > 0 {
			// Read buffer and trying to parse it
			objs, err := fr.objectReader.ReadObjInJSON(string(fr.buffer[:n]))
			if err != nil {
				return err
			}
			results <- objs
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "Read file err")
		}
	}

	close(results)
	return nil
}

func (r *fileReader) runPoolWorkers(results <-chan []string) {
	for i := 0; i < r.workers; i++ {
		go r.worker(results)
	}

}

func (r *fileReader) worker(results <-chan []string) {
	fmt.Println("Run worker")
	for {
		select {
		case obj, ok := <-results:
			if !ok {
				fmt.Println("Stop channel")
				return
			}
			for k, _ := range obj {
				m := models.Place{}
				// TODO: rewritte
				arrStr := strings.SplitAfterN(obj[k], ":", 2)
				err := json.Unmarshal([]byte(arrStr[1]), &m)
				if err != nil {
					log.Printf("Cant unmarshal Place: \n Error: %v,\n JSON: %v", err, obj[k])
				}
				r.entryForResult <- m
			}
		}
	}
}
