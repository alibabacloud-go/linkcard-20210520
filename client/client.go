// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCardToDirectionalGroupRequest struct {
	AddType *string `json:"AddType,omitempty" xml:"AddType,omitempty"`
	// Linkcard
	ApiProduct *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	GroupId    *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IccidList  []*string `json:"IccidList,omitempty" xml:"IccidList,omitempty" type:"Repeated"`
	MsgNotify  *bool     `json:"MsgNotify,omitempty" xml:"MsgNotify,omitempty"`
	SerialNo   *string   `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AddCardToDirectionalGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCardToDirectionalGroupRequest) GoString() string {
	return s.String()
}

func (s *AddCardToDirectionalGroupRequest) SetAddType(v string) *AddCardToDirectionalGroupRequest {
	s.AddType = &v
	return s
}

func (s *AddCardToDirectionalGroupRequest) SetApiProduct(v string) *AddCardToDirectionalGroupRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddCardToDirectionalGroupRequest) SetGroupId(v string) *AddCardToDirectionalGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AddCardToDirectionalGroupRequest) SetIccidList(v []*string) *AddCardToDirectionalGroupRequest {
	s.IccidList = v
	return s
}

func (s *AddCardToDirectionalGroupRequest) SetMsgNotify(v bool) *AddCardToDirectionalGroupRequest {
	s.MsgNotify = &v
	return s
}

func (s *AddCardToDirectionalGroupRequest) SetSerialNo(v string) *AddCardToDirectionalGroupRequest {
	s.SerialNo = &v
	return s
}

type AddCardToDirectionalGroupShrinkRequest struct {
	AddType *string `json:"AddType,omitempty" xml:"AddType,omitempty"`
	// Linkcard
	ApiProduct      *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IccidListShrink *string `json:"IccidList,omitempty" xml:"IccidList,omitempty"`
	MsgNotify       *bool   `json:"MsgNotify,omitempty" xml:"MsgNotify,omitempty"`
	SerialNo        *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AddCardToDirectionalGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCardToDirectionalGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetAddType(v string) *AddCardToDirectionalGroupShrinkRequest {
	s.AddType = &v
	return s
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetApiProduct(v string) *AddCardToDirectionalGroupShrinkRequest {
	s.ApiProduct = &v
	return s
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetGroupId(v string) *AddCardToDirectionalGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetIccidListShrink(v string) *AddCardToDirectionalGroupShrinkRequest {
	s.IccidListShrink = &v
	return s
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetMsgNotify(v bool) *AddCardToDirectionalGroupShrinkRequest {
	s.MsgNotify = &v
	return s
}

func (s *AddCardToDirectionalGroupShrinkRequest) SetSerialNo(v string) *AddCardToDirectionalGroupShrinkRequest {
	s.SerialNo = &v
	return s
}

type AddCardToDirectionalGroupResponseBody struct {
	Code             *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *AddCardToDirectionalGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                                    `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCardToDirectionalGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCardToDirectionalGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddCardToDirectionalGroupResponseBody) SetCode(v string) *AddCardToDirectionalGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddCardToDirectionalGroupResponseBody) SetData(v *AddCardToDirectionalGroupResponseBodyData) *AddCardToDirectionalGroupResponseBody {
	s.Data = v
	return s
}

func (s *AddCardToDirectionalGroupResponseBody) SetErrorMessage(v string) *AddCardToDirectionalGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddCardToDirectionalGroupResponseBody) SetLocalizedMessage(v string) *AddCardToDirectionalGroupResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *AddCardToDirectionalGroupResponseBody) SetRequestId(v string) *AddCardToDirectionalGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCardToDirectionalGroupResponseBody) SetSuccess(v bool) *AddCardToDirectionalGroupResponseBody {
	s.Success = &v
	return s
}

type AddCardToDirectionalGroupResponseBodyData struct {
	Result   *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s AddCardToDirectionalGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddCardToDirectionalGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCardToDirectionalGroupResponseBodyData) SetResult(v bool) *AddCardToDirectionalGroupResponseBodyData {
	s.Result = &v
	return s
}

func (s *AddCardToDirectionalGroupResponseBodyData) SetSerialNo(v string) *AddCardToDirectionalGroupResponseBodyData {
	s.SerialNo = &v
	return s
}

type AddCardToDirectionalGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCardToDirectionalGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCardToDirectionalGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCardToDirectionalGroupResponse) GoString() string {
	return s.String()
}

func (s *AddCardToDirectionalGroupResponse) SetHeaders(v map[string]*string) *AddCardToDirectionalGroupResponse {
	s.Headers = v
	return s
}

func (s *AddCardToDirectionalGroupResponse) SetStatusCode(v int32) *AddCardToDirectionalGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCardToDirectionalGroupResponse) SetBody(v *AddCardToDirectionalGroupResponseBody) *AddCardToDirectionalGroupResponse {
	s.Body = v
	return s
}

type AddDirectionalAddressRequest struct {
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType        *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MsgNotify          *bool   `json:"MsgNotify,omitempty" xml:"MsgNotify,omitempty"`
	SerialNo           *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UrlInsecurityForce *bool   `json:"UrlInsecurityForce,omitempty" xml:"UrlInsecurityForce,omitempty"`
}

func (s AddDirectionalAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalAddressRequest) GoString() string {
	return s.String()
}

func (s *AddDirectionalAddressRequest) SetAddress(v string) *AddDirectionalAddressRequest {
	s.Address = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetAddressType(v string) *AddDirectionalAddressRequest {
	s.AddressType = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetGroupId(v string) *AddDirectionalAddressRequest {
	s.GroupId = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetMsgNotify(v bool) *AddDirectionalAddressRequest {
	s.MsgNotify = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetSerialNo(v string) *AddDirectionalAddressRequest {
	s.SerialNo = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetSource(v string) *AddDirectionalAddressRequest {
	s.Source = &v
	return s
}

func (s *AddDirectionalAddressRequest) SetUrlInsecurityForce(v bool) *AddDirectionalAddressRequest {
	s.UrlInsecurityForce = &v
	return s
}

type AddDirectionalAddressResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDirectionalAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AddDirectionalAddressResponseBody) SetCode(v string) *AddDirectionalAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AddDirectionalAddressResponseBody) SetData(v bool) *AddDirectionalAddressResponseBody {
	s.Data = &v
	return s
}

func (s *AddDirectionalAddressResponseBody) SetErrorMessage(v string) *AddDirectionalAddressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddDirectionalAddressResponseBody) SetLocalizedMessage(v string) *AddDirectionalAddressResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *AddDirectionalAddressResponseBody) SetRequestId(v string) *AddDirectionalAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDirectionalAddressResponseBody) SetSuccess(v bool) *AddDirectionalAddressResponseBody {
	s.Success = &v
	return s
}

type AddDirectionalAddressResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDirectionalAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDirectionalAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalAddressResponse) GoString() string {
	return s.String()
}

func (s *AddDirectionalAddressResponse) SetHeaders(v map[string]*string) *AddDirectionalAddressResponse {
	s.Headers = v
	return s
}

func (s *AddDirectionalAddressResponse) SetStatusCode(v int32) *AddDirectionalAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDirectionalAddressResponse) SetBody(v *AddDirectionalAddressResponseBody) *AddDirectionalAddressResponse {
	s.Body = v
	return s
}

type AddDirectionalCardRequest struct {
	FileUri      *string   `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
	GroupId      *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OrderList    []*string `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
	TagList      []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	UploadMethod *string   `json:"UploadMethod,omitempty" xml:"UploadMethod,omitempty"`
	UploadType   *string   `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s AddDirectionalCardRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalCardRequest) GoString() string {
	return s.String()
}

func (s *AddDirectionalCardRequest) SetFileUri(v string) *AddDirectionalCardRequest {
	s.FileUri = &v
	return s
}

func (s *AddDirectionalCardRequest) SetGroupId(v string) *AddDirectionalCardRequest {
	s.GroupId = &v
	return s
}

func (s *AddDirectionalCardRequest) SetGroupName(v string) *AddDirectionalCardRequest {
	s.GroupName = &v
	return s
}

func (s *AddDirectionalCardRequest) SetOrderList(v []*string) *AddDirectionalCardRequest {
	s.OrderList = v
	return s
}

func (s *AddDirectionalCardRequest) SetTagList(v []*string) *AddDirectionalCardRequest {
	s.TagList = v
	return s
}

func (s *AddDirectionalCardRequest) SetUploadMethod(v string) *AddDirectionalCardRequest {
	s.UploadMethod = &v
	return s
}

func (s *AddDirectionalCardRequest) SetUploadType(v string) *AddDirectionalCardRequest {
	s.UploadType = &v
	return s
}

type AddDirectionalCardShrinkRequest struct {
	FileUri         *string `json:"FileUri,omitempty" xml:"FileUri,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OrderListShrink *string `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	TagListShrink   *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	UploadMethod    *string `json:"UploadMethod,omitempty" xml:"UploadMethod,omitempty"`
	UploadType      *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
}

