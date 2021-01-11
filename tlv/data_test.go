package tlv_test

import "strconv"

type TLVTest struct {
	StructTest        *StructTest         `tlv:"15" json:"structTest"`
	BinaryMarshalTest []BinaryMarshalTest `tlv:"65535" json:"binaryMarshalTest"`
	SliceTest         []uint16            `tlv:"255" json:"sliceTest"`
}

type NumberTest struct {
	Int8Data   int8   `tlv:"1" json:"int8Data"`
	Int16Data  int16  `tlv:"2" json:"int16Data"`
	Int32Data  int32  `tlv:"3" json:"int32Data"`
	Int64Data  int64  `tlv:"4" json:"int64Data"`
	UInt8Data  uint8  `tlv:"8" json:"uInt8Data"`
	UInt16Data uint16 `tlv:"9" json:"uInt16Data"`
	UInt32Data uint32 `tlv:"10" json:"uInt32Data"`
	UInt64Data uint64 `tlv:"15" json:"uInt64Data"`
}

type StructTest struct {
	Name     []byte `tlv:"20" json:"name"`
	Sequence uint16 `tlv:"40" json:"sequence"`
}

type BinaryMarshalTest struct {
	Value int `json:"value"`
}

func (mt *BinaryMarshalTest) MarshalBinary() (data []byte, err error) {
	return []byte(strconv.Itoa(mt.Value)), nil
}

func (mt *BinaryMarshalTest) UnmarshalBinary(data []byte) (err error) {
	mt.Value, err = strconv.Atoi(string(data))
	return err
}
