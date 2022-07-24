// Code generated by xgen. DO NOT EDIT.

package schema

import (
	"encoding/xml"
)

// Document ...
type Document *Document

// AccountIdentification3Choice ...
type AccountIdentification3Choice struct {
	IBAN      string                            `xml:"IBAN"`
	BBAN      string                            `xml:"BBAN"`
	UPIC      string                            `xml:"UPIC"`
	PrtryAcct *SimpleIdentificationInformation2 `xml:"PrtryAcct"`
}

// AddressType2Code ...
type AddressType2Code string

// AmountType2Choice ...
type AmountType2Choice struct {
	InstdAmt *CurrencyAndAmount `xml:"InstdAmt"`
	EqvtAmt  *EquivalentAmount  `xml:"EqvtAmt"`
}

// BBANIdentifier ...
type BBANIdentifier string

// BEIIdentifier ...
type BEIIdentifier string

// BICIdentifier ...
type BICIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BatchBookingIndicator ...
type BatchBookingIndicator bool

// BranchAndFinancialInstitutionIdentification3 ...
type BranchAndFinancialInstitutionIdentification3 struct {
	FinInstnId *FinancialInstitutionIdentification5Choice `xml:"FinInstnId"`
	BrnchId    *BranchData                                `xml:"BrnchId"`
}

// BranchData ...
type BranchData struct {
	Id      string          `xml:"Id"`
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr"`
}

// CHIPSUniversalIdentifier ...
type CHIPSUniversalIdentifier string

// CashAccount7 ...
type CashAccount7 struct {
	Id  *AccountIdentification3Choice `xml:"Id"`
	Tp  *CashAccountType2             `xml:"Tp"`
	Ccy string                        `xml:"Ccy"`
	Nm  string                        `xml:"Nm"`
}

// CashAccountType2 ...
type CashAccountType2 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CashAccountType4Code ...
type CashAccountType4Code string

// ChargeBearerType1Code ...
type ChargeBearerType1Code string

// Cheque5 ...
type Cheque5 struct {
	ChqTp       string                       `xml:"ChqTp"`
	ChqNb       string                       `xml:"ChqNb"`
	ChqFr       *NameAndAddress3             `xml:"ChqFr"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd"`
	DlvrTo      *NameAndAddress3             `xml:"DlvrTo"`
	InstrPrty   string                       `xml:"InstrPrty"`
	ChqMtrtyDt  string                       `xml:"ChqMtrtyDt"`
	FrmsCd      string                       `xml:"FrmsCd"`
	MemoFld     string                       `xml:"MemoFld"`
	RgnlClrZone string                       `xml:"RgnlClrZone"`
	PrtLctn     string                       `xml:"PrtLctn"`
}

// ChequeDelivery1Code ...
type ChequeDelivery1Code string

// ChequeDeliveryMethod1Choice ...
type ChequeDeliveryMethod1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ChequeType2Code ...
type ChequeType2Code string

// ClearingChannel2Code ...
type ClearingChannel2Code string

// ClearingSystemMemberIdentification3Choice ...
type ClearingSystemMemberIdentification3Choice struct {
	Id    string `xml:"Id"`
	Prtry string `xml:"Prtry"`
}

// CountryCode ...
type CountryCode string

