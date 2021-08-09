package x

//数据看板培训人数
func SumNumberTrainingRouter(router *gin.RouterGroup) {

	control := GetSumNumberTraining()
	router.GET("/getsumnumbertraining", control.SumNumberTraining) //获取培训总人数

}
