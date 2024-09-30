package bencode

import (
	"bytes"
	"os"
	"testing"

	qt "github.com/go-quicktest/qt"
	"github.com/stretchr/testify/assert"
)

func loadFile(name string, t *testing.T) []byte {
	data, err := os.ReadFile(name)
	qt.Assert(t, qt.IsNil(err))
	return data
}

func testFileInterface(t *testing.T, filename string) {
	data1 := loadFile(filename, t)

	var iface interface{}
	err := Unmarshal(data1, &iface)
	qt.Assert(t, qt.IsNil(err))

	data2, err := Marshal(iface)
	qt.Assert(t, qt.IsNil(err))

	assert.EqualValues(t, data1, data2)
}

func TestBothInterface(t *testing.T) {
	testFileInterface(t, "testdata/archlinux-2011.08.19-netinstall-i686.iso.torrent")
	testFileInterface(t, "testdata/continuum.torrent")
}

type torrentFile struct {
	Info struct {
		Name        string `bencode:"name"`
		Length      int64  `bencode:"length"`
		MD5Sum      string `bencode:"md5sum,omitempty"`
		PieceLength int64  `bencode:"piece length"`
		Pieces      string `bencode:"pieces"`
		Private     bool   `bencode:"private,omitempty"`
	} `bencode:"info"`

	Announce     string      `bencode:"announce"`
	AnnounceList [][]string  `bencode:"announce-list,omitempty"`
	CreationDate int64       `bencode:"creation date,omitempty"`
	Comment      string      `bencode:"comment,omitempty"`
	CreatedBy    string      `bencode:"created by,omitempty"`
	URLList      interface{} `bencode:"url-list,omitempty"`
}

func testFile(t *testing.T, filename string) {
	data1 := loadFile(filename, t)
	var f torrentFile

	err := Unmarshal(data1, &f)
	if err != nil {
		t.Fatal(err)
	}

	data2, err := Marshal(&f)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data1, data2) {
		println(string(data2))
		t.Fatalf("equality expected")
	}
}

func TestBoth(t *testing.T) {
	testFile(t, "testdata/archlinux-2011.08.19-netinstall-i686.iso.torrent")
}
