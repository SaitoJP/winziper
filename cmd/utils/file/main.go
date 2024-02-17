package file

import (
	"fmt"
	"os"
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
