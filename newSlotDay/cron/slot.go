package cron

import (
	CONSTANT "newslotday/constant"
	DB "newslotday/database"
	UTIL "newslotday/util"
	"strconv"
)

func addSlots() {

	slotDate := UTIL.GetCurrentTime().AddDate(0, 0, 30) // 30th day
	schedules, ok := DB.SelectProcess("select counsellor_id, sum(`0`) as `0`, sum(`1`) as `1`, sum(`2`) as `2`, sum(`3`) as `3`, sum(`4`) as `4`, sum(`5`) as `5`, sum(`6`) as `6`, sum(`7`) as `7`, sum(`8`) as `8`, sum(`9`) as `9`, sum(`10`) as `10`, sum(`11`) as `11`, sum(`12`) as `12`, sum(`13`) as `13`, sum(`14`) as `14`, sum(`15`) as `15`, sum(`16`) as `16`, sum(`17`) as `17`, sum(`18`) as `18`, sum(`19`) as `19`, sum(`20`) as `20`, sum(`21`) as `21`, sum(`22`) as `22`, sum(`23`) as `23`, sum(`24`) as `24`, sum(`25`) as `25`, sum(`26`) as `26`, sum(`27`) as `27`, sum(`28`) as `28`, sum(`29`) as `29`, sum(`30`) as `30`, sum(`31`) as `31`, sum(`32`) as `32`, sum(`33`) as `33`, sum(`34`) as `34`, sum(`35`) as `35`, sum(`36`) as `36`, sum(`37`) as `37`, sum(`38`) as `38`, sum(`39`) as `39`, sum(`40`) as `40`, sum(`41`) as `41`, sum(`42`) as `42`, sum(`43`) as `43`, sum(`44`) as `44`, sum(`45`) as `45`, sum(`46`) as `46`, sum(`47`) as `47` from " + CONSTANT.SchedulesTable + " where weekday = " + strconv.Itoa(int(slotDate.Weekday())) + " and availability_status = 1 and status = 1 group by counsellor_id, weekday")
	if !ok {
		return
	}

	for _, schedule := range schedules {
		for key, value := range schedule {
			slot, err := strconv.Atoi(value)
			if err == nil && slot >= 1 {
				schedule[key] = "1"         // updating to 1, since group sum can give more than 1
				schedule["available"] = "1" // caluculate availibity for a weekday
			}
		}
		schedule["date"] = slotDate.Format("2006-01-02")
		DB.InsertSQL(CONSTANT.SlotsTable, schedule)
	}

	// insert a day for whom there is no schedule
	counsellorIDs, ok := DB.SelectProcess("(select counsellor_id as id from " + CONSTANT.CounsellorsTable + ") union (select listener_id as id from " + CONSTANT.ListenersTable + ") union (select therapist_id as id from " + CONSTANT.TherapistsTable + ")")
	if !ok {
		return
	}
	for _, counsellorID := range counsellorIDs {
		DB.InsertSQL(CONSTANT.SlotsTable, map[string]string{"counsellor_id": counsellorID["id"], "date": slotDate.Format("2006-01-02")})
	}
}
