// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CardStatisticsResponseBody struct {
	Code             *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *CardStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                         `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CardStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBody) SetCode(v string) *CardStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *CardStatisticsResponseBody) SetData(v *CardStatisticsResponseBodyData) *CardStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *CardStatisticsResponseBody) SetErrorMessage(v string) *CardStatisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CardStatisticsResponseBody) SetLocalizedMessage(v string) *CardStatisticsResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *CardStatisticsResponseBody) SetRequestId(v string) *CardStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CardStatisticsResponseBody) SetSuccess(v bool) *CardStatisticsResponseBody {
	s.Success = &v
	return s
}

type CardStatisticsResponseBodyData struct {
	ErrorStopStatisticsDTO       *CardStatisticsResponseBodyDataErrorStopStatisticsDTO       `json:"ErrorStopStatisticsDTO,omitempty" xml:"ErrorStopStatisticsDTO,omitempty" type:"Struct"`
	ExhaustStopStatisticsDTO     *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO     `json:"ExhaustStopStatisticsDTO,omitempty" xml:"ExhaustStopStatisticsDTO,omitempty" type:"Struct"`
	ExpireStopStatisticsDTO      *CardStatisticsResponseBodyDataExpireStopStatisticsDTO      `json:"ExpireStopStatisticsDTO,omitempty" xml:"ExpireStopStatisticsDTO,omitempty" type:"Struct"`
	FlowOutStatisticsDTO         *CardStatisticsResponseBodyDataFlowOutStatisticsDTO         `json:"FlowOutStatisticsDTO,omitempty" xml:"FlowOutStatisticsDTO,omitempty" type:"Struct"`
	ManageStopStatisticsDTO      *CardStatisticsResponseBodyDataManageStopStatisticsDTO      `json:"ManageStopStatisticsDTO,omitempty" xml:"ManageStopStatisticsDTO,omitempty" type:"Struct"`
	RiskWaringStatisticsDTO      *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO      `json:"RiskWaringStatisticsDTO,omitempty" xml:"RiskWaringStatisticsDTO,omitempty" type:"Struct"`
	UnCertifiedStopStatisticsDTO *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO `json:"UnCertifiedStopStatisticsDTO,omitempty" xml:"UnCertifiedStopStatisticsDTO,omitempty" type:"Struct"`
	UnbindResumeStatisticsDTO    *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO    `json:"UnbindResumeStatisticsDTO,omitempty" xml:"UnbindResumeStatisticsDTO,omitempty" type:"Struct"`
	WeekWarnStatisticsDTO        *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO        `json:"WeekWarnStatisticsDTO,omitempty" xml:"WeekWarnStatisticsDTO,omitempty" type:"Struct"`
}

func (s CardStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyData) SetErrorStopStatisticsDTO(v *CardStatisticsResponseBodyDataErrorStopStatisticsDTO) *CardStatisticsResponseBodyData {
	s.ErrorStopStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetExhaustStopStatisticsDTO(v *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) *CardStatisticsResponseBodyData {
	s.ExhaustStopStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetExpireStopStatisticsDTO(v *CardStatisticsResponseBodyDataExpireStopStatisticsDTO) *CardStatisticsResponseBodyData {
	s.ExpireStopStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetFlowOutStatisticsDTO(v *CardStatisticsResponseBodyDataFlowOutStatisticsDTO) *CardStatisticsResponseBodyData {
	s.FlowOutStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetManageStopStatisticsDTO(v *CardStatisticsResponseBodyDataManageStopStatisticsDTO) *CardStatisticsResponseBodyData {
	s.ManageStopStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetRiskWaringStatisticsDTO(v *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) *CardStatisticsResponseBodyData {
	s.RiskWaringStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetUnCertifiedStopStatisticsDTO(v *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) *CardStatisticsResponseBodyData {
	s.UnCertifiedStopStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetUnbindResumeStatisticsDTO(v *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) *CardStatisticsResponseBodyData {
	s.UnbindResumeStatisticsDTO = v
	return s
}

func (s *CardStatisticsResponseBodyData) SetWeekWarnStatisticsDTO(v *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) *CardStatisticsResponseBodyData {
	s.WeekWarnStatisticsDTO = v
	return s
}

type CardStatisticsResponseBodyDataErrorStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataErrorStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataErrorStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataErrorStopStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataErrorStopStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataErrorStopStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataErrorStopStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataExhaustStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataExpireStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataExpireStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataExpireStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataExpireStopStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExpireStopStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExpireStopStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataExpireStopStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataFlowOutStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataFlowOutStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataFlowOutStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataFlowOutStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataFlowOutStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataFlowOutStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataFlowOutStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataManageStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataManageStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataManageStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataManageStopStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataManageStopStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataManageStopStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataManageStopStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataRiskWaringStatisticsDTO struct {
	StopCount        *int64 `json:"StopCount,omitempty" xml:"StopCount,omitempty"`
	WaringTotalCount *int64 `json:"WaringTotalCount,omitempty" xml:"WaringTotalCount,omitempty"`
	WarningCount     *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetStopCount(v int64) *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.StopCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetWaringTotalCount(v int64) *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.WaringTotalCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetWarningCount(v int64) *CardStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.WarningCount = &v
	return s
}

type CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponseBodyDataWeekWarnStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetPoolCount(v int64) *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetSameFlowCardCount(v int64) *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetSingleCardCount(v int64) *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetTotalCount(v int64) *CardStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.TotalCount = &v
	return s
}

type CardStatisticsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CardStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s CardStatisticsResponse) GoString() string {
	return s.String()
}

func (s *CardStatisticsResponse) SetHeaders(v map[string]*string) *CardStatisticsResponse {
	s.Headers = v
	return s
}

func (s *CardStatisticsResponse) SetStatusCode(v int32) *CardStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *CardStatisticsResponse) SetBody(v *CardStatisticsResponseBody) *CardStatisticsResponse {
	s.Body = v
	return s
}

type ForceActivationRequest struct {
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	Iccid    *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s ForceActivationRequest) String() string {
	return tea.Prettify(s)
}

func (s ForceActivationRequest) GoString() string {
	return s.String()
}

func (s *ForceActivationRequest) SetDateType(v string) *ForceActivationRequest {
	s.DateType = &v
	return s
}

func (s *ForceActivationRequest) SetIccid(v string) *ForceActivationRequest {
	s.Iccid = &v
	return s
}

type ForceActivationResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ForceActivationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceActivationResponseBody) GoString() string {
	return s.String()
}

func (s *ForceActivationResponseBody) SetCode(v string) *ForceActivationResponseBody {
	s.Code = &v
	return s
}

func (s *ForceActivationResponseBody) SetData(v bool) *ForceActivationResponseBody {
	s.Data = &v
	return s
}

func (s *ForceActivationResponseBody) SetErrorMessage(v string) *ForceActivationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ForceActivationResponseBody) SetLocalizedMessage(v string) *ForceActivationResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ForceActivationResponseBody) SetRequestId(v string) *ForceActivationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceActivationResponseBody) SetSuccess(v bool) *ForceActivationResponseBody {
	s.Success = &v
	return s
}

type ForceActivationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ForceActivationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ForceActivationResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceActivationResponse) GoString() string {
	return s.String()
}

func (s *ForceActivationResponse) SetHeaders(v map[string]*string) *ForceActivationResponse {
	s.Headers = v
	return s
}

func (s *ForceActivationResponse) SetStatusCode(v int32) *ForceActivationResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceActivationResponse) SetBody(v *ForceActivationResponseBody) *ForceActivationResponse {
	s.Body = v
	return s
}

type GetCardDetailRequest struct {
	DestroyCard *bool   `json:"DestroyCard,omitempty" xml:"DestroyCard,omitempty"`
	Iccid       *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ShowPsim    *bool   `json:"ShowPsim,omitempty" xml:"ShowPsim,omitempty"`
}

func (s GetCardDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCardDetailRequest) SetDestroyCard(v bool) *GetCardDetailRequest {
	s.DestroyCard = &v
	return s
}

func (s *GetCardDetailRequest) SetIccid(v string) *GetCardDetailRequest {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailRequest) SetInstanceId(v string) *GetCardDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCardDetailRequest) SetShowPsim(v bool) *GetCardDetailRequest {
	s.ShowPsim = &v
	return s
}

type GetCardDetailResponseBody struct {
	Code             *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetCardDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                        `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBody) SetCode(v string) *GetCardDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardDetailResponseBody) SetData(v *GetCardDetailResponseBodyData) *GetCardDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetCardDetailResponseBody) SetErrorMessage(v string) *GetCardDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardDetailResponseBody) SetLocalizedMessage(v string) *GetCardDetailResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardDetailResponseBody) SetRequestId(v string) *GetCardDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardDetailResponseBody) SetSuccess(v bool) *GetCardDetailResponseBody {
	s.Success = &v
	return s
}