func (s AddDirectionalCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddDirectionalCardShrinkRequest) SetFileUri(v string) *AddDirectionalCardShrinkRequest {
	s.FileUri = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetGroupId(v string) *AddDirectionalCardShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetGroupName(v string) *AddDirectionalCardShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetOrderListShrink(v string) *AddDirectionalCardShrinkRequest {
	s.OrderListShrink = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetTagListShrink(v string) *AddDirectionalCardShrinkRequest {
	s.TagListShrink = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetUploadMethod(v string) *AddDirectionalCardShrinkRequest {
	s.UploadMethod = &v
	return s
}

func (s *AddDirectionalCardShrinkRequest) SetUploadType(v string) *AddDirectionalCardShrinkRequest {
	s.UploadType = &v
	return s
}

type AddDirectionalCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDirectionalCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalCardResponseBody) GoString() string {
	return s.String()
}

func (s *AddDirectionalCardResponseBody) SetCode(v string) *AddDirectionalCardResponseBody {
	s.Code = &v
	return s
}

func (s *AddDirectionalCardResponseBody) SetData(v string) *AddDirectionalCardResponseBody {
	s.Data = &v
	return s
}

func (s *AddDirectionalCardResponseBody) SetErrorMessage(v string) *AddDirectionalCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddDirectionalCardResponseBody) SetLocalizedMessage(v string) *AddDirectionalCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *AddDirectionalCardResponseBody) SetRequestId(v string) *AddDirectionalCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDirectionalCardResponseBody) SetSuccess(v bool) *AddDirectionalCardResponseBody {
	s.Success = &v
	return s
}

type AddDirectionalCardResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDirectionalCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDirectionalCardResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalCardResponse) GoString() string {
	return s.String()
}

func (s *AddDirectionalCardResponse) SetHeaders(v map[string]*string) *AddDirectionalCardResponse {
	s.Headers = v
	return s
}

func (s *AddDirectionalCardResponse) SetStatusCode(v int32) *AddDirectionalCardResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDirectionalCardResponse) SetBody(v *AddDirectionalCardResponseBody) *AddDirectionalCardResponse {
	s.Body = v
	return s
}

type AddDirectionalGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s AddDirectionalGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDirectionalGroupRequest) SetGroupName(v string) *AddDirectionalGroupRequest {
	s.GroupName = &v
	return s
}

type AddDirectionalGroupResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDirectionalGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDirectionalGroupResponseBody) SetCode(v string) *AddDirectionalGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddDirectionalGroupResponseBody) SetData(v int64) *AddDirectionalGroupResponseBody {
	s.Data = &v
	return s
}

func (s *AddDirectionalGroupResponseBody) SetErrorMessage(v string) *AddDirectionalGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddDirectionalGroupResponseBody) SetLocalizedMessage(v string) *AddDirectionalGroupResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *AddDirectionalGroupResponseBody) SetRequestId(v string) *AddDirectionalGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDirectionalGroupResponseBody) SetSuccess(v bool) *AddDirectionalGroupResponseBody {
	s.Success = &v
	return s
}

type AddDirectionalGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDirectionalGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDirectionalGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDirectionalGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDirectionalGroupResponse) SetHeaders(v map[string]*string) *AddDirectionalGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDirectionalGroupResponse) SetStatusCode(v int32) *AddDirectionalGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDirectionalGroupResponse) SetBody(v *AddDirectionalGroupResponseBody) *AddDirectionalGroupResponse {
	s.Body = v
	return s
}

type AddTagsToCardRequest struct {
	Iccid       *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	TagNameList []*string `json:"TagNameList,omitempty" xml:"TagNameList,omitempty" type:"Repeated"`
}

func (s AddTagsToCardRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToCardRequest) GoString() string {
	return s.String()
}

func (s *AddTagsToCardRequest) SetIccid(v string) *AddTagsToCardRequest {
	s.Iccid = &v
	return s
}

func (s *AddTagsToCardRequest) SetTagNameList(v []*string) *AddTagsToCardRequest {
	s.TagNameList = v
	return s
}

type AddTagsToCardShrinkRequest struct {
	Iccid             *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	TagNameListShrink *string `json:"TagNameList,omitempty" xml:"TagNameList,omitempty"`
}

func (s AddTagsToCardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToCardShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTagsToCardShrinkRequest) SetIccid(v string) *AddTagsToCardShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *AddTagsToCardShrinkRequest) SetTagNameListShrink(v string) *AddTagsToCardShrinkRequest {
	s.TagNameListShrink = &v
	return s
}

type AddTagsToCardResponseBody struct {
	Code         *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*AddTagsToCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTagsToCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToCardResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsToCardResponseBody) SetCode(v string) *AddTagsToCardResponseBody {
	s.Code = &v
	return s
}

func (s *AddTagsToCardResponseBody) SetData(v []*AddTagsToCardResponseBodyData) *AddTagsToCardResponseBody {
	s.Data = v
	return s
}

func (s *AddTagsToCardResponseBody) SetErrorMessage(v string) *AddTagsToCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddTagsToCardResponseBody) SetRequestId(v string) *AddTagsToCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsToCardResponseBody) SetSuccess(v bool) *AddTagsToCardResponseBody {
	s.Success = &v
	return s
}

type AddTagsToCardResponseBodyData struct {
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s AddTagsToCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddTagsToCardResponseBodyData) SetTagId(v string) *AddTagsToCardResponseBodyData {
	s.TagId = &v
	return s
}

func (s *AddTagsToCardResponseBodyData) SetTagName(v string) *AddTagsToCardResponseBodyData {
	s.TagName = &v
	return s
}

type AddTagsToCardResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTagsToCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTagsToCardResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsToCardResponse) GoString() string {
	return s.String()
}

func (s *AddTagsToCardResponse) SetHeaders(v map[string]*string) *AddTagsToCardResponse {
	s.Headers = v
	return s
}

func (s *AddTagsToCardResponse) SetStatusCode(v int32) *AddTagsToCardResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsToCardResponse) SetBody(v *AddTagsToCardResponseBody) *AddTagsToCardResponse {
	s.Body = v
	return s
}

type BatchAddDirectionalAddressRequest struct {
	AddressType *string   `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	GroupId     *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ListAddress []*string `json:"ListAddress,omitempty" xml:"ListAddress,omitempty" type:"Repeated"`
	Source      *string   `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s BatchAddDirectionalAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDirectionalAddressRequest) GoString() string {
	return s.String()
}

func (s *BatchAddDirectionalAddressRequest) SetAddressType(v string) *BatchAddDirectionalAddressRequest {
	s.AddressType = &v
	return s
}

func (s *BatchAddDirectionalAddressRequest) SetGroupId(v int64) *BatchAddDirectionalAddressRequest {
	s.GroupId = &v
	return s
}

func (s *BatchAddDirectionalAddressRequest) SetListAddress(v []*string) *BatchAddDirectionalAddressRequest {
	s.ListAddress = v
	return s
}

func (s *BatchAddDirectionalAddressRequest) SetSource(v string) *BatchAddDirectionalAddressRequest {
	s.Source = &v
	return s
}

type BatchAddDirectionalAddressResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchAddDirectionalAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDirectionalAddressResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddDirectionalAddressResponseBody) SetCode(v string) *BatchAddDirectionalAddressResponseBody {
	s.Code = &v
	return s
}

func (s *BatchAddDirectionalAddressResponseBody) SetData(v bool) *BatchAddDirectionalAddressResponseBody {
	s.Data = &v
	return s
}

func (s *BatchAddDirectionalAddressResponseBody) SetErrorMessage(v string) *BatchAddDirectionalAddressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchAddDirectionalAddressResponseBody) SetLocalizedMessage(v string) *BatchAddDirectionalAddressResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *BatchAddDirectionalAddressResponseBody) SetRequestId(v string) *BatchAddDirectionalAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchAddDirectionalAddressResponseBody) SetSuccess(v bool) *BatchAddDirectionalAddressResponseBody {
	s.Success = &v
	return s
}

type BatchAddDirectionalAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddDirectionalAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddDirectionalAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddDirectionalAddressResponse) GoString() string {
	return s.String()
}

func (s *BatchAddDirectionalAddressResponse) SetHeaders(v map[string]*string) *BatchAddDirectionalAddressResponse {
	s.Headers = v
	return s
}

func (s *BatchAddDirectionalAddressResponse) SetStatusCode(v int32) *BatchAddDirectionalAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddDirectionalAddressResponse) SetBody(v *BatchAddDirectionalAddressResponseBody) *BatchAddDirectionalAddressResponse {
	s.Body = v
	return s
}

type DeleteDirectionalAddressRequest struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MsgNotify *bool   `json:"MsgNotify,omitempty" xml:"MsgNotify,omitempty"`
	SerialNo  *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s DeleteDirectionalAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalAddressRequest) SetAddress(v string) *DeleteDirectionalAddressRequest {
	s.Address = &v
	return s
}