// CreditTransferTransactionInformation1 ...
type CreditTransferTransactionInformation1 struct {
	PmtId           *PaymentIdentification1                       `xml:"PmtId"`
	PmtTpInf        *PaymentTypeInformation1                      `xml:"PmtTpInf"`
	Amt             *AmountType2Choice                            `xml:"Amt"`
	XchgRateInf     *ExchangeRateInformation1                     `xml:"XchgRateInf"`
	ChrgBr          string                                        `xml:"ChrgBr"`
	ChqInstr        *Cheque5                                      `xml:"ChqInstr"`
	UltmtDbtr       *PartyIdentification8                         `xml:"UltmtDbtr"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt1"`
	IntrmyAgt1Acct  *CashAccount7                                 `xml:"IntrmyAgt1Acct"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt2"`
	IntrmyAgt2Acct  *CashAccount7                                 `xml:"IntrmyAgt2Acct"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification3 `xml:"IntrmyAgt3"`
	IntrmyAgt3Acct  *CashAccount7                                 `xml:"IntrmyAgt3Acct"`
	CdtrAgt         *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt"`
	CdtrAgtAcct     *CashAccount7                                 `xml:"CdtrAgtAcct"`
	Cdtr            *PartyIdentification8                         `xml:"Cdtr"`
	CdtrAcct        *CashAccount7                                 `xml:"CdtrAcct"`
	UltmtCdtr       *PartyIdentification8                         `xml:"UltmtCdtr"`
	InstrForCdtrAgt []*InstructionForCreditorAgent1               `xml:"InstrForCdtrAgt"`
	InstrForDbtrAgt string                                        `xml:"InstrForDbtrAgt"`
	Purp            *Purpose1Choice                               `xml:"Purp"`
	RgltryRptg      []*RegulatoryReporting2                       `xml:"RgltryRptg"`
	Tax             *TaxInformation2                              `xml:"Tax"`
	RltdRmtInf      []*RemittanceLocation1                        `xml:"RltdRmtInf"`
	RmtInf          *RemittanceInformation1                       `xml:"RmtInf"`
}

// CreditorReferenceInformation1 ...
type CreditorReferenceInformation1 struct {
	CdtrRefTp *CreditorReferenceType1 `xml:"CdtrRefTp"`
	CdtrRef   string                  `xml:"CdtrRef"`
}

// CreditorReferenceType1 ...
type CreditorReferenceType1 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
	Issr  string `xml:"Issr"`
}

// CurrencyAndAmountSimpleType ...
type CurrencyAndAmountSimpleType float64

// CurrencyAndAmount ...
type CurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// CurrencyCode ...
type CurrencyCode string

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt"`
	PrvcOfBirth string `xml:"PrvcOfBirth"`
	CityOfBirth string `xml:"CityOfBirth"`
	CtryOfBirth string `xml:"CtryOfBirth"`
}

// DecimalNumber ...
type DecimalNumber float64

// DocumentType2Code ...
type DocumentType2Code string

// DocumentType3Code ...
type DocumentType3Code string

// DunsIdentifier ...
type DunsIdentifier string

// EANGLNIdentifier ...
type EANGLNIdentifier string

// EquivalentAmount ...
type EquivalentAmount struct {
	Amt      *CurrencyAndAmount `xml:"Amt"`
	CcyOfTrf string             `xml:"CcyOfTrf"`
}

// ExchangeRateInformation1 ...
type ExchangeRateInformation1 struct {
	XchgRate float64 `xml:"XchgRate"`
	RateTp   string  `xml:"RateTp"`
	CtrctId  string  `xml:"CtrctId"`
}

// ExchangeRateType1Code ...
type ExchangeRateType1Code string

// ExternalClearingSystemMemberCode ...
type ExternalClearingSystemMemberCode string

// ExternalLocalInstrumentCode ...
type ExternalLocalInstrumentCode string

// ExternalPurposeCode ...
type ExternalPurposeCode string

// FinancialInstitutionIdentification3 ...
type FinancialInstitutionIdentification3 struct {
	BIC         string                                     `xml:"BIC"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId"`
	Nm          string                                     `xml:"Nm"`
	PstlAdr     *PostalAddress1                            `xml:"PstlAdr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId"`
}

// FinancialInstitutionIdentification5Choice ...
type FinancialInstitutionIdentification5Choice struct {
	BIC         string                                     `xml:"BIC"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId"`
	NmAndAdr    *NameAndAddress7                           `xml:"NmAndAdr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId"`
	CmbndId     *FinancialInstitutionIdentification3       `xml:"CmbndId"`
}

