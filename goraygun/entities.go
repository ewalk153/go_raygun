package goraygun

//RayGunConfigFile holds all the config items about RayGunApp and the current
//app you are using this recovery package in.
//There is an example config file in the repo 'RayGunConfig.json'.
//This needs to be on your local package folder to run.
type RayGunConfigFile struct {
	RaygunAppName  string
	RaygunAPIKey   string
	RaygunEndpoint string
	ClientName     string
	ClientVersion  string
	ClientUrl      string
}

type entrie struct {
	OccurredOn string `json:"occurredOn"`
	Details    detail `json:"details"`
}

type detail struct {
	MachineName string        `json:"machineName"`
	Version     string        `json:"version"`
	Client      clientDetails `json:"client"`
	Error       errorDetail   `json:"error"`
	//Environment    environmentDetail      `json:"environment"`
	//UserCustomData userCustomerDataDetail `json:"userCustomData"`
	//Request detailRequest `json:"request"`
	//Response       detailResponce         `json:"response"`
	//User           detailUser             `json:"user"`
	Context detailContext `json:"context"`
}

type clientDetails struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	ClientUrl string `json:"clientUrl"`
}

type errorDetail struct {
	//InnerError string            `json:"innerError"`
	Data       string            `json:"data"`
	ClassName  string            `json:"className"`
	Message    string            `json:"message"`
	StackTrace []errorStackTrace `json:"stackTrace"`
}

type errorData struct {
	bytes     []byte `json:"bytes"`
	stringMsg string `json:"stringMsg"`
}

type errorStackTrace struct {
	LineNumber int    `json:"lineNumber"`
	ClassName  string `json:"className"`
	FileName   string `json:"fileName"`
	MethodName string `json:"methodName"`
}

type detailRequest struct {
	HostName   string `json:"hostName"`
	Url        string `json:"url"`
	HttpMethod string `json:"httpMethod"`
	IpAddress  string `json:"ipAddress"`
	// Querystring requestQuerystring `json:"querystring"`
	// Form        requestForm        `json:"form"`
	// Headers     requestHeaders     `json:"headers"`
	// RawData     requestRawData     `json:"rawData"`
	Querystring string `json:"querystring"`
	Form        string `json:"form"`
	Headers     string `json:"headers"`
	RawData     string `json:"rawData"`
}

type requestQuerystring struct {
	data string
}

type requestForm struct {
}

type requestHeaders struct {
}

type requestRawData struct {
}

type detailResponce struct {
	StatusCode int `json:"statusCode"`
}

type detailUser struct {
	Identifier string `json:"identifier"`
}

type detailContext struct {
	Identifier string `json:"identifier"`
}
