package lib

// POST constant identifier
const POST = "POST"

// GET constant identifier
const GET = "GET"

// Response struct
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