// GenericIdentification3 ...
type GenericIdentification3 struct {
	Id   string `xml:"Id"`
	Issr string `xml:"Issr"`
}

// GenericIdentification4 ...
type GenericIdentification4 struct {
	Id   string `xml:"Id"`
	IdTp string `xml:"IdTp"`
}

// GroupHeader1 ...
type GroupHeader1 struct {
	MsgId     string                                        `xml:"MsgId"`
	CreDtTm   string                                        `xml:"CreDtTm"`
	Authstn   []string                                      `xml:"Authstn"`
	BtchBookg bool                                          `xml:"BtchBookg"`
	NbOfTxs   string                                        `xml:"NbOfTxs"`
	CtrlSum   float64                                       `xml:"CtrlSum"`
	Grpg      string                                        `xml:"Grpg"`
	InitgPty  *PartyIdentification8                         `xml:"InitgPty"`
	FwdgAgt   *BranchAndFinancialInstitutionIdentification3 `xml:"FwdgAgt"`
}

// Grouping1Code ...
type Grouping1Code string

// IBANIdentifier ...
type IBANIdentifier string

// IBEIIdentifier ...
type IBEIIdentifier string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// Instruction3Code ...
type Instruction3Code string

// InstructionForCreditorAgent1 ...
type InstructionForCreditorAgent1 struct {
	Cd       string `xml:"Cd"`
	InstrInf string `xml:"InstrInf"`
}

// LocalInstrument1Choice ...
type LocalInstrument1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// Max128Text ...
type Max128Text string

// Max140Text ...
type Max140Text string

// Max15NumericText ...
type Max15NumericText string

// Max16Text ...
type Max16Text string

// Max256Text ...
type Max256Text string

// Max34Text ...
type Max34Text string

// Max35Text ...
type Max35Text string

// Max3Text ...
type Max3Text string

// Max70Text ...
type Max70Text string

// NameAndAddress3 ...
type NameAndAddress3 struct {
	Nm  string          `xml:"Nm"`
	Adr *PostalAddress1 `xml:"Adr"`
}

// NameAndAddress7 ...
type NameAndAddress7 struct {
	Nm      string          `xml:"Nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr"`
}

// OrganisationIdentification2 ...
type OrganisationIdentification2 struct {
	BIC     string                  `xml:"BIC"`
	IBEI    string                  `xml:"IBEI"`
	BEI     string                  `xml:"BEI"`
	EANGLN  string                  `xml:"EANGLN"`
	USCHU   string                  `xml:"USCHU"`
	DUNS    string                  `xml:"DUNS"`
	BkPtyId string                  `xml:"BkPtyId"`
	TaxIdNb string                  `xml:"TaxIdNb"`
	PrtryId *GenericIdentification3 `xml:"PrtryId"`
}

// Party2Choice ...
type Party2Choice struct {
	OrgId  *OrganisationIdentification2 `xml:"OrgId"`
	PrvtId []*PersonIdentification3     `xml:"PrvtId"`
}

// PartyIdentification8 ...
type PartyIdentification8 struct {
	Nm        string          `xml:"Nm"`
	PstlAdr   *PostalAddress1 `xml:"PstlAdr"`
	Id        *Party2Choice   `xml:"Id"`
	CtryOfRes string          `xml:"CtryOfRes"`
}

// PaymentCategoryPurpose1Code ...
type PaymentCategoryPurpose1Code string

// PaymentIdentification1 ...
type PaymentIdentification1 struct {
	InstrId    string `xml:"InstrId"`
	EndToEndId string `xml:"EndToEndId"`
}

