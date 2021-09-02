package graph

import "github.com/HuanLiu-hotstar/search-graphql/graph/model"

var listdata = []model.ListData{
	{
		ID:   "1",
		Type: "List",
		Poster: []*model.Poster{
			{
				ID:   "p1",
				Name: "luka",
				URL:  "www.hotstar.com/movie/luka",
			},

			{
				ID:   "p2",
				Name: "out-guy",
				URL:  "www.hotstar.com/movie/out-guy",
			},
		},
	},
}

func convertListData() []model.SearchTamplate {
	res := []model.SearchTamplate{}
	for _, e := range listdata {
		res = append(res, e)
	}
	return res
}

func convertScoreCardData() []model.SearchTamplate {
	res := []model.SearchTamplate{}
	for _, e := range scorecardData {
		res = append(res, e)
	}
	return res
}

var scorecardData = []model.ScoreCardData{
	{
		ID:   "s1",
		Type: "ScoreCard",
		ScoreCard: []*model.ScoreCard{
			{ID: "one",
				Teams: []*model.Team{
					{
						ID:   "player1",
						Name: "first player",
					},
					{
						ID:   "player2",
						Name: "second player",
					},
				},
			},
		},
	},
}
