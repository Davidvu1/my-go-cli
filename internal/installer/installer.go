package installer

import (
	"fmt"
	"io/fs"
	"path"
)

// takes in the templates.FS from apply.go
func CopySkill(src fs.FS, name string, targetDir string) error {
	start := path.Join("skills", name) // get the path from our embedded skills

	return fs.WalkDir(
		src,
		start,
		func(p string, d fs.DirEntry, err error) error {

			if err != nil {
				return fmt.Errorf("there was an error trying to walk the dir %w", err)
			}
			if d.IsDir() {
				return nil
			}
			fmt.Println(p)
			return nil
		},
	)
}
