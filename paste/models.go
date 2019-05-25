package paste

import (
	"github.com/jinzhu/gorm"
	"hash/fnv"

	"time"
)

var db *gorm.DB

type paste struct {
	ID        int    `gorm:"primary_key"`
	Title     string `gorm:"type:varchar(128)"`
	Hash      uint32 `gorm:"unique_index:hash_idx;"`
	Content   string `gorm:"type:TEXT"`
	Language  string `gorm:"type:varchar(64)"`
	Long      int     `gorm:"type:int"`
	CreatedAt time.Time
}
func CreateTables(){
	if !db.HasTable(&paste{}){
		db.CreateTable(&paste{})
	}
}
func Index()error{
	return nil
}
func Get(hash uint64,p *paste) error{

	return db.Where("hash = ?", hash).First(p).Error
}

func Create(p *paste) error {
	h := fnv.New32a()

	_,err := h.Write([]byte(p.Title + p.Content))
	if err != nil{
		return err
	}
	var count int
	db.Model(&paste{}).Where("hash = ?", h.Sum32()).Count(&count)
	if count != 0{
		db.Where("hash = ?",h.Sum32()).First(p)
		return db.Error
	}
	p.Hash = h.Sum32()
	db.NewRecord(p)
	db.Create(&p)
	return db.Error

}
