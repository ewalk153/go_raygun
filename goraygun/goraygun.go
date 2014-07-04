// Copyright 2013 Marks and Spencer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package goraygun is to be uesd as a recover from panic's and post
	to RayGun (http://raygun.io)
	This can be used in a parent of multiple children, see example.
	There is the ability to use the RaygunRecovery() in multiple nested functions.
	A blog on how to use it can be found here http://mcpugh.com
*/
package goraygun

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var (
	rayGunConfig         RayGunConfigFile
	rayGunConfigFileName = "RayGunConfig.json"
)

type Goraygun struct {
}

//LoadRaygunSettings sets up the enviroment details.
//This needs to be called before using RaygunRecovery
//
//RayGunConfig.json file has to be in your local package with your RayGun Settings
//
//		goraygun := new(goraygun.Goraygun)
//
// 		err := goraygun.LoadRaygunSettings()
// 		if err != nil {
// 			log.Printf("Error loading goraygun config: %v", err)
//		}
//
func (g *Goraygun) LoadRaygunSettings() error {

	data, err := ioutil.ReadFile(rayGunConfigFileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &rayGunConfig)
	if err != nil {
		return err
	}

	return nil
}

//RaygunRecovery will handle any panic that may happen,
//Post to RayGun and then return caller function safely.
//LoadRayGunSettings must have been successfuly called
//before using this function.
// 		defer goraygun.RaygunRecovery()
//
// 		mimicPanic()
//
// 		log.Println("Have recovered and now can continue or fail")
//
func (g *Goraygun) RaygunRecovery() {
	if err := recover(); err != nil {

		//_, filePath, line, _ := runtime.Caller(4)

		stack := make([]byte, 1<<16)
		stack = stack[:runtime.Stack(stack, false)]
		errorMsg := getErrorMessage(err)
		//errorMsg := err.(error).Error()

		sendPanicRayGun(stack, errorMsg)

		log.Println("Recovered from panic:", err)

	}
}

func SendError(err error) {
	stack := make([]byte, 1<<16)
	stack = stack[:runtime.Stack(stack, false)]
	errorMsg := getErrorMessage(err)
	sendPanicRayGun(stack, errorMsg)
}

func getErrorMessage(err interface{}) string {

	t := reflect.TypeOf(err).Kind()
	// check if the type returned from recover is an error
	if t == reflect.TypeOf((*error)(nil)).Kind() {
		return err.(error).Error()
	}
	// check if the type returned from recover is a string
	if t == reflect.String {
		return err.(string)
	}
	// check if the type returned from recover is a uintptr
	if t == reflect.Uintptr {
		return err.(error).Error()
	}

	return "TODO:Need to implement other type: " + t.String()

}

func sendPanicRayGun(exception []byte, errMsg string) {

	var rayError entrie

	rayError.OccurredOn = time.Now().Format("2006-01-02T15:04:05Z")
	rayError.Details = getDetails(errMsg, exception)

	raygunPost(rayError)

	log.Println("Raygun message sent")
}

func raygunPost(rayError entrie) {

	data, err := json.Marshal(rayError)
	if err != nil {
		log.Printf("Error Marshalling RayGun Message: %v:", err)
	}

	req, err := http.NewRequest("POST", rayGunConfig.RaygunEndpoint, strings.NewReader(string(data)))
	if err != nil {
		log.Printf("Error creating POST request: %v", err)
	}

	req.Header.Set("X-ApiKey", rayGunConfig.RaygunAPIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending Request: %v", err)
	}

	if resp.StatusCode != http.StatusAccepted {
		log.Printf("Error status sent back: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}
}
