package gin_controller

import (
	"datathink.top/Waigo/internal/common"
	"datathink.top/Waigo/internal/common/kits"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// GetRootPath 判断当前访问路径是否属于已设置的最短root_path
func (gcl *GinController) GetRootPath(nowDir string) string {
	playAudioListDirKey := "play_audio_list_dirs"
	db := kits.LocalDB{}
	_value, _state := db.LocalDBGetData(playAudioListDirKey)
	if _state != -1 {
		//
		var rootPath string
		var rootPaths []string
		//
		array := strings.Split(_value, "#@")
		for _, thePath := range array {
			if len(thePath) <= len(nowDir) {
				_nowDir := nowDir[0:len(thePath)]
				if thePath == _nowDir {
					rootPaths = append(rootPaths, thePath)
				}
			}
		}
		// 数组中长度最短的就是root_path
		if len(rootPaths) > 0 {
			rootPath = rootPaths[0]
			for _, _min := range rootPaths {
				if len(_min) < len(rootPath) {
					rootPath = _min
				}
			}
		}
		//
		return rootPath
	} else {
		return ""
	}
}

// GetPlayAudioList 获取音频文件播放列表
func (gcl *GinController) GetPlayAudioList(ctx *gin.Context) {
	_nowDir := common.RequestInput(ctx, "now_dir")
	_nowDir = common.ConvertedPath(_nowDir)

	// 目标地址
	var listDirs []map[string]interface{}
	var listFiles []map[string]interface{}
	var rootPath string = ""

	state := 0
	msg := ""

	//
	if len(_nowDir) > 0 {
		rootPath = gcl.GetRootPath(_nowDir)
		if len(rootPath) > 0 {
			if common.IsDir(_nowDir) {
				// 读取目录内容
				entries, err := os.ReadDir(_nowDir)
				if err != nil {
					state = 0
					msg = "Read Error"
				} else {
					//
					for _, entry := range entries {
						info, err := entry.Info()
						if err != nil {
							continue
						}
						// 判断类型
						if entry.IsDir() { // dir
							if strings.Index(entry.Name(), ".") == 0 {
								// 排除
							} else {
								dirInfo := map[string]interface{}{
									"name":        entry.Name(),
									"token":       common.MD5("dir=" + common.URLEncode(_nowDir+"/"+entry.Name())),
									"size":        "",
									"create_time": common.FormatTimeToDate("Y/m/d H:i", info.ModTime()),
								}
								listDirs = append(listDirs, dirInfo)
							}
						} else { // file
							if strings.Index(entry.Name(), ".") == 0 {
								// 排除
							} else {
								fileInfo := map[string]interface{}{
									"name":        entry.Name(),
									"token":       common.MD5("file=" + common.URLEncode(_nowDir+"/"+entry.Name())),
									"size":        common.FormatFileSize(info.Size()),
									"create_time": common.FormatTimeToDate("Y/m/d H:i", info.ModTime()),
								}
								listFiles = append(listFiles, fileInfo)
							}
						}
					}
					//
					state = 1
					msg = "OK"
				}
			} else {
				state = 0
				msg = "No Path"
			}
		} else {
			state = 0
			msg = "Path Error"
		}
	} else {
		playAudioListDirKey := "play_audio_list_dirs"
		db := kits.LocalDB{}
		_value, _state := db.LocalDBGetData(playAudioListDirKey)
		if _state != -1 {
			array := strings.Split(_value, "#@")
			for _, thePath := range array {
				if len(thePath) > 0 {
					fileInfo := map[string]interface{}{
						"name":        thePath,
						"token":       "",
						"size":        "",
						"create_time": "(Added Folder)",
					}
					listDirs = append(listDirs, fileInfo)
				}
			}
			state = 1
			msg = "Default Set Data"
		} else {
			state = 0
			msg = "Null set"
		}
	}

	//
	ctx.JSON(http.StatusOK, gin.H{
		"state": state,
		"msg":   msg,
		"content": map[string]interface{}{
			"list_dirs":  listDirs,
			"list_files": listFiles,
			"view_path":  _nowDir,
			"root_path":  rootPath,
		},
	})

	return
}

// PlayAudio 返回文件
func (gcl *GinController) PlayAudio(ctx *gin.Context) {

	return
}
