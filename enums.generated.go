package enum

// AggregatedBook field enumeration values.
type AggregatedBook string

const (
	AggregatedBook_NO  AggregatedBook = "N"
	AggregatedBook_YES AggregatedBook = "Y"
)

// ApplVerID field enumeration values.
type ApplVerID string

const (
	ApplVerID_FIX27    ApplVerID = "0"
	ApplVerID_FIX30    ApplVerID = "1"
	ApplVerID_FIX40    ApplVerID = "2"
	ApplVerID_FIX41    ApplVerID = "3"
	ApplVerID_FIX42    ApplVerID = "4"
	ApplVerID_FIX43    ApplVerID = "5"
	ApplVerID_FIX44    ApplVerID = "6"
	ApplVerID_FIX50    ApplVerID = "7"
	ApplVerID_FIX50SP1 ApplVerID = "8"
	ApplVerID_FIX50SP2 ApplVerID = "9"
)

// BusinessRejectReason field enumeration values.
type BusinessRejectReason string

const (
	BusinessRejectReason_OTHER                                                BusinessRejectReason = "0"
	BusinessRejectReason_UNKNOWN_ID                                           BusinessRejectReason = "1"
	BusinessRejectReason_THROTTLED_MESSAGES_REJECTED_ON_REQUEST               BusinessRejectReason = "10"
	BusinessRejectReason_INVALID_PRICE_INCREMENT                              BusinessRejectReason = "18"
	BusinessRejectReason_UNKNOWN_SECURITY                                     BusinessRejectReason = "2"
	BusinessRejectReason_UNSUPPORTED_MESSAGE_TYPE                             BusinessRejectReason = "3"
	BusinessRejectReason_APPLICATION_NOT_AVAILABLE                            BusinessRejectReason = "4"
	BusinessRejectReason_CONDITIONALLY_REQUIRED_FIELD_MISSING                 BusinessRejectReason = "5"
	BusinessRejectReason_NOT_AUTHORIZED                                       BusinessRejectReason = "6"
	BusinessRejectReason_DELIVERTO_FIRM_NOT_AVAILABLE_AT_THIS_TIME            BusinessRejectReason = "7"
	BusinessRejectReason_THROTTLE_LIMIT_EXCEEDED                              BusinessRejectReason = "8"
	BusinessRejectReason_THROTTLE_LIMIT_EXCEEDED_SESSION_WILL_BE_DISCONNECTED BusinessRejectReason = "9"
)

// CxlRejReason field enumeration values.
type CxlRejReason string

const (
	CxlRejReason_TOO_LATE_TO_CANCEL                                        CxlRejReason = "0"
	CxlRejReason_UNKNOWN_ORDER                                             CxlRejReason = "1"
	CxlRejReason_INVALID_PRICE_INCREMENT                                   CxlRejReason = "18"
	CxlRejReason_BROKER                                                    CxlRejReason = "2"
	CxlRejReason_ORDER_ALREADY_IN_PENDING_CANCEL_OR_PENDING_REPLACE_STATUS CxlRejReason = "3"
	CxlRejReason_UNABLE_TO_PROCESS_ORDER_MASS_CANCEL_REQUEST               CxlRejReason = "4"
	CxlRejReason_ORIGORDMODTIME                                            CxlRejReason = "5"
	CxlRejReason_DUPLICATE_CLORDID                                         CxlRejReason = "6"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE                               CxlRejReason = "7"
	CxlRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                          CxlRejReason = "8"
	CxlRejReason_OTHER                                                     CxlRejReason = "99"
)

// CxlRejResponseTo field enumeration values.
type CxlRejResponseTo string

const (
	CxlRejResponseTo_ORDER_CANCEL_REQUEST         CxlRejResponseTo = "1"
	CxlRejResponseTo_ORDER_CANCEL_REPLACE_REQUEST CxlRejResponseTo = "2"
)

// EncryptMethod field enumeration values.
type EncryptMethod string

const (
	EncryptMethod_NONE_OTHER  EncryptMethod = "0"
	EncryptMethod_PKCS        EncryptMethod = "1"
	EncryptMethod_DES         EncryptMethod = "2"
	EncryptMethod_PKCS_DES    EncryptMethod = "3"
	EncryptMethod_PGP_DES     EncryptMethod = "4"
	EncryptMethod_PGP_DES_MD5 EncryptMethod = "5"
	EncryptMethod_PEM_DES_MD5 EncryptMethod = "6"
)

// ExecType field enumeration values.
type ExecType string

const (
	ExecType_NEW                                 ExecType = "0"
	ExecType_DONE_FOR_DAY                        ExecType = "3"
	ExecType_CANCELED                            ExecType = "4"
	ExecType_REPLACED                            ExecType = "5"
	ExecType_PENDING_CANCEL                      ExecType = "6"
	ExecType_STOPPED                             ExecType = "7"
	ExecType_REJECTED                            ExecType = "8"
	ExecType_SUSPENDED                           ExecType = "9"
	ExecType_PENDING_NEW                         ExecType = "A"
	ExecType_CALCULATED                          ExecType = "B"
	ExecType_EXPIRED                             ExecType = "C"
	ExecType_RESTATED                            ExecType = "D"
	ExecType_PENDING_REPLACE                     ExecType = "E"
	ExecType_TRADE                               ExecType = "F"
	ExecType_TRADE_CORRECT                       ExecType = "G"
	ExecType_TRADE_CANCEL                        ExecType = "H"
	ExecType_ORDER_STATUS                        ExecType = "I"
	ExecType_TRADE_IN_A_CLEARING_HOLD            ExecType = "J"
	ExecType_TRADE_HAS_BEEN_RELEASED_TO_CLEARING ExecType = "K"
	ExecType_TRIGGERED_OR_ACTIVATED_BY_SYSTEM    ExecType = "L"
	ExecType_LOCKED                              ExecType = "M"
	ExecType_RELEASED                            ExecType = "N"
)

// GapFillFlag field enumeration values.
type GapFillFlag string

const (
	GapFillFlag_NO  GapFillFlag = "N"
	GapFillFlag_YES GapFillFlag = "Y"
)

// LastRptRequested field enumeration values.
type LastRptRequested string

const (
	LastRptRequested_NO  LastRptRequested = "N"
	LastRptRequested_YES LastRptRequested = "Y"
)

// MDEntryType field enumeration values.
type MDEntryType string

const (
	MDEntryType_BID                                             MDEntryType = "0"
	MDEntryType_OFFER                                           MDEntryType = "1"
	MDEntryType_TRADE                                           MDEntryType = "2"
	MDEntryType_INDEX_VALUE                                     MDEntryType = "3"
	MDEntryType_OPENING_PRICE                                   MDEntryType = "4"
	MDEntryType_CLOSING_PRICE                                   MDEntryType = "5"
	MDEntryType_SETTLEMENT_PRICE                                MDEntryType = "6"
	MDEntryType_TRADING_SESSION_HIGH_PRICE                      MDEntryType = "7"
	MDEntryType_TRADING_SESSION_LOW_PRICE                       MDEntryType = "8"
	MDEntryType_VOLUME_WEIGHTED_AVERAGE_PRICE                   MDEntryType = "9"
	MDEntryType_IMBALANCE                                       MDEntryType = "A"
	MDEntryType_TRADE_VOLUME                                    MDEntryType = "B"
	MDEntryType_OPEN_INTEREST                                   MDEntryType = "C"
	MDEntryType_COMPOSITE_UNDERLYING_PRICE                      MDEntryType = "D"
	MDEntryType_SIMULATED_SELL_PRICE                            MDEntryType = "E"
	MDEntryType_SIMULATED_BUY_PRICE                             MDEntryType = "F"
	MDEntryType_MARGIN_RATE                                     MDEntryType = "G"
	MDEntryType_MID_PRICE                                       MDEntryType = "H"
	MDEntryType_EMPTY_BOOK                                      MDEntryType = "J"
	MDEntryType_SETTLE_HIGH_PRICE                               MDEntryType = "K"
	MDEntryType_SETTLE_LOW_PRICE                                MDEntryType = "L"
	MDEntryType_PRIOR_SETTLE_PRICE                              MDEntryType = "M"
	MDEntryType_SESSION_HIGH_BID                                MDEntryType = "N"
	MDEntryType_SESSION_LOW_OFFER                               MDEntryType = "O"
	MDEntryType_EARLY_PRICES                                    MDEntryType = "P"
	MDEntryType_AUCTION_CLEARING_PRICE                          MDEntryType = "Q"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS       MDEntryType = "R"
	MDEntryType_SWAP_VALUE_FACTOR                               MDEntryType = "S"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_LONG_POSITIONS  MDEntryType = "T"
	MDEntryType_DAILY_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS      MDEntryType = "U"
	MDEntryType_CUMULATIVE_VALUE_ADJUSTMENT_FOR_SHORT_POSITIONS MDEntryType = "V"
	MDEntryType_FIXING_PRICE                                    MDEntryType = "W"
	MDEntryType_CASH_RATE                                       MDEntryType = "X"
	MDEntryType_RECOVERY_RATE                                   MDEntryType = "Y"
	MDEntryType_RECOVERY_RATE_FOR_LONG_POSITIONS                MDEntryType = "Z"
	MDEntryType_RECOVERY_RATE_FOR_SHORT_POSITIONS               MDEntryType = "a"
	MDEntryType_MARKET_BID                                      MDEntryType = "b"
	MDEntryType_MARKET_OFFER                                    MDEntryType = "c"
	MDEntryType_SHORT_SALE_MINIMUM_PRICE                        MDEntryType = "d"
	MDEntryType_PREVIOUS_CLOSING_PRICE                          MDEntryType = "e"
	MDEntryType_THRESHOLD_LIMITS_AND_PRICE_BANDING              MDEntryType = "g"
	MDEntryType_DAILY_FINANCING_VALUE                           MDEntryType = "h"
	MDEntryType_ACCRUED_FINANCING_VALUE                         MDEntryType = "i"
	MDEntryType_TIME_WEIGHTED_AVERAGE_PRICE                     MDEntryType = "t"
)

// MDReqRejReason field enumeration values.
type MDReqRejReason string

