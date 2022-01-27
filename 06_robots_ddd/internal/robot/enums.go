package robot

type robotStatusType string

var RobotStatus = struct {
	Active   robotStatusType
	Charged  robotStatusType
	Repaired robotStatusType
}{
	Active:   "active",
	Charged:  "charged",
	Repaired: "repaired",
}

// var Status = newStatus()

// func newStatus() *statusValues {
// 	return &statusValues{
// 		Active:   "active",
// 		Charged:  "charged",
// 		Repaired: "repaired",
// 	}
// }

// type statusValues struct {
// 	Active   string
// 	Charged  string
// 	Repaired string
// }
