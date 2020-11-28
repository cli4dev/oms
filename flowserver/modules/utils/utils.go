package utils

// GetDeliveryOvertime 获取发货超时时间
func GetDeliveryOvertime(orderOvertime, deliveryOvertime int) int {
	if orderOvertime > deliveryOvertime {
		return deliveryOvertime
	}
	return orderOvertime
}