const (
	MDReqRejReason_UNKNOWN_SYMBOL                      MDReqRejReason = "0"
	MDReqRejReason_DUPLICATE_MDREQID                   MDReqRejReason = "1"
	MDReqRejReason_INSUFFICIENT_BANDWIDTH              MDReqRejReason = "2"
	MDReqRejReason_INSUFFICIENT_PERMISSIONS            MDReqRejReason = "3"
	MDReqRejReason_UNSUPPORTED_SUBSCRIPTIONREQUESTTYPE MDReqRejReason = "4"
	MDReqRejReason_UNSUPPORTED_MARKETDEPTH             MDReqRejReason = "5"
	MDReqRejReason_UNSUPPORTED_MDUPDATETYPE            MDReqRejReason = "6"
	MDReqRejReason_UNSUPPORTED_AGGREGATEDBOOK          MDReqRejReason = "7"
	MDReqRejReason_UNSUPPORTED_MDENTRYTYPE             MDReqRejReason = "8"
	MDReqRejReason_UNSUPPORTED_TRADINGSESSIONID        MDReqRejReason = "9"
	MDReqRejReason_UNSUPPORTED_SCOPE                   MDReqRejReason = "A"
	MDReqRejReason_UNSUPPORTED_OPENCLOSESETTLEFLAG     MDReqRejReason = "B"
	MDReqRejReason_UNSUPPORTED_MDIMPLICITDELETE        MDReqRejReason = "C"
	MDReqRejReason_INSUFFICIENT_CREDIT                 MDReqRejReason = "D"
)

// MDUpdateAction field enumeration values.
type MDUpdateAction string

const (
	MDUpdateAction_NEW         MDUpdateAction = "0"
	MDUpdateAction_CHANGE      MDUpdateAction = "1"
	MDUpdateAction_DELETE      MDUpdateAction = "2"
	MDUpdateAction_DELETE_THRU MDUpdateAction = "3"
	MDUpdateAction_DELETE_FROM MDUpdateAction = "4"
	MDUpdateAction_OVERLAY     MDUpdateAction = "5"
)

// MDUpdateType field enumeration values.
type MDUpdateType string

const (
	MDUpdateType_FULL_REFRESH        MDUpdateType = "0"
	MDUpdateType_INCREMENTAL_REFRESH MDUpdateType = "1"
)

// MassCancelRejectReason field enumeration values.
type MassCancelRejectReason string

const (
	MassCancelRejectReason_MASS_CANCEL_NOT_SUPPORTED                        MassCancelRejectReason = "0"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY                      MassCancelRejectReason = "1"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               MassCancelRejectReason = "10"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY MassCancelRejectReason = "11"
	MassCancelRejectReason_INVALID_OR_UNKOWN_UNDERLYING_SECURITY            MassCancelRejectReason = "2"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_PRODUCT                       MassCancelRejectReason = "3"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_CFICODE                       MassCancelRejectReason = "4"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITYTYPE                  MassCancelRejectReason = "5"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_TRADING_SESSION               MassCancelRejectReason = "6"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_MARKET                        MassCancelRejectReason = "7"
	MassCancelRejectReason_INVALID_OR_UNKOWN_MARKET_SEGMENT                 MassCancelRejectReason = "8"
	MassCancelRejectReason_INVALID_OR_UNKNOWN_SECURITY_GROUP                MassCancelRejectReason = "9"
	MassCancelRejectReason_OTHER                                            MassCancelRejectReason = "99"
)

// MassCancelRequestType field enumeration values.
type MassCancelRequestType string

const (
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY             MassCancelRequestType = "1"
	MassCancelRequestType_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY MassCancelRequestType = "2"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_PRODUCT              MassCancelRequestType = "3"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_CFICODE              MassCancelRequestType = "4"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITYTYPE         MassCancelRequestType = "5"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_TRADING_SESSION      MassCancelRequestType = "6"
	MassCancelRequestType_CANCEL_ALL_ORDERS                        MassCancelRequestType = "7"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET               MassCancelRequestType = "8"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT       MassCancelRequestType = "9"
	MassCancelRequestType_CANCEL_ORDERS_FOR_A_SECURITY_GROUP       MassCancelRequestType = "A"
	MassCancelRequestType_CANCEL_FOR_SECURITY_ISSUER               MassCancelRequestType = "B"
	MassCancelRequestType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelRequestType = "C"
)

// MassCancelResponse field enumeration values.
type MassCancelResponse string

const (
	MassCancelResponse_CANCEL_REQUEST_REJECTED                         MassCancelResponse = "0"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY                    MassCancelResponse = "1"
	MassCancelResponse_CANCEL_ORDERS_FOR_AN_UNDERLYING_SECURITY        MassCancelResponse = "2"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_PRODUCT                     MassCancelResponse = "3"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_CFICODE                     MassCancelResponse = "4"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITYTYPE                MassCancelResponse = "5"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_TRADING_SESSION             MassCancelResponse = "6"
	MassCancelResponse_CANCEL_ALL_ORDERS                               MassCancelResponse = "7"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET                      MassCancelResponse = "8"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_MARKET_SEGMENT              MassCancelResponse = "9"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITY_GROUP              MassCancelResponse = "A"
	MassCancelResponse_CANCEL_ORDERS_FOR_A_SECURITIES_ISSUER           MassCancelResponse = "B"
	MassCancelResponse_CANCEL_ORDERS_FOR_ISSUER_OF_UNDERLYING_SECURITY MassCancelResponse = "C"
)

// MassStatusReqType field enumeration values.
type MassStatusReqType string

const (
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITY             MassStatusReqType = "1"
	MassStatusReqType_STATUS_FOR_ISSUER_OF_UNDERLYING_SECURITY     MassStatusReqType = "10"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_AN_UNDERLYING_SECURITY MassStatusReqType = "2"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PRODUCT              MassStatusReqType = "3"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_CFICODE              MassStatusReqType = "4"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_SECURITYTYPE         MassStatusReqType = "5"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_TRADING_SESSION      MassStatusReqType = "6"
	MassStatusReqType_STATUS_FOR_ALL_ORDERS                        MassStatusReqType = "7"
	MassStatusReqType_STATUS_FOR_ORDERS_FOR_A_PARTYID              MassStatusReqType = "8"
	MassStatusReqType_STATUS_FOR_SECURITY_ISSUER                   MassStatusReqType = "9"
)

// MessageEncoding field enumeration values.
type MessageEncoding string

const (
	MessageEncoding_EUC_JP      MessageEncoding = "EUC-JP"
	MessageEncoding_ISO_2022_JP MessageEncoding = "ISO-2022-JP"
	MessageEncoding_SHIFT_JIS   MessageEncoding = "SHIFT_JIS"
	MessageEncoding_UTF_8       MessageEncoding = "UTF-8"
)

// MsgDirection field enumeration values.
type MsgDirection string

const (
	MsgDirection_RECEIVE MsgDirection = "R"
	MsgDirection_SEND    MsgDirection = "S"
)

// MsgType field enumeration values.
type MsgType string

