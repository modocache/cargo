package comparators

type Less func(value, treeValue interface{}) bool

var IntLess Less = func(value, treeValue interface{}) bool {
	return value.(int) < treeValue.(int)
}

var StringLess Less = func(value, treeValue interface{}) bool {
	return value.(string) < treeValue.(string)
}
