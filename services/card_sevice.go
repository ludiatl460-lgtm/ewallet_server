package services
import(
	"backend/models"
	"backend/config"
)

func GetAllCards()(list []models.Card,err error){
	err = config.DB.Select(&list,`SELECT * FROM card ORDER BY id`)
	return list,err
}

func GetCardByID(id int64)(*models.Card,error){
	var c models.Card
	err := config.DB.Get(&c,`SELECT * FROM card WHERE id = $1`,id)
	return &c,err
}

func CreateCard(c *models.Card)(*models.Card,error){
	query := `INSERT INTO card (user_id,card_token,card_brand,holder_name,expiry,card_type,status) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *`
	var created models.Card
	err := config.DB.Get(&created,query,c.UserID,c.CardToken,c.CardBrand,c.HolderName,c.Expiry,c.CardType,c.Status)
	return &created,err
}

func UpdateCard(id int64,c *models.Card)(*models.Card,error){
	query := `UPDATE card SET user_id=$1,card_token=$2,card_brand=$3,holder_name=$4,expiry=$5,card_type=$6,status=$7 WHERE id=$8 RETURNING *`
	var updated models.Card
	err := config.DB.Get(&updated,query,c.UserID,c.CardToken,c.CardBrand,c.HolderName,c.Expiry,c.CardType,c.Status,id)
	return &updated,err
}

func DeleteCard(id int64)(*models.Card,error){
	query := `DELETE FROM card WHERE id=$1 RETURNING *`
	var deleted models.Card
	err := config.DB.Get(&deleted,query,id)
	return &deleted,err
}