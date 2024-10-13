package handlers

type Data struct {
	// пример: [userID][maxUsers]"20"
	UserData map[int64]map[string]string
}

func NewData() *Data {
	return &Data{UserData: make(map[int64]map[string]string)}
}

func (d *Data) Set(userID int64, key string, value string) {

	if d.UserData[userID] == nil {
		d.UserData[userID] = make(map[string]string)
	}

	d.UserData[userID][key] = value
}

func (d *Data) Get(userID int64, key string) (string, bool) {

	if userDataMap, exists := d.UserData[userID]; exists {

		if value, found := userDataMap[key]; found {
			return value, true
		}
	}
	return "", false
}
