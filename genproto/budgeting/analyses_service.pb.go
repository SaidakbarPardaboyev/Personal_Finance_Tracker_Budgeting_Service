// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: analyses_service.proto

package budgeting

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpendingReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName  string  `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`    // from accounts
	AccountType  string  `protobuf:"bytes,2,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`    // from accounts
	CategoryName string  `protobuf:"bytes,3,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"` // from categories
	Amount       float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`                               // from transactions
	Description  string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`                       // from transactions
	Date         string  `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`                                     // from transactions
}

func (x *SpendingReport) Reset() {
	*x = SpendingReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyses_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendingReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendingReport) ProtoMessage() {}

func (x *SpendingReport) ProtoReflect() protoreflect.Message {
	mi := &file_analyses_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendingReport.ProtoReflect.Descriptor instead.
func (*SpendingReport) Descriptor() ([]byte, []int) {
	return file_analyses_service_proto_rawDescGZIP(), []int{0}
}

func (x *SpendingReport) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *SpendingReport) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *SpendingReport) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *SpendingReport) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SpendingReport) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SpendingReport) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type SpendingReports struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpendingReports []*SpendingReport `protobuf:"bytes,1,rep,name=SpendingReports,proto3" json:"SpendingReports,omitempty"`
	UserId          string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Count           int32             `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SpendingReports) Reset() {
	*x = SpendingReports{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyses_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendingReports) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendingReports) ProtoMessage() {}

func (x *SpendingReports) ProtoReflect() protoreflect.Message {
	mi := &file_analyses_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendingReports.ProtoReflect.Descriptor instead.
func (*SpendingReports) Descriptor() ([]byte, []int) {
	return file_analyses_service_proto_rawDescGZIP(), []int{1}
}

func (x *SpendingReports) GetSpendingReports() []*SpendingReport {
	if x != nil {
		return x.SpendingReports
	}
	return nil
}

func (x *SpendingReports) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SpendingReports) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BudgetPerformance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName   string  `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`         // from accounts
	AccountType   string  `protobuf:"bytes,2,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`         // from accounts
	CategoryName  string  `protobuf:"bytes,3,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`      // from categories
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`    // from budgets (amount field)
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"` // from transactions
	Period        string  `protobuf:"bytes,6,opt,name=period,proto3" json:"period,omitempty"`                                      // from budgets
	BudgetType    string  `protobuf:"bytes,7,opt,name=budget_type,json=budgetType,proto3" json:"budget_type,omitempty"`            // from categories
	StartDate     string  `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`               // from budgets
	EndDate       string  `protobuf:"bytes,9,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`                     // from budgets
}

func (x *BudgetPerformance) Reset() {
	*x = BudgetPerformance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyses_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetPerformance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetPerformance) ProtoMessage() {}

func (x *BudgetPerformance) ProtoReflect() protoreflect.Message {
	mi := &file_analyses_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetPerformance.ProtoReflect.Descriptor instead.
func (*BudgetPerformance) Descriptor() ([]byte, []int) {
	return file_analyses_service_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetPerformance) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *BudgetPerformance) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *BudgetPerformance) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *BudgetPerformance) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *BudgetPerformance) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *BudgetPerformance) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *BudgetPerformance) GetBudgetType() string {
	if x != nil {
		return x.BudgetType
	}
	return ""
}

func (x *BudgetPerformance) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BudgetPerformance) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

var File_analyses_service_proto protoreflect.FileDescriptor

var file_analyses_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x0f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x0f, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0f, 0x53,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbd, 0x02,
	0x0a, 0x11, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x32, 0xb4, 0x02,
	0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1a, 0x2e,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x1a, 0x1a, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12,
	0x51, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1c, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x1a, 0x0f, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x6f, 0x61, 0x6c, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analyses_service_proto_rawDescOnce sync.Once
	file_analyses_service_proto_rawDescData = file_analyses_service_proto_rawDesc
)

func file_analyses_service_proto_rawDescGZIP() []byte {
	file_analyses_service_proto_rawDescOnce.Do(func() {
		file_analyses_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_analyses_service_proto_rawDescData)
	})
	return file_analyses_service_proto_rawDescData
}

var file_analyses_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_analyses_service_proto_goTypes = []interface{}{
	(*SpendingReport)(nil),    // 0: budgeting.SpendingReport
	(*SpendingReports)(nil),   // 1: budgeting.SpendingReports
	(*BudgetPerformance)(nil), // 2: budgeting.BudgetPerformance
	(*PrimaryKey)(nil),        // 3: budgeting.PrimaryKey
	(*Goal)(nil),              // 4: budgeting.Goal
}
var file_analyses_service_proto_depIdxs = []int32{
	0, // 0: budgeting.SpendingReports.SpendingReports:type_name -> budgeting.SpendingReport
	3, // 1: budgeting.AnalysesService.GetSpendingReport:input_type -> budgeting.PrimaryKey
	3, // 2: budgeting.AnalysesService.GetIncomeReport:input_type -> budgeting.PrimaryKey
	3, // 3: budgeting.AnalysesService.GetBudgetPerformenceReport:input_type -> budgeting.PrimaryKey
	3, // 4: budgeting.AnalysesService.GetGoalsProgressReport:input_type -> budgeting.PrimaryKey
	1, // 5: budgeting.AnalysesService.GetSpendingReport:output_type -> budgeting.SpendingReports
	1, // 6: budgeting.AnalysesService.GetIncomeReport:output_type -> budgeting.SpendingReports
	2, // 7: budgeting.AnalysesService.GetBudgetPerformenceReport:output_type -> budgeting.BudgetPerformance
	4, // 8: budgeting.AnalysesService.GetGoalsProgressReport:output_type -> budgeting.Goal
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_analyses_service_proto_init() }
func file_analyses_service_proto_init() {
	if File_analyses_service_proto != nil {
		return
	}
	file_budgeting_proto_init()
	file_goals_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_analyses_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendingReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_analyses_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpendingReports); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_analyses_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BudgetPerformance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_analyses_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analyses_service_proto_goTypes,
		DependencyIndexes: file_analyses_service_proto_depIdxs,
		MessageInfos:      file_analyses_service_proto_msgTypes,
	}.Build()
	File_analyses_service_proto = out.File
	file_analyses_service_proto_rawDesc = nil
	file_analyses_service_proto_goTypes = nil
	file_analyses_service_proto_depIdxs = nil
}
