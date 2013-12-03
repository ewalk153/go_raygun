package goraygun

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Entry point for creating RayGun JSON structure returning the detail struct.
func getDetails(errMsg string, exception []byte, filePath string, line int) detail {

	var d detail

	hostname, _ := os.Hostname()

	d.MachineName = hostname
	d.Version = rayGunConfig.ClientVersion
	d.Client = getClientDetails()
	d.Error = getErrorDetails(errMsg, exception, filePath, line)
	//d.Request = getRequestDetials()
	d.Context.Identifier = uuid()

	return d
}

func getClientDetails() clientDetails {

	var c clientDetails

	c.Name = rayGunConfig.ClientName
	c.Version = rayGunConfig.ClientVersion
	c.ClientUrl = rayGunConfig.ClientUrl

	return c
}

func getErrorDetails(errMsg string, exception []byte, filePath string, line int) errorDetail {

	var e errorDetail

	if exception != nil {
		e.Data = string(exception)

		if filePath != "" && line != 0 {
			e.StackTrace = getErrorStackTrace(exception, filePath, line)
		}
	}

	if len(e.StackTrace) > 0 {
		e.ClassName = e.StackTrace[0].ClassName
	}
	// e.ClassName = "error class name"
	e.Message = errMsg

	return e
}

func getErrorStackTrace(exception []byte, filePath string, line int) []errorStackTrace {

	var es errorStackTrace

	packagekName, methodName := processException(exception, line)

	es.ClassName = packagekName
	es.MethodName = methodName

	es.FileName = filePath
	es.LineNumber = line

	return []errorStackTrace{es}
}

func processException(exception []byte, line int) (string, string) {

	var stringExcep string
	var splitstring []string
	var panicLine int
	var pkgAndMehtod []string
	var pkgName string
	var methodName string

	if exception != nil {
		stringExcep = string(exception)
		splitstring = strings.Split(stringExcep, "\n")
	}

	if line != 0 && len(splitstring) != 0 {
		for key, value := range splitstring {
			if strings.Contains(value, strconv.Itoa(line)) {
				panicLine = key
				break
			}
		}
	}

	if panicLine != 0 {
		pkgAndMehtod = strings.Split(splitstring[panicLine-1], ".")
	}

	if len(pkgAndMehtod) >= 2 {
		pkgName = pkgAndMehtod[0]
		methodName = pkgAndMehtod[1]
	}

	return pkgName, methodName
}

func uuid() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "123"
	}
	return strings.TrimSpace(string(out))
}

// func getRequestDetials() detailRequest {
// 	var r detailRequest

// 	r.HostName = "request host name"
// 	r.Url = "request URL"
// 	r.HttpMethod = "request http method"
// 	r.IpAddress = "request IP address"
// 	r.Querystring = "requestQuerystring"
// 	r.Form = "requestForm"
// 	r.Headers = "requestHeaders"
// 	r.RawData = "requestRawData"

// 	return r
// }
