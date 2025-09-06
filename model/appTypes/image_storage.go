package appTypes

import "encoding/json"

// Storage 鍥剧墖瀛樺偍绫诲瀷
type Storage int

const (
	Local Storage = iota // 本地
	Qiniu                // 七牛云
)

// MarshalJSON 瀹炵幇浜?json.Marshaler 鎺ュ彛
func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON 瀹炵幇浜?json.Unmarshaler 鎺ュ彛
func (s *Storage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

// String 鏂规硶杩斿洖 Storage 鐨勫瓧绗︿覆琛ㄧず
func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "鏈湴"
	case Qiniu:
		str = "涓冪墰浜?"
	default:
		str = "鏈煡瀛樺偍"
	}
	return str
}

// ToStorage 鍑芥暟灏嗗瓧绗︿覆杞崲涓?Storage
func ToStorage(str string) Storage {
	switch str {
	case "鏈湴":
		return Local
	case "涓冪墰浜?":
		return Qiniu
	default:
		return -1
	}
}
