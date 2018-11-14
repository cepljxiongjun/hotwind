package handles

// Handle define all func
type Handle interface {
	Get()
	Create()
	Update()
	Delete()
}
