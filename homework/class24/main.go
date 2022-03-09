package main

//本周作业：
//体脂排行榜
//
//面向对象方式编写体脂排行榜
//排行榜
//	注册个人信息
//	获取体脂排行
//
//人
//	注册个人信息（体脂相关必要信息
//	更新身高、体重、年龄
//	获取体脂排行
//
//作业要求：
//	使用冒泡、快速排序分别实现体脂排序功能
//	注册时，保存所有注册信息到文件中
//	使用 JSON 格式保存
//	每行一条记录
func main() {
	p := &PersonalInformationHW{}
	frr := &FatRateRank{}

	//人 注册个人信息
	pi := p.RegisterPersonalInformationFake()

	//排行榜 注册个人信息
	frr.registerPersonalInformationFake(pi)

	//排行榜 获取体脂排行
	frr.getFatRateRank(pi)

	//人 更新身高、体重、年龄
	tall, weight, age := 1.80, 80.0, 33
	p.updatePersonalInformation(tall, weight, int64(age))

	//人 获取体脂排行
	p.getFatRateRank(p.Name)

	// 保存所有注册信息到文件中,使用 JSON 格式保存
	filePath := "homework/class24/file.json"
	p.savePersonalInformation(filePath)
	//fmt.Println(pi, frr)
	//fmt.Println(r, f)
}
