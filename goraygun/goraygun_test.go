package goraygun

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestLoadRayGunSettings(t *testing.T) {

	rayGunConfigFileName = "../testdata/validRayGunConfig.json"

	testgoraygun := new(Goraygun)

	err := testgoraygun.LoadRaygunSettings()
	if err != nil {
		t.Errorf("Error Reading RayGun config: %v", err)
	}

	assert.Equal(t, rayGunConfig.ClientName, "ClientName", "Expected ClientName, recived %v", rayGunConfig.ClientName)
	assert.Equal(t, rayGunConfig.ClientUrl, "ClientUrl", "Expected ClientUrl, recived %v", rayGunConfig.ClientUrl)
	assert.Equal(t, rayGunConfig.ClientVersion, "ClientVersion", "Expected ClientVersion, recived %v", rayGunConfig.ClientVersion)
	assert.Equal(t, rayGunConfig.RaygunAPIKey, "RaygunAPIKey", "Expected RaygunAPIKey, recoved %v", rayGunConfig.RaygunAPIKey)
	assert.Equal(t, rayGunConfig.RaygunAppName, "RaygunAppName", "Expected RaygunAppName, recoved %v", rayGunConfig.RaygunAppName)
	assert.Equal(t, rayGunConfig.RaygunEndpoint, "RaygunEndpoint", "Expected RaygunEndpoint, recoved %v", rayGunConfig.RaygunEndpoint)
}

func readByteException(t *testing.T, exceptionFile string) []byte {

	byteException := exceptionFile

	data, err := ioutil.ReadFile(byteException)
	if err != nil {
		t.Errorf("Error Reading byteException test file: %v", err)
	}

	return data
}

func TestGetPackageMethod_pass(t *testing.T) {

	exceptionLine := "main.method6()"

	expectedPackage := "main"
	expectedMethod := "method6()"

	pack, method := getPackageMethod(exceptionLine)

	assert.Equal(t, pack, expectedPackage, "Recieved %v, Expected %v", pack, expectedPackage)
	assert.Equal(t, method, expectedMethod, "Recieved %v, Expected %v", method, expectedMethod)

}

func TestGetPackageMethod_fail(t *testing.T) {

	exceptionLine := "mainmethod6()"

	expectedPackage := ""
	expectedMethod := ""

	pack, method := getPackageMethod(exceptionLine)

	assert.Equal(t, pack, expectedPackage, "Recieved %v, Expected %v", pack, expectedPackage)
	assert.Equal(t, method, expectedMethod, "Recieved %v, Expected %v", method, expectedMethod)

}

func TestGetFileLineNumber_pass(t *testing.T) {

	exceptionLine := "/Users/Documents/Code/GoRayGun/Server/go/src/testapp/server.go:47 +0xba"

	expectedPath := "/Users/Documents/Code/GoRayGun/Server/go/src/testapp/server.go:"
	expectedLine := "47"

	path, line := getFileLineNumber(exceptionLine)

	assert.Equal(t, path, expectedPath, "Recieved %v, Expected %v", path, expectedPath)
	assert.Equal(t, line, expectedLine, "Recieved %v, Expected %v", line, expectedLine)
}

func TestGetFileLineNumber_fail(t *testing.T) {

	exceptionLine := "/Users/Documents/Code/GoRayGun/Server/go/src/testapp/server:47 +0xba"

	expectedPath := ""
	expectedLine := ""

	path, line := getFileLineNumber(exceptionLine)

	assert.Equal(t, path, expectedPath, "Recieved %v, Expected %v", path, expectedPath)
	assert.Equal(t, line, expectedLine, "Recieved %v, Expected %v", line, expectedLine)
}

func TestGetErrorStackTrace_pass(t *testing.T) {
	testExc := readByteException(t, "../testdata/byteException2")

	expectedCount := 7
	firstFilepath := "        /Users/gregpugh/Documents/Code/GoRayGun/Server/go/src/testapp/server.go:"
	firstPackagename := "main"
	firstMethod := "method6()"
	firstLineNumber := 47

	returnedStackTrace := getErrorStackTrace(testExc)

	assert.True(t, len(returnedStackTrace) == expectedCount, "Returned stack trace count %v Expected %v", len(returnedStackTrace), expectedCount)
	assert.Equal(t, returnedStackTrace[0].FileName, firstFilepath, "Returned %v, Expected %v", returnedStackTrace[0].FileName, firstFilepath)
	assert.Equal(t, returnedStackTrace[0].LineNumber, firstLineNumber, "Returned %v, Expected %v", returnedStackTrace[0].LineNumber, firstLineNumber)
	assert.Equal(t, returnedStackTrace[0].MethodName, firstMethod, "Returned %v, Expected %v", returnedStackTrace[0].MethodName, firstMethod)
	assert.Equal(t, returnedStackTrace[0].ClassName, firstPackagename, "Returned %v, Expected %v", returnedStackTrace[0].ClassName, firstPackagename)
}

func TestGetErrorStackTrace_fail(t *testing.T) {
	testExc := readByteException(t, "../testdata/byteException_fail")

	expectedCount := 0

	returnedStackTrace := getErrorStackTrace(testExc)

	assert.True(t, len(returnedStackTrace) == expectedCount, "Returned stack trace count %v Expected %v", len(returnedStackTrace), expectedCount)
}

func TestGetErrorDetials(t *testing.T) {

	testExc := readByteException(t, "../testdata/byteException")
	errMsg := "errMsg"
	//expectedClassName := "main"

	testErrorDetails := getErrorDetails(errMsg, testExc)

	assert.NotEmpty(t, testErrorDetails.Data, "Error Details Data stack trace should have been set")
	assert.NotEmpty(t, testErrorDetails.StackTrace, "Error Details Stack trace should have been set")
	assert.True(t, len(testErrorDetails.StackTrace) > 1, "There should have been at least one stact trace returned.")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
	//assert.Equal(t, testErrorDetails.ClassName, expectedClassName, "Expected %v, Recived %v", expectedClassName, testErrorDetails.ClassName)
}

func TestGetErrorDetails_noException(t *testing.T) {

	var testExc []byte
	errMsg := "errMsg"

	testErrorDetails := getErrorDetails(errMsg, testExc)

	assert.Empty(t, testErrorDetails.Data, "Error Detials Data stack trace should not have been set")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
}

func TestGetErrorMessage_error(t *testing.T) {
	errorText := "This is an error"
	err := errors.New(errorText)
	returnText := getErrorMessage(err)

	assert.Equal(t, errorText, returnText, "Expected error message: %v", errorText)
}

func TestGetErrorMessage_string(t *testing.T) {
	errorText := "This is a string"
	returnText := getErrorMessage(errorText)

	assert.Equal(t, errorText, returnText, "Expected error message: %v", errorText)
}
