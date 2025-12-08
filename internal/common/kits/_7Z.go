package kits

// Unzip 解压zip到对应文件夹
func Unzip(_file string, savePath string) error {
	var back error = nil
	// 打开 zip 文件
	unfile, err1 := zip.OpenReader(_file)
	if err1 != nil {
		back = err1
		fmt.Println("Unzip不能正确读取目标文件：", _file)
	} else {
		defer unfile.Close()
		// 遍历 zip 中的文件
		for _, f := range unfile.File {
			filePath := savePath + f.Name
			if f.FileInfo().IsDir() {
				_ = os.MkdirAll(filePath, os.ModePerm)
				continue
			}
			// 创建对应文件夹
			err2 := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err2 != nil {
				back = err2
			} else {
				// 解压到的目标文件
				dstFile, err3 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err3 != nil {
					back = err3
				} else {
					file, err4 := f.Open()
					if err4 != nil {
						back = err4
					} else {
						// 写入到解压到的目标文件
						_, err5 := io.Copy(dstFile, file)
						if err5 != nil {
							back = err5
						} else {
							back = nil
						}
					}
					file.Close()
				}
				dstFile.Close()
			}
		}
		//fmt.Println("Unzip完成：", _file)
	}
	return back
}

// Un7z 解压7z
func Un7z(_file string, savePath string) error {
	var back error = nil
	// 打开 7z 文件
	unfile, err1 := sevenzip.OpenReader(_file)
	if err1 != nil {
		back = err1
		fmt.Println("Un7z不能正确读取目标文件：", _file, back)
	} else {
		defer unfile.Close()
		// 遍历 zip 中的文件
		for _, f := range unfile.File {
			filePath := savePath + f.Name
			if f.FileInfo().IsDir() {
				_ = os.MkdirAll(filePath, os.ModePerm)
				continue
			}
			// 创建对应文件夹
			err2 := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err2 != nil {
				back = err2
			} else {
				// 解压到的目标文件
				dstFile, err3 := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err3 != nil {
					back = err3
				} else {
					file, err4 := f.Open()
					if err4 != nil {
						back = err4
					} else {
						// 写入到解压到的目标文件
						_, err5 := io.Copy(dstFile, file)
						if err5 != nil {
							back = err5
						} else {
							back = nil
						}
					}
					file.Close()
				}
				dstFile.Close()
			}
		}
		//fmt.Println("Un7z完成：", _file)
	}
	return back
}
