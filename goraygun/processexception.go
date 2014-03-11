package goraygun

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Entry point for creating RayGun JSON structure returning the detail struct.
func getDetails(errMsg string, exception []byte) detail {

	var d detail

	hostname, _ := os.Hostname()

	d.MachineName = hostname
	d.Version = rayGunConfig.ClientVersion
	d.Client = getClientDetails()
	d.Error = getErrorDetails(errMsg, exception)
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

func getErrorDetails(errMsg string, exception []byte) errorDetail {

	var e errorDetail

	if exception != nil {
		e.Data = string(exception)
		e.StackTrace = getErrorStackTrace(exception)
	}

	if len(e.StackTrace) > 0 {
		e.ClassName = e.StackTrace[0].ClassName
	}
	e.Message = errMsg

	return e
}

func getErrorStackTrace(exception []byte) []errorStackTrace {

	var stackTraces []errorStackTrace

	var stringExcep string
	var splitstring []string

	if exception != nil {
		stringExcep = string(exception)
		splitstring = strings.Split(stringExcep, "\n")
	}

	if len(splitstring) >= 5 {

		for i := 4; i < len(splitstring); i++ {

			filepath, linenum := getFileLineNumber(splitstring[i])

			packagename, methodname := getPackageMethod(splitstring[i-1])

			if filepath != "" || linenum != "" || packagename != "" || methodname != "" {

				var tempStack errorStackTrace

				errorLineNum, err := strconv.Atoi(linenum)
				if err == nil {
					tempStack.LineNumber = errorLineNum
				}
				tempStack.FileName = filepath
				tempStack.ClassName = packagename
				tempStack.MethodName = methodname

				stackTraces = append(stackTraces, tempStack)
			}

			i++
		}
	}

	return stackTraces
}

func getFileLineNumber(exLine string) (string, string) {

	var filePath string
	var lineNum string

	if strings.Contains(exLine, ".go:") {

		line := strings.SplitAfter(exLine, ".go:")

		if len(line) > 1 {
			filePath = line[0]
			if strings.Contains(line[1], " ") {
				linesplit := strings.Split(line[1], " ")
				if len(linesplit) > 0 {
					lineNum = linesplit[0]
				}
			}
		}
	}

	return filePath, lineNum
}

func getPackageMethod(exLine string) (string, string) {

	var pkgName string
	var methodName string

	pkgAndMehtod := strings.Split(exLine, ".")

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
