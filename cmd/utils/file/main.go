package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsDir(path string) int {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("The specified path does not exist: %s\n", path)
		return -1
	} else if err != nil {
		// その他のエラーを処理する
		fmt.Printf("An error occurred: %s\n", err)
		return -1
	}

	if info.IsDir() {
		// ディレクトリを処理するコードをここに追加する
		return 0
	} else {
		// ファイルを処理するコードをここに追加する
		return 1
	}
}

// 指定されたパスと拡張子に基づいてユニークなファイルパスを生成します。
func UniqueFilePathWithExtension(path string, newExt string) (string, error) {
	if newExt[0] != '.' {
		newExt = "." + newExt // Ensure the extension starts with a dot
	}

	// 拡張子を除いた元のファイル名を取得
	dir, base := filepath.Split(path)
	ext := filepath.Ext(base)
	// if ext == "" {
	// 	return "", fmt.Errorf("original path has no extension")
	// }

	base = base[:len(base)-len(ext)]
	newPath := filepath.Join(dir, base+newExt)

	// ファイルが既に存在するかチェック
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		// 存在しない場合はそのパスをそのまま使用
		return newPath, nil
	}

	// 既に存在する場合はユニークな名前を探す
	for i := 1; ; i++ {
		newPathWithIndex := filepath.Join(dir, fmt.Sprintf("%s(%d)%s", base, i, newExt))
		if _, err := os.Stat(newPathWithIndex); os.IsNotExist(err) {
			return newPathWithIndex, nil
		}
	}
}
