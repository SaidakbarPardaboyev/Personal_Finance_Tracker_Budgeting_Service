package grpc

import (
	pb "budgeting_service/genproto/budgeting"
	"budgeting_service/pkg/logger"
	"budgeting_service/services"
	"budgeting_service/storage"

	"google.golang.org/grpc"
)

func SetUpServer(iServiceManager services.IServiceManager, storage storage.IStorage, log logger.ILogger) *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterAccountsServiceServer(grpcServer, iServiceManager.AccountService())
	pb.RegisterBudgetsServiceServer(grpcServer, iServiceManager.BudgetService())
	pb.RegisterCategoriesServiceServer(grpcServer, iServiceManager.CategoryService())
	pb.RegisterGoalsServiceServer(grpcServer, iServiceManager.GoalService())
	pb.RegisterTransactionsServiceServer(grpcServer, iServiceManager.TransactionService())

	return grpcServer
}
