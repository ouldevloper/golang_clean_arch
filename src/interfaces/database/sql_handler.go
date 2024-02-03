package database

type SqlHandler interface {
	Create(obj interface{})
	Find(obj interface{})
	Delete(obj interface{}, id int)
}