type GetCardDetailResponseBodyData struct {
	ListPsimCards []*GetCardDetailResponseBodyDataListPsimCards `json:"ListPsimCards,omitempty" xml:"ListPsimCards,omitempty" type:"Repeated"`
	VsimCardInfo  *GetCardDetailResponseBodyDataVsimCardInfo    `json:"VsimCardInfo,omitempty" xml:"VsimCardInfo,omitempty" type:"Struct"`
}

func (s GetCardDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyData) SetListPsimCards(v []*GetCardDetailResponseBodyDataListPsimCards) *GetCardDetailResponseBodyData {
	s.ListPsimCards = v
	return s
}

func (s *GetCardDetailResponseBodyData) SetVsimCardInfo(v *GetCardDetailResponseBodyDataVsimCardInfo) *GetCardDetailResponseBodyData {
	s.VsimCardInfo = v
	return s
}

type GetCardDetailResponseBodyDataListPsimCards struct {
	ApnName               *string   `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyStatus         *string   `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	Iccid                 *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                  []*string `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	Ip                    []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Msisdn                []*string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	OpenSms               *bool     `json:"OpenSms,omitempty" xml:"OpenSms,omitempty"`
	OsStatus              *string   `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	PeriodAddFlow         *string   `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodSmsUse          *string   `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment *string   `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	Status                *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Vendor                *string   `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCardDetailResponseBodyDataListPsimCards) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataListPsimCards) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetApnName(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.ApnName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetCertifyStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.CertifyStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetIccid(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetImsi(v []*string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Imsi = v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetIp(v []*string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Ip = v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetMsisdn(v []*string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Msisdn = v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetOpenSms(v bool) *GetCardDetailResponseBodyDataListPsimCards {
	s.OpenSms = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetOsStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.OsStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPeriodAddFlow(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PeriodAddFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPeriodSmsUse(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PeriodSmsUse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetPrivateNetworkSegment(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.PrivateNetworkSegment = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetStatus(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Status = &v
	return s
}

func (s *GetCardDetailResponseBodyDataListPsimCards) SetVendor(v string) *GetCardDetailResponseBodyDataListPsimCards {
	s.Vendor = &v
	return s
}

type GetCardDetailResponseBodyDataVsimCardInfo struct {
	ActiveTime                    *string                                             `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	ActiveType                    *string                                             `json:"ActiveType,omitempty" xml:"ActiveType,omitempty"`
	AliFee                        *string                                             `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId                 *string                                             `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName                       *string                                             `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	AutoLimitResume               *bool                                               `json:"AutoLimitResume,omitempty" xml:"AutoLimitResume,omitempty"`
	AutoRebindReuse               *bool                                               `json:"AutoRebindReuse,omitempty" xml:"AutoRebindReuse,omitempty"`
	CardLimitSpeedThreshold       *int32                                              `json:"CardLimitSpeedThreshold,omitempty" xml:"CardLimitSpeedThreshold,omitempty"`
	CardLimitStopThreshold        *int32                                              `json:"CardLimitStopThreshold,omitempty" xml:"CardLimitStopThreshold,omitempty"`
	CertifyStatus                 *string                                             `json:"CertifyStatus,omitempty" xml:"CertifyStatus,omitempty"`
	CertifyType                   *string                                             `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialInstanceId          *string                                             `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialLimitSpeedThreshold *int32                                              `json:"CredentialLimitSpeedThreshold,omitempty" xml:"CredentialLimitSpeedThreshold,omitempty"`
	CredentialLimitStopThreshold  *int32                                              `json:"CredentialLimitStopThreshold,omitempty" xml:"CredentialLimitStopThreshold,omitempty"`
	CredentialNo                  *string                                             `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType                *string                                             `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DataLevel                     *string                                             `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType                      *string                                             `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DeviceImei                    *string                                             `json:"DeviceImei,omitempty" xml:"DeviceImei,omitempty"`
	DirectionalGroupId            *string                                             `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	DirectionalGroupName          *string                                             `json:"DirectionalGroupName,omitempty" xml:"DirectionalGroupName,omitempty"`
	ExpireTime                    *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FlowThresholdUnit             *string                                             `json:"FlowThresholdUnit,omitempty" xml:"FlowThresholdUnit,omitempty"`
	Iccid                         *string                                             `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                          []*string                                           `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	Ip                            []*string                                           `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	IsAutoRecharge                *bool                                               `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	Msisdn                        []*string                                           `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	NotifyId                      *string                                             `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OpenAccountTime               *string                                             `json:"OpenAccountTime,omitempty" xml:"OpenAccountTime,omitempty"`
	OpenSms                       *bool                                               `json:"OpenSms,omitempty" xml:"OpenSms,omitempty"`
	OsStatus                      *string                                             `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	Period                        *string                                             `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodAddFlow                 *string                                             `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodRestFlow                *string                                             `json:"PeriodRestFlow,omitempty" xml:"PeriodRestFlow,omitempty"`
	PeriodSmsUse                  *string                                             `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment         *string                                             `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	SimType                       *string                                             `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status                        *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TagList                       []*GetCardDetailResponseBodyDataVsimCardInfoTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Vendor                        *string                                             `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	VsimInstanceId                *int32                                              `json:"VsimInstanceId,omitempty" xml:"VsimInstanceId,omitempty"`
}

func (s GetCardDetailResponseBodyDataVsimCardInfo) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataVsimCardInfo) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetActiveTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ActiveTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetActiveType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ActiveType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAliFee(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AliFee = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAliyunOrderId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AliyunOrderId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetApnName(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ApnName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAutoLimitResume(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AutoLimitResume = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetAutoRebindReuse(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.AutoRebindReuse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCardLimitSpeedThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CardLimitSpeedThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCardLimitStopThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CardLimitStopThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCertifyStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CertifyStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCertifyType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CertifyType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialInstanceId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialInstanceId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialLimitSpeedThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialLimitSpeedThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialLimitStopThreshold(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialLimitStopThreshold = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialNo(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialNo = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetCredentialType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.CredentialType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDataLevel(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DataLevel = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDataType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DataType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDeviceImei(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DeviceImei = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDirectionalGroupId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DirectionalGroupId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetDirectionalGroupName(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.DirectionalGroupName = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetExpireTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.ExpireTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetFlowThresholdUnit(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.FlowThresholdUnit = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetIccid(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Iccid = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetImsi(v []*string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Imsi = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetIp(v []*string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Ip = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetIsAutoRecharge(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.IsAutoRecharge = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetMsisdn(v []*string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Msisdn = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetNotifyId(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.NotifyId = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetOpenAccountTime(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.OpenAccountTime = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetOpenSms(v bool) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.OpenSms = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetOsStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.OsStatus = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriod(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Period = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodAddFlow(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodAddFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodRestFlow(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodRestFlow = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPeriodSmsUse(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PeriodSmsUse = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetPrivateNetworkSegment(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.PrivateNetworkSegment = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetSimType(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.SimType = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetStatus(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Status = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetTagList(v []*GetCardDetailResponseBodyDataVsimCardInfoTagList) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.TagList = v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetVendor(v string) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.Vendor = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfo) SetVsimInstanceId(v int32) *GetCardDetailResponseBodyDataVsimCardInfo {
	s.VsimInstanceId = &v
	return s
}

type GetCardDetailResponseBodyDataVsimCardInfoTagList struct {
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetCardDetailResponseBodyDataVsimCardInfoTagList) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponseBodyDataVsimCardInfoTagList) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponseBodyDataVsimCardInfoTagList) SetId(v int64) *GetCardDetailResponseBodyDataVsimCardInfoTagList {
	s.Id = &v
	return s
}

func (s *GetCardDetailResponseBodyDataVsimCardInfoTagList) SetTagName(v string) *GetCardDetailResponseBodyDataVsimCardInfoTagList {
	s.TagName = &v
	return s
}

type GetCardDetailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCardDetailResponse) SetHeaders(v map[string]*string) *GetCardDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCardDetailResponse) SetStatusCode(v int32) *GetCardDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardDetailResponse) SetBody(v *GetCardDetailResponseBody) *GetCardDetailResponse {
	s.Body = v
	return s
}

type GetCardFlowInfoRequest struct {
	DateList []*string `json:"DateList,omitempty" xml:"DateList,omitempty" type:"Repeated"`
	Iccid    *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s GetCardFlowInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoRequest) SetDateList(v []*string) *GetCardFlowInfoRequest {
	s.DateList = v
	return s
}

func (s *GetCardFlowInfoRequest) SetIccid(v string) *GetCardFlowInfoRequest {
	s.Iccid = &v
	return s
}

type GetCardFlowInfoResponseBody struct {
	Code             *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetCardFlowInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                          `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardFlowInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBody) SetCode(v string) *GetCardFlowInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetData(v *GetCardFlowInfoResponseBodyData) *GetCardFlowInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetErrorMessage(v string) *GetCardFlowInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetLocalizedMessage(v string) *GetCardFlowInfoResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetRequestId(v string) *GetCardFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardFlowInfoResponseBody) SetSuccess(v bool) *GetCardFlowInfoResponseBody {
	s.Success = &v
	return s
}

type GetCardFlowInfoResponseBodyData struct {
	ListCardMonthFlow []*GetCardFlowInfoResponseBodyDataListCardMonthFlow `json:"ListCardMonthFlow,omitempty" xml:"ListCardMonthFlow,omitempty" type:"Repeated"`
	ListPackageDTO    []*GetCardFlowInfoResponseBodyDataListPackageDTO    `json:"ListPackageDTO,omitempty" xml:"ListPackageDTO,omitempty" type:"Repeated"`
	ListVendorDetail  []*GetCardFlowInfoResponseBodyDataListVendorDetail  `json:"ListVendorDetail,omitempty" xml:"ListVendorDetail,omitempty" type:"Repeated"`
}

func (s GetCardFlowInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyData) SetListCardMonthFlow(v []*GetCardFlowInfoResponseBodyDataListCardMonthFlow) *GetCardFlowInfoResponseBodyData {
	s.ListCardMonthFlow = v
	return s
}

func (s *GetCardFlowInfoResponseBodyData) SetListPackageDTO(v []*GetCardFlowInfoResponseBodyDataListPackageDTO) *GetCardFlowInfoResponseBodyData {
	s.ListPackageDTO = v
	return s
}

func (s *GetCardFlowInfoResponseBodyData) SetListVendorDetail(v []*GetCardFlowInfoResponseBodyDataListVendorDetail) *GetCardFlowInfoResponseBodyData {
	s.ListVendorDetail = v
	return s
}

type GetCardFlowInfoResponseBodyDataListCardMonthFlow struct {
	FlowCount   *string                                                        `json:"FlowCount,omitempty" xml:"FlowCount,omitempty"`
	ListDayFlow []*GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow `json:"ListDayFlow,omitempty" xml:"ListDayFlow,omitempty" type:"Repeated"`
	Month       *string                                                        `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlow) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlow) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetFlowCount(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.FlowCount = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetListDayFlow(v []*GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.ListDayFlow = v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlow) SetMonth(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlow {
	s.Month = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow struct {
	Day  *string `json:"Day,omitempty" xml:"Day,omitempty"`
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) SetDay(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow {
	s.Day = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow) SetFlow(v string) *GetCardFlowInfoResponseBodyDataListCardMonthFlowListDayFlow {
	s.Flow = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListPackageDTO struct {
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PackageName   *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListPackageDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListPackageDTO) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetEffectiveTime(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.EffectiveTime = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetExpireTime(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.ExpireTime = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetPackageName(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.PackageName = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListPackageDTO) SetRemark(v string) *GetCardFlowInfoResponseBodyDataListPackageDTO {
	s.Remark = &v
	return s
}

type GetCardFlowInfoResponseBodyDataListVendorDetail struct {
	NetWorkDelay   *string `json:"NetWorkDelay,omitempty" xml:"NetWorkDelay,omitempty"`
	Ratio          *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	SignalStrength *string `json:"SignalStrength,omitempty" xml:"SignalStrength,omitempty"`
	UsedFlow       *string `json:"UsedFlow,omitempty" xml:"UsedFlow,omitempty"`
	Vendor         *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCardFlowInfoResponseBodyDataListVendorDetail) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponseBodyDataListVendorDetail) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetNetWorkDelay(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.NetWorkDelay = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetRatio(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.Ratio = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetSignalStrength(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.SignalStrength = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetUsedFlow(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.UsedFlow = &v
	return s
}

func (s *GetCardFlowInfoResponseBodyDataListVendorDetail) SetVendor(v string) *GetCardFlowInfoResponseBodyDataListVendorDetail {
	s.Vendor = &v
	return s
}

type GetCardFlowInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCardFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCardFlowInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardFlowInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCardFlowInfoResponse) SetHeaders(v map[string]*string) *GetCardFlowInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCardFlowInfoResponse) SetStatusCode(v int32) *GetCardFlowInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardFlowInfoResponse) SetBody(v *GetCardFlowInfoResponseBody) *GetCardFlowInfoResponse {
	s.Body = v
	return s
}

type GetCredentialPoolStatisticsRequest struct {
	CredentialNO *string `json:"CredentialNO,omitempty" xml:"CredentialNO,omitempty"`
	Date         *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetCredentialPoolStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsRequest) SetCredentialNO(v string) *GetCredentialPoolStatisticsRequest {
	s.CredentialNO = &v
	return s
}

func (s *GetCredentialPoolStatisticsRequest) SetDate(v string) *GetCredentialPoolStatisticsRequest {
	s.Date = &v
	return s
}

type GetCredentialPoolStatisticsResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetCredentialPoolStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCredentialPoolStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponseBody) SetCode(v string) *GetCredentialPoolStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetData(v *GetCredentialPoolStatisticsResponseBodyData) *GetCredentialPoolStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetErrorMessage(v string) *GetCredentialPoolStatisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetRequestId(v string) *GetCredentialPoolStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBody) SetSuccess(v bool) *GetCredentialPoolStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetCredentialPoolStatisticsResponseBodyData struct {
	CardActiveNum          *int64  `json:"CardActiveNum,omitempty" xml:"CardActiveNum,omitempty"`
	CardTotalNum           *int64  `json:"CardTotalNum,omitempty" xml:"CardTotalNum,omitempty"`
	CredentialInstanceId   *string `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialNO           *string `json:"CredentialNO,omitempty" xml:"CredentialNO,omitempty"`
	CredentialType         *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	EffectiveAvailableFlow *string `json:"EffectiveAvailableFlow,omitempty" xml:"EffectiveAvailableFlow,omitempty"`
	EffectiveTotalFlow     *string `json:"EffectiveTotalFlow,omitempty" xml:"EffectiveTotalFlow,omitempty"`
	MonthFeatureFee        *int64  `json:"MonthFeatureFee,omitempty" xml:"MonthFeatureFee,omitempty"`
	PoolAvaiable           *string `json:"PoolAvaiable,omitempty" xml:"PoolAvaiable,omitempty"`
	PoolGrandTotal         *string `json:"PoolGrandTotal,omitempty" xml:"PoolGrandTotal,omitempty"`
	PoolGrandTotalUsed     *string `json:"PoolGrandTotalUsed,omitempty" xml:"PoolGrandTotalUsed,omitempty"`
	PoolOutUsed            *string `json:"PoolOutUsed,omitempty" xml:"PoolOutUsed,omitempty"`
	PoolUsed               *string `json:"PoolUsed,omitempty" xml:"PoolUsed,omitempty"`
	SmsUsed                *int64  `json:"SmsUsed,omitempty" xml:"SmsUsed,omitempty"`
}

func (s GetCredentialPoolStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCardActiveNum(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.CardActiveNum = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCardTotalNum(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.CardTotalNum = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialInstanceId(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialInstanceId = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialNO(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialNO = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetCredentialType(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetEffectiveAvailableFlow(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.EffectiveAvailableFlow = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetEffectiveTotalFlow(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.EffectiveTotalFlow = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetMonthFeatureFee(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.MonthFeatureFee = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolAvaiable(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolAvaiable = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolGrandTotal(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolGrandTotal = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolGrandTotalUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolGrandTotalUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolOutUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolOutUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetPoolUsed(v string) *GetCredentialPoolStatisticsResponseBodyData {
	s.PoolUsed = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetSmsUsed(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.SmsUsed = &v
	return s
}

type GetCredentialPoolStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCredentialPoolStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCredentialPoolStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialPoolStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialPoolStatisticsResponse) SetHeaders(v map[string]*string) *GetCredentialPoolStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialPoolStatisticsResponse) SetStatusCode(v int32) *GetCredentialPoolStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponse) SetBody(v *GetCredentialPoolStatisticsResponseBody) *GetCredentialPoolStatisticsResponse {
	s.Body = v
	return s
}

type ListCardInfoRequest struct {
	ActiveTimeEnd      *string `json:"ActiveTimeEnd,omitempty" xml:"ActiveTimeEnd,omitempty"`
	ActiveTimeStart    *string `json:"ActiveTimeStart,omitempty" xml:"ActiveTimeStart,omitempty"`
	AliFee             *string `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId      *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName            *string `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyType        *string `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialNo       *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	DataLevel          *string `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType           *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DirectionalGroupId *string `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	ExpireTimeEnd      *string `json:"ExpireTimeEnd,omitempty" xml:"ExpireTimeEnd,omitempty"`
	ExpireTimeStart    *string `json:"ExpireTimeStart,omitempty" xml:"ExpireTimeStart,omitempty"`
	Iccid              *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi               *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	IsAutoRecharge     *bool   `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	Msisdn             *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	NotifyId           *string `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OsStatus           *string `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	PageNo             *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PoolId             *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	SimType            *string `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagName            *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Vendor             *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListCardInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoRequest) GoString() string {
	return s.String()
}

func (s *ListCardInfoRequest) SetActiveTimeEnd(v string) *ListCardInfoRequest {
	s.ActiveTimeEnd = &v
	return s
}

func (s *ListCardInfoRequest) SetActiveTimeStart(v string) *ListCardInfoRequest {
	s.ActiveTimeStart = &v
	return s
}

func (s *ListCardInfoRequest) SetAliFee(v string) *ListCardInfoRequest {
	s.AliFee = &v
	return s
}

func (s *ListCardInfoRequest) SetAliyunOrderId(v string) *ListCardInfoRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *ListCardInfoRequest) SetApnName(v string) *ListCardInfoRequest {
	s.ApnName = &v
	return s
}

func (s *ListCardInfoRequest) SetCertifyType(v string) *ListCardInfoRequest {
	s.CertifyType = &v
	return s
}

func (s *ListCardInfoRequest) SetCredentialNo(v string) *ListCardInfoRequest {
	s.CredentialNo = &v
	return s
}

func (s *ListCardInfoRequest) SetDataLevel(v string) *ListCardInfoRequest {
	s.DataLevel = &v
	return s
}

func (s *ListCardInfoRequest) SetDataType(v string) *ListCardInfoRequest {
	s.DataType = &v
	return s
}

func (s *ListCardInfoRequest) SetDirectionalGroupId(v string) *ListCardInfoRequest {
	s.DirectionalGroupId = &v
	return s
}

func (s *ListCardInfoRequest) SetExpireTimeEnd(v string) *ListCardInfoRequest {
	s.ExpireTimeEnd = &v
	return s
}

func (s *ListCardInfoRequest) SetExpireTimeStart(v string) *ListCardInfoRequest {
	s.ExpireTimeStart = &v
	return s
}

func (s *ListCardInfoRequest) SetIccid(v string) *ListCardInfoRequest {
	s.Iccid = &v
	return s
}

func (s *ListCardInfoRequest) SetImsi(v string) *ListCardInfoRequest {
	s.Imsi = &v
	return s
}

func (s *ListCardInfoRequest) SetIsAutoRecharge(v bool) *ListCardInfoRequest {
	s.IsAutoRecharge = &v
	return s
}

func (s *ListCardInfoRequest) SetMsisdn(v string) *ListCardInfoRequest {
	s.Msisdn = &v
	return s
}

func (s *ListCardInfoRequest) SetNotifyId(v string) *ListCardInfoRequest {
	s.NotifyId = &v
	return s
}

func (s *ListCardInfoRequest) SetOsStatus(v string) *ListCardInfoRequest {
	s.OsStatus = &v
	return s
}

func (s *ListCardInfoRequest) SetPageNo(v int32) *ListCardInfoRequest {
	s.PageNo = &v
	return s
}

func (s *ListCardInfoRequest) SetPageSize(v int32) *ListCardInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListCardInfoRequest) SetPeriod(v string) *ListCardInfoRequest {
	s.Period = &v
	return s
}

func (s *ListCardInfoRequest) SetPoolId(v string) *ListCardInfoRequest {
	s.PoolId = &v
	return s
}

func (s *ListCardInfoRequest) SetSimType(v string) *ListCardInfoRequest {
	s.SimType = &v
	return s
}

func (s *ListCardInfoRequest) SetStatus(v string) *ListCardInfoRequest {
	s.Status = &v
	return s
}

func (s *ListCardInfoRequest) SetTagName(v string) *ListCardInfoRequest {
	s.TagName = &v
	return s
}

func (s *ListCardInfoRequest) SetVendor(v string) *ListCardInfoRequest {
	s.Vendor = &v
	return s
}

type ListCardInfoResponseBody struct {
	Code             *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *ListCardInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                       `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCardInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListCardInfoResponseBody) SetCode(v string) *ListCardInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListCardInfoResponseBody) SetData(v *ListCardInfoResponseBodyData) *ListCardInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListCardInfoResponseBody) SetErrorMessage(v string) *ListCardInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCardInfoResponseBody) SetLocalizedMessage(v string) *ListCardInfoResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ListCardInfoResponseBody) SetRequestId(v string) *ListCardInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCardInfoResponseBody) SetSuccess(v bool) *ListCardInfoResponseBody {
	s.Success = &v
	return s
}

type ListCardInfoResponseBodyData struct {
	List      []*ListCardInfoResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageCount *int32                              `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNo    *int32                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCardInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCardInfoResponseBodyData) SetList(v []*ListCardInfoResponseBodyDataList) *ListCardInfoResponseBodyData {
	s.List = v
	return s
}

func (s *ListCardInfoResponseBodyData) SetPageCount(v int32) *ListCardInfoResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *ListCardInfoResponseBodyData) SetPageNo(v int32) *ListCardInfoResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListCardInfoResponseBodyData) SetPageSize(v int32) *ListCardInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCardInfoResponseBodyData) SetTotal(v int32) *ListCardInfoResponseBodyData {
	s.Total = &v
	return s
}

type ListCardInfoResponseBodyDataList struct {
	ActiveTime            *string                                    `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	ActiveType            *string                                    `json:"ActiveType,omitempty" xml:"ActiveType,omitempty"`
	AliFee                *string                                    `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId         *string                                    `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName               *string                                    `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyType           *string                                    `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialInstanceId  *string                                    `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialNo          *string                                    `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType        *string                                    `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DataLevel             *string                                    `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType              *string                                    `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DirectionalGroupName  *string                                    `json:"DirectionalGroupName,omitempty" xml:"DirectionalGroupName,omitempty"`
	ExpireTime            *string                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Iccid                 *string                                    `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                  []*string                                  `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	IsAutoRecharge        *bool                                      `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	Msisdn                []*string                                  `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	NotifyId              *string                                    `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OpenAccountTime       *string                                    `json:"OpenAccountTime,omitempty" xml:"OpenAccountTime,omitempty"`
	OsStatus              *string                                    `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	Period                *string                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodAddFlow         *string                                    `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodRestFlow        *string                                    `json:"PeriodRestFlow,omitempty" xml:"PeriodRestFlow,omitempty"`
	PeriodSmsUse          *string                                    `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment *string                                    `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	Remark                *string                                    `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SimType               *string                                    `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status                *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TagList               []*ListCardInfoResponseBodyDataListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Vendor                *string                                    `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	VsimInstanceId        *int64                                     `json:"VsimInstanceId,omitempty" xml:"VsimInstanceId,omitempty"`
}

func (s ListCardInfoResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCardInfoResponseBodyDataList) SetActiveTime(v string) *ListCardInfoResponseBodyDataList {
	s.ActiveTime = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetActiveType(v string) *ListCardInfoResponseBodyDataList {
	s.ActiveType = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetAliFee(v string) *ListCardInfoResponseBodyDataList {
	s.AliFee = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetAliyunOrderId(v string) *ListCardInfoResponseBodyDataList {
	s.AliyunOrderId = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetApnName(v string) *ListCardInfoResponseBodyDataList {
	s.ApnName = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetCertifyType(v string) *ListCardInfoResponseBodyDataList {
	s.CertifyType = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetCredentialInstanceId(v string) *ListCardInfoResponseBodyDataList {
	s.CredentialInstanceId = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetCredentialNo(v string) *ListCardInfoResponseBodyDataList {
	s.CredentialNo = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetCredentialType(v string) *ListCardInfoResponseBodyDataList {
	s.CredentialType = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetDataLevel(v string) *ListCardInfoResponseBodyDataList {
	s.DataLevel = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetDataType(v string) *ListCardInfoResponseBodyDataList {
	s.DataType = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetDirectionalGroupName(v string) *ListCardInfoResponseBodyDataList {
	s.DirectionalGroupName = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetExpireTime(v string) *ListCardInfoResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetIccid(v string) *ListCardInfoResponseBodyDataList {
	s.Iccid = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetImsi(v []*string) *ListCardInfoResponseBodyDataList {
	s.Imsi = v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetIsAutoRecharge(v bool) *ListCardInfoResponseBodyDataList {
	s.IsAutoRecharge = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetMsisdn(v []*string) *ListCardInfoResponseBodyDataList {
	s.Msisdn = v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetNotifyId(v string) *ListCardInfoResponseBodyDataList {
	s.NotifyId = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetOpenAccountTime(v string) *ListCardInfoResponseBodyDataList {
	s.OpenAccountTime = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetOsStatus(v string) *ListCardInfoResponseBodyDataList {
	s.OsStatus = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetPeriod(v string) *ListCardInfoResponseBodyDataList {
	s.Period = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetPeriodAddFlow(v string) *ListCardInfoResponseBodyDataList {
	s.PeriodAddFlow = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetPeriodRestFlow(v string) *ListCardInfoResponseBodyDataList {
	s.PeriodRestFlow = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetPeriodSmsUse(v string) *ListCardInfoResponseBodyDataList {
	s.PeriodSmsUse = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetPrivateNetworkSegment(v string) *ListCardInfoResponseBodyDataList {
	s.PrivateNetworkSegment = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetRemark(v string) *ListCardInfoResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetSimType(v string) *ListCardInfoResponseBodyDataList {
	s.SimType = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetStatus(v string) *ListCardInfoResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetTagList(v []*ListCardInfoResponseBodyDataListTagList) *ListCardInfoResponseBodyDataList {
	s.TagList = v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetVendor(v string) *ListCardInfoResponseBodyDataList {
	s.Vendor = &v
	return s
}

func (s *ListCardInfoResponseBodyDataList) SetVsimInstanceId(v int64) *ListCardInfoResponseBodyDataList {
	s.VsimInstanceId = &v
	return s
}

type ListCardInfoResponseBodyDataListTagList struct {
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListCardInfoResponseBodyDataListTagList) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoResponseBodyDataListTagList) GoString() string {
	return s.String()
}

func (s *ListCardInfoResponseBodyDataListTagList) SetId(v int64) *ListCardInfoResponseBodyDataListTagList {
	s.Id = &v
	return s
}

func (s *ListCardInfoResponseBodyDataListTagList) SetTagName(v string) *ListCardInfoResponseBodyDataListTagList {
	s.TagName = &v
	return s
}

type ListCardInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCardInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCardInfoResponse) GoString() string {
	return s.String()
}

func (s *ListCardInfoResponse) SetHeaders(v map[string]*string) *ListCardInfoResponse {
	s.Headers = v
	return s
}

func (s *ListCardInfoResponse) SetStatusCode(v int32) *ListCardInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCardInfoResponse) SetBody(v *ListCardInfoResponseBody) *ListCardInfoResponse {
	s.Body = v
	return s
}

type ListOrderRequest struct {
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType   *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrderRequest) GoString() string {
	return s.String()
}

func (s *ListOrderRequest) SetEndDate(v string) *ListOrderRequest {
	s.EndDate = &v
	return s
}

func (s *ListOrderRequest) SetOrderId(v string) *ListOrderRequest {
	s.OrderId = &v
	return s
}

func (s *ListOrderRequest) SetOrderStatus(v string) *ListOrderRequest {
	s.OrderStatus = &v
	return s
}

func (s *ListOrderRequest) SetOrderType(v string) *ListOrderRequest {
	s.OrderType = &v
	return s
}

func (s *ListOrderRequest) SetPageNo(v int32) *ListOrderRequest {
	s.PageNo = &v
	return s
}

func (s *ListOrderRequest) SetPageSize(v int32) *ListOrderRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrderRequest) SetStartDate(v string) *ListOrderRequest {
	s.StartDate = &v
	return s
}

type ListOrderResponseBody struct {
	Code             *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *ListOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                    `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBody) SetCode(v string) *ListOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListOrderResponseBody) SetData(v *ListOrderResponseBodyData) *ListOrderResponseBody {
	s.Data = v
	return s
}

func (s *ListOrderResponseBody) SetErrorMessage(v string) *ListOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListOrderResponseBody) SetLocalizedMessage(v string) *ListOrderResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ListOrderResponseBody) SetRequestId(v string) *ListOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrderResponseBody) SetSuccess(v bool) *ListOrderResponseBody {
	s.Success = &v
	return s
}

type ListOrderResponseBodyData struct {
	List      []*ListOrderResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageCount *int32                           `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNo    *int32                           `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBodyData) SetList(v []*ListOrderResponseBodyDataList) *ListOrderResponseBodyData {
	s.List = v
	return s
}

func (s *ListOrderResponseBodyData) SetPageCount(v int32) *ListOrderResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *ListOrderResponseBodyData) SetPageNo(v int32) *ListOrderResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListOrderResponseBodyData) SetPageSize(v int32) *ListOrderResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListOrderResponseBodyData) SetTotal(v int32) *ListOrderResponseBodyData {
	s.Total = &v
	return s
}

type ListOrderResponseBodyDataList struct {
	AliFee            *string                                    `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	BillingCycle      *string                                    `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	BuyNum            *int32                                     `json:"BuyNum,omitempty" xml:"BuyNum,omitempty"`
	CardPayCount      *int32                                     `json:"CardPayCount,omitempty" xml:"CardPayCount,omitempty"`
	CredentialNo      *string                                    `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialPackage *string                                    `json:"CredentialPackage,omitempty" xml:"CredentialPackage,omitempty"`
	DataLevel         *string                                    `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DeliveryInfo      *ListOrderResponseBodyDataListDeliveryInfo `json:"DeliveryInfo,omitempty" xml:"DeliveryInfo,omitempty" type:"Struct"`
	ExpressNoList     []*string                                  `json:"ExpressNoList,omitempty" xml:"ExpressNoList,omitempty" type:"Repeated"`
	FlowType          *string                                    `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	FunctionFee       *int32                                     `json:"FunctionFee,omitempty" xml:"FunctionFee,omitempty"`
	OrderDetailUrl    *string                                    `json:"OrderDetailUrl,omitempty" xml:"OrderDetailUrl,omitempty"`
	OrderId           *string                                    `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderInfo         *string                                    `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty"`
	OrderStatus       *string                                    `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType         *string                                    `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PayDuration       *string                                    `json:"PayDuration,omitempty" xml:"PayDuration,omitempty"`
	PayTime           *string                                    `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	PoolCapacity      *string                                    `json:"PoolCapacity,omitempty" xml:"PoolCapacity,omitempty"`
	PoolCapacityUnit  *string                                    `json:"PoolCapacityUnit,omitempty" xml:"PoolCapacityUnit,omitempty"`
	PoolNo            *string                                    `json:"PoolNo,omitempty" xml:"PoolNo,omitempty"`
	Vendor            *string                                    `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListOrderResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBodyDataList) SetAliFee(v string) *ListOrderResponseBodyDataList {
	s.AliFee = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetBillingCycle(v string) *ListOrderResponseBodyDataList {
	s.BillingCycle = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetBuyNum(v int32) *ListOrderResponseBodyDataList {
	s.BuyNum = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetCardPayCount(v int32) *ListOrderResponseBodyDataList {
	s.CardPayCount = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetCredentialNo(v string) *ListOrderResponseBodyDataList {
	s.CredentialNo = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetCredentialPackage(v string) *ListOrderResponseBodyDataList {
	s.CredentialPackage = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetDataLevel(v string) *ListOrderResponseBodyDataList {
	s.DataLevel = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetDeliveryInfo(v *ListOrderResponseBodyDataListDeliveryInfo) *ListOrderResponseBodyDataList {
	s.DeliveryInfo = v
	return s
}

func (s *ListOrderResponseBodyDataList) SetExpressNoList(v []*string) *ListOrderResponseBodyDataList {
	s.ExpressNoList = v
	return s
}

func (s *ListOrderResponseBodyDataList) SetFlowType(v string) *ListOrderResponseBodyDataList {
	s.FlowType = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetFunctionFee(v int32) *ListOrderResponseBodyDataList {
	s.FunctionFee = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetOrderDetailUrl(v string) *ListOrderResponseBodyDataList {
	s.OrderDetailUrl = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetOrderId(v string) *ListOrderResponseBodyDataList {
	s.OrderId = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetOrderInfo(v string) *ListOrderResponseBodyDataList {
	s.OrderInfo = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetOrderStatus(v string) *ListOrderResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetOrderType(v string) *ListOrderResponseBodyDataList {
	s.OrderType = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetPayDuration(v string) *ListOrderResponseBodyDataList {
	s.PayDuration = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetPayTime(v string) *ListOrderResponseBodyDataList {
	s.PayTime = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetPoolCapacity(v string) *ListOrderResponseBodyDataList {
	s.PoolCapacity = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetPoolCapacityUnit(v string) *ListOrderResponseBodyDataList {
	s.PoolCapacityUnit = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetPoolNo(v string) *ListOrderResponseBodyDataList {
	s.PoolNo = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetVendor(v string) *ListOrderResponseBodyDataList {
	s.Vendor = &v
	return s
}

type ListOrderResponseBodyDataListDeliveryInfo struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	BuyerMessage *string `json:"BuyerMessage,omitempty" xml:"BuyerMessage,omitempty"`
	Mail         *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	Receiver     *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	ZipCode      *string `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

func (s ListOrderResponseBodyDataListDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponseBodyDataListDeliveryInfo) GoString() string {
	return s.String()
}

func (s *ListOrderResponseBodyDataListDeliveryInfo) SetAddress(v string) *ListOrderResponseBodyDataListDeliveryInfo {
	s.Address = &v
	return s
}

func (s *ListOrderResponseBodyDataListDeliveryInfo) SetBuyerMessage(v string) *ListOrderResponseBodyDataListDeliveryInfo {
	s.BuyerMessage = &v
	return s
}

func (s *ListOrderResponseBodyDataListDeliveryInfo) SetMail(v string) *ListOrderResponseBodyDataListDeliveryInfo {
	s.Mail = &v
	return s
}

func (s *ListOrderResponseBodyDataListDeliveryInfo) SetReceiver(v string) *ListOrderResponseBodyDataListDeliveryInfo {
	s.Receiver = &v
	return s
}

func (s *ListOrderResponseBodyDataListDeliveryInfo) SetZipCode(v string) *ListOrderResponseBodyDataListDeliveryInfo {
	s.ZipCode = &v
	return s
}

type ListOrderResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrderResponse) GoString() string {
	return s.String()
}

func (s *ListOrderResponse) SetHeaders(v map[string]*string) *ListOrderResponse {
	s.Headers = v
	return s
}

func (s *ListOrderResponse) SetStatusCode(v int32) *ListOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrderResponse) SetBody(v *ListOrderResponseBody) *ListOrderResponse {
	s.Body = v
	return s
}

type RebindResumeSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s RebindResumeSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardRequest) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardRequest) SetIccid(v string) *RebindResumeSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *RebindResumeSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *RebindResumeSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type RebindResumeSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s RebindResumeSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardShrinkRequest) SetIccid(v string) *RebindResumeSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *RebindResumeSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *RebindResumeSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type RebindResumeSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RebindResumeSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardResponseBody) SetCode(v string) *RebindResumeSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetData(v bool) *RebindResumeSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetErrorMessage(v string) *RebindResumeSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetLocalizedMessage(v string) *RebindResumeSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetRequestId(v string) *RebindResumeSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebindResumeSingleCardResponseBody) SetSuccess(v bool) *RebindResumeSingleCardResponseBody {
	s.Success = &v
	return s
}

type RebindResumeSingleCardResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebindResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebindResumeSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RebindResumeSingleCardResponse) GoString() string {
	return s.String()
}

func (s *RebindResumeSingleCardResponse) SetHeaders(v map[string]*string) *RebindResumeSingleCardResponse {
	s.Headers = v
	return s
}

func (s *RebindResumeSingleCardResponse) SetStatusCode(v int32) *RebindResumeSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RebindResumeSingleCardResponse) SetBody(v *RebindResumeSingleCardResponseBody) *RebindResumeSingleCardResponse {
	s.Body = v
	return s
}

type RenewRequest struct {
	ApiProduct   *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ApiRevision  *string `json:"ApiRevision,omitempty" xml:"ApiRevision,omitempty"`
	BuyNum       *int32  `json:"BuyNum,omitempty" xml:"BuyNum,omitempty"`
	Iccid        *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OfferCode    *string `json:"OfferCode,omitempty" xml:"OfferCode,omitempty"`
	RechargeType *string `json:"RechargeType,omitempty" xml:"RechargeType,omitempty"`
	SerialNo     *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s RenewRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewRequest) GoString() string {
	return s.String()
}

func (s *RenewRequest) SetApiProduct(v string) *RenewRequest {
	s.ApiProduct = &v
	return s
}

func (s *RenewRequest) SetApiRevision(v string) *RenewRequest {
	s.ApiRevision = &v
	return s
}

func (s *RenewRequest) SetBuyNum(v int32) *RenewRequest {
	s.BuyNum = &v
	return s
}

func (s *RenewRequest) SetIccid(v string) *RenewRequest {
	s.Iccid = &v
	return s
}

func (s *RenewRequest) SetOfferCode(v string) *RenewRequest {
	s.OfferCode = &v
	return s
}

func (s *RenewRequest) SetRechargeType(v string) *RenewRequest {
	s.RechargeType = &v
	return s
}

func (s *RenewRequest) SetSerialNo(v string) *RenewRequest {
	s.SerialNo = &v
	return s
}

type RenewResponseBody struct {
	Code             *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *RenewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewResponseBody) GoString() string {
	return s.String()
}

func (s *RenewResponseBody) SetCode(v string) *RenewResponseBody {
	s.Code = &v
	return s
}

func (s *RenewResponseBody) SetData(v *RenewResponseBodyData) *RenewResponseBody {
	s.Data = v
	return s
}

func (s *RenewResponseBody) SetErrorMessage(v string) *RenewResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RenewResponseBody) SetLocalizedMessage(v string) *RenewResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *RenewResponseBody) SetRequestId(v string) *RenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewResponseBody) SetSuccess(v bool) *RenewResponseBody {
	s.Success = &v
	return s
}

type RenewResponseBodyData struct {
	OrderNo  *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s RenewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenewResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewResponseBodyData) SetOrderNo(v string) *RenewResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *RenewResponseBodyData) SetSerialNo(v string) *RenewResponseBodyData {
	s.SerialNo = &v
	return s
}

type RenewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewResponse) GoString() string {
	return s.String()
}

func (s *RenewResponse) SetHeaders(v map[string]*string) *RenewResponse {
	s.Headers = v
	return s
}

func (s *RenewResponse) SetStatusCode(v int32) *RenewResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewResponse) SetBody(v *RenewResponseBody) *RenewResponse {
	s.Body = v
	return s
}

type ResumeSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s ResumeSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardRequest) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardRequest) SetIccid(v string) *ResumeSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *ResumeSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *ResumeSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type ResumeSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s ResumeSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardShrinkRequest) SetIccid(v string) *ResumeSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *ResumeSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *ResumeSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type ResumeSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardResponseBody) SetCode(v string) *ResumeSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetData(v bool) *ResumeSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetErrorMessage(v string) *ResumeSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetLocalizedMessage(v string) *ResumeSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetRequestId(v string) *ResumeSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeSingleCardResponseBody) SetSuccess(v bool) *ResumeSingleCardResponseBody {
	s.Success = &v
	return s
}

type ResumeSingleCardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeSingleCardResponse) GoString() string {
	return s.String()
}

func (s *ResumeSingleCardResponse) SetHeaders(v map[string]*string) *ResumeSingleCardResponse {
	s.Headers = v
	return s
}

func (s *ResumeSingleCardResponse) SetStatusCode(v int32) *ResumeSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSingleCardResponse) SetBody(v *ResumeSingleCardResponseBody) *ResumeSingleCardResponse {
	s.Body = v
	return s
}

type SetCardStopRuleRequest struct {
	AutoRestore *bool   `json:"AutoRestore,omitempty" xml:"AutoRestore,omitempty"`
	FlowLimit   *int64  `json:"FlowLimit,omitempty" xml:"FlowLimit,omitempty"`
	Iccid       *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s SetCardStopRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCardStopRuleRequest) GoString() string {
	return s.String()
}

func (s *SetCardStopRuleRequest) SetAutoRestore(v bool) *SetCardStopRuleRequest {
	s.AutoRestore = &v
	return s
}

func (s *SetCardStopRuleRequest) SetFlowLimit(v int64) *SetCardStopRuleRequest {
	s.FlowLimit = &v
	return s
}

func (s *SetCardStopRuleRequest) SetIccid(v string) *SetCardStopRuleRequest {
	s.Iccid = &v
	return s
}

type SetCardStopRuleResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCardStopRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCardStopRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetCardStopRuleResponseBody) SetCode(v string) *SetCardStopRuleResponseBody {
	s.Code = &v
	return s
}

func (s *SetCardStopRuleResponseBody) SetData(v bool) *SetCardStopRuleResponseBody {
	s.Data = &v
	return s
}

func (s *SetCardStopRuleResponseBody) SetErrorMessage(v string) *SetCardStopRuleResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SetCardStopRuleResponseBody) SetLocalizedMessage(v string) *SetCardStopRuleResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *SetCardStopRuleResponseBody) SetRequestId(v string) *SetCardStopRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCardStopRuleResponseBody) SetSuccess(v bool) *SetCardStopRuleResponseBody {
	s.Success = &v
	return s
}

type SetCardStopRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetCardStopRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCardStopRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCardStopRuleResponse) GoString() string {
	return s.String()
}

func (s *SetCardStopRuleResponse) SetHeaders(v map[string]*string) *SetCardStopRuleResponse {
	s.Headers = v
	return s
}

func (s *SetCardStopRuleResponse) SetStatusCode(v int32) *SetCardStopRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCardStopRuleResponse) SetBody(v *SetCardStopRuleResponseBody) *SetCardStopRuleResponse {
	s.Body = v
	return s
}

type StopSingleCardRequest struct {
	Iccid      *string                `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns map[string]interface{} `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s StopSingleCardRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardRequest) GoString() string {
	return s.String()
}

func (s *StopSingleCardRequest) SetIccid(v string) *StopSingleCardRequest {
	s.Iccid = &v
	return s
}

func (s *StopSingleCardRequest) SetOptMsisdns(v map[string]interface{}) *StopSingleCardRequest {
	s.OptMsisdns = v
	return s
}

type StopSingleCardShrinkRequest struct {
	Iccid            *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdnsShrink *string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty"`
}

func (s StopSingleCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopSingleCardShrinkRequest) SetIccid(v string) *StopSingleCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *StopSingleCardShrinkRequest) SetOptMsisdnsShrink(v string) *StopSingleCardShrinkRequest {
	s.OptMsisdnsShrink = &v
	return s
}

type StopSingleCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopSingleCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardResponseBody) GoString() string {
	return s.String()
}

func (s *StopSingleCardResponseBody) SetCode(v string) *StopSingleCardResponseBody {
	s.Code = &v
	return s
}

func (s *StopSingleCardResponseBody) SetData(v bool) *StopSingleCardResponseBody {
	s.Data = &v
	return s
}

func (s *StopSingleCardResponseBody) SetErrorMessage(v string) *StopSingleCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopSingleCardResponseBody) SetLocalizedMessage(v string) *StopSingleCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *StopSingleCardResponseBody) SetRequestId(v string) *StopSingleCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSingleCardResponseBody) SetSuccess(v bool) *StopSingleCardResponseBody {
	s.Success = &v
	return s
}

type StopSingleCardResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopSingleCardResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSingleCardResponse) GoString() string {
	return s.String()
}

func (s *StopSingleCardResponse) SetHeaders(v map[string]*string) *StopSingleCardResponse {
	s.Headers = v
	return s
}

func (s *StopSingleCardResponse) SetStatusCode(v int32) *StopSingleCardResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSingleCardResponse) SetBody(v *StopSingleCardResponseBody) *StopSingleCardResponse {
	s.Body = v
	return s
}

type UpdateAutoRechargeSwitchRequest struct {
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Open  *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s UpdateAutoRechargeSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoRechargeSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoRechargeSwitchRequest) SetIccid(v string) *UpdateAutoRechargeSwitchRequest {
	s.Iccid = &v
	return s
}

func (s *UpdateAutoRechargeSwitchRequest) SetOpen(v bool) *UpdateAutoRechargeSwitchRequest {
	s.Open = &v
	return s
}

type UpdateAutoRechargeSwitchResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutoRechargeSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoRechargeSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetCode(v string) *UpdateAutoRechargeSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetData(v bool) *UpdateAutoRechargeSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetErrorMessage(v string) *UpdateAutoRechargeSwitchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetLocalizedMessage(v string) *UpdateAutoRechargeSwitchResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetRequestId(v string) *UpdateAutoRechargeSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponseBody) SetSuccess(v bool) *UpdateAutoRechargeSwitchResponseBody {
	s.Success = &v
	return s
}

type UpdateAutoRechargeSwitchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAutoRechargeSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAutoRechargeSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoRechargeSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoRechargeSwitchResponse) SetHeaders(v map[string]*string) *UpdateAutoRechargeSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoRechargeSwitchResponse) SetStatusCode(v int32) *UpdateAutoRechargeSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoRechargeSwitchResponse) SetBody(v *UpdateAutoRechargeSwitchResponseBody) *UpdateAutoRechargeSwitchResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkcard"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CardStatisticsWithOptions(runtime *util.RuntimeOptions) (_result *CardStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CardStatistics"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CardStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CardStatistics() (_result *CardStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CardStatisticsResponse{}
	_body, _err := client.CardStatisticsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForceActivationWithOptions(request *ForceActivationRequest, runtime *util.RuntimeOptions) (_result *ForceActivationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateType)) {
		query["DateType"] = request.DateType
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ForceActivation"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ForceActivationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForceActivation(request *ForceActivationRequest) (_result *ForceActivationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForceActivationResponse{}
	_body, _err := client.ForceActivationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardDetailWithOptions(request *GetCardDetailRequest, runtime *util.RuntimeOptions) (_result *GetCardDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestroyCard)) {
		query["DestroyCard"] = request.DestroyCard
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowPsim)) {
		query["ShowPsim"] = request.ShowPsim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardDetail"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardDetail(request *GetCardDetailRequest) (_result *GetCardDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardDetailResponse{}
	_body, _err := client.GetCardDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardFlowInfoWithOptions(request *GetCardFlowInfoRequest, runtime *util.RuntimeOptions) (_result *GetCardFlowInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DateList)) {
		query["DateList"] = request.DateList
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardFlowInfo"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardFlowInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardFlowInfo(request *GetCardFlowInfoRequest) (_result *GetCardFlowInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardFlowInfoResponse{}
	_body, _err := client.GetCardFlowInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCredentialPoolStatisticsWithOptions(request *GetCredentialPoolStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetCredentialPoolStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialNO)) {
		query["CredentialNO"] = request.CredentialNO
	}

	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["Date"] = request.Date
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCredentialPoolStatistics"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCredentialPoolStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCredentialPoolStatistics(request *GetCredentialPoolStatisticsRequest) (_result *GetCredentialPoolStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCredentialPoolStatisticsResponse{}
	_body, _err := client.GetCredentialPoolStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCardInfoWithOptions(request *ListCardInfoRequest, runtime *util.RuntimeOptions) (_result *ListCardInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveTimeEnd)) {
		query["ActiveTimeEnd"] = request.ActiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveTimeStart)) {
		query["ActiveTimeStart"] = request.ActiveTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.AliFee)) {
		query["AliFee"] = request.AliFee
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunOrderId)) {
		query["AliyunOrderId"] = request.AliyunOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.ApnName)) {
		query["ApnName"] = request.ApnName
	}

	if !tea.BoolValue(util.IsUnset(request.CertifyType)) {
		query["CertifyType"] = request.CertifyType
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialNo)) {
		query["CredentialNo"] = request.CredentialNo
	}

	if !tea.BoolValue(util.IsUnset(request.DataLevel)) {
		query["DataLevel"] = request.DataLevel
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.DirectionalGroupId)) {
		query["DirectionalGroupId"] = request.DirectionalGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTimeEnd)) {
		query["ExpireTimeEnd"] = request.ExpireTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTimeStart)) {
		query["ExpireTimeStart"] = request.ExpireTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.Imsi)) {
		query["Imsi"] = request.Imsi
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRecharge)) {
		query["IsAutoRecharge"] = request.IsAutoRecharge
	}

	if !tea.BoolValue(util.IsUnset(request.Msisdn)) {
		query["Msisdn"] = request.Msisdn
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyId)) {
		query["NotifyId"] = request.NotifyId
	}

	if !tea.BoolValue(util.IsUnset(request.OsStatus)) {
		query["OsStatus"] = request.OsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PoolId)) {
		query["PoolId"] = request.PoolId
	}

	if !tea.BoolValue(util.IsUnset(request.SimType)) {
		query["SimType"] = request.SimType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		query["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCardInfo"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCardInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCardInfo(request *ListCardInfoRequest) (_result *ListCardInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCardInfoResponse{}
	_body, _err := client.ListCardInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrderWithOptions(request *ListOrderRequest, runtime *util.RuntimeOptions) (_result *ListOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStatus)) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrder"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrder(request *ListOrderRequest) (_result *ListOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrderResponse{}
	_body, _err := client.ListOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebindResumeSingleCardWithOptions(tmpReq *RebindResumeSingleCardRequest, runtime *util.RuntimeOptions) (_result *RebindResumeSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RebindResumeSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebindResumeSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebindResumeSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebindResumeSingleCard(request *RebindResumeSingleCardRequest) (_result *RebindResumeSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebindResumeSingleCardResponse{}
	_body, _err := client.RebindResumeSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewWithOptions(request *RenewRequest, runtime *util.RuntimeOptions) (_result *RenewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuyNum)) {
		query["BuyNum"] = request.BuyNum
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OfferCode)) {
		query["OfferCode"] = request.OfferCode
	}

	if !tea.BoolValue(util.IsUnset(request.RechargeType)) {
		query["RechargeType"] = request.RechargeType
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiProduct)) {
		body["ApiProduct"] = request.ApiProduct
	}

	if !tea.BoolValue(util.IsUnset(request.ApiRevision)) {
		body["ApiRevision"] = request.ApiRevision
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Renew"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Renew(request *RenewRequest) (_result *RenewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewResponse{}
	_body, _err := client.RenewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeSingleCardWithOptions(tmpReq *ResumeSingleCardRequest, runtime *util.RuntimeOptions) (_result *ResumeSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ResumeSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeSingleCard(request *ResumeSingleCardRequest) (_result *ResumeSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeSingleCardResponse{}
	_body, _err := client.ResumeSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCardStopRuleWithOptions(request *SetCardStopRuleRequest, runtime *util.RuntimeOptions) (_result *SetCardStopRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRestore)) {
		query["AutoRestore"] = request.AutoRestore
	}

	if !tea.BoolValue(util.IsUnset(request.FlowLimit)) {
		query["FlowLimit"] = request.FlowLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetCardStopRule"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetCardStopRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCardStopRule(request *SetCardStopRuleRequest) (_result *SetCardStopRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCardStopRuleResponse{}
	_body, _err := client.SetCardStopRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopSingleCardWithOptions(tmpReq *StopSingleCardRequest, runtime *util.RuntimeOptions) (_result *StopSingleCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopSingleCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OptMsisdns)) {
		request.OptMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OptMsisdns, tea.String("OptMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.OptMsisdnsShrink)) {
		query["OptMsisdns"] = request.OptMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopSingleCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSingleCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopSingleCard(request *StopSingleCardRequest) (_result *StopSingleCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopSingleCardResponse{}
	_body, _err := client.StopSingleCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutoRechargeSwitchWithOptions(request *UpdateAutoRechargeSwitchRequest, runtime *util.RuntimeOptions) (_result *UpdateAutoRechargeSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.Open)) {
		query["Open"] = request.Open
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAutoRechargeSwitch"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutoRechargeSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutoRechargeSwitch(request *UpdateAutoRechargeSwitchRequest) (_result *UpdateAutoRechargeSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAutoRechargeSwitchResponse{}
	_body, _err := client.UpdateAutoRechargeSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
