package metainfo

import (
	"testing"

	"github.com/anacrolix/torrent/internal/qtnew"
	qt "github.com/go-quicktest/qt"
)

func TestParseMagnetV2(t *testing.T) {
	c := qtnew.New(t)

	const v2Only = "magnet:?xt=urn:btmh:1220caf1e1c30e81cb361b9ee167c4aa64228a7fa4fa9f6105232b28ad099f3a302e&dn=bittorrent-v2-test"

	m2, err := ParseMagnetV2Uri(v2Only)
	qt.Assert(t, qt.IsNil(err))
	qt.Check(c, qt.IsFalse(m2.InfoHash.Ok))
	qt.Check(c, qt.IsTrue(m2.V2InfoHash.Ok))
	qt.Check(qt, qt.Equals(m2.V2InfoHash.Value.HexString(), "caf1e1c30e81cb361b9ee167c4aa64228a7fa4fa9f6105232b28ad099f3a302e")(c))
	qt.Check(qt, qt.HasLen(m2.Params, 0)(c))

	_, err = ParseMagnetUri(v2Only)
	qt.Check(c, qt.IsNotNil(err))

	const hybrid = "magnet:?xt=urn:btih:631a31dd0a46257d5078c0dee4e66e26f73e42ac&xt=urn:btmh:1220d8dd32ac93357c368556af3ac1d95c9d76bd0dff6fa9833ecdac3d53134efabb&dn=bittorrent-v1-v2-hybrid-test"

	m2, err = ParseMagnetV2Uri(hybrid)
	qt.Assert(t, qt.IsNil(err))
	qt.Check(c, qt.IsTrue(m2.InfoHash.Ok))
	qt.Check(qt, qt.Equals(m2.InfoHash.Value.HexString(), "631a31dd0a46257d5078c0dee4e66e26f73e42ac")(c))
	qt.Check(c, qt.IsTrue(m2.V2InfoHash.Ok))
	qt.Check(qt, qt.Equals(m2.V2InfoHash.Value.HexString(), "d8dd32ac93357c368556af3ac1d95c9d76bd0dff6fa9833ecdac3d53134efabb")(c))
	qt.Check(qt, qt.HasLen(m2.Params, 0)(c))

	m, err := ParseMagnetUri(hybrid)
	qt.Assert(t, qt.IsNil(err))
	qt.Check(qt, qt.Equals(m.InfoHash.HexString(), "631a31dd0a46257d5078c0dee4e66e26f73e42ac")(c))
	qt.Check(qt, qt.HasLen(m.Params["xt"], 1)(c))
}
