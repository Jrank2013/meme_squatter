package web

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

var memes = [4]string{
	"https://api.memegen.link/images/buzz/memes/memes_everywhere.webp",
	"https://api.memegen.link/images/fine/breaking_production_at_4:30_pm/this_is_fine.webp?token=40gzthhk9isnk93mm482.webp",
	"https://api.memegen.link/images/right/senior_developer/junior_developer/Put_it_in_the_backlog./So_we_can_fix_it_later,_right~q/So_we_can_fix_it_later,_right~q.jpg?token=0mombk2a9y830pj22koi.webp",
	"https://api.memegen.link/images/headaches/Breaking_Production.jpg?token=kxdztwp3cq0vshbuucga.webp",
}

func index(c *gin.Context) {
	meme := memes[rand.Int()%4]
	c.HTML(200, "index.html", gin.H{"Meme": meme, "Redirect_url": c.GetString("host")})

}
