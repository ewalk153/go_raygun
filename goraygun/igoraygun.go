package goraygun

type Igoraygun interface {
	LoadRaygunSettings() error
	RaygunRecovery()
}