func (s *DeleteDirectionalAddressRequest) SetGroupId(v string) *DeleteDirectionalAddressRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDirectionalAddressRequest) SetMsgNotify(v bool) *DeleteDirectionalAddressRequest {
	s.MsgNotify = &v
	return s
}

func (s *DeleteDirectionalAddressRequest) SetSerialNo(v string) *DeleteDirectionalAddressRequest {
	s.SerialNo = &v
	return s
}

type DeleteDirectionalAddressResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDirectionalAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalAddressResponseBody) SetCode(v string) *DeleteDirectionalAddressResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDirectionalAddressResponseBody) SetData(v bool) *DeleteDirectionalAddressResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDirectionalAddressResponseBody) SetErrorMessage(v string) *DeleteDirectionalAddressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDirectionalAddressResponseBody) SetLocalizedMessage(v string) *DeleteDirectionalAddressResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *DeleteDirectionalAddressResponseBody) SetRequestId(v string) *DeleteDirectionalAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectionalAddressResponseBody) SetSuccess(v bool) *DeleteDirectionalAddressResponseBody {
	s.Success = &v
	return s
}

type DeleteDirectionalAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectionalAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectionalAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalAddressResponse) SetHeaders(v map[string]*string) *DeleteDirectionalAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectionalAddressResponse) SetStatusCode(v int32) *DeleteDirectionalAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectionalAddressResponse) SetBody(v *DeleteDirectionalAddressResponseBody) *DeleteDirectionalAddressResponse {
	s.Body = v
	return s
}

type DeleteDirectionalGroupRequest struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteDirectionalGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalGroupRequest) SetGroupId(v string) *DeleteDirectionalGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteDirectionalGroupResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDirectionalGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalGroupResponseBody) SetCode(v string) *DeleteDirectionalGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDirectionalGroupResponseBody) SetData(v bool) *DeleteDirectionalGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDirectionalGroupResponseBody) SetErrorMessage(v string) *DeleteDirectionalGroupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDirectionalGroupResponseBody) SetLocalizedMessage(v string) *DeleteDirectionalGroupResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *DeleteDirectionalGroupResponseBody) SetRequestId(v string) *DeleteDirectionalGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectionalGroupResponseBody) SetSuccess(v bool) *DeleteDirectionalGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteDirectionalGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirectionalGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirectionalGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirectionalGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirectionalGroupResponse) SetHeaders(v map[string]*string) *DeleteDirectionalGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirectionalGroupResponse) SetStatusCode(v int32) *DeleteDirectionalGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirectionalGroupResponse) SetBody(v *DeleteDirectionalGroupResponseBody) *DeleteDirectionalGroupResponse {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceActivationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardFlowInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetCardLatestFlowRequest struct {
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s GetCardLatestFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardLatestFlowRequest) GoString() string {
	return s.String()
}

func (s *GetCardLatestFlowRequest) SetIccid(v string) *GetCardLatestFlowRequest {
	s.Iccid = &v
	return s
}

type GetCardLatestFlowResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardLatestFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardLatestFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardLatestFlowResponseBody) SetCode(v string) *GetCardLatestFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardLatestFlowResponseBody) SetData(v string) *GetCardLatestFlowResponseBody {
	s.Data = &v
	return s
}

func (s *GetCardLatestFlowResponseBody) SetErrorMessage(v string) *GetCardLatestFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardLatestFlowResponseBody) SetLocalizedMessage(v string) *GetCardLatestFlowResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardLatestFlowResponseBody) SetRequestId(v string) *GetCardLatestFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardLatestFlowResponseBody) SetSuccess(v bool) *GetCardLatestFlowResponseBody {
	s.Success = &v
	return s
}

type GetCardLatestFlowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardLatestFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardLatestFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardLatestFlowResponse) GoString() string {
	return s.String()
}

func (s *GetCardLatestFlowResponse) SetHeaders(v map[string]*string) *GetCardLatestFlowResponse {
	s.Headers = v
	return s
}

func (s *GetCardLatestFlowResponse) SetStatusCode(v int32) *GetCardLatestFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardLatestFlowResponse) SetBody(v *GetCardLatestFlowResponseBody) *GetCardLatestFlowResponse {
	s.Body = v
	return s
}

type GetCardRealStatusRequest struct {
	Iccid    *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Msisdn   *string   `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	SerialNo []*string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty" type:"Repeated"`
}

func (s GetCardRealStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardRealStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCardRealStatusRequest) SetIccid(v string) *GetCardRealStatusRequest {
	s.Iccid = &v
	return s
}

func (s *GetCardRealStatusRequest) SetMsisdn(v string) *GetCardRealStatusRequest {
	s.Msisdn = &v
	return s
}

func (s *GetCardRealStatusRequest) SetSerialNo(v []*string) *GetCardRealStatusRequest {
	s.SerialNo = v
	return s
}

type GetCardRealStatusShrinkRequest struct {
	Iccid          *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Msisdn         *string `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	SerialNoShrink *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s GetCardRealStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCardRealStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCardRealStatusShrinkRequest) SetIccid(v string) *GetCardRealStatusShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *GetCardRealStatusShrinkRequest) SetMsisdn(v string) *GetCardRealStatusShrinkRequest {
	s.Msisdn = &v
	return s
}

func (s *GetCardRealStatusShrinkRequest) SetSerialNoShrink(v string) *GetCardRealStatusShrinkRequest {
	s.SerialNoShrink = &v
	return s
}

type GetCardRealStatusResponseBody struct {
	Code             *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             []*GetCardRealStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage     *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                              `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardRealStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardRealStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardRealStatusResponseBody) SetCode(v string) *GetCardRealStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardRealStatusResponseBody) SetData(v []*GetCardRealStatusResponseBodyData) *GetCardRealStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetCardRealStatusResponseBody) SetErrorMessage(v string) *GetCardRealStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardRealStatusResponseBody) SetLocalizedMessage(v string) *GetCardRealStatusResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardRealStatusResponseBody) SetRequestId(v string) *GetCardRealStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardRealStatusResponseBody) SetSuccess(v bool) *GetCardRealStatusResponseBody {
	s.Success = &v
	return s
}

type GetCardRealStatusResponseBodyData struct {
	Gprs     *bool   `json:"Gprs,omitempty" xml:"Gprs,omitempty"`
	Iccid    *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Online   *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCardRealStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardRealStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardRealStatusResponseBodyData) SetGprs(v bool) *GetCardRealStatusResponseBodyData {
	s.Gprs = &v
	return s
}

func (s *GetCardRealStatusResponseBodyData) SetIccid(v string) *GetCardRealStatusResponseBodyData {
	s.Iccid = &v
	return s
}

func (s *GetCardRealStatusResponseBodyData) SetOnline(v bool) *GetCardRealStatusResponseBodyData {
	s.Online = &v
	return s
}

func (s *GetCardRealStatusResponseBodyData) SetSerialNo(v string) *GetCardRealStatusResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *GetCardRealStatusResponseBodyData) SetStatus(v string) *GetCardRealStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetCardRealStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardRealStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardRealStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardRealStatusResponse) GoString() string {
	return s.String()
}

func (s *GetCardRealStatusResponse) SetHeaders(v map[string]*string) *GetCardRealStatusResponse {
	s.Headers = v
	return s
}

func (s *GetCardRealStatusResponse) SetStatusCode(v int32) *GetCardRealStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardRealStatusResponse) SetBody(v *GetCardRealStatusResponseBody) *GetCardRealStatusResponse {
	s.Body = v
	return s
}