const (
	MsgType_HEARTBEAT                             MsgType = "0"
	MsgType_TEST_REQUEST                          MsgType = "1"
	MsgType_RESEND_REQUEST                        MsgType = "2"
	MsgType_REJECT                                MsgType = "3"
	MsgType_SEQUENCE_RESET                        MsgType = "4"
	MsgType_LOGOUT                                MsgType = "5"
	MsgType_INDICATION_OF_INTEREST                MsgType = "6"
	MsgType_ADVERTISEMENT                         MsgType = "7"
	MsgType_EXECUTION_REPORT                      MsgType = "8"
	MsgType_ORDER_CANCEL_REJECT                   MsgType = "9"
	MsgType_LOGON                                 MsgType = "A"
	MsgType_DERIVATIVE_SECURITY_LIST              MsgType = "AA"
	MsgType_NEW_ORDER_MULTILEG                    MsgType = "AB"
	MsgType_MULTILEG_ORDER_CANCEL_REPLACE         MsgType = "AC"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST          MsgType = "AD"
	MsgType_TRADE_CAPTURE_REPORT                  MsgType = "AE"
	MsgType_ORDER_MASS_STATUS_REQUEST             MsgType = "AF"
	MsgType_QUOTE_REQUEST_REJECT                  MsgType = "AG"
	MsgType_RFQ_REQUEST                           MsgType = "AH"
	MsgType_QUOTE_STATUS_REPORT                   MsgType = "AI"
	MsgType_QUOTE_RESPONSE                        MsgType = "AJ"
	MsgType_CONFIRMATION                          MsgType = "AK"
	MsgType_POSITION_MAINTENANCE_REQUEST          MsgType = "AL"
	MsgType_POSITION_MAINTENANCE_REPORT           MsgType = "AM"
	MsgType_REQUEST_FOR_POSITIONS                 MsgType = "AN"
	MsgType_REQUEST_FOR_POSITIONS_ACK             MsgType = "AO"
	MsgType_POSITION_REPORT                       MsgType = "AP"
	MsgType_TRADE_CAPTURE_REPORT_REQUEST_ACK      MsgType = "AQ"
	MsgType_TRADE_CAPTURE_REPORT_ACK              MsgType = "AR"
	MsgType_ALLOCATION_REPORT                     MsgType = "AS"
	MsgType_ALLOCATION_REPORT_ACK                 MsgType = "AT"
	MsgType_CONFIRMATION_ACK                      MsgType = "AU"
	MsgType_SETTLEMENT_INSTRUCTION_REQUEST        MsgType = "AV"
	MsgType_ASSIGNMENT_REPORT                     MsgType = "AW"
	MsgType_COLLATERAL_REQUEST                    MsgType = "AX"
	MsgType_COLLATERAL_ASSIGNMENT                 MsgType = "AY"
	MsgType_COLLATERAL_RESPONSE                   MsgType = "AZ"
	MsgType_NEWS                                  MsgType = "B"
	MsgType_COLLATERAL_REPORT                     MsgType = "BA"
	MsgType_COLLATERAL_INQUIRY                    MsgType = "BB"
	MsgType_NETWORK_STATUS_REQUEST                MsgType = "BC"
	MsgType_NETWORK_STATUS_RESPONSE               MsgType = "BD"
	MsgType_USER_REQUEST                          MsgType = "BE"
	MsgType_USER_RESPONSE                         MsgType = "BF"
	MsgType_COLLATERAL_INQUIRY_ACK                MsgType = "BG"
	MsgType_CONFIRMATION_REQUEST                  MsgType = "BH"
	MsgType_TRADING_SESSION_LIST_REQUEST          MsgType = "BI"
	MsgType_TRADING_SESSION_LIST                  MsgType = "BJ"
	MsgType_SECURITY_LIST_UPDATE_REPORT           MsgType = "BK"
	MsgType_ADJUSTED_POSITION_REPORT              MsgType = "BL"
	MsgType_ALLOCATION_INSTRUCTION_ALERT          MsgType = "BM"
	MsgType_EXECUTION_ACKNOWLEDGEMENT             MsgType = "BN"
	MsgType_CONTRARY_INTENTION_REPORT             MsgType = "BO"
	MsgType_SECURITY_DEFINITION_UPDATE_REPORT     MsgType = "BP"
	MsgType_SETTLEMENTOBLIGATIONREPORT            MsgType = "BQ"
	MsgType_DERIVATIVESECURITYLISTUPDATEREPORT    MsgType = "BR"
	MsgType_TRADINGSESSIONLISTUPDATEREPORT        MsgType = "BS"
	MsgType_MARKETDEFINITIONREQUEST               MsgType = "BT"
	MsgType_MARKETDEFINITION                      MsgType = "BU"
	MsgType_MARKETDEFINITIONUPDATEREPORT          MsgType = "BV"
	MsgType_APPLICATIONMESSAGEREQUEST             MsgType = "BW"
	MsgType_APPLICATIONMESSAGEREQUESTACK          MsgType = "BX"
	MsgType_APPLICATIONMESSAGEREPORT              MsgType = "BY"
	MsgType_ORDERMASSACTIONREPORT                 MsgType = "BZ"
	MsgType_EMAIL                                 MsgType = "C"
	MsgType_ORDERMASSACTIONREQUEST                MsgType = "CA"
	MsgType_USERNOTIFICATION                      MsgType = "CB"
	MsgType_STREAMASSIGNMENTREQUEST               MsgType = "CC"
	MsgType_STREAMASSIGNMENTREPORT                MsgType = "CD"
	MsgType_STREAMASSIGNMENTREPORTACK             MsgType = "CE"
	MsgType_PARTYDETAILSLISTREQUEST               MsgType = "CF"
	MsgType_PARTYDETAILSLISTREPORT                MsgType = "CG"
	MsgType_MARGINREQUIREMENTINQUIRY              MsgType = "CH"
	MsgType_MARGINREQUIREMENTINQUIRYACK           MsgType = "CI"
	MsgType_MARGINREQUIREMENTREPORT               MsgType = "CJ"
	MsgType_PARTYDETAILSLISTUPDATEREPORT          MsgType = "CK"
	MsgType_PARTYRISKLIMITSREQUEST                MsgType = "CL"
	MsgType_PARTYRISKLIMITSREPORT                 MsgType = "CM"
	MsgType_SECURITYMASSSTATUSREQUEST             MsgType = "CN"
	MsgType_SECURITYMASSSTATUS                    MsgType = "CO"
	MsgType_ACCOUNTSUMMARYREPORT                  MsgType = "CQ"
	MsgType_PARTYRISKLIMITSUPDATEREPORT           MsgType = "CR"
	MsgType_PARTYRISKLIMITSDEFINITIONREQUEST      MsgType = "CS"
	MsgType_PARTYRISKLIMITSDEFINITIONREQUESTACK   MsgType = "CT"
	MsgType_PARTYENTITLEMENTSREQUEST              MsgType = "CU"
	MsgType_PARTYENTITLEMENTSREPORT               MsgType = "CV"
	MsgType_QUOTEACK                              MsgType = "CW"
	MsgType_PARTYDETAILSDEFINITIONREQUEST         MsgType = "CX"
	MsgType_PARTYDETAILSDEFINITIONREQUESTACK      MsgType = "CY"
	MsgType_PARTYENTITLEMENTSUPDATEREPORT         MsgType = "CZ"
	MsgType_ORDER_SINGLE                          MsgType = "D"
	MsgType_PARTYENTITLEMENTSDEFINITIONREQUEST    MsgType = "DA"
	MsgType_PARTYENTITLEMENTSDEFINITIONREQUESTACK MsgType = "DB"
	MsgType_TRADEMATCHREPORT                      MsgType = "DC"
	MsgType_TRADEMATCHREPORTACK                   MsgType = "DD"
	MsgType_PARTYRISKLIMITSREPORTACK              MsgType = "DE"
	MsgType_PARTYRISKLIMITCHECKREQUEST            MsgType = "DF"
	MsgType_PARTYRISKLIMITCHECKREQUESTACK         MsgType = "DG"
	MsgType_PARTYACTIONREQUEST                    MsgType = "DH"
	MsgType_PARTYACTIONREPORT                     MsgType = "DI"
	MsgType_MASSORDER                             MsgType = "DJ"
	MsgType_MASSORDERACK                          MsgType = "DK"
	MsgType_POSITIONTRANSFERINSTRUCTION           MsgType = "DL"
	MsgType_POSITIONTRANSFERINSTRUCTIONACK        MsgType = "DM"
	MsgType_POSITIONTRANSFERREPORT                MsgType = "DN"
	MsgType_MARKETDATASTATISTICSREQUEST           MsgType = "DO"
	MsgType_MARKETDATASTATISTICSREPORT            MsgType = "DP"
	MsgType_COLLATERALREPORTACK                   MsgType = "DQ"
	MsgType_MARKETDATAREPORT                      MsgType = "DR"
	MsgType_CROSSREQUEST                          MsgType = "DS"
	MsgType_CROSSREQUESTACK                       MsgType = "DT"
	MsgType_ALLOCATIONINSTRUCTIONALERTREQUEST     MsgType = "DU"
	MsgType_ALLOCATIONINSTRUCTIONALERTREQUESTACK  MsgType = "DV"
	MsgType_TRADEAGGREGATIONREQUEST               MsgType = "DW"
	MsgType_TRADEAGGREGATIONREPORT                MsgType = "DX"
	MsgType_PAYMANAGEMENTREQUEST                  MsgType = "DY"
	MsgType_PAYMANAGEMENTREQUESTACK               MsgType = "DZ"
	MsgType_ORDER_LIST                            MsgType = "E"
	MsgType_PAYMANAGEMENTREPORT                   MsgType = "EA"
	MsgType_PAYMANAGEMENTREPORTACK                MsgType = "EB"
	MsgType_ORDER_CANCEL_REQUEST                  MsgType = "F"
	MsgType_ORDER_CANCEL_REPLACE_REQUEST          MsgType = "G"
	MsgType_ORDER_STATUS_REQUEST                  MsgType = "H"
	MsgType_ALLOCATION_INSTRUCTION                MsgType = "J"
	MsgType_LIST_CANCEL_REQUEST                   MsgType = "K"
	MsgType_LIST_EXECUTE                          MsgType = "L"
	MsgType_LIST_STATUS_REQUEST                   MsgType = "M"
	MsgType_LIST_STATUS                           MsgType = "N"
	MsgType_ALLOCATION_INSTRUCTION_ACK            MsgType = "P"
	MsgType_DONT_KNOW_TRADE                       MsgType = "Q"
	MsgType_QUOTE_REQUEST                         MsgType = "R"
	MsgType_QUOTE                                 MsgType = "S"
	MsgType_SETTLEMENT_INSTRUCTIONS               MsgType = "T"
	MsgType_MARKET_DATA_REQUEST                   MsgType = "V"
	MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH     MsgType = "W"
	MsgType_MARKET_DATA_INCREMENTAL_REFRESH       MsgType = "X"
	MsgType_MARKET_DATA_REQUEST_REJECT            MsgType = "Y"
	MsgType_QUOTE_CANCEL                          MsgType = "Z"
	MsgType_QUOTE_STATUS_REQUEST                  MsgType = "a"
	MsgType_MASS_QUOTE_ACKNOWLEDGEMENT            MsgType = "b"
	MsgType_SECURITY_DEFINITION_REQUEST           MsgType = "c"
	MsgType_SECURITY_DEFINITION                   MsgType = "d"
	MsgType_SECURITY_STATUS_REQUEST               MsgType = "e"
	MsgType_SECURITY_STATUS                       MsgType = "f"
	MsgType_TRADING_SESSION_STATUS_REQUEST        MsgType = "g"
	MsgType_TRADING_SESSION_STATUS                MsgType = "h"
	MsgType_MASS_QUOTE                            MsgType = "i"
	MsgType_BUSINESS_MESSAGE_REJECT               MsgType = "j"
	MsgType_BID_REQUEST                           MsgType = "k"
	MsgType_BID_RESPONSE                          MsgType = "l"
	MsgType_LIST_STRIKE_PRICE                     MsgType = "m"
	MsgType_XML_MESSAGE                           MsgType = "n"
	MsgType_REGISTRATION_INSTRUCTIONS             MsgType = "o"
	MsgType_REGISTRATION_INSTRUCTIONS_RESPONSE    MsgType = "p"
	MsgType_ORDER_MASS_CANCEL_REQUEST             MsgType = "q"
	MsgType_ORDER_MASS_CANCEL_REPORT              MsgType = "r"
	MsgType_NEW_ORDER_CROSS                       MsgType = "s"
	MsgType_CROSS_ORDER_CANCEL_REPLACE_REQUEST    MsgType = "t"
	MsgType_CROSS_ORDER_CANCEL_REQUEST            MsgType = "u"
	MsgType_SECURITY_TYPE_REQUEST                 MsgType = "v"
	MsgType_SECURITY_TYPES                        MsgType = "w"
	MsgType_SECURITY_LIST_REQUEST                 MsgType = "x"
	MsgType_SECURITY_LIST                         MsgType = "y"
	MsgType_DERIVATIVE_SECURITY_LIST_REQUEST      MsgType = "z"
)

// OpenCloseSettlFlag field enumeration values.
type OpenCloseSettlFlag string

const (
	OpenCloseSettlFlag_DAILY_OPEN                       OpenCloseSettlFlag = "0"
	OpenCloseSettlFlag_SESSION_OPEN                     OpenCloseSettlFlag = "1"
	OpenCloseSettlFlag_DELIVERY_SETTLEMENT_ENTRY        OpenCloseSettlFlag = "2"
	OpenCloseSettlFlag_EXPECTED_ENTRY                   OpenCloseSettlFlag = "3"
	OpenCloseSettlFlag_ENTRY_FROM_PREVIOUS_BUSINESS_DAY OpenCloseSettlFlag = "4"
	OpenCloseSettlFlag_THEORETICAL_PRICE_VALUE          OpenCloseSettlFlag = "5"
)

// OrdRejReason field enumeration values.
type OrdRejReason string

