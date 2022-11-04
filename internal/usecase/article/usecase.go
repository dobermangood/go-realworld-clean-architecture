package article

//go:generate mockgen -source=./registry.go -destination=./registry_mock.go -package=registry
type Repo interface {
	GetArticlesRepo
}

type UseCase struct {
	repo Repo
}

func New(repo Repo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