type GetCardStatusStatisticsResponseBody struct {
	Code             *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetCardStatusStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                                  `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCardStatusStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBody) SetCode(v string) *GetCardStatusStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBody) SetData(v *GetCardStatusStatisticsResponseBodyData) *GetCardStatusStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetCardStatusStatisticsResponseBody) SetErrorMessage(v string) *GetCardStatusStatisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBody) SetLocalizedMessage(v string) *GetCardStatusStatisticsResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBody) SetRequestId(v string) *GetCardStatusStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBody) SetSuccess(v bool) *GetCardStatusStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetCardStatusStatisticsResponseBodyData struct {
	ErrorStopStatisticsDTO        *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO        `json:"ErrorStopStatisticsDTO,omitempty" xml:"ErrorStopStatisticsDTO,omitempty" type:"Struct"`
	ExhaustStopStatisticsDTO      *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO      `json:"ExhaustStopStatisticsDTO,omitempty" xml:"ExhaustStopStatisticsDTO,omitempty" type:"Struct"`
	ExpireStopStatisticsDTO       *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO       `json:"ExpireStopStatisticsDTO,omitempty" xml:"ExpireStopStatisticsDTO,omitempty" type:"Struct"`
	FlowOutStatisticsDTO          *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO          `json:"FlowOutStatisticsDTO,omitempty" xml:"FlowOutStatisticsDTO,omitempty" type:"Struct"`
	ManageStopStatisticsDTO       *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO       `json:"ManageStopStatisticsDTO,omitempty" xml:"ManageStopStatisticsDTO,omitempty" type:"Struct"`
	RiskWaringStatisticsDTO       *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO       `json:"RiskWaringStatisticsDTO,omitempty" xml:"RiskWaringStatisticsDTO,omitempty" type:"Struct"`
	SingCardPeriodLeftFlowWarnDTO *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO `json:"SingCardPeriodLeftFlowWarnDTO,omitempty" xml:"SingCardPeriodLeftFlowWarnDTO,omitempty" type:"Struct"`
	UnCertifiedStopStatisticsDTO  *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO  `json:"UnCertifiedStopStatisticsDTO,omitempty" xml:"UnCertifiedStopStatisticsDTO,omitempty" type:"Struct"`
	UnbindResumeStatisticsDTO     *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO     `json:"UnbindResumeStatisticsDTO,omitempty" xml:"UnbindResumeStatisticsDTO,omitempty" type:"Struct"`
	WeekWarnStatisticsDTO         *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO         `json:"WeekWarnStatisticsDTO,omitempty" xml:"WeekWarnStatisticsDTO,omitempty" type:"Struct"`
}

func (s GetCardStatusStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyData) SetErrorStopStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.ErrorStopStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetExhaustStopStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.ExhaustStopStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetExpireStopStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.ExpireStopStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetFlowOutStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.FlowOutStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetManageStopStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.ManageStopStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetRiskWaringStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.RiskWaringStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetSingCardPeriodLeftFlowWarnDTO(v *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO) *GetCardStatusStatisticsResponseBodyData {
	s.SingCardPeriodLeftFlowWarnDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetUnCertifiedStopStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.UnCertifiedStopStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetUnbindResumeStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.UnbindResumeStatisticsDTO = v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyData) SetWeekWarnStatisticsDTO(v *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) *GetCardStatusStatisticsResponseBodyData {
	s.WeekWarnStatisticsDTO = v
	return s
}

type GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataErrorStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataExhaustStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataExpireStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataFlowOutStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataManageStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO struct {
	LeftFlowPercentageWarnCount *int64 `json:"LeftFlowPercentageWarnCount,omitempty" xml:"LeftFlowPercentageWarnCount,omitempty"`
	StopCount                   *int64 `json:"StopCount,omitempty" xml:"StopCount,omitempty"`
	WaringTotalCount            *int64 `json:"WaringTotalCount,omitempty" xml:"WaringTotalCount,omitempty"`
	WarningCount                *int64 `json:"WarningCount,omitempty" xml:"WarningCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetLeftFlowPercentageWarnCount(v int64) *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.LeftFlowPercentageWarnCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetStopCount(v int64) *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.StopCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetWaringTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.WaringTotalCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO) SetWarningCount(v int64) *GetCardStatusStatisticsResponseBodyDataRiskWaringStatisticsDTO {
	s.WarningCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO struct {
	LessFlowPercentage10Count *int64 `json:"LessFlowPercentage10Count,omitempty" xml:"LessFlowPercentage10Count,omitempty"`
	LessFlowPercentage30Count *int64 `json:"LessFlowPercentage30Count,omitempty" xml:"LessFlowPercentage30Count,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO) SetLessFlowPercentage10Count(v int64) *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO {
	s.LessFlowPercentage10Count = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO) SetLessFlowPercentage30Count(v int64) *GetCardStatusStatisticsResponseBodyDataSingCardPeriodLeftFlowWarnDTO {
	s.LessFlowPercentage30Count = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnCertifiedStopStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataUnbindResumeStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO struct {
	PoolCount         *int64 `json:"PoolCount,omitempty" xml:"PoolCount,omitempty"`
	SameFlowCardCount *int64 `json:"SameFlowCardCount,omitempty" xml:"SameFlowCardCount,omitempty"`
	SingleCardCount   *int64 `json:"SingleCardCount,omitempty" xml:"SingleCardCount,omitempty"`
	TotalCount        *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetPoolCount(v int64) *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.PoolCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetSameFlowCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.SameFlowCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetSingleCardCount(v int64) *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.SingleCardCount = &v
	return s
}

func (s *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO) SetTotalCount(v int64) *GetCardStatusStatisticsResponseBodyDataWeekWarnStatisticsDTO {
	s.TotalCount = &v
	return s
}

type GetCardStatusStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCardStatusStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCardStatusStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCardStatusStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCardStatusStatisticsResponse) SetHeaders(v map[string]*string) *GetCardStatusStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetCardStatusStatisticsResponse) SetStatusCode(v int32) *GetCardStatusStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCardStatusStatisticsResponse) SetBody(v *GetCardStatusStatisticsResponseBody) *GetCardStatusStatisticsResponse {
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
	MonthExceedFee         *int64  `json:"MonthExceedFee,omitempty" xml:"MonthExceedFee,omitempty"`
	MonthFeatureFee        *int64  `json:"MonthFeatureFee,omitempty" xml:"MonthFeatureFee,omitempty"`
	MonthUsedAmount        *int64  `json:"MonthUsedAmount,omitempty" xml:"MonthUsedAmount,omitempty"`
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

func (s *GetCredentialPoolStatisticsResponseBodyData) SetMonthExceedFee(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.MonthExceedFee = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetMonthFeatureFee(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.MonthFeatureFee = &v
	return s
}

func (s *GetCredentialPoolStatisticsResponseBodyData) SetMonthUsedAmount(v int64) *GetCredentialPoolStatisticsResponseBodyData {
	s.MonthUsedAmount = &v
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCredentialPoolStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetOperateResultRequest struct {
	ApiProduct *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	ResId      *string `json:"ResId,omitempty" xml:"ResId,omitempty"`
	SerialNo   *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s GetOperateResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperateResultRequest) GoString() string {
	return s.String()
}

func (s *GetOperateResultRequest) SetApiProduct(v string) *GetOperateResultRequest {
	s.ApiProduct = &v
	return s
}

func (s *GetOperateResultRequest) SetResId(v string) *GetOperateResultRequest {
	s.ResId = &v
	return s
}

func (s *GetOperateResultRequest) SetSerialNo(v string) *GetOperateResultRequest {
	s.SerialNo = &v
	return s
}

type GetOperateResultResponseBody struct {
	Code             *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetOperateResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                           `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperateResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperateResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperateResultResponseBody) SetCode(v string) *GetOperateResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperateResultResponseBody) SetData(v *GetOperateResultResponseBodyData) *GetOperateResultResponseBody {
	s.Data = v
	return s
}

func (s *GetOperateResultResponseBody) SetErrorMessage(v string) *GetOperateResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOperateResultResponseBody) SetLocalizedMessage(v string) *GetOperateResultResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetOperateResultResponseBody) SetRequestId(v string) *GetOperateResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperateResultResponseBody) SetSuccess(v bool) *GetOperateResultResponseBody {
	s.Success = &v
	return s
}

type GetOperateResultResponseBodyData struct {
	ExecuteResult *string `json:"ExecuteResult,omitempty" xml:"ExecuteResult,omitempty"`
	OperateType   *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Result        *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOperateResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOperateResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOperateResultResponseBodyData) SetExecuteResult(v string) *GetOperateResultResponseBodyData {
	s.ExecuteResult = &v
	return s
}

func (s *GetOperateResultResponseBodyData) SetOperateType(v string) *GetOperateResultResponseBodyData {
	s.OperateType = &v
	return s
}

func (s *GetOperateResultResponseBodyData) SetResult(v bool) *GetOperateResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetOperateResultResponseBodyData) SetStatus(v string) *GetOperateResultResponseBodyData {
	s.Status = &v
	return s
}

type GetOperateResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOperateResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOperateResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperateResultResponse) GoString() string {
	return s.String()
}

func (s *GetOperateResultResponse) SetHeaders(v map[string]*string) *GetOperateResultResponse {
	s.Headers = v
	return s
}

func (s *GetOperateResultResponse) SetStatusCode(v int32) *GetOperateResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOperateResultResponse) SetBody(v *GetOperateResultResponseBody) *GetOperateResultResponse {
	s.Body = v
	return s
}

type GetRealNameStatusRequest struct {
	Iccid       *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	ListMsisdns []*string `json:"ListMsisdns,omitempty" xml:"ListMsisdns,omitempty" type:"Repeated"`
}

func (s GetRealNameStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealNameStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRealNameStatusRequest) SetIccid(v string) *GetRealNameStatusRequest {
	s.Iccid = &v
	return s
}

func (s *GetRealNameStatusRequest) SetListMsisdns(v []*string) *GetRealNameStatusRequest {
	s.ListMsisdns = v
	return s
}

