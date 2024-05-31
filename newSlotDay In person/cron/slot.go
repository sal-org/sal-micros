package cron

import (
	CONSTANT "newslotday/constant"
	DB "newslotday/database"
	UTIL "newslotday/util"
)

func addSlots() {

	slotDate := UTIL.GetCurrentTime().AddDate(0, 0, 30) // 30th day
	// schedules, ok := DB.SelectProcess("select counsellor_id, sum(`0`) as `0`, sum(`1`) as `1`, sum(`2`) as `2`, sum(`3`) as `3`, sum(`4`) as `4`, sum(`5`) as `5`, sum(`6`) as `6`, sum(`7`) as `7`, sum(`8`) as `8`, sum(`9`) as `9`, sum(`10`) as `10`, sum(`11`) as `11`, sum(`12`) as `12`, sum(`13`) as `13`, sum(`14`) as `14`, sum(`15`) as `15`, sum(`16`) as `16`, sum(`17`) as `17`, sum(`18`) as `18`, sum(`19`) as `19`, sum(`20`) as `20`, sum(`21`) as `21`, sum(`22`) as `22`, sum(`23`) as `23`, sum(`24`) as `24`, sum(`25`) as `25`, sum(`26`) as `26`, sum(`27`) as `27`, sum(`28`) as `28`, sum(`29`) as `29`, sum(`30`) as `30`, sum(`31`) as `31`, sum(`32`) as `32`, sum(`33`) as `33`, sum(`34`) as `34`, sum(`35`) as `35`, sum(`36`) as `36`, sum(`37`) as `37`, sum(`38`) as `38`, sum(`39`) as `39`, sum(`40`) as `40`, sum(`41`) as `41`, sum(`42`) as `42`, sum(`43`) as `43`, sum(`44`) as `44`, sum(`45`) as `45`, sum(`46`) as `46`, sum(`47`) as `47` from " + CONSTANT.SchedulesTable + " where weekday = " + strconv.Itoa(int(slotDate.Weekday())) + " and availability_status = 1 and status = 1 group by counsellor_id, weekday")
	// if !ok {
	// 	return
	// }

	counsellorIDs, ok := DB.SelectProcess("(select counsellor_id as id from " + CONSTANT.CounsellorsTable + ") union (select therapist_id as id from " + CONSTANT.TherapistsTable + ")")
	if !ok {
		return
	}

	// for _, schedule := range schedules {
	// 	for key, value := range schedule {
	// 		slot, err := strconv.Atoi(value)
	// 		if err == nil && slot >= 1 {
	// 			schedule[key] = "1"         // updating to 1, since group sum can give more than 1
	// 			schedule["available"] = "1" // caluculate availibity for a weekday
	// 		}
	// 	}
	// 	schedule["date"] = slotDate.Format("2006-01-02")
	// 	DB.InsertSQL(CONSTANT.SlotsTable, schedule)
	// }

	// insert a day for whom there is no schedule

	for _, counsellorID := range counsellorIDs {

		slots := map[string]string{}

		slots["counsellor_id"] = counsellorID["id"]
		slots["date"] = slotDate.Format("2006-01-02")
		slots["available"] = "0"
		slots["0"] = "0"
		slots["1"] = "0"
		slots["2"] = "0"
		slots["3"] = "0"
		slots["4"] = "0"
		slots["5"] = "0"
		slots["6"] = "0"
		slots["7"] = "0"
		slots["8"] = "0"
		slots["9"] = "0"
		slots["10"] = "0"
		slots["11"] = "0"
		slots["12"] = "0"
		slots["13"] = "0"
		slots["14"] = "0"
		slots["15"] = "0"
		slots["16"] = "0"
		slots["17"] = "0"
		slots["18"] = "0"
		slots["19"] = "0"
		slots["20"] = "0"
		slots["21"] = "0"
		slots["22"] = "0"
		slots["23"] = "0"
		slots["24"] = "0"
		slots["25"] = "0"
		slots["26"] = "0"
		slots["27"] = "0"
		slots["28"] = "0"
		slots["29"] = "0"
		slots["30"] = "0"
		slots["31"] = "0"
		slots["32"] = "0"
		slots["33"] = "0"
		slots["34"] = "0"
		slots["35"] = "0"
		slots["36"] = "0"
		slots["37"] = "0"
		slots["38"] = "0"
		slots["39"] = "0"
		slots["40"] = "0"
		slots["41"] = "0"
		slots["42"] = "0"
		slots["43"] = "0"
		slots["44"] = "0"
		slots["45"] = "0"
		slots["46"] = "0"
		slots["47"] = "0"

		// fmt.Println(counsellorID["id"])

		slotsWith, _ := DB.SelectProcess(" select * from "+CONSTANT.InPersonSLotsTable+" where counsellor_id = ? and date = ? ", counsellorID["id"], UTIL.GetCurrentTime().Format("2006-01-02"))
		if len(slotsWith) == 0 {
			for i := 0; i < 30; i++ {

				slotsWithEmpty := map[string]string{}

				slotsWithEmpty["counsellor_id"] = counsellorID["id"]
				slotsWithEmpty["date"] = UTIL.GetCurrentTime().AddDate(0, 0, i).Format("2006-01-02")
				slotsWithEmpty["available"] = "0"
				slotsWithEmpty["0"] = "0"
				slotsWithEmpty["1"] = "0"
				slotsWithEmpty["2"] = "0"
				slotsWithEmpty["3"] = "0"
				slotsWithEmpty["4"] = "0"
				slotsWithEmpty["5"] = "0"
				slotsWithEmpty["6"] = "0"
				slotsWithEmpty["7"] = "0"
				slotsWithEmpty["8"] = "0"
				slotsWithEmpty["9"] = "0"
				slotsWithEmpty["10"] = "0"
				slotsWithEmpty["11"] = "0"
				slotsWithEmpty["12"] = "0"
				slotsWithEmpty["13"] = "0"
				slotsWithEmpty["14"] = "0"
				slotsWithEmpty["15"] = "0"
				slotsWithEmpty["16"] = "0"
				slotsWithEmpty["17"] = "0"
				slotsWithEmpty["18"] = "0"
				slotsWithEmpty["19"] = "0"
				slotsWithEmpty["20"] = "0"
				slotsWithEmpty["21"] = "0"
				slotsWithEmpty["22"] = "0"
				slotsWithEmpty["23"] = "0"
				slotsWithEmpty["24"] = "0"
				slotsWithEmpty["25"] = "0"
				slotsWithEmpty["26"] = "0"
				slotsWithEmpty["27"] = "0"
				slotsWithEmpty["28"] = "0"
				slotsWithEmpty["29"] = "0"
				slotsWithEmpty["30"] = "0"
				slotsWithEmpty["31"] = "0"
				slotsWithEmpty["32"] = "0"
				slotsWithEmpty["33"] = "0"
				slotsWithEmpty["34"] = "0"
				slotsWithEmpty["35"] = "0"
				slotsWithEmpty["36"] = "0"
				slotsWithEmpty["37"] = "0"
				slotsWithEmpty["38"] = "0"
				slotsWithEmpty["39"] = "0"
				slotsWithEmpty["40"] = "0"
				slotsWithEmpty["41"] = "0"
				slotsWithEmpty["42"] = "0"
				slotsWithEmpty["43"] = "0"
				slotsWithEmpty["44"] = "0"
				slotsWithEmpty["45"] = "0"
				slotsWithEmpty["46"] = "0"
				slotsWithEmpty["47"] = "0"

				DB.InsertSQL(CONSTANT.InPersonSLotsTable, slotsWithEmpty)
			}
		}

		DB.InsertSQL(CONSTANT.InPersonSLotsTable, slots)
	}
}