// PaymentInstructionInformation1 ...
type PaymentInstructionInformation1 struct {
	PmtInfId        string                                        `xml:"PmtInfId"`
	PmtMtd          string                                        `xml:"PmtMtd"`
	PmtTpInf        *PaymentTypeInformation1                      `xml:"PmtTpInf"`
	ReqdExctnDt     string                                        `xml:"ReqdExctnDt"`
	PoolgAdjstmntDt string                                        `xml:"PoolgAdjstmntDt"`
	Dbtr            *PartyIdentification8                         `xml:"Dbtr"`
	DbtrAcct        *CashAccount7                                 `xml:"DbtrAcct"`
	DbtrAgt         *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`
	DbtrAgtAcct     *CashAccount7                                 `xml:"DbtrAgtAcct"`
	UltmtDbtr       *PartyIdentification8                         `xml:"UltmtDbtr"`
	ChrgBr          string                                        `xml:"ChrgBr"`
	ChrgsAcct       *CashAccount7                                 `xml:"ChrgsAcct"`
	ChrgsAcctAgt    *BranchAndFinancialInstitutionIdentification3 `xml:"ChrgsAcctAgt"`
	CdtTrfTxInf     []*CreditTransferTransactionInformation1      `xml:"CdtTrfTxInf"`
}

// PaymentMethod3Code ...
type PaymentMethod3Code string

// PaymentTypeInformation1 ...
type PaymentTypeInformation1 struct {
	InstrPrty string                  `xml:"InstrPrty"`
	SvcLvl    *ServiceLevel2Choice    `xml:"SvcLvl"`
	ClrChanl  string                  `xml:"ClrChanl"`
	LclInstrm *LocalInstrument1Choice `xml:"LclInstrm"`
	CtgyPurp  string                  `xml:"CtgyPurp"`
}

// PercentageRate ...
type PercentageRate float64

// PersonIdentification3 ...
type PersonIdentification3 struct {
	DrvrsLicNb      string                  `xml:"DrvrsLicNb"`
	CstmrNb         string                  `xml:"CstmrNb"`
	SclSctyNb       string                  `xml:"SclSctyNb"`
	AlnRegnNb       string                  `xml:"AlnRegnNb"`
	PsptNb          string                  `xml:"PsptNb"`
	TaxIdNb         string                  `xml:"TaxIdNb"`
	IdntyCardNb     string                  `xml:"IdntyCardNb"`
	MplyrIdNb       string                  `xml:"MplyrIdNb"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth    `xml:"DtAndPlcOfBirth"`
	OthrId          *GenericIdentification4 `xml:"OthrId"`
	Issr            string                  `xml:"Issr"`
}

// PostalAddress1 ...
type PostalAddress1 struct {
	AdrTp       string   `xml:"AdrTp"`
	AdrLine     []string `xml:"AdrLine"`
	StrtNm      string   `xml:"StrtNm"`
	BldgNb      string   `xml:"BldgNb"`
	PstCd       string   `xml:"PstCd"`
	TwnNm       string   `xml:"TwnNm"`
	CtrySubDvsn string   `xml:"CtrySubDvsn"`
	Ctry        string   `xml:"Ctry"`
}

// Priority2Code ...
type Priority2Code string

// Purpose1Choice ...
type Purpose1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ReferredDocumentAmount1Choice ...
type ReferredDocumentAmount1Choice struct {
	DuePyblAmt   *CurrencyAndAmount `xml:"DuePyblAmt"`
	DscntApldAmt *CurrencyAndAmount `xml:"DscntApldAmt"`
	RmtdAmt      *CurrencyAndAmount `xml:"RmtdAmt"`
	CdtNoteAmt   *CurrencyAndAmount `xml:"CdtNoteAmt"`
	TaxAmt       *CurrencyAndAmount `xml:"TaxAmt"`
}

// ReferredDocumentInformation1 ...
type ReferredDocumentInformation1 struct {
	RfrdDocTp *ReferredDocumentType1 `xml:"RfrdDocTp"`
	RfrdDocNb string                 `xml:"RfrdDocNb"`
}

// ReferredDocumentType1 ...
type ReferredDocumentType1 struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
	Issr  string `xml:"Issr"`
}

// RegulatoryAuthority ...
type RegulatoryAuthority struct {
	AuthrtyNm   string `xml:"AuthrtyNm"`
	AuthrtyCtry string `xml:"AuthrtyCtry"`
}

