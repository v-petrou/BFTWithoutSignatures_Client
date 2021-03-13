package variables

var (
	// N - Number of processors
	N int

	// F - Number of faulty processors
	F int

	// ID - This processor's id.
	ID int

	// Clients - Size of Clients Set
	Clients int

	// Remote - If we are running locally or remotely
	Remote bool
)

// Initialize - Variables initializer method
func Initialize(id int, n int, c int) {
	ID = id
	N = n
	F = (N - 1) / 3
	Clients = c
	Remote = false
}
