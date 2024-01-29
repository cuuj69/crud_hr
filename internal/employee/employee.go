//usr/bin/env go run "$0" "$@"; exit "$?"

package employee

import (
	
	
	"time"

)

// type Employee struct {

// 	gorm.Model

// 	Name string `json:"name"`
// 	Email string `json:"email"`


// }

type Employee struct{
	
	ID    		uint `gorm:"primaryKey;autoIncrement"`
	Name		string
	Email		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	
}
//the gorm.Model provides basic fields like `ID`, `CreatedAt`, and `DeletedAt`


