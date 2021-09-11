package main

const HOME_WIN_SCORE = 1

func TournamentWinner(competitions [][]string, results []int) string {
	teamScoreHash := map[string]int{"": 0}
	winner := ""
	winnerScore := 0
	for idx, comp := range competitions {
		homeTeam := comp[0]
		awayTeam := comp[1]
		if results[idx] == 1 {
			teamScoreHash[homeTeam] += 3
			if teamScoreHash[homeTeam] > winnerScore {
				winnerScore = teamScoreHash[homeTeam]
				winner = homeTeam
			}
		} else {
			teamScoreHash[awayTeam] += 3
			if teamScoreHash[awayTeam] > winnerScore {
				winnerScore = teamScoreHash[awayTeam]
				winner = awayTeam
			}
		}

	}
	return winner
}
