package article

type GetArticlesRepo interface {
	GetArticles()
}

func (u *UseCase) GetArticles() {
	u.repo.GetArticles()
}
