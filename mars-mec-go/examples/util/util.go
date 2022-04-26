package util

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//字符串转数组int32
func StringToArray32(str string) []int32 {
	var arr []int32
	json.Unmarshal([]byte(str), &arr)

	return arr
}

//数组int32转字符串
func Array32ToString(list []int32) string {
	var str string
	marshal, _ := json.Marshal(list)
	str = string(marshal)
	return str
}

//字符串转int32
func StringToInt32(str string) int32 {
	val, _ := strconv.Atoi(str)
	return int32(val)
}

//proto转byte
func TransProtoToByte(pb proto.Message) []byte {
	bytes, _ := json.Marshal(pb)
	return bytes
}

func WeightedRandomIndex(seveList []int32, count int32) []int32 {
	rand.Seed(time.Now().UnixNano())

	var tmp *[]int32
	tmp = &seveList
	var res []int32
	for i := 0; i < int(count); i++ {
		leng := len(*(tmp))
		r := RandInt(1, leng) - 1
		a := (*tmp)[r]
		res = append(res, a)

		for j := leng - 1; j >= 0; j-- {
			if (*tmp)[j] == a {
				(*tmp) = append((*tmp)[:j], (*tmp)[j+1:]...)
			}
		}
	}
	return res
}
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func GetRandWithoutExists(originData []int32, existsId []int32, count int32) []int32 {
	data := make([]int32, 0)
	data = append(data, originData...)
	for _, id := range existsId {
		for i := len(data) - 1; i >= 0; i-- {
			if data[i] == id {
				data = append(data[:i], data[i+1:]...)
			}
		}
	}
	if int(count) >= len(data) {
		return data
	}
	res := WeightedRandomIndex(data, count)
	return res
}

func GetUidFromRoomId(roomId string) (string, string) {
	split := strings.Split(roomId, "_")
	if len(split) > 2 {
		return split[0], split[1]
	}
	return "", ""
}
