package actions

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/afero"
	"net/http"
	"os/user"
	"path"
	"path/filepath"
)

var AppFs = afero.NewOsFs()
var diskFs = afero.Afero{Fs: AppFs}

type PathInfo struct {
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	IsDir    bool        `json:"is_dir"`
	Items    int64       `json:"items"`
	Files    int64       `json:"files"`
	Size     int64       `json:"size"`
	Children []*PathInfo `json:"children"`
}

type PathInfoFilter struct {
	Name     string           `json:"name"`
	Path     string           `json:"path"`
	IsDir    bool             `json:"is_dir"`
	Items    int64            `json:"items"`
	Files    int64            `json:"files"`
	Size     int64            `json:"size"`
	Children []PathInfoFilter `json:"child"`
}

var rootDirData *PathInfo

func Scan(c echo.Context) error {
	ScanDir(UserDir())
	data := rootDirData
	result := DataFilter(data)
	return c.JSON(http.StatusOK, result)
}

func DataFilter(data *PathInfo) PathInfoFilter {
	result := PathInfoFilter{data.Name, data.Path, data.IsDir, data.Items, data.Files, data.Size, []PathInfoFilter{}}
	for _, v := range data.Children {
		result.Children = append(result.Children, PathInfoFilter{v.Name, v.Path, v.IsDir, v.Items, v.Files, v.Size, []PathInfoFilter{}})
	}
	return result
}

func Get(c echo.Context) error {
	type PostData struct {
		Path []string `json:"path"`
	}

	pathData := new(PostData)

	_ = c.Bind(pathData)
	if rootDirData == nil {
		ScanDir(UserDir())
	}
	data := rootDirData

	if len(pathData.Path) != 0 {
		data = Find(data, pathData.Path)
	}
	result := DataFilter(data)

	return c.JSON(http.StatusOK, result)
}

func Find(data *PathInfo, s []string) *PathInfo {
	for _, f0 := range s {
		for _, ar := range data.Children {
			if ar.Name == f0 {
				data = ar
			}
		}
	}
	return data
}
func ScanDir(s string) *PathInfo {
	rootDirInfo, _ := diskFs.Stat(s)
	rootDirData = &PathInfo{Name: rootDirInfo.Name(), Path: s, IsDir: true}
	children := ScanDirCh(rootDirData)
	rootDirData.Children = children.Children
	return nil
}

func ScanDirCh(parent *PathInfo) *PathInfo {
	files, _ := diskFs.ReadDir(parent.Path)

	for _, file := range files {
		if file.IsDir() && file.Name() != "proc" {
			f := &PathInfo{Path: parent.JoinPath(file.Name()), Name: file.Name(), IsDir: true}
			f.Children = ScanDirCh(f).Children
			parent.Children = append(parent.Children, f)
			parent.Size = parent.Size + f.Size
			parent.Items = parent.Items + f.Items + 1
			parent.Files = parent.Files + f.Files
		} else {
			f := &PathInfo{Path: parent.JoinPath(file.Name()), Name: file.Name(), Size: file.Size(), IsDir: false}
			f.Children = nil
			parent.Children = append(parent.Children, f)
			parent.Size = parent.Size + f.Size
			parent.Items = parent.Items + 1
			parent.Files = parent.Files + 1
		}
	}
	return parent
}

func (d PathInfo) JoinPath(Name string) string {
	return filepath.Join(d.Path, Name)
}

func DirTree(c echo.Context) error {
	rootPath := UserDir()

	if len(c.QueryParam("path")) != 0 {
		rootPath = c.QueryParam("path")
	}

	type dirName struct {
		Name string `json:"name"`
		Path string `json:"path"`
	}

	dirList := []dirName{}

	dirInfo, err := diskFs.ReadDir(rootPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, d := range dirInfo {
		if d.IsDir() == true {
			dirList = append(dirList, dirName{d.Name(), path.Join(rootPath, d.Name())})
		}

	}

	return c.JSON(http.StatusOK, dirList)
}

func UserDir() string {
	userInfo, _ := user.Current()
	if userInfo.Username == "root" {
		return "/"
	}

	return "."
}