const (
	OrdRejReason_BROKER                                                     OrdRejReason = "0"
	OrdRejReason_UNKNOWN_SYMBOL                                             OrdRejReason = "1"
	OrdRejReason_INVALID_INVESTOR_ID                                        OrdRejReason = "10"
	OrdRejReason_UNSUPPORTED_ORDER_CHARACTERISTIC                           OrdRejReason = "11"
	OrdRejReason_SURVEILLANCE_OPTION                                        OrdRejReason = "12"
	OrdRejReason_INCORRECT_QUANTITY                                         OrdRejReason = "13"
	OrdRejReason_INCORRECT_ALLOCATED_QUANTITY                               OrdRejReason = "14"
	OrdRejReason_UNKNOWN_ACCOUNT                                            OrdRejReason = "15"
	OrdRejReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND                           OrdRejReason = "16"
	OrdRejReason_INVALID_PRICE_INCREMENT                                    OrdRejReason = "18"
	OrdRejReason_REFERENCE_PRICE_NOT_AVAILABLE                              OrdRejReason = "19"
	OrdRejReason_EXCHANGE_CLOSED                                            OrdRejReason = "2"
	OrdRejReason_NOTIONAL_VALUE_EXCEEDS_THRESHOLD                           OrdRejReason = "20"
	OrdRejReason_ALGORITHM_RISK_THRESHOLD_BREACHED                          OrdRejReason = "21"
	OrdRejReason_SHORT_SELL_NOT_PERMITTED                                   OrdRejReason = "22"
	OrdRejReason_SHORT_SELL_REJECTED_DUE_TO_SECURITY_PRE_BORROW_RESTRICTION OrdRejReason = "23"
	OrdRejReason_SHORT_SELL_REJECTED_DUE_TO_ACCOUNT_PRE_BORROW_RESTRICTION  OrdRejReason = "24"
	OrdRejReason_INSUFFICIENT_CREDIT_LIMIT                                  OrdRejReason = "25"
	OrdRejReason_EXCEEDED_CLIP_SIZE_LIMIT                                   OrdRejReason = "26"
	OrdRejReason_EXCEEDED_MAXIMUM_NOTIONAL_ORDER_AMOUNT                     OrdRejReason = "27"
	OrdRejReason_EXCEEDED_DV01_PV01_LIMIT                                   OrdRejReason = "28"
	OrdRejReason_EXCEEDED_CS01_LIMIT                                        OrdRejReason = "29"
	OrdRejReason_ORDER_EXCEEDS_LIMIT                                        OrdRejReason = "3"
	OrdRejReason_TOO_LATE_TO_ENTER                                          OrdRejReason = "4"
	OrdRejReason_UNKNOWN_ORDER                                              OrdRejReason = "5"
	OrdRejReason_DUPLICATE_ORDER                                            OrdRejReason = "6"
	OrdRejReason_DUPLICATE_OF_A_VERBALLY_COMMUNICATED_ORDER                 OrdRejReason = "7"
	OrdRejReason_STALE_ORDER                                                OrdRejReason = "8"
	OrdRejReason_TRADE_ALONG_REQUIRED                                       OrdRejReason = "9"
	OrdRejReason_OTHER                                                      OrdRejReason = "99"
)

// OrdStatus field enumeration values.
type OrdStatus string

const (
	OrdStatus_NEW                  OrdStatus = "0"
	OrdStatus_PARTIALLY_FILLED     OrdStatus = "1"
	OrdStatus_FILLED               OrdStatus = "2"
	OrdStatus_DONE_FOR_DAY         OrdStatus = "3"
	OrdStatus_CANCELED             OrdStatus = "4"
	OrdStatus_REPLACED             OrdStatus = "5"
	OrdStatus_PENDING_CANCEL       OrdStatus = "6"
	OrdStatus_STOPPED              OrdStatus = "7"
	OrdStatus_REJECTED             OrdStatus = "8"
	OrdStatus_SUSPENDED            OrdStatus = "9"
	OrdStatus_PENDING_NEW          OrdStatus = "A"
	OrdStatus_CALCULATED           OrdStatus = "B"
	OrdStatus_EXPIRED              OrdStatus = "C"
	OrdStatus_ACCEPTED_FOR_BIDDING OrdStatus = "D"
	OrdStatus_PENDING_REPLACE      OrdStatus = "E"
)

// OrdType field enumeration values.
type OrdType string

const (
	OrdType_MARKET                         OrdType = "1"
	OrdType_LIMIT                          OrdType = "2"
	OrdType_STOP_STOP_LOSS                 OrdType = "3"
	OrdType_STOP_LIMIT                     OrdType = "4"
	OrdType_MARKET_ON_CLOSE                OrdType = "5"
	OrdType_WITH_OR_WITHOUT                OrdType = "6"
	OrdType_LIMIT_OR_BETTER                OrdType = "7"
	OrdType_LIMIT_WITH_OR_WITHOUT          OrdType = "8"
	OrdType_ON_BASIS                       OrdType = "9"
	OrdType_ON_CLOSE                       OrdType = "A"
	OrdType_LIMIT_ON_CLOSE                 OrdType = "B"
	OrdType_FOREX_MARKET                   OrdType = "C"
	OrdType_PREVIOUSLY_QUOTED              OrdType = "D"
	OrdType_PREVIOUSLY_INDICATED           OrdType = "E"
	OrdType_FOREX_LIMIT                    OrdType = "F"
	OrdType_FOREX_SWAP                     OrdType = "G"
	OrdType_FOREX_PREVIOUSLY_QUOTED        OrdType = "H"
	OrdType_FUNARI                         OrdType = "I"
	OrdType_MARKET_IF_TOUCHED              OrdType = "J"
	OrdType_MARKET_WITH_LEFT_OVER_AS_LIMIT OrdType = "K"
	OrdType_PREVIOUS_FUND_VALUATION_POINT  OrdType = "L"
	OrdType_NEXT_FUND_VALUATION_POINT      OrdType = "M"
	OrdType_PEGGED                         OrdType = "P"
	OrdType_COUNTER_ORDER_SELECTION        OrdType = "Q"
	OrdType_STOP_ON_BID_OR_OFFER           OrdType = "R"
	OrdType_STOP_LIMIT_ON_BID_OR_OFFER     OrdType = "S"
)

// OrderOrigination field enumeration values.
type OrderOrigination string

const (
	OrderOrigination_ORDER_RECEIVED_FROM_A_CUSTOMER                                    OrderOrigination = "1"
	OrderOrigination_ORDER_RECEIVED_FROM_WITHIN_THE_FIRM                               OrderOrigination = "2"
	OrderOrigination_ORDER_RECEIVED_FROM_ANOTHER_BROKER_DEALER                         OrderOrigination = "3"
	OrderOrigination_ORDER_RECEIVED_FROM_A_CUSTOMER_OR_ORIGINATED_FROM_WITHIN_THE_FIRM OrderOrigination = "4"
	OrderOrigination_ORDER_RECEIVED_FROM_A_DIRECT_ACCESS_OR_SPONSORED_ACCESS_CUSTOMER  OrderOrigination = "5"
	OrderOrigination_ORDER_RECEIVED_FROM_A_FOREIGN_DEALER_EQUIVALENT                   OrderOrigination = "6"
	OrderOrigination_ORDER_RECEIVED_FROM_AN_EXECUTION_ONLY_SERVICE                     OrderOrigination = "7"
)

// PartyIDSource field enumeration values.
type PartyIDSource string

const (
	PartyIDSource_KOREAN_INVESTOR_ID                                                                                PartyIDSource = "1"
	PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID                                                  PartyIDSource = "2"
	PartyIDSource_TAIWANESE_TRADING_ACCT                                                                            PartyIDSource = "3"
	PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY                                                                      PartyIDSource = "4"
	PartyIDSource_CHINESE_INVESTOR_ID                                                                               PartyIDSource = "5"
	PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER                                                           PartyIDSource = "6"
	PartyIDSource_US_SOCIAL_SECURITY_NUMBER                                                                         PartyIDSource = "7"
	PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER                                                                      PartyIDSource = "8"
	PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER                                                                        PartyIDSource = "9"
	PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER                                                                        PartyIDSource = "A"
	PartyIDSource_BIC                                                                                               PartyIDSource = "B"
	PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER                                                  PartyIDSource = "C"
	PartyIDSource_PROPRIETARY                                                                                       PartyIDSource = "D"
	PartyIDSource_ISO_COUNTRY_CODE                                                                                  PartyIDSource = "E"
	PartyIDSource_SETTLEMENT_ENTITY_LOCATION                                                                        PartyIDSource = "F"
	PartyIDSource_MARKET_IDENTIFIER_CODE                                                                            PartyIDSource = "G"
	PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE                                                                       PartyIDSource = "H"
	PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT PartyIDSource = "I"
	PartyIDSource_TAX_ID                                                                                            PartyIDSource = "J"
	PartyIDSource_AUSTRALIAN_COMPANY_NUMBER                                                                         PartyIDSource = "K"
	PartyIDSource_AUSTRALIAN_REGISTERED_BODY_NUMBER                                                                 PartyIDSource = "L"
	PartyIDSource_CFTC_REPORTING_FIRM_IDENTIFIER                                                                    PartyIDSource = "M"
	PartyIDSource_LEGAL_ENTITY_IDENTIFIER                                                                           PartyIDSource = "N"
	PartyIDSource_INTERIM_IDENTIFIER                                                                                PartyIDSource = "O"
	PartyIDSource_SHORT_CODE_IDENTIFIER                                                                             PartyIDSource = "P"
	PartyIDSource_NATIONAL_ID_OF_NATURAL_PERSON                                                                     PartyIDSource = "Q"
	PartyIDSource_INDIA_PERMANENT_ACCOUNT_NUMBER                                                                    PartyIDSource = "R"
	PartyIDSource_FIRM_DESIGNATED_IDENTIFIER                                                                        PartyIDSource = "S"
	PartyIDSource_SPECIAL_SEGREGATED_ACCOUNT_ID                                                                     PartyIDSource = "T"
	PartyIDSource_MASTER_SPECIAL_SEGREGATED_ACCOUNT_ID                                                              PartyIDSource = "U"
)

// PartyRole field enumeration values.
type PartyRole string

