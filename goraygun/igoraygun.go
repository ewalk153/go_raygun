package goraygun

//Interface provied for mocking and extensibility
type Igoraygun interface {
	LoadRaygunSettings() error
	RaygunRecovery()
}
