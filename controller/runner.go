package controller

/*
	Runner is an abstract that triggers snake and
	field interaction and provides controllers' signals handling
*/
// fixme: should return a error
type Runner interface {
	Run() chan error
}