const (
	PartyRole_EXECUTING_FIRM                                                        PartyRole = "1"
	PartyRole_SETTLEMENT_LOCATION                                                   PartyRole = "10"
	PartyRole_MARGIN_ACCOUNT                                                        PartyRole = "100"
	PartyRole_COLLATERAL_ASSET_ACCOUNT                                              PartyRole = "101"
	PartyRole_DATA_REPOSITORY                                                       PartyRole = "102"
	PartyRole_CALCULATION_AGENT                                                     PartyRole = "103"
	PartyRole_SENDER_OF_EXERCISE_NOTICE                                             PartyRole = "104"
	PartyRole_RECEIVER_OF_EXERCISE_NOTICE                                           PartyRole = "105"
	PartyRole_RATE_REFERENCE_BANK                                                   PartyRole = "106"
	PartyRole_CORRESPONDENT                                                         PartyRole = "107"
	PartyRole_BENEFICIARYS_BANK_OR_DEPOSITORY_INSTITUTION                           PartyRole = "109"
	PartyRole_ORDER_ORIGINATION_TRADER                                              PartyRole = "11"
	PartyRole_BORROWER                                                              PartyRole = "110"
	PartyRole_PRIMARY_OBLIGATOR                                                     PartyRole = "111"
	PartyRole_GUARANTOR                                                             PartyRole = "112"
	PartyRole_EXCLUDED_REFERENCE_ENTITY                                             PartyRole = "113"
	PartyRole_DETERMINING_PARTY                                                     PartyRole = "114"
	PartyRole_HEDGING_PARTY                                                         PartyRole = "115"
	PartyRole_REPORTING_ENTITY                                                      PartyRole = "116"
	PartyRole_SALES_PERSON                                                          PartyRole = "117"
	PartyRole_OPERATOR                                                              PartyRole = "118"
	PartyRole_CENTRAL_SECURITIES_DEPOSITORY_119                                     PartyRole = "119"
	PartyRole_EXECUTING_TRADER                                                      PartyRole = "12"
	PartyRole_INTERNATIONAL_CENTRAL_SECURITIES_DEPOSITORY                           PartyRole = "120"
	PartyRole_TRADING_SUB_ACCOUNT                                                   PartyRole = "121"
	PartyRole_INVESTMENT_DECISION_MAKER                                             PartyRole = "122"
	PartyRole_PUBLISHING_INTERMEDIARY                                               PartyRole = "123"
	PartyRole_CENTRAL_SECURITIES_DEPOSITORY_124                                     PartyRole = "124"
	PartyRole_ISSUER                                                                PartyRole = "125"
	PartyRole_CONTRA_CUSTOMER_ACCOUNT                                               PartyRole = "126"
	PartyRole_CONTRA_INVESTMENT_DECISION_MAKER                                      PartyRole = "127"
	PartyRole_ORDER_ORIGINATION_FIRM                                                PartyRole = "13"
	PartyRole_GIVEUP_CLEARING_FIRM                                                  PartyRole = "14"
	PartyRole_CORRESPONDANT_CLEARING_FIRM                                           PartyRole = "15"
	PartyRole_EXECUTING_SYSTEM                                                      PartyRole = "16"
	PartyRole_CONTRA_FIRM                                                           PartyRole = "17"
	PartyRole_CONTRA_CLEARING_FIRM                                                  PartyRole = "18"
	PartyRole_SPONSORING_FIRM                                                       PartyRole = "19"
	PartyRole_BROKER_OF_CREDIT                                                      PartyRole = "2"
	PartyRole_UNDERLYING_CONTRA_FIRM                                                PartyRole = "20"
	PartyRole_CLEARING_ORGANIZATION                                                 PartyRole = "21"
	PartyRole_EXCHANGE                                                              PartyRole = "22"
	PartyRole_CUSTOMER_ACCOUNT                                                      PartyRole = "24"
	PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION                                   PartyRole = "25"
	PartyRole_CORRESPONDENT_BROKER                                                  PartyRole = "26"
	PartyRole_BUYER_SELLER                                                          PartyRole = "27"
	PartyRole_CUSTODIAN                                                             PartyRole = "28"
	PartyRole_INTERMEDIARY                                                          PartyRole = "29"
	PartyRole_CLIENT_ID                                                             PartyRole = "3"
	PartyRole_AGENT                                                                 PartyRole = "30"
	PartyRole_SUB_CUSTODIAN                                                         PartyRole = "31"
	PartyRole_BENEFICIARY                                                           PartyRole = "32"
	PartyRole_INTERESTED_PARTY                                                      PartyRole = "33"
	PartyRole_REGULATORY_BODY                                                       PartyRole = "34"
	PartyRole_LIQUIDITY_PROVIDER                                                    PartyRole = "35"
	PartyRole_ENTERING_TRADER                                                       PartyRole = "36"
	PartyRole_CONTRA_TRADER                                                         PartyRole = "37"
	PartyRole_POSITION_ACCOUNT                                                      PartyRole = "38"
	PartyRole_CONTRA_INVESTOR_ID                                                    PartyRole = "39"
	PartyRole_CLEARING_FIRM                                                         PartyRole = "4"
	PartyRole_TRANSFER_TO_FIRM                                                      PartyRole = "40"
	PartyRole_CONTRA_POSITION_ACCOUNT                                               PartyRole = "41"
	PartyRole_CONTRA_EXCHANGE                                                       PartyRole = "42"
	PartyRole_INTERNAL_CARRY_ACCOUNT                                                PartyRole = "43"
	PartyRole_ORDER_ENTRY_OPERATOR_ID                                               PartyRole = "44"
	PartyRole_SECONDARY_ACCOUNT_NUMBER                                              PartyRole = "45"
	PartyRole_FOREIGN_FIRM                                                          PartyRole = "46"
	PartyRole_THIRD_PARTY_ALLOCATION_FIRM                                           PartyRole = "47"
	PartyRole_CLAIMING_ACCOUNT                                                      PartyRole = "48"
	PartyRole_ASSET_MANAGER                                                         PartyRole = "49"
	PartyRole_INVESTOR_ID                                                           PartyRole = "5"
	PartyRole_PLEDGOR_ACCOUNT                                                       PartyRole = "50"
	PartyRole_PLEDGEE_ACCOUNT                                                       PartyRole = "51"
	PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT                                       PartyRole = "52"
	PartyRole_TRADER_MNEMONIC                                                       PartyRole = "53"
	PartyRole_SENDER_LOCATION                                                       PartyRole = "54"
	PartyRole_SESSION_ID                                                            PartyRole = "55"
	PartyRole_ACCEPTABLE_COUNTERPARTY                                               PartyRole = "56"
	PartyRole_UNACCEPTABLE_COUNTERPARTY                                             PartyRole = "57"
	PartyRole_ENTERING_UNIT                                                         PartyRole = "58"
	PartyRole_EXECUTING_UNIT                                                        PartyRole = "59"
	PartyRole_INTRODUCING_FIRM                                                      PartyRole = "6"
	PartyRole_INTRODUCING_BROKER                                                    PartyRole = "60"
	PartyRole_QUOTE_ORIGINATOR                                                      PartyRole = "61"
	PartyRole_REPORT_ORIGINATOR                                                     PartyRole = "62"
	PartyRole_SYSTEMATIC_INTERNALISER                                               PartyRole = "63"
	PartyRole_MULTILATERAL_TRADING_FACILITY                                         PartyRole = "64"
	PartyRole_REGULATED_MARKET                                                      PartyRole = "65"
	PartyRole_MARKET_MAKER                                                          PartyRole = "66"
	PartyRole_INVESTMENT_FIRM                                                       PartyRole = "67"
	PartyRole_HOST_COMPETENT_AUTHORITY                                              PartyRole = "68"
	PartyRole_HOME_COMPETENT_AUTHORITY                                              PartyRole = "69"
	PartyRole_ENTERING_FIRM                                                         PartyRole = "7"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY PartyRole = "70"
	PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION                                PartyRole = "71"
	PartyRole_REPORTING_INTERMEDIARY                                                PartyRole = "72"
	PartyRole_EXECUTION_VENUE                                                       PartyRole = "73"
	PartyRole_MARKET_DATA_ENTRY_ORIGINATOR                                          PartyRole = "74"
	PartyRole_LOCATION_ID                                                           PartyRole = "75"
	PartyRole_DESK_ID                                                               PartyRole = "76"
	PartyRole_MARKET_DATA_MARKET                                                    PartyRole = "77"
	PartyRole_ALLOCATION_ENTITY                                                     PartyRole = "78"
	PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES                         PartyRole = "79"
	PartyRole_LOCATE                                                                PartyRole = "8"
	PartyRole_STEP_OUT_FIRM                                                         PartyRole = "80"
	PartyRole_BROKER_CIENT_ID                                                       PartyRole = "81"
	PartyRole_CENTRAL_REGISTRATION_DEPOSITORY                                       PartyRole = "82"
	PartyRole_CLEARING_ACCOUNT                                                      PartyRole = "83"
	PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY                                      PartyRole = "84"
	PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY                                    PartyRole = "85"
	PartyRole_CLS_MEMBER_BANK                                                       PartyRole = "86"
	PartyRole_IN_CONCERT_GROUP                                                      PartyRole = "87"
	PartyRole_IN_CONCERT_CONTROLLING_ENTITY                                         PartyRole = "88"
	PartyRole_LARGE_POSITIONS_REPORTING_ACCOUNT                                     PartyRole = "89"
	PartyRole_FUND_MANAGER_CLIENT_ID                                                PartyRole = "9"
	PartyRole_SETTLEMENT_FIRM                                                       PartyRole = "90"
	PartyRole_SETTLEMENT_ACCOUNT                                                    PartyRole = "91"
	PartyRole_REPORTING_MARKET_CENTER                                               PartyRole = "92"
	PartyRole_RELATED_REPORTING_MARKET_CENTER                                       PartyRole = "93"
	PartyRole_AWAY_MARKET                                                           PartyRole = "94"
	PartyRole_GIVE_UP                                                               PartyRole = "95"
	PartyRole_TAKE_UP                                                               PartyRole = "96"
	PartyRole_GIVE_UP_CLEARING_FIRM                                                 PartyRole = "97"
	PartyRole_TAKE_UP_CLEARING_FIRM                                                 PartyRole = "98"
	PartyRole_ORIGINATING_MARKET                                                    PartyRole = "99"
)

// PartyRoleQualifier field enumeration values.
type PartyRoleQualifier string

