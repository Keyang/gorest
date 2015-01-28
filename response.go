package gorest

type ERRResponse struct {
	//Not status code but business error code
	Code int
	Err  string
}
