package user

import (
	"main/db"
	"main/utils"
	"net/http"
)

func Leaderboard(w http.ResponseWriter, r *http.Request)  {
	var userLeaderboard []LeaderboardUser
	results, err := db.Connection.Query("SELECT id, name, surname, email,(SELECT COUNT(*) FROM answers WHERE author_id=users.id) as answers,(SELECT COUNT(*) FROM questions WHERE author_id=users.id) as questions FROM users ORDER BY answers DESC")
	if err!=nil {
		utils.WriteError(w,"Unable to fetch user leaderboard",err,http.StatusInternalServerError)
	}
	for results.Next(){
		var user LeaderboardUser
		err=results.Scan(&user.Id,&user.Name,&user.Surname,&user.Email,&user.Answers,&user.Questions)
		if err!=nil{
			utils.WriteError(w,"Unable to scan user leaderboard",err,http.StatusInternalServerError)
		}
		userLeaderboard=append(userLeaderboard,user)
	}
	utils.WriteSuccess(w,userLeaderboard,true)

}
