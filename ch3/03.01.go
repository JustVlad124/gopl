package main

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID int
	Status    string
}

// func main() {
// 	healthCheckSlice := GenerateCheck()
// 	for _, val := range healthCheckSlice {
// 		if val.Status == PassStatus {
// 			fmt.Println(val.ServiceID)
// 		}
// 	}
// }

func GenerateCheck() []HealthCheck {
	var res []HealthCheck
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			res = append(res, HealthCheck{ServiceID: i, Status: PassStatus})
		} else {
			res = append(res, HealthCheck{ServiceID: i, Status: FailStatus})
		}
	}
	return res
}