type GetRealNameStatusShrinkRequest struct {
	Iccid             *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	ListMsisdnsShrink *string `json:"ListMsisdns,omitempty" xml:"ListMsisdns,omitempty"`
}

func (s GetRealNameStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealNameStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRealNameStatusShrinkRequest) SetIccid(v string) *GetRealNameStatusShrinkRequest {
	s.Iccid = &v
	return s
}

func (s *GetRealNameStatusShrinkRequest) SetListMsisdnsShrink(v string) *GetRealNameStatusShrinkRequest {
	s.ListMsisdnsShrink = &v
	return s
}

type GetRealNameStatusResponseBody struct {
	Code             *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *GetRealNameStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                            `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRealNameStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealNameStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealNameStatusResponseBody) SetCode(v string) *GetRealNameStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealNameStatusResponseBody) SetData(v *GetRealNameStatusResponseBodyData) *GetRealNameStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetRealNameStatusResponseBody) SetErrorMessage(v string) *GetRealNameStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRealNameStatusResponseBody) SetLocalizedMessage(v string) *GetRealNameStatusResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *GetRealNameStatusResponseBody) SetRequestId(v string) *GetRealNameStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealNameStatusResponseBody) SetSuccess(v bool) *GetRealNameStatusResponseBody {
	s.Success = &v
	return s
}

type GetRealNameStatusResponseBodyData struct {
	Desc           *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	RealNameStatus *string `json:"RealNameStatus,omitempty" xml:"RealNameStatus,omitempty"`
}

func (s GetRealNameStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealNameStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealNameStatusResponseBodyData) SetDesc(v string) *GetRealNameStatusResponseBodyData {
	s.Desc = &v
	return s
}

func (s *GetRealNameStatusResponseBodyData) SetRealNameStatus(v string) *GetRealNameStatusResponseBodyData {
	s.RealNameStatus = &v
	return s
}

type GetRealNameStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealNameStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealNameStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealNameStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRealNameStatusResponse) SetHeaders(v map[string]*string) *GetRealNameStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRealNameStatusResponse) SetStatusCode(v int32) *GetRealNameStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealNameStatusResponse) SetBody(v *GetRealNameStatusResponseBody) *GetRealNameStatusResponse {
	s.Body = v
	return s
}

type GetSimCardStateDistributionRequest struct {
	CredentialNO *string `json:"CredentialNO,omitempty" xml:"CredentialNO,omitempty"`
	Date         *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetSimCardStateDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimCardStateDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetSimCardStateDistributionRequest) SetCredentialNO(v string) *GetSimCardStateDistributionRequest {
	s.CredentialNO = &v
	return s
}

func (s *GetSimCardStateDistributionRequest) SetDate(v string) *GetSimCardStateDistributionRequest {
	s.Date = &v
	return s
}

type GetSimCardStateDistributionResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GetSimCardStateDistributionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSimCardStateDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimCardStateDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimCardStateDistributionResponseBody) SetCode(v string) *GetSimCardStateDistributionResponseBody {
	s.Code = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBody) SetData(v *GetSimCardStateDistributionResponseBodyData) *GetSimCardStateDistributionResponseBody {
	s.Data = v
	return s
}

func (s *GetSimCardStateDistributionResponseBody) SetErrorMessage(v string) *GetSimCardStateDistributionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBody) SetRequestId(v string) *GetSimCardStateDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBody) SetSuccess(v bool) *GetSimCardStateDistributionResponseBody {
	s.Success = &v
	return s
}

type GetSimCardStateDistributionResponseBodyData struct {
	CardCount      *int64 `json:"CardCount,omitempty" xml:"CardCount,omitempty"`
	DestoryedCount *int64 `json:"DestoryedCount,omitempty" xml:"DestoryedCount,omitempty"`
	ShutDownCount  *int64 `json:"ShutDownCount,omitempty" xml:"ShutDownCount,omitempty"`
	StopCount      *int64 `json:"StopCount,omitempty" xml:"StopCount,omitempty"`
	TestCount      *int64 `json:"TestCount,omitempty" xml:"TestCount,omitempty"`
	UnusedCount    *int64 `json:"UnusedCount,omitempty" xml:"UnusedCount,omitempty"`
	UsingCount     *int64 `json:"UsingCount,omitempty" xml:"UsingCount,omitempty"`
}

func (s GetSimCardStateDistributionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSimCardStateDistributionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSimCardStateDistributionResponseBodyData) SetCardCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.CardCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetDestoryedCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.DestoryedCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetShutDownCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.ShutDownCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetStopCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.StopCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetTestCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.TestCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetUnusedCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.UnusedCount = &v
	return s
}

func (s *GetSimCardStateDistributionResponseBodyData) SetUsingCount(v int64) *GetSimCardStateDistributionResponseBodyData {
	s.UsingCount = &v
	return s
}

type GetSimCardStateDistributionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSimCardStateDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSimCardStateDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimCardStateDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetSimCardStateDistributionResponse) SetHeaders(v map[string]*string) *GetSimCardStateDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetSimCardStateDistributionResponse) SetStatusCode(v int32) *GetSimCardStateDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSimCardStateDistributionResponse) SetBody(v *GetSimCardStateDistributionResponseBody) *GetSimCardStateDistributionResponse {
	s.Body = v
	return s
}

type ListCardInfoRequest struct {
	ActiveTimeEnd         *string  `json:"ActiveTimeEnd,omitempty" xml:"ActiveTimeEnd,omitempty"`
	ActiveTimeStart       *string  `json:"ActiveTimeStart,omitempty" xml:"ActiveTimeStart,omitempty"`
	AliFee                *string  `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId         *string  `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName               *string  `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyType           *string  `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialNo          *string  `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	DataLevel             *string  `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType              *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DirectionalGroupId    *string  `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	ExpireTimeEnd         *string  `json:"ExpireTimeEnd,omitempty" xml:"ExpireTimeEnd,omitempty"`
	ExpireTimeStart       *string  `json:"ExpireTimeStart,omitempty" xml:"ExpireTimeStart,omitempty"`
	Iccid                 *string  `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                  *string  `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	IsAutoRecharge        *bool    `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	MaxFlow               *string  `json:"MaxFlow,omitempty" xml:"MaxFlow,omitempty"`
	MaxRestFlowPercentage *float64 `json:"MaxRestFlowPercentage,omitempty" xml:"MaxRestFlowPercentage,omitempty"`
	MinFlow               *string  `json:"MinFlow,omitempty" xml:"MinFlow,omitempty"`
	Msisdn                *string  `json:"Msisdn,omitempty" xml:"Msisdn,omitempty"`
	NetworkType           *string  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NotifyId              *string  `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OsStatus              *string  `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	PageNo                *int32   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize              *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Period                *string  `json:"Period,omitempty" xml:"Period,omitempty"`
	PoolId                *string  `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	SimType               *string  `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status                *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TagName               *string  `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Vendor                *string  `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
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

func (s *ListCardInfoRequest) SetMaxFlow(v string) *ListCardInfoRequest {
	s.MaxFlow = &v
	return s
}

func (s *ListCardInfoRequest) SetMaxRestFlowPercentage(v float64) *ListCardInfoRequest {
	s.MaxRestFlowPercentage = &v
	return s
}

func (s *ListCardInfoRequest) SetMinFlow(v string) *ListCardInfoRequest {
	s.MinFlow = &v
	return s
}

func (s *ListCardInfoRequest) SetMsisdn(v string) *ListCardInfoRequest {
	s.Msisdn = &v
	return s
}

func (s *ListCardInfoRequest) SetNetworkType(v string) *ListCardInfoRequest {
	s.NetworkType = &v
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
	ActiveTime             *string                                    `json:"ActiveTime,omitempty" xml:"ActiveTime,omitempty"`
	ActiveType             *string                                    `json:"ActiveType,omitempty" xml:"ActiveType,omitempty"`
	AliFee                 *string                                    `json:"AliFee,omitempty" xml:"AliFee,omitempty"`
	AliyunOrderId          *string                                    `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	ApnName                *string                                    `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	CertifyType            *string                                    `json:"CertifyType,omitempty" xml:"CertifyType,omitempty"`
	CredentialInstanceId   *string                                    `json:"CredentialInstanceId,omitempty" xml:"CredentialInstanceId,omitempty"`
	CredentialNo           *string                                    `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType         *string                                    `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	DataLevel              *string                                    `json:"DataLevel,omitempty" xml:"DataLevel,omitempty"`
	DataType               *string                                    `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DirectionalGroupId     *int64                                     `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	DirectionalGroupName   *string                                    `json:"DirectionalGroupName,omitempty" xml:"DirectionalGroupName,omitempty"`
	ExpireTime             *string                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FlowLatestModifiedTime *string                                    `json:"FlowLatestModifiedTime,omitempty" xml:"FlowLatestModifiedTime,omitempty"`
	Iccid                  *string                                    `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	Imsi                   []*string                                  `json:"Imsi,omitempty" xml:"Imsi,omitempty" type:"Repeated"`
	IsAutoRecharge         *bool                                      `json:"IsAutoRecharge,omitempty" xml:"IsAutoRecharge,omitempty"`
	Msisdn                 []*string                                  `json:"Msisdn,omitempty" xml:"Msisdn,omitempty" type:"Repeated"`
	NetworkType            *string                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NotifyId               *string                                    `json:"NotifyId,omitempty" xml:"NotifyId,omitempty"`
	OpenAccountTime        *string                                    `json:"OpenAccountTime,omitempty" xml:"OpenAccountTime,omitempty"`
	OsStatus               *string                                    `json:"OsStatus,omitempty" xml:"OsStatus,omitempty"`
	Period                 *string                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodAddFlow          *string                                    `json:"PeriodAddFlow,omitempty" xml:"PeriodAddFlow,omitempty"`
	PeriodRestFlow         *string                                    `json:"PeriodRestFlow,omitempty" xml:"PeriodRestFlow,omitempty"`
	PeriodSmsUse           *string                                    `json:"PeriodSmsUse,omitempty" xml:"PeriodSmsUse,omitempty"`
	PrivateNetworkSegment  *string                                    `json:"PrivateNetworkSegment,omitempty" xml:"PrivateNetworkSegment,omitempty"`
	Remark                 *string                                    `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SimType                *string                                    `json:"SimType,omitempty" xml:"SimType,omitempty"`
	Status                 *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	TagList                []*ListCardInfoResponseBodyDataListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	Vendor                 *string                                    `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	VsimInstanceId         *int64                                     `json:"VsimInstanceId,omitempty" xml:"VsimInstanceId,omitempty"`
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

func (s *ListCardInfoResponseBodyDataList) SetDirectionalGroupId(v int64) *ListCardInfoResponseBodyDataList {
	s.DirectionalGroupId = &v
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

func (s *ListCardInfoResponseBodyDataList) SetFlowLatestModifiedTime(v string) *ListCardInfoResponseBodyDataList {
	s.FlowLatestModifiedTime = &v
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

func (s *ListCardInfoResponseBodyDataList) SetNetworkType(v string) *ListCardInfoResponseBodyDataList {
	s.NetworkType = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListDirectionalAddressRequest struct {
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDirectionalAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalAddressRequest) GoString() string {
	return s.String()
}

func (s *ListDirectionalAddressRequest) SetGroupId(v string) *ListDirectionalAddressRequest {
	s.GroupId = &v
	return s
}

func (s *ListDirectionalAddressRequest) SetPageNo(v int32) *ListDirectionalAddressRequest {
	s.PageNo = &v
	return s
}

func (s *ListDirectionalAddressRequest) SetPageSize(v int32) *ListDirectionalAddressRequest {
	s.PageSize = &v
	return s
}

type ListDirectionalAddressResponseBody struct {
	Code             *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *ListDirectionalAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                                 `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDirectionalAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectionalAddressResponseBody) SetCode(v string) *ListDirectionalAddressResponseBody {
	s.Code = &v
	return s
}

func (s *ListDirectionalAddressResponseBody) SetData(v *ListDirectionalAddressResponseBodyData) *ListDirectionalAddressResponseBody {
	s.Data = v
	return s
}

func (s *ListDirectionalAddressResponseBody) SetErrorMessage(v string) *ListDirectionalAddressResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDirectionalAddressResponseBody) SetLocalizedMessage(v string) *ListDirectionalAddressResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ListDirectionalAddressResponseBody) SetRequestId(v string) *ListDirectionalAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectionalAddressResponseBody) SetSuccess(v bool) *ListDirectionalAddressResponseBody {
	s.Success = &v
	return s
}

type ListDirectionalAddressResponseBodyData struct {
	List      []*ListDirectionalAddressResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageCount *int32                                        `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNo    *int32                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDirectionalAddressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDirectionalAddressResponseBodyData) SetList(v []*ListDirectionalAddressResponseBodyDataList) *ListDirectionalAddressResponseBodyData {
	s.List = v
	return s
}

func (s *ListDirectionalAddressResponseBodyData) SetPageCount(v int32) *ListDirectionalAddressResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyData) SetPageNo(v int32) *ListDirectionalAddressResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyData) SetPageSize(v int32) *ListDirectionalAddressResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyData) SetTotal(v int32) *ListDirectionalAddressResponseBodyData {
	s.Total = &v
	return s
}

type ListDirectionalAddressResponseBodyDataList struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	State       *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDirectionalAddressResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalAddressResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDirectionalAddressResponseBodyDataList) SetAddress(v string) *ListDirectionalAddressResponseBodyDataList {
	s.Address = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyDataList) SetAddressType(v string) *ListDirectionalAddressResponseBodyDataList {
	s.AddressType = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyDataList) SetGroupId(v string) *ListDirectionalAddressResponseBodyDataList {
	s.GroupId = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyDataList) SetSource(v string) *ListDirectionalAddressResponseBodyDataList {
	s.Source = &v
	return s
}

