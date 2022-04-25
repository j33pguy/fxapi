package warapi

func Baseurl(param string) *string {
	const host = "https://war-service-live.foxholeservices.com/api"
	baseurl := host + param
	return &baseurl
}
