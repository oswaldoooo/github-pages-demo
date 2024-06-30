package route

import "strconv"

type Status int8

func (s Status) String() string {
	if _, ok := statusMap[s]; !ok {
		return "invalid status " + strconv.Itoa(int(s))
	}
	return statusMap[s]
}

// go-macro:enum2str(statusMap,*,*)
const (
	Ok Status = iota + 1
	ParamInvalid
	NotFound
	ServerError
)

type BaseResponse struct {
	Status `json:"status"`
	Data   any `json:"data,omitempty"`
}

var statusMap = map[Status]string{
	Ok:           "Ok",
	ParamInvalid: "Param Invalid",
	NotFound:     "Not Found",
	ServerError:  "Server Error",
}