func (s *ListDirectionalAddressResponseBodyDataList) SetState(v int32) *ListDirectionalAddressResponseBodyDataList {
	s.State = &v
	return s
}

type ListDirectionalAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectionalAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectionalAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalAddressResponse) GoString() string {
	return s.String()
}

func (s *ListDirectionalAddressResponse) SetHeaders(v map[string]*string) *ListDirectionalAddressResponse {
	s.Headers = v
	return s
}

func (s *ListDirectionalAddressResponse) SetStatusCode(v int32) *ListDirectionalAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectionalAddressResponse) SetBody(v *ListDirectionalAddressResponseBody) *ListDirectionalAddressResponse {
	s.Body = v
	return s
}

type ListDirectionalDetailRequest struct {
	Iccid    *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDirectionalDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailRequest) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailRequest) SetIccid(v string) *ListDirectionalDetailRequest {
	s.Iccid = &v
	return s
}

func (s *ListDirectionalDetailRequest) SetPageNo(v int32) *ListDirectionalDetailRequest {
	s.PageNo = &v
	return s
}

func (s *ListDirectionalDetailRequest) SetPageSize(v int32) *ListDirectionalDetailRequest {
	s.PageSize = &v
	return s
}

type ListDirectionalDetailResponseBody struct {
	Code             *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *ListDirectionalDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage     *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string                                `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDirectionalDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailResponseBody) SetCode(v string) *ListDirectionalDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListDirectionalDetailResponseBody) SetData(v *ListDirectionalDetailResponseBodyData) *ListDirectionalDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListDirectionalDetailResponseBody) SetErrorMessage(v string) *ListDirectionalDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDirectionalDetailResponseBody) SetLocalizedMessage(v string) *ListDirectionalDetailResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *ListDirectionalDetailResponseBody) SetRequestId(v string) *ListDirectionalDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectionalDetailResponseBody) SetSuccess(v bool) *ListDirectionalDetailResponseBody {
	s.Success = &v
	return s
}

type ListDirectionalDetailResponseBodyData struct {
	DirectionalGroupId *int64                                                 `json:"DirectionalGroupId,omitempty" xml:"DirectionalGroupId,omitempty"`
	DirectionalName    *string                                                `json:"DirectionalName,omitempty" xml:"DirectionalName,omitempty"`
	PaginationResult   *ListDirectionalDetailResponseBodyDataPaginationResult `json:"PaginationResult,omitempty" xml:"PaginationResult,omitempty" type:"Struct"`
}

func (s ListDirectionalDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailResponseBodyData) SetDirectionalGroupId(v int64) *ListDirectionalDetailResponseBodyData {
	s.DirectionalGroupId = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyData) SetDirectionalName(v string) *ListDirectionalDetailResponseBodyData {
	s.DirectionalName = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyData) SetPaginationResult(v *ListDirectionalDetailResponseBodyDataPaginationResult) *ListDirectionalDetailResponseBodyData {
	s.PaginationResult = v
	return s
}

type ListDirectionalDetailResponseBodyDataPaginationResult struct {
	List      []*ListDirectionalDetailResponseBodyDataPaginationResultList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageCount *int32                                                       `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNo    *int32                                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDirectionalDetailResponseBodyDataPaginationResult) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailResponseBodyDataPaginationResult) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResult) SetList(v []*ListDirectionalDetailResponseBodyDataPaginationResultList) *ListDirectionalDetailResponseBodyDataPaginationResult {
	s.List = v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResult) SetPageCount(v int32) *ListDirectionalDetailResponseBodyDataPaginationResult {
	s.PageCount = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResult) SetPageNo(v int32) *ListDirectionalDetailResponseBodyDataPaginationResult {
	s.PageNo = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResult) SetPageSize(v int32) *ListDirectionalDetailResponseBodyDataPaginationResult {
	s.PageSize = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResult) SetTotal(v int32) *ListDirectionalDetailResponseBodyDataPaginationResult {
	s.Total = &v
	return s
}

