package service

type service interface {
	Page(page uint, size uint) interface{}
	List() interface{}
	Get(id string) (interface{}, error)
	Create(interface{})
	DeleteByID(id string)
	Delete(interface{})
	Update(interface{})
}