const (
	PartyRoleQualifier_AGENCY                                    PartyRoleQualifier = "0"
	PartyRoleQualifier_PRINCIPAL                                 PartyRoleQualifier = "1"
	PartyRoleQualifier_ORIGINAL_TRADE_REPOSITORY                 PartyRoleQualifier = "10"
	PartyRoleQualifier_ADDITIONAL_INTERNATIONAL_TRADE_REPOSITORY PartyRoleQualifier = "11"
	PartyRoleQualifier_ADDITIONAL_DOMESTIC_TRADE_REPOSITORY      PartyRoleQualifier = "12"
	PartyRoleQualifier_RELATED_EXCHANGE                          PartyRoleQualifier = "13"
	PartyRoleQualifier_OPTIONS_EXCHANGE                          PartyRoleQualifier = "14"
	PartyRoleQualifier_SPECIFIED_EXCHANGE                        PartyRoleQualifier = "15"
	PartyRoleQualifier_CONSTITUENT_EXCHANGE                      PartyRoleQualifier = "16"
	PartyRoleQualifier_EXEMPT_FROM_TRADE_REPORTING               PartyRoleQualifier = "17"
	PartyRoleQualifier_CURRENT                                   PartyRoleQualifier = "18"
	PartyRoleQualifier_NEW                                       PartyRoleQualifier = "19"
	PartyRoleQualifier_RISKLESS_PRINCIPAL                        PartyRoleQualifier = "2"
	PartyRoleQualifier_DESIGNATED_SPONSOR                        PartyRoleQualifier = "20"
	PartyRoleQualifier_SPECIALIST                                PartyRoleQualifier = "21"
	PartyRoleQualifier_ALGORITHM                                 PartyRoleQualifier = "22"
	PartyRoleQualifier_FIRM_OR_LEGAL_ENTITY                      PartyRoleQualifier = "23"
	PartyRoleQualifier_NATURAL_PERSON                            PartyRoleQualifier = "24"
	PartyRoleQualifier_REGULAR_TRADER                            PartyRoleQualifier = "25"
	PartyRoleQualifier_HEAD_TRADER                               PartyRoleQualifier = "26"
	PartyRoleQualifier_SUPERVISOR                                PartyRoleQualifier = "27"
	PartyRoleQualifier_TRI_PARTY                                 PartyRoleQualifier = "28"
	PartyRoleQualifier_LENDER                                    PartyRoleQualifier = "29"
	PartyRoleQualifier_GENERAL_CLEARING_MEMBER                   PartyRoleQualifier = "3"
	PartyRoleQualifier_INDIVIDUAL_CLEARING_MEMBER                PartyRoleQualifier = "4"
	PartyRoleQualifier_PREFERRED_MARKET_MAKER                    PartyRoleQualifier = "5"
	PartyRoleQualifier_DIRECTED_MARKET_MAKER                     PartyRoleQualifier = "6"
	PartyRoleQualifier_BANK                                      PartyRoleQualifier = "7"
	PartyRoleQualifier_HUB                                       PartyRoleQualifier = "8"
	PartyRoleQualifier_PRIMARY_TRADE_REPOSITORY                  PartyRoleQualifier = "9"
)

// PartySubIDType field enumeration values.
type PartySubIDType string

const (
	PartySubIDType_FIRM                         PartySubIDType = "1"
	PartySubIDType_SECURITIES_ACCOUNT_NUMBER    PartySubIDType = "10"
	PartySubIDType_REGISTRATION_NUMBER          PartySubIDType = "11"
	PartySubIDType_REGISTERED_ADDRESS_12        PartySubIDType = "12"
	PartySubIDType_REGULATORY_STATUS            PartySubIDType = "13"
	PartySubIDType_REGISTRATION_NAME            PartySubIDType = "14"
	PartySubIDType_CASH_ACCOUNT_NUMBER          PartySubIDType = "15"
	PartySubIDType_BIC                          PartySubIDType = "16"
	PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE  PartySubIDType = "17"
	PartySubIDType_REGISTERED_ADDRESS_18        PartySubIDType = "18"
	PartySubIDType_FUND_ACCOUNT_NAME            PartySubIDType = "19"
	PartySubIDType_PERSON                       PartySubIDType = "2"
	PartySubIDType_TELEX_NUMBER                 PartySubIDType = "20"
	PartySubIDType_FAX_NUMBER                   PartySubIDType = "21"
	PartySubIDType_SECURITIES_ACCOUNT_NAME      PartySubIDType = "22"
	PartySubIDType_CASH_ACCOUNT_NAME            PartySubIDType = "23"
	PartySubIDType_DEPARTMENT                   PartySubIDType = "24"
	PartySubIDType_LOCATION_DESK                PartySubIDType = "25"
	PartySubIDType_POSITION_ACCOUNT_TYPE        PartySubIDType = "26"
	PartySubIDType_SECURITY_LOCATE_ID           PartySubIDType = "27"
	PartySubIDType_MARKET_MAKER                 PartySubIDType = "28"
	PartySubIDType_ELIGIBLE_COUNTERPARTY        PartySubIDType = "29"
	PartySubIDType_SYSTEM                       PartySubIDType = "3"
	PartySubIDType_PROFESSIONAL_CLIENT          PartySubIDType = "30"
	PartySubIDType_LOCATION                     PartySubIDType = "31"
	PartySubIDType_EXECUTION_VENUE              PartySubIDType = "32"
	PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER PartySubIDType = "33"
	PartySubIDType_APPLICATION                  PartySubIDType = "4"
	PartySubIDType_FULL_LEGAL_NAME_OF_FIRM      PartySubIDType = "5"
	PartySubIDType_POSTAL_ADDRESS               PartySubIDType = "6"
	PartySubIDType_PHONE_NUMBER                 PartySubIDType = "7"
	PartySubIDType_EMAIL_ADDRESS                PartySubIDType = "8"
	PartySubIDType_CONTACT_NAME                 PartySubIDType = "9"
)

// PossDupFlag field enumeration values.
type PossDupFlag string

const (
	PossDupFlag_NO  PossDupFlag = "N"
	PossDupFlag_YES PossDupFlag = "Y"
)

// PossResend field enumeration values.
type PossResend string

const (
	PossResend_NO  PossResend = "N"
	PossResend_YES PossResend = "Y"
)

// QuoteCancelType field enumeration values.
type QuoteCancelType string

const (
	QuoteCancelType_CANCEL_FOR_ONE_OR_MORE_SECURITIES        QuoteCancelType = "1"
	QuoteCancelType_CANCEL_FOR_SECURITY_TYPE                 QuoteCancelType = "2"
	QuoteCancelType_CANCEL_FOR_UNDERLYING_SECURITY           QuoteCancelType = "3"
	QuoteCancelType_CANCEL_ALL_QUOTES                        QuoteCancelType = "4"
	QuoteCancelType_CANCEL_SPECIFIED_SINGLE_QUOTE            QuoteCancelType = "5"
	QuoteCancelType_CANCEL_BY_TYPE_OF_QUOTE                  QuoteCancelType = "6"
	QuoteCancelType_CANCEL_FOR_SECURITY_ISSUER               QuoteCancelType = "7"
	QuoteCancelType_CANCEL_FOR_ISSUER_OF_UNDERLYING_SECURITY QuoteCancelType = "8"
)

// QuoteRejectReason field enumeration values.
type QuoteRejectReason string

const (
	QuoteRejectReason_UNKNOWN_SYMBOL                                   QuoteRejectReason = "1"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND_10              QuoteRejectReason = "10"
	QuoteRejectReason_QUOTE_LOCKED                                     QuoteRejectReason = "11"
	QuoteRejectReason_INVALID_OR_UNKNOWN_SECURITY_ISSUER               QuoteRejectReason = "12"
	QuoteRejectReason_INVALID_OR_UNKNOWN_ISSUER_OF_UNDERLYING_SECURITY QuoteRejectReason = "13"
	QuoteRejectReason_NOTIONAL_VALUE_EXCEEDS_THRESHOLD                 QuoteRejectReason = "14"
	QuoteRejectReason_PRICE_EXCEEDS_CURRENT_PRICE_BAND_15              QuoteRejectReason = "15"
	QuoteRejectReason_REFERENCE_PRICE_NOT_AVAILABLE                    QuoteRejectReason = "16"
	QuoteRejectReason_INSUFFICIENT_CREDIT_LIMIT                        QuoteRejectReason = "17"
	QuoteRejectReason_EXCEEDED_CLIP_SIZE_LIMIT                         QuoteRejectReason = "18"
	QuoteRejectReason_EXCEEDED_MAXIMUM_NOTIONAL_ORDER_AMOUNT           QuoteRejectReason = "19"
	QuoteRejectReason_EXCHANGE                                         QuoteRejectReason = "2"
	QuoteRejectReason_EXCEEDED_DV01_PV01_LIMIT                         QuoteRejectReason = "20"
	QuoteRejectReason_EXCEEDED_CS01_LIMIT                              QuoteRejectReason = "21"
	QuoteRejectReason_QUOTE_REQUEST_EXCEEDS_LIMIT                      QuoteRejectReason = "3"
	QuoteRejectReason_TOO_LATE_TO_ENTER                                QuoteRejectReason = "4"
	QuoteRejectReason_UNKNOWN_QUOTE                                    QuoteRejectReason = "5"
	QuoteRejectReason_DUPLICATE_QUOTE                                  QuoteRejectReason = "6"
	QuoteRejectReason_INVALID_BID_ASK_SPREAD                           QuoteRejectReason = "7"
	QuoteRejectReason_INVALID_PRICE                                    QuoteRejectReason = "8"
	QuoteRejectReason_NOT_AUTHORIZED_TO_QUOTE_SECURITY                 QuoteRejectReason = "9"
	QuoteRejectReason_OTHER                                            QuoteRejectReason = "99"
)

// QuoteStatus field enumeration values.
type QuoteStatus string

const (
	QuoteStatus_ACCEPTED                            QuoteStatus = "0"
	QuoteStatus_CANCELED_FOR_SPECIFIC_SECURITIES    QuoteStatus = "1"
	QuoteStatus_PENDING                             QuoteStatus = "10"
	QuoteStatus_PASS                                QuoteStatus = "11"
	QuoteStatus_LOCKED_MARKET_WARNING               QuoteStatus = "12"
	QuoteStatus_CROSSED_MARKET_WARNING              QuoteStatus = "13"
	QuoteStatus_CANCELED_DUE_TO_LOCKED_MARKET       QuoteStatus = "14"
	QuoteStatus_CANCELED_DUE_TO_CROSSED_MARKET      QuoteStatus = "15"
	QuoteStatus_ACTIVE                              QuoteStatus = "16"
	QuoteStatus_CANCELED                            QuoteStatus = "17"
	QuoteStatus_UNSOLICITED_QUOTE_REPLENISHMENT     QuoteStatus = "18"
	QuoteStatus_PENDING_END_TRADE                   QuoteStatus = "19"
	QuoteStatus_CANCELED_FOR_SPECIFIC_SECURITYTYPES QuoteStatus = "2"
	QuoteStatus_TOO_LATE_TO_END                     QuoteStatus = "20"
	QuoteStatus_TRADED                              QuoteStatus = "21"
	QuoteStatus_TRADED_AND_REMOVED                  QuoteStatus = "22"
	QuoteStatus_CONTRACT_TERMINATED                 QuoteStatus = "23"
	QuoteStatus_CANCELED_FOR_UNDERLYING             QuoteStatus = "3"
	QuoteStatus_CANCELED_ALL                        QuoteStatus = "4"
	QuoteStatus_REJECTED                            QuoteStatus = "5"
	QuoteStatus_REMOVED_FROM_MARKET                 QuoteStatus = "6"
	QuoteStatus_EXPIRED                             QuoteStatus = "7"
	QuoteStatus_QUERY                               QuoteStatus = "8"
	QuoteStatus_QUOTE_NOT_FOUND                     QuoteStatus = "9"
)

