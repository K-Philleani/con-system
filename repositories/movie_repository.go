package repositories

import "con-system/datamodels"

type MovieRepository interface {
	GetMovieName() string
}


type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// 模拟赋值给模型
	movie := &datamodels.Movie{
		Name: "电影视频",
	}
	return movie.Name
}