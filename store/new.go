package store

import "os"

type StoreFS interface {
	Stat(string) (os.FileInfo, error)
}

type StoreInput struct {
	StoreFS StoreFS
	StoreDir string
}

type Store struct {
	input *StoreInput
}

func New(storeInput *StoreInput) *Store {
	return &Store{input: storeInput}
}

func (self *Store)StoreDir() string {
	if self.input.StoreDir == "" {
		return ".aws"
	}
	return self.input.StoreDir
}