type ListDirectionalDetailResponseBodyDataPaginationResultList struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDirectionalDetailResponseBodyDataPaginationResultList) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailResponseBodyDataPaginationResultList) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResultList) SetAddress(v string) *ListDirectionalDetailResponseBodyDataPaginationResultList {
	s.Address = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResultList) SetAddressType(v string) *ListDirectionalDetailResponseBodyDataPaginationResultList {
	s.AddressType = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResultList) SetGroupId(v string) *ListDirectionalDetailResponseBodyDataPaginationResultList {
	s.GroupId = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResultList) SetSource(v string) *ListDirectionalDetailResponseBodyDataPaginationResultList {
	s.Source = &v
	return s
}

func (s *ListDirectionalDetailResponseBodyDataPaginationResultList) SetState(v string) *ListDirectionalDetailResponseBodyDataPaginationResultList {
	s.State = &v
	return s
}

type ListDirectionalDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDirectionalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDirectionalDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDirectionalDetailResponse) GoString() string {
	return s.String()
}

func (s *ListDirectionalDetailResponse) SetHeaders(v map[string]*string) *ListDirectionalDetailResponse {
	s.Headers = v
	return s
}

func (s *ListDirectionalDetailResponse) SetStatusCode(v int32) *ListDirectionalDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDirectionalDetailResponse) SetBody(v *ListDirectionalDetailResponseBody) *ListDirectionalDetailResponse {
	s.Body = v
	return s
}

type ListOrderRequest struct {
	CredentialNo *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus  *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrderRequest) GoString() string {
	return s.String()
}

func (s *ListOrderRequest) SetCredentialNo(v string) *ListOrderRequest {
	s.CredentialNo = &v
	return s
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
	ApnName           *string                                    `json:"ApnName,omitempty" xml:"ApnName,omitempty"`
	ApnRegion         *string                                    `json:"ApnRegion,omitempty" xml:"ApnRegion,omitempty"`
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
	NetworkType       *string                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
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
	ResourceQuantity  *int64                                     `json:"ResourceQuantity,omitempty" xml:"ResourceQuantity,omitempty"`
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

func (s *ListOrderResponseBodyDataList) SetApnName(v string) *ListOrderResponseBodyDataList {
	s.ApnName = &v
	return s
}

func (s *ListOrderResponseBodyDataList) SetApnRegion(v string) *ListOrderResponseBodyDataList {
	s.ApnRegion = &v
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

func (s *ListOrderResponseBodyDataList) SetNetworkType(v string) *ListOrderResponseBodyDataList {
	s.NetworkType = &v
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

func (s *ListOrderResponseBodyDataList) SetResourceQuantity(v int64) *ListOrderResponseBodyDataList {
	s.ResourceQuantity = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Iccid      *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns []*string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty" type:"Repeated"`
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

func (s *RebindResumeSingleCardRequest) SetOptMsisdns(v []*string) *RebindResumeSingleCardRequest {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebindResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	RequestId        *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Iccid      *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns []*string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty" type:"Repeated"`
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

func (s *ResumeSingleCardRequest) SetOptMsisdns(v []*string) *ResumeSingleCardRequest {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SendMessageRequest struct {
	// Linkcard
	ApiProduct           *string   `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	MessageSendTime      *int64    `json:"MessageSendTime,omitempty" xml:"MessageSendTime,omitempty"`
	MessageTemplateId    *int64    `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	MessageVariableParam *string   `json:"MessageVariableParam,omitempty" xml:"MessageVariableParam,omitempty"`
	Msisdns              []*string `json:"Msisdns,omitempty" xml:"Msisdns,omitempty" type:"Repeated"`
	TaskName             *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetApiProduct(v string) *SendMessageRequest {
	s.ApiProduct = &v
	return s
}

func (s *SendMessageRequest) SetMessageSendTime(v int64) *SendMessageRequest {
	s.MessageSendTime = &v
	return s
}

func (s *SendMessageRequest) SetMessageTemplateId(v int64) *SendMessageRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *SendMessageRequest) SetMessageVariableParam(v string) *SendMessageRequest {
	s.MessageVariableParam = &v
	return s
}

func (s *SendMessageRequest) SetMsisdns(v []*string) *SendMessageRequest {
	s.Msisdns = v
	return s
}

func (s *SendMessageRequest) SetTaskName(v string) *SendMessageRequest {
	s.TaskName = &v
	return s
}

type SendMessageShrinkRequest struct {
	// Linkcard
	ApiProduct           *string `json:"ApiProduct,omitempty" xml:"ApiProduct,omitempty"`
	MessageSendTime      *int64  `json:"MessageSendTime,omitempty" xml:"MessageSendTime,omitempty"`
	MessageTemplateId    *int64  `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	MessageVariableParam *string `json:"MessageVariableParam,omitempty" xml:"MessageVariableParam,omitempty"`
	MsisdnsShrink        *string `json:"Msisdns,omitempty" xml:"Msisdns,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SendMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageShrinkRequest) SetApiProduct(v string) *SendMessageShrinkRequest {
	s.ApiProduct = &v
	return s
}

func (s *SendMessageShrinkRequest) SetMessageSendTime(v int64) *SendMessageShrinkRequest {
	s.MessageSendTime = &v
	return s
}

func (s *SendMessageShrinkRequest) SetMessageTemplateId(v int64) *SendMessageShrinkRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *SendMessageShrinkRequest) SetMessageVariableParam(v string) *SendMessageShrinkRequest {
	s.MessageVariableParam = &v
	return s
}

func (s *SendMessageShrinkRequest) SetMsisdnsShrink(v string) *SendMessageShrinkRequest {
	s.MsisdnsShrink = &v
	return s
}

func (s *SendMessageShrinkRequest) SetTaskName(v string) *SendMessageShrinkRequest {
	s.TaskName = &v
	return s
}

type SendMessageResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	DynamicCode      *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage   *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCode(v string) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetData(v int64) *SendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *SendMessageResponseBody) SetDynamicCode(v string) *SendMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SendMessageResponseBody) SetDynamicMessage(v string) *SendMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SendMessageResponseBody) SetErrorMessage(v string) *SendMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendMessageResponseBody) SetLocalizedMessage(v string) *SendMessageResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBody) SetSuccess(v bool) *SendMessageResponseBody {
	s.Success = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCardStopRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Iccid      *string   `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
	OptMsisdns []*string `json:"OptMsisdns,omitempty" xml:"OptMsisdns,omitempty" type:"Repeated"`
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

func (s *StopSingleCardRequest) SetOptMsisdns(v []*string) *StopSingleCardRequest {
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSingleCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoRechargeSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type VerifyIotCardRequest struct {
	Iccid *string `json:"Iccid,omitempty" xml:"Iccid,omitempty"`
}

func (s VerifyIotCardRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyIotCardRequest) GoString() string {
	return s.String()
}

func (s *VerifyIotCardRequest) SetIccid(v string) *VerifyIotCardRequest {
	s.Iccid = &v
	return s
}

type VerifyIotCardResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data             *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorMessage     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LocalizedMessage *string `json:"LocalizedMessage,omitempty" xml:"LocalizedMessage,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyIotCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyIotCardResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyIotCardResponseBody) SetCode(v string) *VerifyIotCardResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyIotCardResponseBody) SetData(v bool) *VerifyIotCardResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyIotCardResponseBody) SetErrorMessage(v string) *VerifyIotCardResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *VerifyIotCardResponseBody) SetLocalizedMessage(v string) *VerifyIotCardResponseBody {
	s.LocalizedMessage = &v
	return s
}

func (s *VerifyIotCardResponseBody) SetRequestId(v string) *VerifyIotCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyIotCardResponseBody) SetSuccess(v bool) *VerifyIotCardResponseBody {
	s.Success = &v
	return s
}

type VerifyIotCardResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyIotCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyIotCardResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyIotCardResponse) GoString() string {
	return s.String()
}

func (s *VerifyIotCardResponse) SetHeaders(v map[string]*string) *VerifyIotCardResponse {
	s.Headers = v
	return s
}

func (s *VerifyIotCardResponse) SetStatusCode(v int32) *VerifyIotCardResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyIotCardResponse) SetBody(v *VerifyIotCardResponseBody) *VerifyIotCardResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
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

