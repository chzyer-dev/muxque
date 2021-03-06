package utils

import (
	"encoding/binary"
	"io"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/logex.v1"
)

var (
	pathReplacer = strings.NewReplacer(":", "_")
	letters      = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randSource   = rand.New(rand.NewSource(time.Now().Unix()))
)

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randSource.Intn(len(letters))]
	}
	return string(b)
}

func PathEncode(p string) string {
	return pathReplacer.Replace(p)
}

func BinaryWriteMulti(w io.Writer, objs []interface{}) (err error) {
	for i := 0; i < len(objs); i++ {
		err = binary.Write(w, binary.LittleEndian, objs[i])
		if err != nil {
			return logex.Trace(err, i)
		}
	}
	return nil
}
