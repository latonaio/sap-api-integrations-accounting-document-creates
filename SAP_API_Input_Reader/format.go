package sap_api_input_reader

import (
	"sap-api-integrations-accounting-document-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToItem() *requests.Item {
	data := sdc.AccountingDocument
	return &requests.Item{
		CompanyCode:                    data.CompanyCode,
		FiscalYear:                     data.FiscalYear,
		AccountingDocument:             data.AccountingDocument,
		AccountingDocumentItem:         data.AccountingDocumentItem,
		CompanyCodeName:                data.CompanyCodeName,
		ChartOfAccounts:                data.ChartOfAccounts,
		AccountingDocumentItemType:     data.AccountingDocumentItemType,
		ClearingDate:                   data.ClearingDate,
		ClearingCreationDate:           data.ClearingCreationDate,
		ClearingAccountingDocument:     data.ClearingAccountingDocument,
		IsCleared:                      data.IsCleared,
		PostingKey:                     data.PostingKey,
		FinancialAccountType:           data.FinancialAccountType,
		SpecialGLCode:                  data.SpecialGLCode,
		SpecialGLTransactionType:       data.SpecialGLTransactionType,
		DebitCreditCode:                data.DebitCreditCode,
		BusinessArea:                   data.BusinessArea,
		BusinessAreaName:               data.BusinessAreaName,
		PartnerBusinessArea:            data.PartnerBusinessArea,
		TaxCode:                        data.TaxCode,
		WithholdingTaxCode:             data.WithholdingTaxCode,
		TaxType:                        data.TaxType,
		TransactionTypeDetermination:   data.TransactionTypeDetermination,
		ValueDate:                      data.ValueDate,
		AssignmentReference:            data.AssignmentReference,
		DocumentItemText:               data.DocumentItemText,
		PartnerCompany:                 data.PartnerCompany,
		FinancialTransactionType:       data.FinancialTransactionType,
		CorporateGroupAccount:          data.CorporateGroupAccount,
		PlanningLevel:                  data.PlanningLevel,
		ControllingArea:                data.ControllingArea,
		ControllingAreaName:            data.ControllingAreaName,
		CostCenter:                     data.CostCenter,
		CostCenterName:                 data.CostCenterName,
		Project:                        data.Project,
		OrderID:                        data.OrderID,
		BillingDocument:                data.BillingDocument,
		SalesDocument:                  data.SalesDocument,
		SalesDocumentItem:              data.SalesDocumentItem,
		ScheduleLine:                   data.ScheduleLine,
		MasterFixedAsset:               data.MasterFixedAsset,
		FixedAsset:                     data.FixedAsset,
		AssetTransactionType:           data.AssetTransactionType,
		AssetValueDate:                 data.AssetValueDate,
		PersonnelNumber:                data.PersonnelNumber,
		IsSalesRelated:                 data.IsSalesRelated,
		LineItemDisplayIsEnabled:       data.LineItemDisplayIsEnabled,
		IsOpenItemManaged:              data.IsOpenItemManaged,
		IsNotCashDiscountLiable:        data.IsNotCashDiscountLiable,
		IsAutomaticallyCreated:         data.IsAutomaticallyCreated,
		IsUsedInPaymentTransaction:     data.IsUsedInPaymentTransaction,
		OperationalGLAccount:           data.OperationalGLAccount,
		GLAccount:                      data.GLAccount,
		GLAccountName:                  data.GLAccountName,
		GLAccountLongName:              data.GLAccountLongName,
		Customer:                       data.Customer,
		CustomerName:                   data.CustomerName,
		Supplier:                       data.Supplier,
		SupplierName:                   data.SupplierName,
		BranchAccount:                  data.BranchAccount,
		IsBalanceSheetAccount:          data.IsBalanceSheetAccount,
		ProfitLossAccountType:          data.ProfitLossAccountType,
		SpecialGLAccountAssignment:     data.SpecialGLAccountAssignment,
		DueCalculationBaseDate:         data.DueCalculationBaseDate,
		PaymentTerms:                   data.PaymentTerms,
		CashDiscount1Days:              data.CashDiscount1Days,
		CashDiscount2Days:              data.CashDiscount2Days,
		NetPaymentDays:                 data.NetPaymentDays,
		CashDiscount1Percent:           data.CashDiscount1Percent,
		CashDiscount2Percent:           data.CashDiscount2Percent,
		PaymentMethod:                  data.PaymentMethod,
		PaymentBlockingReason:          data.PaymentBlockingReason,
		FixedCashDiscount:              data.FixedCashDiscount,
		HouseBank:                      data.HouseBank,
		BPBankAccountInternalID:        data.BPBankAccountInternalID,
		TaxDistributionCode1:           data.TaxDistributionCode1,
		TaxDistributionCode2:           data.TaxDistributionCode2,
		TaxDistributionCode3:           data.TaxDistributionCode3,
		InvoiceReference:               data.InvoiceReference,
		InvoiceReferenceFiscalYear:     data.InvoiceReferenceFiscalYear,
		InvoiceItemReference:           data.InvoiceItemReference,
		FollowOnDocumentType:           data.FollowOnDocumentType,
		StateCentralBankPaymentReason:  data.StateCentralBankPaymentReason,
		SupplyingCountry:               data.SupplyingCountry,
		InvoiceList:                    data.InvoiceList,
		BillOfExchangeUsage:            data.BillOfExchangeUsage,
		DunningKey:                     data.DunningKey,
		DunningBlockingReason:          data.DunningBlockingReason,
		LastDunningDate:                data.LastDunningDate,
		DunningLevel:                   data.DunningLevel,
		DunningArea:                    data.DunningArea,
		WithholdingTaxCertificate:      data.WithholdingTaxCertificate,
		Material:                       data.Material,
		Product:                        data.Product,
		Plant:                          data.Plant,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		AccountAssignmentNumber:        data.AccountAssignmentNumber,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		MaterialPriceControl:           data.MaterialPriceControl,
		ValuationArea:                  data.ValuationArea,
		InventoryValuationType:         data.InventoryValuationType,
		VATRegistration:                data.VATRegistration,
		DelivOfGoodsDestCountry:        data.DelivOfGoodsDestCountry,
		PaymentDifferenceReason:        data.PaymentDifferenceReason,
		ProfitCenter:                   data.ProfitCenter,
		ProfitCenterName:               data.ProfitCenterName,
		JointVenture:                   data.JointVenture,
		JointVentureCostRecoveryCode:   data.JointVentureCostRecoveryCode,
		JointVentureEquityGroup:        data.JointVentureEquityGroup,
		TreasuryContractType:           data.TreasuryContractType,
		AssetContract:                  data.AssetContract,
		CashFlowType:                   data.CashFlowType,
		TaxJurisdiction:                data.TaxJurisdiction,
		RealEstateObject:               data.RealEstateObject,
		SettlementReferenceDate:        data.SettlementReferenceDate,
		CommitmentItem:                 data.CommitmentItem,
		CostObject:                     data.CostObject,
		ProjectNetwork:                 data.ProjectNetwork,
		OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
		OrderIntBillOfOperationsItem:   data.OrderIntBillOfOperationsItem,
		WBSElementInternalID:           data.WBSElementInternalID,
		ProfitabilitySegment:           data.ProfitabilitySegment,
		JointVentureEquityType:         data.JointVentureEquityType,
		IsEUTriangularDeal:             data.IsEUTriangularDeal,
		CostOriginGroup:                data.CostOriginGroup,
		CompanyCodeCurrencyDetnMethod:  data.CompanyCodeCurrencyDetnMethod,
		ClearingIsReversed:             data.ClearingIsReversed,
		PaymentMethodSupplement:        data.PaymentMethodSupplement,
		AlternativeGLAccount:           data.AlternativeGLAccount,
		FundsCenter:                    data.FundsCenter,
		Fund:                           data.Fund,
		PartnerProfitCenter:            data.PartnerProfitCenter,
		Reference1IDByBusinessPartner:  data.Reference1IDByBusinessPartner,
		Reference2IDByBusinessPartner:  data.Reference2IDByBusinessPartner,
		IsNegativePosting:              data.IsNegativePosting,
		PaymentCardItem:                data.PaymentCardItem,
		PaymentCardPaymentSettlement:   data.PaymentCardPaymentSettlement,
		CreditControlArea:              data.CreditControlArea,
		Reference3IDByBusinessPartner:  data.Reference3IDByBusinessPartner,
		DataExchangeInstruction1:       data.DataExchangeInstruction1,
		DataExchangeInstruction2:       data.DataExchangeInstruction2,
		DataExchangeInstruction3:       data.DataExchangeInstruction3,
		DataExchangeInstruction4:       data.DataExchangeInstruction4,
		Region:                         data.Region,
		HasPaymentOrder:                data.HasPaymentOrder,
		PaymentReference:               data.PaymentReference,
		TaxDeterminationDate:           data.TaxDeterminationDate,
		ClearingItem:                   data.ClearingItem,
		BusinessPlace:                  data.BusinessPlace,
		TaxSection:                     data.TaxSection,
		CostCtrActivityType:            data.CostCtrActivityType,
		AccountsReceivableIsPledged:    data.AccountsReceivableIsPledged,
		BusinessProcess:                data.BusinessProcess,
		GrantID:                        data.GrantID,
		FunctionalArea:                 data.FunctionalArea,
		FunctionalAreaName:             data.FunctionalAreaName,
		CustomerIsInExecution:          data.CustomerIsInExecution,
		FundedProgram:                  data.FundedProgram,
		ClearingDocFiscalYear:          data.ClearingDocFiscalYear,
		LedgerGLLineItem:               data.LedgerGLLineItem,
		Segment:                        data.Segment,
		SegmentName:                    data.SegmentName,
		PartnerSegment:                 data.PartnerSegment,
		PartnerFunctionalArea:          data.PartnerFunctionalArea,
		HouseBankAccount:               data.HouseBankAccount,
		CostElement:                    data.CostElement,
		SEPAMandate:                    data.SEPAMandate,
		ReferenceDocumentType:          data.ReferenceDocumentType,
		OriginalReferenceDocument:      data.OriginalReferenceDocument,
		ReferenceDocumentLogicalSystem: data.ReferenceDocumentLogicalSystem,
		AccountingDocumentItemRef:      data.AccountingDocumentItemRef,
		FiscalPeriod:                   data.FiscalPeriod,
		AccountingDocumentCategory:     data.AccountingDocumentCategory,
		AccountingDocumentCategoryName: data.AccountingDocumentCategoryName,
		PostingDate:                    data.PostingDate,
		DocumentDate:                   data.DocumentDate,
		AccountingDocumentType:         data.AccountingDocumentType,
		AccountingDocumentTypeName:     data.AccountingDocumentTypeName,
		NetDueDate:                     data.NetDueDate,
		CashDiscount1DueDate:           data.CashDiscount1DueDate,
		CashDiscount2DueDate:           data.CashDiscount2DueDate,
		OffsettingAccount:              data.OffsettingAccount,
		OffsettingAccountType:          data.OffsettingAccountType,
		PartnerFund:                    data.PartnerFund,
		PartnerGrant:                   data.PartnerGrant,
		BudgetPeriod:                   data.BudgetPeriod,
		PartnerBudgetPeriod:            data.PartnerBudgetPeriod,
		BranchCode:                     data.BranchCode,
		CompanyCodeCurrency:            data.CompanyCodeCurrency,
		AmountInCompanyCodeCurrency:    data.AmountInCompanyCodeCurrency,
		TaxAmountInCoCodeCrcy:          data.TaxAmountInCoCodeCrcy,
		TaxBaseAmountInCoCodeCrcy:      data.TaxBaseAmountInCoCodeCrcy,
		ValuationDiffAmtInCoCodeCrcy:   data.ValuationDiffAmtInCoCodeCrcy,
		CashDiscountAmtInCoCodeCrcy:    data.CashDiscountAmtInCoCodeCrcy,
		InvoiceAmtInCoCodeCrcy:         data.InvoiceAmtInCoCodeCrcy,
		TransactionCurrency:            data.TransactionCurrency,
		AmountInTransactionCurrency:    data.AmountInTransactionCurrency,
		OriginalTaxBaseAmount:          data.OriginalTaxBaseAmount,
		TaxAmount:                      data.TaxAmount,
		TaxBaseAmountInTransCrcy:       data.TaxBaseAmountInTransCrcy,
		WithholdingTaxBaseAmount:       data.WithholdingTaxBaseAmount,
		PlannedAmtInTransactionCrcy:    data.PlannedAmtInTransactionCrcy,
		CashDiscountBaseAmount:         data.CashDiscountBaseAmount,
		CashDiscountAmount:             data.CashDiscountAmount,
		NetPaymentAmount:               data.NetPaymentAmount,
		WithholdingTaxAmount:           data.WithholdingTaxAmount,
		WithholdingTaxExemptionAmt:     data.WithholdingTaxExemptionAmt,
		InvoiceAmountInFrgnCurrency:    data.InvoiceAmountInFrgnCurrency,
		BalanceTransactionCurrency:     data.BalanceTransactionCurrency,
		AmountInBalanceTransacCrcy:     data.AmountInBalanceTransacCrcy,
		AdditionalCurrency1:            data.AdditionalCurrency1,
		ValuationDiffAmtInAddlCrcy1:    data.ValuationDiffAmtInAddlCrcy1,
		AmountInAdditionalCurrency1:    data.AmountInAdditionalCurrency1,
		AdditionalCurrency2:            data.AdditionalCurrency2,
		AmountInAdditionalCurrency2:    data.AmountInAdditionalCurrency2,
		ValuationDiffAmtInAddlCrcy2:    data.ValuationDiffAmtInAddlCrcy2,
		PaymentCurrency:                data.PaymentCurrency,
		AmountInPaymentCurrency:        data.AmountInPaymentCurrency,
		CreditControlAreaCurrency:      data.CreditControlAreaCurrency,
		HedgedAmount:                   data.HedgedAmount,
		BaseUnit:                       data.BaseUnit,
		Quantity:                       data.Quantity,
		GoodsMovementEntryUnit:         data.GoodsMovementEntryUnit,
		QuantityInEntryUnit:            data.QuantityInEntryUnit,
		PurchasingDocumentPriceUnit:    data.PurchasingDocumentPriceUnit,
		PurchaseOrderQty:               data.PurchaseOrderQty,
		MaterialPriceUnitQty:           data.MaterialPriceUnitQty,
		NumberOfItems:                  data.NumberOfItems,
		AccountingDocumentCreationDate: data.AccountingDocumentCreationDate,
		CreationTime:                   data.CreationTime,
		LastChangeDate:                 data.LastChangeDate,
		ExchangeRateDate:               data.ExchangeRateDate,
		AccountingDocCreatedByUser:     data.AccountingDocCreatedByUser,
		TransactionCode:                data.TransactionCode,
		IntercompanyTransaction:        data.IntercompanyTransaction,
		DocumentReferenceID:            data.DocumentReferenceID,
		RecurringAccountingDocument:    data.RecurringAccountingDocument,
		ReverseDocument:                data.ReverseDocument,
		ReverseDocumentFiscalYear:      data.ReverseDocumentFiscalYear,
		AccountingDocumentHeaderText:   data.AccountingDocumentHeaderText,
		ExchangeRate:                   data.ExchangeRate,
		BusinessTransactionType:        data.BusinessTransactionType,
		BatchInputSession:              data.BatchInputSession,
		FinancialManagementArea:        data.FinancialManagementArea,
		ReversalIsPlanned:              data.ReversalIsPlanned,
		PlannedReversalDate:            data.PlannedReversalDate,
		TaxIsCalculatedAutomatically:   data.TaxIsCalculatedAutomatically,
		TaxBaseAmountIsNetAmount:       data.TaxBaseAmountIsNetAmount,
		SourceCompanyCode:              data.SourceCompanyCode,
		LogicalSystem:                  data.LogicalSystem,
		TaxExchangeRate:                data.TaxExchangeRate,
		ReversalReason:                 data.ReversalReason,
		Branch:                         data.Branch,
		Reference1InDocumentHeader:     data.Reference1InDocumentHeader,
		Reference2InDocumentHeader:     data.Reference2InDocumentHeader,
		InvoiceReceiptDate:             data.InvoiceReceiptDate,
		Ledger:                         data.Ledger,
		LedgerGroup:                    data.LedgerGroup,
		AlternativeReferenceDocument:   data.AlternativeReferenceDocument,
		TaxReportingDate:               data.TaxReportingDate,
		AccountingDocumentClass:        data.AccountingDocumentClass,
		ExchangeRateType:               data.ExchangeRateType,
		LatePaymentReason:              data.LatePaymentReason,
		SalesDocumentCondition:         data.SalesDocumentCondition,
		IsReversal:                     data.IsReversal,
		IsReversed:                     data.IsReversed,
	}
}
