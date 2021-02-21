package todos

import "github.com/ahsan/todo/dao"

type Todos struct {
	dao dao.TodoDAO
}

func TodosFactory() Todos {
	return Todos{dao: dao.DaoFactory()}
}
