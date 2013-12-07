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

func readByteException(t *testing.T) []byte {

	byteException := "../testdata/byteException"

	data, err := ioutil.ReadFile(byteException)
	if err != nil {
		t.Errorf("Error Reading byteException test file: %v", err)
	}

	return data
}

func TestProcessException(t *testing.T) {

	testExc := readByteException(t)
	testLine := 266

	expectedPackageName := "main"
	expectedFuncName := "otherMethod"

	packageName, funcName := processException(testExc, testLine)

	assert.Equal(t, packageName, expectedPackageName, "Expected %v, Recived %v", expectedPackageName, packageName)
	assert.Contains(t, funcName, expectedFuncName, "Expected %v, Recived %v", expectedFuncName, funcName)
}

func TestProcessException_noException(t *testing.T) {

	var testExc []byte
	testLine := 222

	packageName, funcName := processException(testExc, testLine)

	assert.Empty(t, packageName, "Package name should not have been returned")
	assert.Empty(t, funcName, "func name should not have been returned")
}

func TestProcessException_hasException_incorrectLineNum(t *testing.T) {

	testExc := readByteException(t)
	testLine := 222

	packageName, funcName := processException(testExc, testLine)

	assert.Empty(t, packageName, "Package name should not have been returned")
	assert.Empty(t, funcName, "func name should not have been returned")
}

func TestGetErrorDetials(t *testing.T) {

	testExc := readByteException(t)
	errMsg := "errMsg"
	testLine := 266
	filePath := "filePath"
	expectedClassName := "main"

	testErrorDetails := getErrorDetails(errMsg, testExc, filePath, testLine)

	assert.NotEmpty(t, testErrorDetails.Data, "Error Details Data stack trace should have been set")
	assert.NotEmpty(t, testErrorDetails.StackTrace, "Error Details Stack trace should have been set")
	assert.Equal(t, 1, len(testErrorDetails.StackTrace), "There should have been a stact traces returned.")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
	assert.Equal(t, testErrorDetails.ClassName, expectedClassName, "Expected %v, Recived %v", expectedClassName, testErrorDetails.ClassName)
}

func TestGetErrorDetails_noException(t *testing.T) {

	var testExc []byte
	errMsg := "errMsg"
	filePath := "filePath"

	testErrorDetails := getErrorDetails(errMsg, testExc, filePath, 1)

	assert.Empty(t, testErrorDetails.Data, "Error Detials Data stack trace should not have been set")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
}

func TestGetErrorDetails_hasException_noline(t *testing.T) {

	testExc := readByteException(t)
	errMsg := "errMsg"
	filePath := "filePath"
	var testLine int

	testErrorDetails := getErrorDetails(errMsg, testExc, filePath, testLine)

	assert.NotEmpty(t, testErrorDetails.Data, "Error Details Data stack trace should have been set")
	assert.Empty(t, testErrorDetails.StackTrace, "Error Details Stack trace should be empty")
	assert.Equal(t, 0, len(testErrorDetails.StackTrace), "There should not have been any stact traces returned.")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
	assert.Empty(t, testErrorDetails.ClassName, "Class name should not have been set")
}

func TestGetErrorDetials_hasExceptionAndline_noFilePath(t *testing.T) {

	testExc := readByteException(t)
	errMsg := "errMsg"
	testLine := 266
	var filePath string

	testErrorDetails := getErrorDetails(errMsg, testExc, filePath, testLine)

	assert.NotEmpty(t, testErrorDetails.Data, "Error Details Data stack trace should have been set")
	assert.Empty(t, testErrorDetails.StackTrace, "Error Details Stack trace should be empty")
	assert.Equal(t, 0, len(testErrorDetails.StackTrace), "There should not have been any stact traces returned.")
	assert.Equal(t, errMsg, testErrorDetails.Message, "Expected %v, Recived %v", errMsg, testErrorDetails.Message)
	assert.Empty(t, testErrorDetails.ClassName, "Class name should not have been set")
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
