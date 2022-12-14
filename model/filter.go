package model

import (
	"rosswilson/usercapacity/utility"
)

type FilterModel struct {
	users    map[int]Userable
	previous Modeler
}

func (f *FilterModel) buildModel() {
	targetUsers := utility.GetConfigProxy().GetUsers()
	previousUsers := f.GetPrevious().GetUsers()

	f.users = map[int]Userable{}
	for _, t := range targetUsers {
		id := t.GetEverhourId()

		f.users[id] = User{
			name:        previousUsers[id].GetName(),
			trackedTime: previousUsers[id].GetTimeTracked(),
			averageTime: previousUsers[id].GetAvgTime(),
			daysHadOff:  previousUsers[id].GetDaysHadOff(),
			daysHaveOff: previousUsers[id].GetDaysHaveOff(),
		}
	}

	utility.GetLogger().Write("model: built FilterModel")
}

func (f *FilterModel) GetPrevious() Modeler {
	return f.previous
}

func (f *FilterModel) GetUsers() map[int]Userable {
	return f.users
}

func CreateFilterModel(previous Modeler) *FilterModel {
	return &FilterModel{
		previous: previous,
	}
}
