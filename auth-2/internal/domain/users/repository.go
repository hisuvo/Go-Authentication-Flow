package users

type Repository interface {
	Create(req User) error
	Get() (*User, error)
}

type repository struct{} // here includ database

func NewRepository() Repository {
	return &repository{
		// db: db,
	}
}

func (r *repository) Create(req User) error {
	return nil
}

func (r *repository) Get() (*User, error) {
	return nil, nil
}