// ResetSeqNumFlag field enumeration values.
type ResetSeqNumFlag string

const (
	ResetSeqNumFlag_NO  ResetSeqNumFlag = "N"
	ResetSeqNumFlag_YES ResetSeqNumFlag = "Y"
)

// SecurityIDSource field enumeration values.
type SecurityIDSource string

const (
	SecurityIDSource_CUSIP                                  SecurityIDSource = "1"
	SecurityIDSource_SEDOL                                  SecurityIDSource = "2"
	SecurityIDSource_QUIK                                   SecurityIDSource = "3"
	SecurityIDSource_ISIN                                   SecurityIDSource = "4"
	SecurityIDSource_RIC                                    SecurityIDSource = "5"
	SecurityIDSource_ISO_CURRENCY_CODE                      SecurityIDSource = "6"
	SecurityIDSource_ISO_COUNTRY_CODE                       SecurityIDSource = "7"
	SecurityIDSource_EXCHANGE_SYMBOL                        SecurityIDSource = "8"
	SecurityIDSource_CONSOLIDATED_TAPE_ASSOCIATION          SecurityIDSource = "9"
	SecurityIDSource_BLOOMBERG_SYMBOL                       SecurityIDSource = "A"
	SecurityIDSource_WERTPAPIER                             SecurityIDSource = "B"
	SecurityIDSource_DUTCH                                  SecurityIDSource = "C"
	SecurityIDSource_VALOREN                                SecurityIDSource = "D"
	SecurityIDSource_SICOVAM                                SecurityIDSource = "E"
	SecurityIDSource_BELGIAN                                SecurityIDSource = "F"
	SecurityIDSource_COMMON                                 SecurityIDSource = "G"
	SecurityIDSource_CLEARING_HOUSE                         SecurityIDSource = "H"
	SecurityIDSource_ISDA_FPML_PRODUCT_SPECIFICATION        SecurityIDSource = "I"
	SecurityIDSource_OPTION_PRICE_REPORTING_AUTHORITY       SecurityIDSource = "J"
	SecurityIDSource_ISDA_FPML_PRODUCT_URL                  SecurityIDSource = "K"
	SecurityIDSource_LETTER_OF_CREDIT                       SecurityIDSource = "L"
	SecurityIDSource_MARKETPLACE_ASSIGNED_IDENTIFIER        SecurityIDSource = "M"
	SecurityIDSource_MARKIT_RED_ENTITY_CLIP                 SecurityIDSource = "N"
	SecurityIDSource_MARKIT_RED_PAIR_CLIP                   SecurityIDSource = "P"
	SecurityIDSource_CFTC_COMMODITY_CODE                    SecurityIDSource = "Q"
	SecurityIDSource_ISDA_COMMODITY_REFERENCE_PRICE         SecurityIDSource = "R"
	SecurityIDSource_FINANCIAL_INSTRUMENT_GLOBAL_IDENTIFIER SecurityIDSource = "S"
	SecurityIDSource_LEGAL_ENTITY_IDENTIFIER                SecurityIDSource = "T"
	SecurityIDSource_SYNTHETIC                              SecurityIDSource = "U"
	SecurityIDSource_FIDESSA_INSTRUMENT_MNEMONIC            SecurityIDSource = "V"
	SecurityIDSource_INDEX_NAME                             SecurityIDSource = "W"
	SecurityIDSource_UNIFORM_SYMBOL                         SecurityIDSource = "X"
)

// SecurityListRequestType field enumeration values.
type SecurityListRequestType string

const (
	SecurityListRequestType_SYMBOL                                    SecurityListRequestType = "0"
	SecurityListRequestType_SECURITYTYPE_AND_OR_CFICODE               SecurityListRequestType = "1"
	SecurityListRequestType_PRODUCT                                   SecurityListRequestType = "2"
	SecurityListRequestType_TRADINGSESSIONID                          SecurityListRequestType = "3"
	SecurityListRequestType_ALL_SECURITIES                            SecurityListRequestType = "4"
	SecurityListRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID SecurityListRequestType = "5"
)

// SecurityRequestResult field enumeration values.
type SecurityRequestResult string

const (
	SecurityRequestResult_VALID_REQUEST                                      SecurityRequestResult = "0"
	SecurityRequestResult_INVALID_OR_UNSUPPORTED_REQUEST                     SecurityRequestResult = "1"
	SecurityRequestResult_NO_INSTRUMENTS_FOUND_THAT_MATCH_SELECTION_CRITERIA SecurityRequestResult = "2"
	SecurityRequestResult_NOT_AUTHORIZED_TO_RETRIEVE_INSTRUMENT_DATA         SecurityRequestResult = "3"
	SecurityRequestResult_INSTRUMENT_DATA_TEMPORARILY_UNAVAILABLE            SecurityRequestResult = "4"
	SecurityRequestResult_REQUEST_FOR_INSTRUMENT_DATA_NOT_SUPPORTED          SecurityRequestResult = "5"
)

// SecurityRequestType field enumeration values.
type SecurityRequestType string

const (
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_AND_SPECIFICATIONS              SecurityRequestType = "0"
	SecurityRequestType_REQUEST_SECURITY_IDENTITY_FOR_THE_SPECIFICATIONS_PROVIDED SecurityRequestType = "1"
	SecurityRequestType_REQUEST_LIST_SECURITY_TYPES                               SecurityRequestType = "2"
	SecurityRequestType_REQUEST_LIST_SECURITIES                                   SecurityRequestType = "3"
	SecurityRequestType_SYMBOL                                                    SecurityRequestType = "4"
	SecurityRequestType_SECURITYTYPE_AND_OR_CFICODE                               SecurityRequestType = "5"
	SecurityRequestType_PRODUCT                                                   SecurityRequestType = "6"
	SecurityRequestType_TRADINGSESSIONID                                          SecurityRequestType = "7"
	SecurityRequestType_ALL_SECURITIES                                            SecurityRequestType = "8"
	SecurityRequestType_MARKETID_OR_MARKETID_PLUS_MARKETSEGMENTID                 SecurityRequestType = "9"
)

// SecurityTradingStatus field enumeration values.
type SecurityTradingStatus string

const (
	SecurityTradingStatus_OPENING_DELAY                  SecurityTradingStatus = "1"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_SELL SecurityTradingStatus = "10"
	SecurityTradingStatus_11                             SecurityTradingStatus = "11"
	SecurityTradingStatus_NO_MARKET_IMBALANCE            SecurityTradingStatus = "12"
	SecurityTradingStatus_NO_MARKET_ON_CLOSE_IMBALANCE   SecurityTradingStatus = "13"
	SecurityTradingStatus_ITS_PRE_OPENING                SecurityTradingStatus = "14"
	SecurityTradingStatus_NEW_PRICE_INDICATION           SecurityTradingStatus = "15"
	SecurityTradingStatus_TRADE_DISSEMINATION_TIME       SecurityTradingStatus = "16"
	SecurityTradingStatus_READY_TO_TRADE                 SecurityTradingStatus = "17"
	SecurityTradingStatus_NOT_AVAILABLE_FOR_TRADING      SecurityTradingStatus = "18"
	SecurityTradingStatus_NOT_TRADED_ON_THIS_MARKET      SecurityTradingStatus = "19"
	SecurityTradingStatus_TRADING_HALT                   SecurityTradingStatus = "2"
	SecurityTradingStatus_UNKNOWN_OR_INVALID             SecurityTradingStatus = "20"
	SecurityTradingStatus_PRE_OPEN                       SecurityTradingStatus = "21"
	SecurityTradingStatus_OPENING_ROTATION               SecurityTradingStatus = "22"
	SecurityTradingStatus_FAST_MARKET                    SecurityTradingStatus = "23"
	SecurityTradingStatus_PRE_CROSS                      SecurityTradingStatus = "24"
	SecurityTradingStatus_CROSS                          SecurityTradingStatus = "25"
	SecurityTradingStatus_POST_CLOSE                     SecurityTradingStatus = "26"
	SecurityTradingStatus_RESUME                         SecurityTradingStatus = "3"
	SecurityTradingStatus_NO_OPEN                        SecurityTradingStatus = "4"
	SecurityTradingStatus_PRICE_INDICATION               SecurityTradingStatus = "5"
	SecurityTradingStatus_TRADING_RANGE_INDICATION       SecurityTradingStatus = "6"
	SecurityTradingStatus_MARKET_IMBALANCE_BUY           SecurityTradingStatus = "7"
	SecurityTradingStatus_MARKET_IMBALANCE_SELL          SecurityTradingStatus = "8"
	SecurityTradingStatus_MARKET_ON_CLOSE_IMBALANCE_BUY  SecurityTradingStatus = "9"
)

// SessionRejectReason field enumeration values.
type SessionRejectReason string

const (
	SessionRejectReason_INVALID_TAG_NUMBER                             SessionRejectReason = "0"
	SessionRejectReason_REQUIRED_TAG_MISSING                           SessionRejectReason = "1"
	SessionRejectReason_SENDINGTIME_ACCURACY_PROBLEM                   SessionRejectReason = "10"
	SessionRejectReason_INVALID_MSGTYPE                                SessionRejectReason = "11"
	SessionRejectReason_XML_VALIDATION_ERROR                           SessionRejectReason = "12"
	SessionRejectReason_TAG_APPEARS_MORE_THAN_ONCE                     SessionRejectReason = "13"
	SessionRejectReason_TAG_SPECIFIED_OUT_OF_REQUIRED_ORDER            SessionRejectReason = "14"
	SessionRejectReason_REPEATING_GROUP_FIELDS_OUT_OF_ORDER            SessionRejectReason = "15"
	SessionRejectReason_INCORRECT_NUMINGROUP_COUNT_FOR_REPEATING_GROUP SessionRejectReason = "16"
	SessionRejectReason_NON_DATA_VALUE_INCLUDES_FIELD_DELIMITER        SessionRejectReason = "17"
	SessionRejectReason_INVALID_UNSUPPORTED_APPLICATION_VERSION        SessionRejectReason = "18"
	SessionRejectReason_TAG_NOT_DEFINED_FOR_THIS_MESSAGE_TYPE          SessionRejectReason = "2"
	SessionRejectReason_UNDEFINED_TAG                                  SessionRejectReason = "3"
	SessionRejectReason_TAG_SPECIFIED_WITHOUT_A_VALUE                  SessionRejectReason = "4"
	SessionRejectReason_VALUE_IS_INCORRECT                             SessionRejectReason = "5"
	SessionRejectReason_INCORRECT_DATA_FORMAT_FOR_VALUE                SessionRejectReason = "6"
	SessionRejectReason_DECRYPTION_PROBLEM                             SessionRejectReason = "7"
	SessionRejectReason_SIGNATURE_PROBLEM                              SessionRejectReason = "8"
	SessionRejectReason_COMPID_PROBLEM                                 SessionRejectReason = "9"
	SessionRejectReason_OTHER                                          SessionRejectReason = "99"
)