// RegulatoryReporting2 ...
type RegulatoryReporting2 struct {
	DbtCdtRptgInd string                          `xml:"DbtCdtRptgInd"`
	Authrty       *RegulatoryAuthority            `xml:"Authrty"`
	RgltryDtls    *StructuredRegulatoryReporting2 `xml:"RgltryDtls"`
}

// RegulatoryReportingType1Code ...
type RegulatoryReportingType1Code string

// RemittanceInformation1 ...
type RemittanceInformation1 struct {
	Ustrd []string                            `xml:"Ustrd"`
	Strd  []*StructuredRemittanceInformation6 `xml:"Strd"`
}

// RemittanceLocation1 ...
type RemittanceLocation1 struct {
	RmtId             string           `xml:"RmtId"`
	RmtLctnMtd        string           `xml:"RmtLctnMtd"`
	RmtLctnElctrncAdr string           `xml:"RmtLctnElctrncAdr"`
	RmtLctnPstlAdr    *NameAndAddress3 `xml:"RmtLctnPstlAdr"`
}

// RemittanceLocationMethod1Code ...
type RemittanceLocationMethod1Code string

// ServiceLevel1Code ...
type ServiceLevel1Code string

// ServiceLevel2Choice ...
type ServiceLevel2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// SimpleIdentificationInformation2 ...
type SimpleIdentificationInformation2 struct {
	Id string `xml:"Id"`
}

// StructuredRegulatoryReporting2 ...
type StructuredRegulatoryReporting2 struct {
	Cd  string             `xml:"Cd"`
	Amt *CurrencyAndAmount `xml:"Amt"`
	Inf string             `xml:"Inf"`
}

// StructuredRemittanceInformation6 ...
type StructuredRemittanceInformation6 struct {
	RfrdDocInf    *ReferredDocumentInformation1    `xml:"RfrdDocInf"`
	RfrdDocRltdDt string                           `xml:"RfrdDocRltdDt"`
	RfrdDocAmt    []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt"`
	CdtrRefInf    *CreditorReferenceInformation1   `xml:"CdtrRefInf"`
	Invcr         *PartyIdentification8            `xml:"Invcr"`
	Invcee        *PartyIdentification8            `xml:"Invcee"`
	AddtlRmtInf   string                           `xml:"AddtlRmtInf"`
}

// TaxDetails ...
type TaxDetails struct {
	CertId string   `xml:"CertId"`
	TaxTp  *TaxType `xml:"TaxTp"`
}

// TaxInformation2 ...
type TaxInformation2 struct {
	CdtrTaxId       string             `xml:"CdtrTaxId"`
	CdtrTaxTp       string             `xml:"CdtrTaxTp"`
	DbtrTaxId       string             `xml:"DbtrTaxId"`
	TaxRefNb        string             `xml:"TaxRefNb"`
	TtlTaxblBaseAmt *CurrencyAndAmount `xml:"TtlTaxblBaseAmt"`
	TtlTaxAmt       *CurrencyAndAmount `xml:"TtlTaxAmt"`
	TaxDt           string             `xml:"TaxDt"`
	TaxTpInf        []*TaxDetails      `xml:"TaxTpInf"`
}

// TaxType ...
type TaxType struct {
	CtgyDesc     string             `xml:"CtgyDesc"`
	Rate         float64            `xml:"Rate"`
	TaxblBaseAmt *CurrencyAndAmount `xml:"TaxblBaseAmt"`
	Amt          *CurrencyAndAmount `xml:"Amt"`
}

// UPICIdentifier ...
type UPICIdentifier string

// Pain00100102 ...
type Pain00100102 struct {
	XMLName xml.Name                          `xml:"pain.001.001.02"`
	GrpHdr  *GroupHeader1                     `xml:"GrpHdr"`
	PmtInf  []*PaymentInstructionInformation1 `xml:"PmtInf"`
}