package staticfy

type UseCase interface {
	Upload(asset Assets) (Assets, error)
	Delete(asset Assets) (Assets, error)
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Delete(asset Assets) (Assets, error) {
	data, err := u.Repo.DeleteFile(asset)
	data.UserID = ""
	return data, err
}

func (u *useCase) Upload(asset Assets) (Assets, error) {
	data, err := u.Repo.CreateFile(asset)
	data.UserID = ""
	return data, err
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
