package tar

import (
	"archive/tar"
	"fmt"
	"gopl/ch10/ex10.2/arch"
	"io"
	"os"
)

func Open(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	tr := tar.NewReader(f)
	for {
		trf, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", trf.Name)
	}
	return nil
}

func init() {
	arch.Register("tar", Open)
}
