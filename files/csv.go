package files

import (
	"github.com/jszwec/csvutil"
	"os"
)

func WriteSamples(filename string, data interface{}) error {
	if b, err := csvutil.Marshal(data); err != nil {
		return err
	} else {
		if f, err := os.Create("samples/" + filename + ".csv"); err != nil {
			return err
		} else {
			defer f.Close()
			if _, err := f.WriteString(string(b)); err != nil {
				return err
			}
		}
	}

	return nil
}
