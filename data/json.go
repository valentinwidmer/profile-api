package data

import (
	"fmt"
	"io"
	"io/ioutil"
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

func FromJSON(r io.ReadCloser) (*Profile, error) {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read data from body: %v", err)
	}
	p := &Profile{}
	err = json.Unmarshal(c, p)
	if err != nil {
		return nil, fmt.Errorf("Couldn't unmarshal data to struct: %v", err)
	}

	return p, nil
}
