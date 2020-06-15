package data

import (
	"io"
	"log"

	"k8s.io/apimachinery/pkg/util/json"
)

func ToJSON(v interface{}, w io.Writer) error {
	payload, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Couldn't marshal struct to JSON", err)
	}
	_, err = w.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
