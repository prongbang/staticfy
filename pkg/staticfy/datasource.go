package staticfy

// DataSource is the interface
type DataSource interface {
	Find(asset Assets) (Assets, error)
	Create(asset Assets) (Assets, error)
	Delete(asset Assets) error
}
