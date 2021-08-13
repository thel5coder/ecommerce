package view_models

type FileVm struct {
	Name string `json:"name"`
	Ext  string `json:"ext"`
	Path string `json:"path"`
}

func NewFileVm(name, ext, path string) FileVm {
	return FileVm{
		Name: name,
		Ext:  ext,
		Path: path,
	}
}
