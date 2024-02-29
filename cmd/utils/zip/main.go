package zip

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/SaitoJP/winziper/cmd/utils/file"
	"github.com/SaitoJP/winziper/cmd/utils/str"
	"github.com/yeka/zip"
)

// パスワードなし。isWindowsがtrueの場合、Mac特有のファイルを排除し、ShiftJISに変換
func Write(source string, isWindows bool) error {
	return write(source, "", isWindows)
}

// パスワード付き。isWindowsがtrueの場合、Mac特有のファイルを排除し、ShiftJISに変換
func WriteEncrypted(source, password string, isWindows bool) error {
	return write(source, password, isWindows)
}

func write(source, password string, isWindows bool) error {
	target, err := uniqueZipPath(source)
	if err != nil {
		return err
	}

	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ソースのパスは不要なためスキップする
		if source == path {
			return nil
		}

		// Mac特有のファイルを除外
		if isWindows && (strings.Contains(path, "__MACOSX") || strings.Contains(path, ".DS_Store")) {
			return nil
		}

		if password == "" {
			return addToZip(archive, source, baseDir, path, isWindows)
		} else {
			return addToEncryptedZip(archive, source, baseDir, path, password)
		}
	})

	return err
}

func addToZip(archive *zip.Writer, source, baseDir, path string, isWindows bool) error {
	return addToZipInternal(archive, source, baseDir, path, "", isWindows)
}

func addToEncryptedZip(archive *zip.Writer, source, baseDir, path, password string) error {
	return addToZipInternal(archive, source, baseDir, path, password, false)
}

func addToZipInternal(archive *zip.Writer, source, baseDir, path, password string, isWindows bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// path と source が不一致の場合はディレクトリ内のファイルであるため、
	// path(ファイルの絶対パス)からsource(ディレクトリの絶対パス)を削除する。
	// 逆に一致する場合は1つのファイルを圧縮しているということ。
	trimmedPath := path
	if path != source {
		trimmedPath = strings.TrimPrefix(path, source)
	}

	relPath := filepath.Join(baseDir, trimmedPath)
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// ZIPファイル内でのパス名。Windowsの場合はShiftJISに変換
	if isWindows {
		encodedName, err2 := str.EncodeShiftJIS(relPath)
		if err2 != nil {
			return err2
		}
		header.Name = encodedName
	} else {
		header.Name = relPath
	}

	if info.IsDir() {
		header.Name += "/"
	} else {
		header.Method = zip.Deflate
	}

	var writer io.Writer
	if password == "" {
		writer, err = archive.CreateHeader(header)
	} else {
		writer, err = archive.Encrypt(header.Name, password, zip.AES256Encryption)
	}
	if err != nil {
		return err
	}

	if !info.IsDir() {
		file, err3 := os.Open(path) // ここではオリジナルのパス（変換されていない）を使用
		if err3 != nil {
			return err3
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
	}

	return err
}

// uniqueZipPath は、指定されたパスに基づいてユニークなZIPファイルパスを生成します。
func uniqueZipPath(path string) (string, error) {
	newPath, err := file.UniqueFilePathWithExtension(path, "zip")
	if err != nil {
		return "", err
	}

	return newPath, nil
}
