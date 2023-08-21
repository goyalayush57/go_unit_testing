package external

import (
	"fmt"
	caches "gophercon_2023/unit_testing/externaldependency/cache"
	database "gophercon_2023/unit_testing/externaldependency/db"
	"gophercon_2023/unit_testing/externaldependency/model"
)

//https://ports.macports.org/port/go-mockgen/
type RegistrationService interface {
	Register(name, email, pass string) error
}

type DatastoreService interface {
	Insert(model.Row) (model.Row, error)
	Query(query string) []model.Row
	Delete(primaryKey, table, schema string)
}

type CacheService interface {
	Get(key string) string
	Put(key string, value int) error
	Delete(key string)
}
type register struct {
	db    DatastoreService
	cache CacheService
}

/* Register
1. verify the same name does not exist
2. put it in db
3. put it in cache
4. return success/err
*/
func (a *register) Register(name, email, encryptPass string) error {
	if !isValidName(name) || !isValidEmail(email) || !isValidPass(encryptPass) {
		return fmt.Errorf("The provided field is not valid")
	}
	if a.cache.Get(name) != "" {
		return fmt.Errorf("Provided User Name is already taken")
	}
	r, err := a.db.Insert(model.Row{Name: name, Pass: encryptPass, Email: email})
	if err != nil {
		//log error
		return fmt.Errorf("Internal Error,please retry!!")
	}
	a.cache.Put(r.Name, r.ID)
	return nil
}

/* Register
1. verify the same name does not exist
2. put it in db
3. put it in cache
4. return success/err
*/
func Register_noInterface(name, email, encryptPass string) error {
	if !isValidName(name) || !isValidEmail(email) || !isValidPass(encryptPass) {
		return fmt.Errorf("The provided field is not valid")
	}

	if caches.Get(name) != "" {
		return fmt.Errorf("Provided User Name is already taken")
	}
	r, err := database.Insert(model.Row{Name: name, Pass: encryptPass, Email: email})
	if err != nil {
		//log error
		return fmt.Errorf("Internal Error,please retry!!")
	}
	caches.Put(r.Name, r.ID)
	return nil
}

func isValidName(string) bool {
	return true
}
func isValidPass(string) bool {
	return true
}

func isValidEmail(string) bool {
	return true
}
