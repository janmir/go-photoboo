package photoboo

import (
	"os"
	"path/filepath"

	"github.com/tj/go-dropbox"
)

var (
	box *dropbox.Client
)

const (
	_token = _secretToken

	_path = "https://www.dropbox.com/sh/wb5thm9h8qu4tfm/AAAlrfWeXQXxIN64gDL4gGQEa/Photoboo?dl=0&preview="
)

func upload(filename string) string {
	file, err := os.Open(filename)
	catch(err)
	defer file.Close()

	//extraction
	path := "/" + filepath.Base(filename)

	meta, err := box.Files.Upload(&dropbox.UploadInput{
		Path:   path,
		Reader: file,
		Mute:   false,
		Mode:   "add",
	})
	catch(err)

	return _path + meta.Name
}
