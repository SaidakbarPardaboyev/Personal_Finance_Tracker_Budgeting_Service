package services

import (
	pb "budgeting_service/genproto/budgeting"
	"budgeting_service/pkg/logger"
	"budgeting_service/storage"
)

type IServiceManager interface {
	AccountService() pb.AccountsServiceServer
	BudgetService() pb.BudgetsServiceServer
	CategoryService() pb.CategoriesServiceServer
	GoalService() pb.GoalsServiceServer
	TransactionService() pb.TransactionsServiceServer
}

type serviceManager struct {
	accountService     pb.AccountsServiceServer
	budgetService      pb.BudgetsServiceServer
	categoryService    pb.CategoriesServiceServer
	goalService        pb.GoalsServiceServer
	transactionService pb.TransactionsServiceServer
}

func NewIServiceManager(storage storage.IStorage, log logger.ILogger) IServiceManager {

	return &serviceManager{
		accountService:     NewAccountService(storage, log),
		budgetService:      NewBudgetService(storage, log),
		categoryService:    NewCategoryService(storage, log),
		goalService:        NewGoalService(storage, log),
		transactionService: NewTransactionService(storage, log),
	}
}

func (s *serviceManager) AccountService() pb.AccountsServiceServer {
	return s.accountService
}

func (s *serviceManager) BudgetService() pb.BudgetsServiceServer {
	return s.budgetService
}

func (s *serviceManager) CategoryService() pb.CategoriesServiceServer {
	return s.categoryService
}

func (s *serviceManager) GoalService() pb.GoalsServiceServer {
	return s.goalService
}

func (s *serviceManager) TransactionService() pb.TransactionsServiceServer {
	return s.transactionService
}
