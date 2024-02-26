package str

import (
	"bytes"

	"github.com/tomtwinkle/garbledreplacer"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// UTF8の文字列をShiftJISに変換する
func EncodeShiftJIS(value string) (string, error) {
	// ShiftJISに変換できない文字を?に変換する
	var buf bytes.Buffer
	w := transform.NewWriter(&buf, garbledreplacer.NewTransformer(japanese.ShiftJIS, '?'))
	if _, err := w.Write([]byte(norm.NFC.String(value))); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
