package pkg

import (
	"bytes"
	"embed"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

func RenderTemplate(src string, out string, fs embed.FS, obj any) error {
	content, err := fs.ReadFile(src + ".tpl")
	if err != nil {
		return err
	}

	tpl, err := template.New(src).Parse(string(content))
	if err != nil {
		return err
	}

	err = EnsureDir(filepath.Dir(out))
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer
	err = tpl.Execute(&buf, obj)
	if err != nil {
		return err
	}

	var res []byte
	if filepath.Ext(out) == ".go" {
		res, err = format.Source(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		res = buf.Bytes()
	}

	_, err = f.Write(res)
	return err
}
