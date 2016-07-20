package main

import (
	"github.com/threevi/vinz/keydb"
)

func main() {
	sshk := keydb.SSHPubKey{}
	sshk.Parse("beenz.pub")
}