func (client *Client) AddCardToDirectionalGroupWithOptions(tmpReq *AddCardToDirectionalGroupRequest, runtime *util.RuntimeOptions) (_result *AddCardToDirectionalGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddCardToDirectionalGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IccidList)) {
		request.IccidListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IccidList, tea.String("IccidList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddType)) {
		query["AddType"] = request.AddType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IccidListShrink)) {
		query["IccidList"] = request.IccidListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MsgNotify)) {
		query["MsgNotify"] = request.MsgNotify
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiProduct)) {
		body["ApiProduct"] = request.ApiProduct
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCardToDirectionalGroup"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCardToDirectionalGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCardToDirectionalGroup(request *AddCardToDirectionalGroupRequest) (_result *AddCardToDirectionalGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCardToDirectionalGroupResponse{}
	_body, _err := client.AddCardToDirectionalGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDirectionalAddressWithOptions(request *AddDirectionalAddressRequest, runtime *util.RuntimeOptions) (_result *AddDirectionalAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgNotify)) {
		query["MsgNotify"] = request.MsgNotify
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.UrlInsecurityForce)) {
		query["UrlInsecurityForce"] = request.UrlInsecurityForce
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDirectionalAddress"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDirectionalAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDirectionalAddress(request *AddDirectionalAddressRequest) (_result *AddDirectionalAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDirectionalAddressResponse{}
	_body, _err := client.AddDirectionalAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDirectionalCardWithOptions(tmpReq *AddDirectionalCardRequest, runtime *util.RuntimeOptions) (_result *AddDirectionalCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddDirectionalCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderList)) {
		request.OrderListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderList, tea.String("OrderList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagList)) {
		request.TagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagList, tea.String("TagList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileUri)) {
		query["FileUri"] = request.FileUri
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderListShrink)) {
		query["OrderList"] = request.OrderListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagListShrink)) {
		query["TagList"] = request.TagListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UploadMethod)) {
		query["UploadMethod"] = request.UploadMethod
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		query["UploadType"] = request.UploadType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDirectionalCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDirectionalCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDirectionalCard(request *AddDirectionalCardRequest) (_result *AddDirectionalCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDirectionalCardResponse{}
	_body, _err := client.AddDirectionalCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddDirectionalGroupWithOptions(request *AddDirectionalGroupRequest, runtime *util.RuntimeOptions) (_result *AddDirectionalGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDirectionalGroup"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDirectionalGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddDirectionalGroup(request *AddDirectionalGroupRequest) (_result *AddDirectionalGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddDirectionalGroupResponse{}
	_body, _err := client.AddDirectionalGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTagsToCardWithOptions(tmpReq *AddTagsToCardRequest, runtime *util.RuntimeOptions) (_result *AddTagsToCardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddTagsToCardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TagNameList)) {
		request.TagNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagNameList, tea.String("TagNameList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.TagNameListShrink)) {
		query["TagNameList"] = request.TagNameListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTagsToCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTagsToCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTagsToCard(request *AddTagsToCardRequest) (_result *AddTagsToCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsToCardResponse{}
	_body, _err := client.AddTagsToCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAddDirectionalAddressWithOptions(request *BatchAddDirectionalAddressRequest, runtime *util.RuntimeOptions) (_result *BatchAddDirectionalAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressType)) {
		query["AddressType"] = request.AddressType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ListAddress)) {
		query["ListAddress"] = request.ListAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAddDirectionalAddress"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddDirectionalAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAddDirectionalAddress(request *BatchAddDirectionalAddressRequest) (_result *BatchAddDirectionalAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAddDirectionalAddressResponse{}
	_body, _err := client.BatchAddDirectionalAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectionalAddressWithOptions(request *DeleteDirectionalAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectionalAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgNotify)) {
		query["MsgNotify"] = request.MsgNotify
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectionalAddress"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectionalAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectionalAddress(request *DeleteDirectionalAddressRequest) (_result *DeleteDirectionalAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectionalAddressResponse{}
	_body, _err := client.DeleteDirectionalAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDirectionalGroupWithOptions(request *DeleteDirectionalGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDirectionalGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDirectionalGroup"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirectionalGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDirectionalGroup(request *DeleteDirectionalGroupRequest) (_result *DeleteDirectionalGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDirectionalGroupResponse{}
	_body, _err := client.DeleteDirectionalGroupWithOptions(request, runtime)
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

func (client *Client) GetCardLatestFlowWithOptions(request *GetCardLatestFlowRequest, runtime *util.RuntimeOptions) (_result *GetCardLatestFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardLatestFlow"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardLatestFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardLatestFlow(request *GetCardLatestFlowRequest) (_result *GetCardLatestFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardLatestFlowResponse{}
	_body, _err := client.GetCardLatestFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardRealStatusWithOptions(tmpReq *GetCardRealStatusRequest, runtime *util.RuntimeOptions) (_result *GetCardRealStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetCardRealStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SerialNo)) {
		request.SerialNoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SerialNo, tea.String("SerialNo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.Msisdn)) {
		query["Msisdn"] = request.Msisdn
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNoShrink)) {
		query["SerialNo"] = request.SerialNoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCardRealStatus"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardRealStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardRealStatus(request *GetCardRealStatusRequest) (_result *GetCardRealStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardRealStatusResponse{}
	_body, _err := client.GetCardRealStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCardStatusStatisticsWithOptions(runtime *util.RuntimeOptions) (_result *GetCardStatusStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetCardStatusStatistics"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCardStatusStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCardStatusStatistics() (_result *GetCardStatusStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCardStatusStatisticsResponse{}
	_body, _err := client.GetCardStatusStatisticsWithOptions(runtime)
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

func (client *Client) GetOperateResultWithOptions(request *GetOperateResultRequest, runtime *util.RuntimeOptions) (_result *GetOperateResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiProduct)) {
		query["ApiProduct"] = request.ApiProduct
	}

	if !tea.BoolValue(util.IsUnset(request.ResId)) {
		query["ResId"] = request.ResId
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNo)) {
		query["SerialNo"] = request.SerialNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOperateResult"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOperateResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOperateResult(request *GetOperateResultRequest) (_result *GetOperateResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOperateResultResponse{}
	_body, _err := client.GetOperateResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealNameStatusWithOptions(tmpReq *GetRealNameStatusRequest, runtime *util.RuntimeOptions) (_result *GetRealNameStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetRealNameStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ListMsisdns)) {
		request.ListMsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListMsisdns, tea.String("ListMsisdns"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.ListMsisdnsShrink)) {
		query["ListMsisdns"] = request.ListMsisdnsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealNameStatus"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealNameStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealNameStatus(request *GetRealNameStatusRequest) (_result *GetRealNameStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealNameStatusResponse{}
	_body, _err := client.GetRealNameStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimCardStateDistributionWithOptions(request *GetSimCardStateDistributionRequest, runtime *util.RuntimeOptions) (_result *GetSimCardStateDistributionResponse, _err error) {
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
		Action:      tea.String("GetSimCardStateDistribution"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimCardStateDistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimCardStateDistribution(request *GetSimCardStateDistributionRequest) (_result *GetSimCardStateDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSimCardStateDistributionResponse{}
	_body, _err := client.GetSimCardStateDistributionWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.MaxFlow)) {
		query["MaxFlow"] = request.MaxFlow
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRestFlowPercentage)) {
		query["MaxRestFlowPercentage"] = request.MaxRestFlowPercentage
	}

	if !tea.BoolValue(util.IsUnset(request.MinFlow)) {
		query["MinFlow"] = request.MinFlow
	}

	if !tea.BoolValue(util.IsUnset(request.Msisdn)) {
		query["Msisdn"] = request.Msisdn
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["NetworkType"] = request.NetworkType
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

func (client *Client) ListDirectionalAddressWithOptions(request *ListDirectionalAddressRequest, runtime *util.RuntimeOptions) (_result *ListDirectionalAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDirectionalAddress"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectionalAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectionalAddress(request *ListDirectionalAddressRequest) (_result *ListDirectionalAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectionalAddressResponse{}
	_body, _err := client.ListDirectionalAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDirectionalDetailWithOptions(request *ListDirectionalDetailRequest, runtime *util.RuntimeOptions) (_result *ListDirectionalDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDirectionalDetail"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDirectionalDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDirectionalDetail(request *ListDirectionalDetailRequest) (_result *ListDirectionalDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDirectionalDetailResponse{}
	_body, _err := client.ListDirectionalDetailWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.CredentialNo)) {
		query["CredentialNo"] = request.CredentialNo
	}

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

func (client *Client) SendMessageWithOptions(tmpReq *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Msisdns)) {
		request.MsisdnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Msisdns, tea.String("Msisdns"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiProduct)) {
		body["ApiProduct"] = request.ApiProduct
	}

	if !tea.BoolValue(util.IsUnset(request.MessageSendTime)) {
		body["MessageSendTime"] = request.MessageSendTime
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTemplateId)) {
		body["MessageTemplateId"] = request.MessageTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageVariableParam)) {
		body["MessageVariableParam"] = request.MessageVariableParam
	}

	if !tea.BoolValue(util.IsUnset(request.MsisdnsShrink)) {
		body["Msisdns"] = request.MsisdnsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
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

func (client *Client) VerifyIotCardWithOptions(request *VerifyIotCardRequest, runtime *util.RuntimeOptions) (_result *VerifyIotCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Iccid)) {
		query["Iccid"] = request.Iccid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyIotCard"),
		Version:     tea.String("2021-05-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyIotCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyIotCard(request *VerifyIotCardRequest) (_result *VerifyIotCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyIotCardResponse{}
	_body, _err := client.VerifyIotCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
