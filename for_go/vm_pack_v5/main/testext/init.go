package testext

/*
func init() {
	msgpack.RegisterExt(1, (*MyTime)(nil))
}

type MyTime struct {
	TT time.Time
}

var (
	_ msgpack.Marshaler   = (*MyTime)(nil)
	_ msgpack.Unmarshaler = (*MyTime)(nil)
)

func (tm *MyTime) MarshalMsgpack() ([]byte, error) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint32(b, uint32(tm.TT.Unix()))
	binary.BigEndian.PutUint32(b[4:], uint32(tm.TT.Nanosecond()))
	return b, nil
}

func (tm *MyTime) UnmarshalMsgpack(b []byte) error {
	if len(b) != 8 {
		return fmt.Errorf("invalid data length: got %d, wanted 8", len(b))
	}
	sec := binary.BigEndian.Uint32(b)
	usec := binary.BigEndian.Uint32(b[4:])
	fmt.Println(sec,usec)
	//tm.Time = time.Unix(int64(sec), int64(usec))
	return nil
}

 */