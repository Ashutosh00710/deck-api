package deck

import (
	"github.com/gin-gonic/gin"
)

func (d *Deck) Open() gin.H {
	return gin.H{
		"deck_id":   d.ID,
		"shuffled":  d.Shuffle,
		"remaining": len(d.Cards),
		"cards":     d.Cards,
	}
}
