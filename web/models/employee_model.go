//usr/bin/env go run "$0" "$@"; exit "$?"

package models

type Employee struct {

    ID int  `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}



