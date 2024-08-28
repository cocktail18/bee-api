package enum

type BeePeiSongFeeType int32

const (
	BeePeiSongFeeTypeFixed BeePeiSongFeeType = 0
	BeePeiSongFeeTypeRate  BeePeiSongFeeType = 1
)

var BeePeiSongFeeTypeMap = map[BeePeiSongFeeType]string{
	BeePeiSongFeeTypeFixed: "固定运费",
	BeePeiSongFeeTypeRate:  "比例运费",
}
