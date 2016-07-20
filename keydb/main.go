package keydb

import (
	"fmt"
	"io/ioutil"
)

type SSHPubKey struct {
	comment string
	keybody string
	keytype string
	options string
}

func (key *SSHPubKey) Parse(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))
}
