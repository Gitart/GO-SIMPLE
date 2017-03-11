
// Scan dir
func scanDir(app *app, userID int64, baseDir, dirname string) {
	
	fileList, err := ioutil.ReadDir(dirname)
	
	if err != nil {
		log.Println(err)
	}

	for _, info := range fileList {
		name := info.Name()
		if info.IsDir() {
			scanDir(app, userID, baseDir, filepath.Join(dirname, name))
		} else {
			
			fullPath := filepath.Join(dirname, name)
			tags     := filepath.SplitList(dirname[len(baseDir):])
			ext      := strings.ToLower(filepath.Ext(name))

			if ext != ".jpg" && ext != ".png" {
				continue
			}
			title := name[:len(name)-len(ext)]

			var contentType string
			if ext == ".jpg" {
				contentType = "image/jpeg"
			} else {
				contentType = "image/png"
			}

			if err := storeFile(app, fullPath, title, contentType, tags, userID); err != nil {
			   log.Println(err)
			}
		}
	}
}
