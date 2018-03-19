package msg

import (
	"bytes"
	"encoding/binary"

	. "SPVWallet/core"
	"SPVWallet/core/serialization"
)

type DataReq struct {
	Header
	Type uint8
	Hash Uint256
}

func NewDataReqMsg(invType uint8, hash Uint256) ([]byte, error) {
	msg := new(DataReq)
	msg.Type = invType
	msg.Hash = hash

	body, err := msg.Serialize()
	if err != nil {
		return nil, err
	}

	return BuildMessage("getdata", body)
}

func (dr *DataReq) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := serialization.WriteUint8(buf, dr.Type)
	if err != nil {
		return nil, err
	}

	_, err = dr.Hash.Serialize(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (dr *DataReq) Deserialize(msg []byte) error {
	buf := bytes.NewReader(msg)
	err := binary.Read(buf, binary.LittleEndian, &dr.Header)
	if err != nil {
		return err
	}

	dr.Type, err = serialization.ReadUint8(buf)
	if err != nil {
		return err
	}

	err = dr.Hash.Deserialize(buf)
	if err != nil {
		return err
	}

	return nil
}