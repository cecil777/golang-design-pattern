package abstractfactory

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func ExampleRdbFactory() {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func ExampleXmlFactory() {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
