package redisclient

func convertInterfaceArrayToStringArray(dataArray []interface{}) []string {
	stringArray := make([]string, 0, len(dataArray))
	for _, dataArray := range dataArray {
		stringArray = append(stringArray, dataArray.(string))
	}
	return stringArray
}