// Side field enumeration values.
type Side string

const (
	Side_BUY                Side = "1"
	Side_SELL               Side = "2"
	Side_BUY_MINUS          Side = "3"
	Side_SELL_PLUS          Side = "4"
	Side_SELL_SHORT         Side = "5"
	Side_SELL_SHORT_EXEMPT  Side = "6"
	Side_UNDISCLOSED        Side = "7"
	Side_CROSS              Side = "8"
	Side_CROSS_SHORT        Side = "9"
	Side_CROSS_SHORT_EXEMPT Side = "A"
	Side_AS_DEFINED         Side = "B"
	Side_OPPOSITE           Side = "C"
	Side_SUBSCRIBE          Side = "D"
	Side_REDEEM             Side = "E"
	Side_LEND               Side = "F"
	Side_BORROW             Side = "G"
	Side_SELL_UNDISCLOSED   Side = "H"
)

// SubscriptionRequestType field enumeration values.
type SubscriptionRequestType string

const (
	SubscriptionRequestType_SNAPSHOT                                      SubscriptionRequestType = "0"
	SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES                         SubscriptionRequestType = "1"
	SubscriptionRequestType_DISABLE_PREVIOUS_SNAPSHOT_PLUS_UPDATE_REQUEST SubscriptionRequestType = "2"
)

// TimeInForce field enumeration values.
type TimeInForce string

const (
	TimeInForce_DAY                   TimeInForce = "0"
	TimeInForce_GOOD_TILL_CANCEL      TimeInForce = "1"
	TimeInForce_AT_THE_OPENING        TimeInForce = "2"
	TimeInForce_IMMEDIATE_OR_CANCEL   TimeInForce = "3"
	TimeInForce_FILL_OR_KILL          TimeInForce = "4"
	TimeInForce_GOOD_TILL_CROSSING    TimeInForce = "5"
	TimeInForce_GOOD_TILL_DATE        TimeInForce = "6"
	TimeInForce_AT_THE_CLOSE          TimeInForce = "7"
	TimeInForce_GOOD_THROUGH_CROSSING TimeInForce = "8"
	TimeInForce_AT_CROSSING           TimeInForce = "9"
	TimeInForce_GOOD_FOR_TIME         TimeInForce = "A"
	TimeInForce_GOOD_FOR_AUCTION      TimeInForce = "B"
	TimeInForce_GOOD_FOR_THIS_MONTH   TimeInForce = "C"
)

// TradeCondition field enumeration values.
type TradeCondition string

const (
	TradeCondition_CANCEL                                  TradeCondition = "0"
	TradeCondition_IMPLIED_TRADE                           TradeCondition = "1"
	TradeCondition_MARKETPLACE_ENTERED_TRADE               TradeCondition = "2"
	TradeCondition_MULT_ASSET_CLASS_MULTILEG_TRADE         TradeCondition = "3"
	TradeCondition_MULTILEG_TO_MULTILEG_TRADE              TradeCondition = "4"
	TradeCondition_CASH                                    TradeCondition = "A"
	TradeCondition_SPREAD                                  TradeCondition = "AA"
	TradeCondition_SPREAD_ETH                              TradeCondition = "AB"
	TradeCondition_STRADDLE                                TradeCondition = "AC"
	TradeCondition_STRADDLE_ETH                            TradeCondition = "AD"
	TradeCondition_STOPPED                                 TradeCondition = "AE"
	TradeCondition_STOPPED_ETH                             TradeCondition = "AF"
	TradeCondition_REGULAR_ETH                             TradeCondition = "AG"
	TradeCondition_COMBO                                   TradeCondition = "AH"
	TradeCondition_COMBO_ETH                               TradeCondition = "AI"
	TradeCondition_OFFICIAL_CLOSING_PRICE                  TradeCondition = "AJ"
	TradeCondition_PRIOR_REFERENCE_PRICE                   TradeCondition = "AK"
	TradeCondition_STOPPED_SOLD_LAST                       TradeCondition = "AL"
	TradeCondition_STOPPED_OUT_OF_SEQUENCE                 TradeCondition = "AM"
	TradeCondition_OFFICAL_CLOSING_PRICE                   TradeCondition = "AN"
	TradeCondition_CROSSED_AO                              TradeCondition = "AO"
	TradeCondition_FAST_MARKET                             TradeCondition = "AP"
	TradeCondition_AUTOMATIC_EXECUTION                     TradeCondition = "AQ"
	TradeCondition_FORM_T                                  TradeCondition = "AR"
	TradeCondition_BASKET_INDEX                            TradeCondition = "AS"
	TradeCondition_BURST_BASKET                            TradeCondition = "AT"
	TradeCondition_OUTSIDE_SPREAD                          TradeCondition = "AV"
	TradeCondition_AVERAGE_PRICE_TRADE                     TradeCondition = "B"
	TradeCondition_CASH_TRADE                              TradeCondition = "C"
	TradeCondition_NEXT_DAY                                TradeCondition = "D"
	TradeCondition_OPENING_REOPENING_TRADE_DETAIL          TradeCondition = "E"
	TradeCondition_INTRADAY_TRADE_DETAIL                   TradeCondition = "F"
	TradeCondition_RULE_127_TRADE                          TradeCondition = "G"
	TradeCondition_RULE_155_TRADE                          TradeCondition = "H"
	TradeCondition_SOLD_LAST                               TradeCondition = "I"
	TradeCondition_NEXT_DAY_TRADE                          TradeCondition = "J"
	TradeCondition_OPENED                                  TradeCondition = "K"
	TradeCondition_SELLER                                  TradeCondition = "L"
	TradeCondition_SOLD                                    TradeCondition = "M"
	TradeCondition_STOPPED_STOCK                           TradeCondition = "N"
	TradeCondition_IMBALANCE_MORE_BUYERS                   TradeCondition = "P"
	TradeCondition_IMBALANCE_MORE_SELLERS                  TradeCondition = "Q"
	TradeCondition_OPENING_PRICE                           TradeCondition = "R"
	TradeCondition_BARGAIN_CONDITION                       TradeCondition = "S"
	TradeCondition_CONVERTED_PRICE_INDICATOR               TradeCondition = "T"
	TradeCondition_EXCHANGE_LAST                           TradeCondition = "U"
	TradeCondition_FINAL_PRICE_OF_SESSION                  TradeCondition = "V"
	TradeCondition_EX_PIT                                  TradeCondition = "W"
	TradeCondition_CROSSED_X                               TradeCondition = "X"
	TradeCondition_TRADES_RESULTING_FROM_MANUAL_SLOW_QUOTE TradeCondition = "Y"
	TradeCondition_TRADES_RESULTING_FROM_INTERMARKET_SWEEP TradeCondition = "Z"
	TradeCondition_VOLUME_ONLY                             TradeCondition = "a"
	TradeCondition_DIRECT_PLUS                             TradeCondition = "b"
	TradeCondition_ACQUISITION                             TradeCondition = "c"
	TradeCondition_BUNCHED                                 TradeCondition = "d"
	TradeCondition_DISTRIBUTION                            TradeCondition = "e"
	TradeCondition_BUNCHED_SALE                            TradeCondition = "f"
	TradeCondition_SPLIT_TRADE                             TradeCondition = "g"
	TradeCondition_CANCEL_STOPPED                          TradeCondition = "h"
	TradeCondition_CANCEL_ETH                              TradeCondition = "i"
	TradeCondition_CANCEL_STOPPED_ETH                      TradeCondition = "j"
	TradeCondition_OUT_OF_SEQUENCE_ETH                     TradeCondition = "k"
	TradeCondition_CANCEL_LAST_ETH                         TradeCondition = "l"
	TradeCondition_SOLD_LAST_SALE_ETH                      TradeCondition = "m"
	TradeCondition_CANCEL_LAST                             TradeCondition = "n"
	TradeCondition_SOLD_LAST_SALE                          TradeCondition = "o"
	TradeCondition_CANCEL_OPEN                             TradeCondition = "p"
	TradeCondition_CANCEL_OPEN_ETH                         TradeCondition = "q"
	TradeCondition_OPENED_SALE_ETH                         TradeCondition = "r"
	TradeCondition_CANCEL_ONLY                             TradeCondition = "s"
	TradeCondition_CANCEL_ONLY_ETH                         TradeCondition = "t"
	TradeCondition_LATE_OPEN_ETH                           TradeCondition = "u"
	TradeCondition_AUTO_EXECUTION_ETH                      TradeCondition = "v"
	TradeCondition_REOPEN                                  TradeCondition = "w"
	TradeCondition_REOPEN_ETH                              TradeCondition = "x"
	TradeCondition_ADJUSTED                                TradeCondition = "y"
	TradeCondition_ADJUSTED_ETH                            TradeCondition = "z"
)

// TradingSessionID field enumeration values.
type TradingSessionID string

const (
	TradingSessionID_DAY                 TradingSessionID = "1"
	TradingSessionID_HALFDAY             TradingSessionID = "2"
	TradingSessionID_MORNING             TradingSessionID = "3"
	TradingSessionID_AFTERNOON           TradingSessionID = "4"
	TradingSessionID_EVENING             TradingSessionID = "5"
	TradingSessionID_AFTER_HOURS         TradingSessionID = "6"
	TradingSessionID_BREAK               TradingSessionID = "BK"
	TradingSessionID_CLOSING_CALL        TradingSessionID = "CC"
	TradingSessionID_CONTINUOUS_TRADING  TradingSessionID = "CO"
	TradingSessionID_CLOSING_UNCROSSING  TradingSessionID = "CX"
	TradingSessionID_OPENING_CALL_PERIOD TradingSessionID = "OC"
	TradingSessionID_OFF_HOURS           TradingSessionID = "OH"
	TradingSessionID_OPENING_UNCROSSING  TradingSessionID = "OX"
)

// Urgency field enumeration values.
type Urgency string

const (
	Urgency_NORMAL     Urgency = "0"
	Urgency_FLASH      Urgency = "1"
	Urgency_BACKGROUND Urgency = "2"
)
