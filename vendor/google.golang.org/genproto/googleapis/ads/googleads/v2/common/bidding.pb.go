// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/bidding.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Commission is an automatic bidding strategy in which the advertiser pays a
// certain portion of the conversion value.
type Commission struct {
	// Commission rate defines the portion of the conversion value that the
	// advertiser will be billed. A commission rate of x should be passed into
	// this field as (x * 1,000,000). For example, 106,000 represents a commission
	// rate of 0.106 (10.6%).
	CommissionRateMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=commission_rate_micros,json=commissionRateMicros,proto3" json:"commission_rate_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Commission) Reset()         { *m = Commission{} }
func (m *Commission) String() string { return proto.CompactTextString(m) }
func (*Commission) ProtoMessage()    {}
func (*Commission) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{0}
}

func (m *Commission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commission.Unmarshal(m, b)
}
func (m *Commission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commission.Marshal(b, m, deterministic)
}
func (m *Commission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commission.Merge(m, src)
}
func (m *Commission) XXX_Size() int {
	return xxx_messageInfo_Commission.Size(m)
}
func (m *Commission) XXX_DiscardUnknown() {
	xxx_messageInfo_Commission.DiscardUnknown(m)
}

var xxx_messageInfo_Commission proto.InternalMessageInfo

func (m *Commission) GetCommissionRateMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CommissionRateMicros
	}
	return nil
}

// An automated bidding strategy that raises bids for clicks
// that seem more likely to lead to a conversion and lowers
// them for clicks where they seem less likely.
type EnhancedCpc struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnhancedCpc) Reset()         { *m = EnhancedCpc{} }
func (m *EnhancedCpc) String() string { return proto.CompactTextString(m) }
func (*EnhancedCpc) ProtoMessage()    {}
func (*EnhancedCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{1}
}

func (m *EnhancedCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnhancedCpc.Unmarshal(m, b)
}
func (m *EnhancedCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnhancedCpc.Marshal(b, m, deterministic)
}
func (m *EnhancedCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnhancedCpc.Merge(m, src)
}
func (m *EnhancedCpc) XXX_Size() int {
	return xxx_messageInfo_EnhancedCpc.Size(m)
}
func (m *EnhancedCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_EnhancedCpc.DiscardUnknown(m)
}

var xxx_messageInfo_EnhancedCpc proto.InternalMessageInfo

// Manual click-based bidding where user pays per click.
type ManualCpc struct {
	// Whether bids are to be enhanced based on conversion optimizer data.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ManualCpc) Reset()         { *m = ManualCpc{} }
func (m *ManualCpc) String() string { return proto.CompactTextString(m) }
func (*ManualCpc) ProtoMessage()    {}
func (*ManualCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{2}
}

func (m *ManualCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpc.Unmarshal(m, b)
}
func (m *ManualCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpc.Marshal(b, m, deterministic)
}
func (m *ManualCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpc.Merge(m, src)
}
func (m *ManualCpc) XXX_Size() int {
	return xxx_messageInfo_ManualCpc.Size(m)
}
func (m *ManualCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpc.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpc proto.InternalMessageInfo

func (m *ManualCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

// Manual impression-based bidding where user pays per thousand impressions.
type ManualCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpm) Reset()         { *m = ManualCpm{} }
func (m *ManualCpm) String() string { return proto.CompactTextString(m) }
func (*ManualCpm) ProtoMessage()    {}
func (*ManualCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{3}
}

func (m *ManualCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpm.Unmarshal(m, b)
}
func (m *ManualCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpm.Marshal(b, m, deterministic)
}
func (m *ManualCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpm.Merge(m, src)
}
func (m *ManualCpm) XXX_Size() int {
	return xxx_messageInfo_ManualCpm.Size(m)
}
func (m *ManualCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpm.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpm proto.InternalMessageInfo

// View based bidding where user pays per video view.
type ManualCpv struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualCpv) Reset()         { *m = ManualCpv{} }
func (m *ManualCpv) String() string { return proto.CompactTextString(m) }
func (*ManualCpv) ProtoMessage()    {}
func (*ManualCpv) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{4}
}

func (m *ManualCpv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualCpv.Unmarshal(m, b)
}
func (m *ManualCpv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualCpv.Marshal(b, m, deterministic)
}
func (m *ManualCpv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualCpv.Merge(m, src)
}
func (m *ManualCpv) XXX_Size() int {
	return xxx_messageInfo_ManualCpv.Size(m)
}
func (m *ManualCpv) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualCpv.DiscardUnknown(m)
}

var xxx_messageInfo_ManualCpv proto.InternalMessageInfo

// An automated bidding strategy that sets bids to help get the most conversions
// for your campaign while spending your budget.
type MaximizeConversions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaximizeConversions) Reset()         { *m = MaximizeConversions{} }
func (m *MaximizeConversions) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversions) ProtoMessage()    {}
func (*MaximizeConversions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{5}
}

func (m *MaximizeConversions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversions.Unmarshal(m, b)
}
func (m *MaximizeConversions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversions.Marshal(b, m, deterministic)
}
func (m *MaximizeConversions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversions.Merge(m, src)
}
func (m *MaximizeConversions) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversions.Size(m)
}
func (m *MaximizeConversions) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversions.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversions proto.InternalMessageInfo

// An automated bidding strategy which tries to maximize conversion value
// given a daily budget.
type MaximizeConversionValue struct {
	// The target return on ad spend (ROAS) option. If set, the bid strategy will
	// maximize revenue while averaging the target return on ad spend. If the
	// target ROAS is high, the bid strategy may not be able to spend the full
	// budget. If the target ROAS is not set, the bid strategy will aim to
	// achieve the highest possible ROAS for the budget.
	TargetRoas           *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MaximizeConversionValue) Reset()         { *m = MaximizeConversionValue{} }
func (m *MaximizeConversionValue) String() string { return proto.CompactTextString(m) }
func (*MaximizeConversionValue) ProtoMessage()    {}
func (*MaximizeConversionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{6}
}

func (m *MaximizeConversionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaximizeConversionValue.Unmarshal(m, b)
}
func (m *MaximizeConversionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaximizeConversionValue.Marshal(b, m, deterministic)
}
func (m *MaximizeConversionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaximizeConversionValue.Merge(m, src)
}
func (m *MaximizeConversionValue) XXX_Size() int {
	return xxx_messageInfo_MaximizeConversionValue.Size(m)
}
func (m *MaximizeConversionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MaximizeConversionValue.DiscardUnknown(m)
}

var xxx_messageInfo_MaximizeConversionValue proto.InternalMessageInfo

func (m *MaximizeConversionValue) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

// An automated bidding strategy which sets CPC bids to target impressions on
// page one, or page one promoted slots on google.com.
// This strategy is deprecated.
type PageOnePromoted struct {
	// The strategy goal of where impressions are desired to be shown on
	// search result pages.
	StrategyGoal enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal `protobuf:"varint,1,opt,name=strategy_goal,json=strategyGoal,proto3,enum=google.ads.googleads.v2.enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal" json:"strategy_goal,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Bid multiplier to be applied to the relevant bid estimate (depending on
	// the `strategy_goal`) in determining a keyword's new CPC bid.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether the strategy should always follow bid estimate changes, or only
	// increase.
	// If false, always sets a keyword's new bid to the current bid estimate.
	// If true, only updates a keyword's bid if the current bid estimate is
	// greater than the current bid.
	OnlyRaiseCpcBids *wrappers.BoolValue `protobuf:"bytes,4,opt,name=only_raise_cpc_bids,json=onlyRaiseCpcBids,proto3" json:"only_raise_cpc_bids,omitempty"`
	// Whether the strategy is allowed to raise bids when the throttling
	// rate of the budget it is serving out of rises above a threshold.
	RaiseCpcBidWhenBudgetConstrained *wrappers.BoolValue `protobuf:"bytes,5,opt,name=raise_cpc_bid_when_budget_constrained,json=raiseCpcBidWhenBudgetConstrained,proto3" json:"raise_cpc_bid_when_budget_constrained,omitempty"`
	// Whether the strategy is allowed to raise bids on keywords with
	// lower-range quality scores.
	RaiseCpcBidWhenQualityScoreIsLow *wrappers.BoolValue `protobuf:"bytes,6,opt,name=raise_cpc_bid_when_quality_score_is_low,json=raiseCpcBidWhenQualityScoreIsLow,proto3" json:"raise_cpc_bid_when_quality_score_is_low,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}            `json:"-"`
	XXX_unrecognized                 []byte              `json:"-"`
	XXX_sizecache                    int32               `json:"-"`
}

func (m *PageOnePromoted) Reset()         { *m = PageOnePromoted{} }
func (m *PageOnePromoted) String() string { return proto.CompactTextString(m) }
func (*PageOnePromoted) ProtoMessage()    {}
func (*PageOnePromoted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{7}
}

func (m *PageOnePromoted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageOnePromoted.Unmarshal(m, b)
}
func (m *PageOnePromoted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageOnePromoted.Marshal(b, m, deterministic)
}
func (m *PageOnePromoted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageOnePromoted.Merge(m, src)
}
func (m *PageOnePromoted) XXX_Size() int {
	return xxx_messageInfo_PageOnePromoted.Size(m)
}
func (m *PageOnePromoted) XXX_DiscardUnknown() {
	xxx_messageInfo_PageOnePromoted.DiscardUnknown(m)
}

var xxx_messageInfo_PageOnePromoted proto.InternalMessageInfo

func (m *PageOnePromoted) GetStrategyGoal() enums.PageOnePromotedStrategyGoalEnum_PageOnePromotedStrategyGoal {
	if m != nil {
		return m.StrategyGoal
	}
	return enums.PageOnePromotedStrategyGoalEnum_UNSPECIFIED
}

func (m *PageOnePromoted) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *PageOnePromoted) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *PageOnePromoted) GetOnlyRaiseCpcBids() *wrappers.BoolValue {
	if m != nil {
		return m.OnlyRaiseCpcBids
	}
	return nil
}

func (m *PageOnePromoted) GetRaiseCpcBidWhenBudgetConstrained() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenBudgetConstrained
	}
	return nil
}

func (m *PageOnePromoted) GetRaiseCpcBidWhenQualityScoreIsLow() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenQualityScoreIsLow
	}
	return nil
}

// An automated bid strategy that sets bids to help get as many conversions as
// possible at the target cost-per-acquisition (CPA) you set.
type TargetCpa struct {
	// Average CPA target.
	// This target should be greater than or equal to minimum billable unit based
	// on the currency for the account.
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetCpa) Reset()         { *m = TargetCpa{} }
func (m *TargetCpa) String() string { return proto.CompactTextString(m) }
func (*TargetCpa) ProtoMessage()    {}
func (*TargetCpa) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{8}
}

func (m *TargetCpa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpa.Unmarshal(m, b)
}
func (m *TargetCpa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpa.Marshal(b, m, deterministic)
}
func (m *TargetCpa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpa.Merge(m, src)
}
func (m *TargetCpa) XXX_Size() int {
	return xxx_messageInfo_TargetCpa.Size(m)
}
func (m *TargetCpa) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpa.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpa proto.InternalMessageInfo

func (m *TargetCpa) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetCpa) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// Target CPM (cost per thousand impressions) is an automated bidding strategy
// that sets bids to optimize performance given the target CPM you set.
type TargetCpm struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetCpm) Reset()         { *m = TargetCpm{} }
func (m *TargetCpm) String() string { return proto.CompactTextString(m) }
func (*TargetCpm) ProtoMessage()    {}
func (*TargetCpm) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{9}
}

func (m *TargetCpm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpm.Unmarshal(m, b)
}
func (m *TargetCpm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpm.Marshal(b, m, deterministic)
}
func (m *TargetCpm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpm.Merge(m, src)
}
func (m *TargetCpm) XXX_Size() int {
	return xxx_messageInfo_TargetCpm.Size(m)
}
func (m *TargetCpm) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpm.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpm proto.InternalMessageInfo

// An automated bidding strategy that sets bids so that a certain percentage of
// search ads are shown at the top of the first page (or other targeted
// location).
// Next Id = 4
type TargetImpressionShare struct {
	// The targeted location on the search results page.
	Location enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation `protobuf:"varint,1,opt,name=location,proto3,enum=google.ads.googleads.v2.enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation" json:"location,omitempty"`
	// The desired fraction of ads to be shown in the targeted location in micros.
	// E.g. 1% equals 10,000.
	LocationFractionMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=location_fraction_micros,json=locationFractionMicros,proto3" json:"location_fraction_micros,omitempty"`
	// The highest CPC bid the automated bidding system is permitted to specify.
	// This is a required field entered by the advertiser that sets the ceiling
	// and specified in local micros.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetImpressionShare) Reset()         { *m = TargetImpressionShare{} }
func (m *TargetImpressionShare) String() string { return proto.CompactTextString(m) }
func (*TargetImpressionShare) ProtoMessage()    {}
func (*TargetImpressionShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{10}
}

func (m *TargetImpressionShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetImpressionShare.Unmarshal(m, b)
}
func (m *TargetImpressionShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetImpressionShare.Marshal(b, m, deterministic)
}
func (m *TargetImpressionShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetImpressionShare.Merge(m, src)
}
func (m *TargetImpressionShare) XXX_Size() int {
	return xxx_messageInfo_TargetImpressionShare.Size(m)
}
func (m *TargetImpressionShare) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetImpressionShare.DiscardUnknown(m)
}

var xxx_messageInfo_TargetImpressionShare proto.InternalMessageInfo

func (m *TargetImpressionShare) GetLocation() enums.TargetImpressionShareLocationEnum_TargetImpressionShareLocation {
	if m != nil {
		return m.Location
	}
	return enums.TargetImpressionShareLocationEnum_UNSPECIFIED
}

func (m *TargetImpressionShare) GetLocationFractionMicros() *wrappers.Int64Value {
	if m != nil {
		return m.LocationFractionMicros
	}
	return nil
}

func (m *TargetImpressionShare) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// An automated bidding strategy that sets bids based on the target fraction of
// auctions where the advertiser should outrank a specific competitor.
// This strategy is deprecated.
type TargetOutrankShare struct {
	// The target fraction of auctions where the advertiser should outrank the
	// competitor.
	// The advertiser outranks the competitor in an auction if either the
	// advertiser appears above the competitor in the search results, or appears
	// in the search results when the competitor does not.
	// Value must be between 1 and 1000000, inclusive.
	TargetOutrankShareMicros *wrappers.Int32Value `protobuf:"bytes,1,opt,name=target_outrank_share_micros,json=targetOutrankShareMicros,proto3" json:"target_outrank_share_micros,omitempty"`
	// Competitor's visible domain URL.
	CompetitorDomain *wrappers.StringValue `protobuf:"bytes,2,opt,name=competitor_domain,json=competitorDomain,proto3" json:"competitor_domain,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Whether the strategy should always follow bid estimate changes,
	// or only increase.
	// If false, always set a keyword's new bid to the current bid estimate.
	// If true, only updates a keyword's bid if the current bid estimate is
	// greater than the current bid.
	OnlyRaiseCpcBids *wrappers.BoolValue `protobuf:"bytes,4,opt,name=only_raise_cpc_bids,json=onlyRaiseCpcBids,proto3" json:"only_raise_cpc_bids,omitempty"`
	// Whether the strategy is allowed to raise bids on keywords with
	// lower-range quality scores.
	RaiseCpcBidWhenQualityScoreIsLow *wrappers.BoolValue `protobuf:"bytes,5,opt,name=raise_cpc_bid_when_quality_score_is_low,json=raiseCpcBidWhenQualityScoreIsLow,proto3" json:"raise_cpc_bid_when_quality_score_is_low,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}            `json:"-"`
	XXX_unrecognized                 []byte              `json:"-"`
	XXX_sizecache                    int32               `json:"-"`
}

func (m *TargetOutrankShare) Reset()         { *m = TargetOutrankShare{} }
func (m *TargetOutrankShare) String() string { return proto.CompactTextString(m) }
func (*TargetOutrankShare) ProtoMessage()    {}
func (*TargetOutrankShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{11}
}

func (m *TargetOutrankShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetOutrankShare.Unmarshal(m, b)
}
func (m *TargetOutrankShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetOutrankShare.Marshal(b, m, deterministic)
}
func (m *TargetOutrankShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetOutrankShare.Merge(m, src)
}
func (m *TargetOutrankShare) XXX_Size() int {
	return xxx_messageInfo_TargetOutrankShare.Size(m)
}
func (m *TargetOutrankShare) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetOutrankShare.DiscardUnknown(m)
}

var xxx_messageInfo_TargetOutrankShare proto.InternalMessageInfo

func (m *TargetOutrankShare) GetTargetOutrankShareMicros() *wrappers.Int32Value {
	if m != nil {
		return m.TargetOutrankShareMicros
	}
	return nil
}

func (m *TargetOutrankShare) GetCompetitorDomain() *wrappers.StringValue {
	if m != nil {
		return m.CompetitorDomain
	}
	return nil
}

func (m *TargetOutrankShare) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetOutrankShare) GetOnlyRaiseCpcBids() *wrappers.BoolValue {
	if m != nil {
		return m.OnlyRaiseCpcBids
	}
	return nil
}

func (m *TargetOutrankShare) GetRaiseCpcBidWhenQualityScoreIsLow() *wrappers.BoolValue {
	if m != nil {
		return m.RaiseCpcBidWhenQualityScoreIsLow
	}
	return nil
}

// An automated bidding strategy that helps you maximize revenue while
// averaging a specific target return on ad spend (ROAS).
type TargetRoas struct {
	// Required. The desired revenue (based on conversion data) per unit of spend.
	// Value must be between 0.01 and 1000.0, inclusive.
	TargetRoas *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=target_roas,json=targetRoas,proto3" json:"target_roas,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Minimum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidFloorMicros    *wrappers.Int64Value `protobuf:"bytes,3,opt,name=cpc_bid_floor_micros,json=cpcBidFloorMicros,proto3" json:"cpc_bid_floor_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetRoas) Reset()         { *m = TargetRoas{} }
func (m *TargetRoas) String() string { return proto.CompactTextString(m) }
func (*TargetRoas) ProtoMessage()    {}
func (*TargetRoas) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{12}
}

func (m *TargetRoas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRoas.Unmarshal(m, b)
}
func (m *TargetRoas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRoas.Marshal(b, m, deterministic)
}
func (m *TargetRoas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRoas.Merge(m, src)
}
func (m *TargetRoas) XXX_Size() int {
	return xxx_messageInfo_TargetRoas.Size(m)
}
func (m *TargetRoas) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRoas.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRoas proto.InternalMessageInfo

func (m *TargetRoas) GetTargetRoas() *wrappers.DoubleValue {
	if m != nil {
		return m.TargetRoas
	}
	return nil
}

func (m *TargetRoas) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *TargetRoas) GetCpcBidFloorMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidFloorMicros
	}
	return nil
}

// An automated bid strategy that sets your bids to help get as many clicks
// as possible within your budget.
type TargetSpend struct {
	// The spend target under which to maximize clicks.
	// A TargetSpend bidder will attempt to spend the smaller of this value
	// or the natural throttling spend amount.
	// If not specified, the budget is used as the spend target.
	TargetSpendMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_spend_micros,json=targetSpendMicros,proto3" json:"target_spend_micros,omitempty"`
	// Maximum bid limit that can be set by the bid strategy.
	// The limit applies to all keywords managed by the strategy.
	CpcBidCeilingMicros  *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetSpend) Reset()         { *m = TargetSpend{} }
func (m *TargetSpend) String() string { return proto.CompactTextString(m) }
func (*TargetSpend) ProtoMessage()    {}
func (*TargetSpend) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{13}
}

func (m *TargetSpend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetSpend.Unmarshal(m, b)
}
func (m *TargetSpend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetSpend.Marshal(b, m, deterministic)
}
func (m *TargetSpend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetSpend.Merge(m, src)
}
func (m *TargetSpend) XXX_Size() int {
	return xxx_messageInfo_TargetSpend.Size(m)
}
func (m *TargetSpend) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetSpend.DiscardUnknown(m)
}

var xxx_messageInfo_TargetSpend proto.InternalMessageInfo

func (m *TargetSpend) GetTargetSpendMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetSpendMicros
	}
	return nil
}

func (m *TargetSpend) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

// A bidding strategy where bids are a fraction of the advertised price for
// some good or service.
type PercentCpc struct {
	// Maximum bid limit that can be set by the bid strategy. This is
	// an optional field entered by the advertiser and specified in local micros.
	// Note: A zero value is interpreted in the same way as having bid_ceiling
	// undefined.
	CpcBidCeilingMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=cpc_bid_ceiling_micros,json=cpcBidCeilingMicros,proto3" json:"cpc_bid_ceiling_micros,omitempty"`
	// Adjusts the bid for each auction upward or downward, depending on the
	// likelihood of a conversion. Individual bids may exceed
	// cpc_bid_ceiling_micros, but the average bid amount for a campaign should
	// not.
	EnhancedCpcEnabled   *wrappers.BoolValue `protobuf:"bytes,2,opt,name=enhanced_cpc_enabled,json=enhancedCpcEnabled,proto3" json:"enhanced_cpc_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PercentCpc) Reset()         { *m = PercentCpc{} }
func (m *PercentCpc) String() string { return proto.CompactTextString(m) }
func (*PercentCpc) ProtoMessage()    {}
func (*PercentCpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa9bf865024275f, []int{14}
}

func (m *PercentCpc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PercentCpc.Unmarshal(m, b)
}
func (m *PercentCpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PercentCpc.Marshal(b, m, deterministic)
}
func (m *PercentCpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PercentCpc.Merge(m, src)
}
func (m *PercentCpc) XXX_Size() int {
	return xxx_messageInfo_PercentCpc.Size(m)
}
func (m *PercentCpc) XXX_DiscardUnknown() {
	xxx_messageInfo_PercentCpc.DiscardUnknown(m)
}

var xxx_messageInfo_PercentCpc proto.InternalMessageInfo

func (m *PercentCpc) GetCpcBidCeilingMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidCeilingMicros
	}
	return nil
}

func (m *PercentCpc) GetEnhancedCpcEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.EnhancedCpcEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*Commission)(nil), "google.ads.googleads.v2.common.Commission")
	proto.RegisterType((*EnhancedCpc)(nil), "google.ads.googleads.v2.common.EnhancedCpc")
	proto.RegisterType((*ManualCpc)(nil), "google.ads.googleads.v2.common.ManualCpc")
	proto.RegisterType((*ManualCpm)(nil), "google.ads.googleads.v2.common.ManualCpm")
	proto.RegisterType((*ManualCpv)(nil), "google.ads.googleads.v2.common.ManualCpv")
	proto.RegisterType((*MaximizeConversions)(nil), "google.ads.googleads.v2.common.MaximizeConversions")
	proto.RegisterType((*MaximizeConversionValue)(nil), "google.ads.googleads.v2.common.MaximizeConversionValue")
	proto.RegisterType((*PageOnePromoted)(nil), "google.ads.googleads.v2.common.PageOnePromoted")
	proto.RegisterType((*TargetCpa)(nil), "google.ads.googleads.v2.common.TargetCpa")
	proto.RegisterType((*TargetCpm)(nil), "google.ads.googleads.v2.common.TargetCpm")
	proto.RegisterType((*TargetImpressionShare)(nil), "google.ads.googleads.v2.common.TargetImpressionShare")
	proto.RegisterType((*TargetOutrankShare)(nil), "google.ads.googleads.v2.common.TargetOutrankShare")
	proto.RegisterType((*TargetRoas)(nil), "google.ads.googleads.v2.common.TargetRoas")
	proto.RegisterType((*TargetSpend)(nil), "google.ads.googleads.v2.common.TargetSpend")
	proto.RegisterType((*PercentCpc)(nil), "google.ads.googleads.v2.common.PercentCpc")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/bidding.proto", fileDescriptor_4aa9bf865024275f)
}

var fileDescriptor_4aa9bf865024275f = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0x80, 0x65, 0xa7, 0xad, 0xe8, 0x3a, 0xa5, 0xe4, 0x92, 0x06, 0x2b, 0xad, 0xaa, 0xea, 0x24,
	0x04, 0x0f, 0xe8, 0x4e, 0x72, 0x11, 0x0f, 0x46, 0x08, 0xd9, 0x97, 0x34, 0xb2, 0x48, 0x54, 0xd7,
	0x2e, 0x01, 0xa2, 0x88, 0xd5, 0x7a, 0x77, 0x73, 0x59, 0xf5, 0x6e, 0xe7, 0xd8, 0xdd, 0x4b, 0x48,
	0x5f, 0xf8, 0x2f, 0x3c, 0x22, 0x7e, 0x02, 0xef, 0x48, 0x88, 0xff, 0x81, 0x04, 0xaf, 0xfc, 0x00,
	0x74, 0xbb, 0x7b, 0x76, 0x8b, 0x93, 0xb8, 0x09, 0x41, 0xe2, 0xc9, 0xbb, 0xde, 0x99, 0x6f, 0x66,
	0x67, 0x66, 0xe7, 0x06, 0x7d, 0x98, 0x02, 0xa4, 0x19, 0x8f, 0x09, 0xd3, 0xb1, 0x5b, 0x56, 0xab,
	0xe3, 0x4e, 0x4c, 0x21, 0xcf, 0x41, 0xc6, 0x13, 0xc1, 0x98, 0x90, 0x69, 0x54, 0x28, 0x30, 0x10,
	0x3c, 0x74, 0x22, 0x11, 0x61, 0x3a, 0x9a, 0x4a, 0x47, 0xc7, 0x9d, 0xc8, 0x49, 0x6f, 0x24, 0xe7,
	0xd1, 0xb8, 0x2c, 0x73, 0x1d, 0x17, 0x24, 0xe5, 0x18, 0x24, 0xc7, 0x85, 0x82, 0x1c, 0x0c, 0x67,
	0x58, 0x1b, 0x45, 0x0c, 0x4f, 0x4f, 0x71, 0x0a, 0x24, 0x73, 0x46, 0x36, 0x36, 0x2f, 0x86, 0x18,
	0xa2, 0x52, 0x6e, 0xb0, 0xc8, 0x0b, 0xc5, 0xb5, 0x16, 0x20, 0xb1, 0x3e, 0x22, 0x8a, 0xe3, 0x0c,
	0x28, 0x31, 0x02, 0xa4, 0xa7, 0x78, 0x57, 0x63, 0xbb, 0x9b, 0x94, 0x87, 0xf1, 0x89, 0x22, 0x45,
	0xc1, 0x95, 0xf6, 0xe7, 0x0f, 0x6a, 0x2b, 0x85, 0x88, 0x89, 0x94, 0x60, 0xac, 0xb2, 0x3f, 0x0d,
	0x31, 0x42, 0x09, 0xe4, 0xb9, 0xb0, 0x06, 0x82, 0x67, 0x68, 0x9d, 0x4e, 0x77, 0xb8, 0xf2, 0x18,
	0xe7, 0x82, 0x2a, 0xd0, 0xed, 0xc6, 0xa3, 0xc6, 0x07, 0xad, 0xce, 0x7d, 0x1f, 0x8c, 0xa8, 0x36,
	0x16, 0x0d, 0xa4, 0xf9, 0xf8, 0xa3, 0x3d, 0x92, 0x95, 0x7c, 0xb4, 0x36, 0x53, 0x1d, 0x11, 0xc3,
	0x77, 0xad, 0x62, 0x78, 0x07, 0xb5, 0xb6, 0xe4, 0x11, 0x91, 0x94, 0xb3, 0xa4, 0xa0, 0xe1, 0xd7,
	0xe8, 0xf6, 0x2e, 0x91, 0x25, 0xc9, 0x92, 0x82, 0x06, 0x3b, 0x68, 0x8d, 0xfb, 0x33, 0x4c, 0x0b,
	0x8a, 0xb9, 0x24, 0x93, 0x8c, 0x33, 0x6f, 0x6c, 0x63, 0xce, 0x58, 0x1f, 0x20, 0x73, 0xb6, 0x02,
	0x3e, 0x63, 0x6e, 0x39, 0xad, 0xb0, 0x35, 0x43, 0xe7, 0xaf, 0x6e, 0x8e, 0xc3, 0x7b, 0x68, 0x75,
	0x97, 0x7c, 0x27, 0x72, 0xf1, 0x92, 0x27, 0x20, 0x8f, 0xb9, 0xaa, 0x7c, 0xd4, 0xe1, 0x57, 0xe8,
	0xdd, 0xf9, 0xbf, 0x2d, 0x3f, 0xf8, 0x14, 0xb5, 0x7c, 0xf8, 0x15, 0x90, 0xfa, 0xf6, 0x0f, 0xe6,
	0x1c, 0xda, 0x84, 0x72, 0x92, 0x71, 0xe7, 0x12, 0x72, 0x0a, 0x23, 0x20, 0x3a, 0xfc, 0xed, 0x06,
	0xba, 0x3b, 0x24, 0x29, 0x7f, 0x2a, 0xf9, 0xd0, 0x57, 0x40, 0xf0, 0x3d, 0xba, 0xf3, 0x5a, 0x11,
	0x58, 0xe8, 0xdb, 0x9d, 0xfd, 0xe8, 0xbc, 0x52, 0xb3, 0x55, 0x10, 0xfd, 0x03, 0x33, 0xf6, 0x88,
	0x6d, 0x20, 0xd9, 0x96, 0x2c, 0xf3, 0x8b, 0xce, 0x47, 0xcb, 0xfa, 0x95, 0x5d, 0x30, 0x44, 0xeb,
	0x55, 0x90, 0x27, 0x82, 0x61, 0xca, 0x45, 0x26, 0x64, 0x5a, 0x27, 0xb7, 0xb9, 0x38, 0xb9, 0xab,
	0xb4, 0xa0, 0x7d, 0xc1, 0x12, 0xa7, 0xe8, 0x72, 0x1b, 0x7c, 0x86, 0x96, 0x2b, 0x5a, 0x0e, 0x4c,
	0x1c, 0x0a, 0xae, 0xda, 0x4b, 0x6f, 0x10, 0xa6, 0xd6, 0x44, 0xb0, 0x5d, 0xaf, 0x10, 0x0c, 0xd0,
	0x2a, 0xc8, 0xec, 0x14, 0x2b, 0x22, 0x34, 0xc7, 0xde, 0x3b, 0xdd, 0xbe, 0xb1, 0x30, 0xff, 0xef,
	0x54, 0x6a, 0xa3, 0x4a, 0x2b, 0xb1, 0x7e, 0xe9, 0xe0, 0x05, 0x7a, 0xef, 0x35, 0x0a, 0x3e, 0x39,
	0xe2, 0x12, 0x4f, 0x4a, 0x56, 0x25, 0x91, 0x82, 0xac, 0x42, 0x21, 0x24, 0x67, 0xed, 0x9b, 0x0b,
	0xe1, 0x8f, 0xd4, 0x0c, 0xfc, 0xe5, 0x11, 0x97, 0x7d, 0x0b, 0x49, 0x66, 0x8c, 0x20, 0x47, 0xef,
	0x9f, 0x61, 0xec, 0xdb, 0x92, 0x64, 0xc2, 0x9c, 0x62, 0x4d, 0x41, 0x71, 0x2c, 0x34, 0xce, 0xe0,
	0xa4, 0x7d, 0xeb, 0xd2, 0xe6, 0x9e, 0x39, 0xcc, 0xb8, 0xa2, 0x0c, 0xf4, 0x0e, 0x9c, 0x84, 0x7f,
	0x35, 0xd0, 0xed, 0xe7, 0xb6, 0xba, 0x92, 0x82, 0x04, 0xdb, 0x68, 0xc5, 0xd7, 0x26, 0x2d, 0xc8,
	0x25, 0xde, 0xe7, 0x5d, 0x53, 0x23, 0x7c, 0xfa, 0xae, 0xbf, 0x20, 0x76, 0xd0, 0x5a, 0x4d, 0x3c,
	0xcc, 0x00, 0x54, 0xcd, 0x5b, 0x5a, 0xcc, 0x5b, 0x71, 0xbc, 0x27, 0x95, 0x9a, 0x6f, 0x1d, 0xad,
	0xd9, 0xad, 0xf3, 0xf0, 0xe7, 0x26, 0xba, 0xe7, 0x76, 0x83, 0x69, 0x43, 0x1c, 0x57, 0xfd, 0x30,
	0x78, 0x89, 0xde, 0xaa, 0x5b, 0xa2, 0x7f, 0x53, 0xdf, 0x2c, 0x78, 0x53, 0x67, 0x72, 0x76, 0x3c,
	0xc3, 0xbe, 0xaa, 0x0b, 0x25, 0x46, 0x53, 0x7b, 0xc1, 0x17, 0xa8, 0x5d, 0xaf, 0xf1, 0xa1, 0x22,
	0xd4, 0x2e, 0xde, 0x3c, 0x88, 0xeb, 0xb5, 0xf2, 0x13, 0xaf, 0xbb, 0x30, 0x33, 0x4b, 0x57, 0xcb,
	0x4c, 0xf8, 0xcb, 0x12, 0x0a, 0xdc, 0xa5, 0x9e, 0x96, 0x46, 0x11, 0xf9, 0xc2, 0xc5, 0x6e, 0x1f,
	0xdd, 0xf7, 0xb5, 0x04, 0xee, 0x6f, 0xff, 0x8d, 0x59, 0x5c, 0x55, 0x8f, 0x3b, 0xce, 0x5a, 0xdb,
	0xcc, 0x51, 0xfd, 0x25, 0x06, 0x68, 0x85, 0x42, 0x5e, 0x70, 0x23, 0x0c, 0x28, 0xcc, 0x20, 0x27,
	0x42, 0xfa, 0xa0, 0xcc, 0xb7, 0x88, 0xb1, 0x51, 0x42, 0xa6, 0xfe, 0x71, 0xcf, 0xd4, 0x36, 0xad,
	0xd6, 0xf5, 0xc7, 0xe3, 0x3a, 0x3b, 0xcf, 0x25, 0x9a, 0xc1, 0xcd, 0x6b, 0x68, 0x06, 0x7f, 0x36,
	0x10, 0x7a, 0x3e, 0xfd, 0xd4, 0xfc, 0xcb, 0x2f, 0xd5, 0xff, 0xbe, 0x07, 0xfc, 0xd4, 0x40, 0x2d,
	0x77, 0xdb, 0x71, 0xc1, 0x25, 0x0b, 0x3e, 0x47, 0xab, 0xfe, 0xba, 0xba, 0xda, 0x5f, 0xa2, 0xfd,
	0xf9, 0xa6, 0x69, 0x31, 0xff, 0x55, 0x03, 0xac, 0xdc, 0x45, 0x43, 0xae, 0x28, 0x97, 0xa6, 0x1a,
	0x70, 0xce, 0x37, 0xd0, 0xb8, 0x7a, 0x74, 0xcf, 0x1c, 0x99, 0x9a, 0x57, 0x19, 0x99, 0xfa, 0xbf,
	0x37, 0x50, 0x48, 0x21, 0x8f, 0x2e, 0x9e, 0x76, 0xfb, 0xcb, 0x7d, 0x37, 0x1c, 0x0f, 0x2b, 0xea,
	0xb0, 0xb1, 0xef, 0x07, 0xd7, 0x28, 0x85, 0x8c, 0xc8, 0x34, 0x02, 0x95, 0xc6, 0x29, 0x97, 0xd6,
	0x66, 0x3d, 0xc8, 0x16, 0x42, 0x9f, 0x37, 0x6a, 0x7f, 0xe2, 0x7e, 0x7e, 0x68, 0x2e, 0x6d, 0xf7,
	0x7a, 0x3f, 0x36, 0x1f, 0x6e, 0x3b, 0x58, 0x8f, 0xe9, 0xc8, 0x2d, 0xab, 0xd5, 0x5e, 0x27, 0x4a,
	0xac, 0xd8, 0xaf, 0xb5, 0xc0, 0x41, 0x8f, 0xe9, 0x83, 0xa9, 0xc0, 0xc1, 0x5e, 0xe7, 0xc0, 0x09,
	0xfc, 0xd1, 0x0c, 0xdd, 0xbf, 0xdd, 0x6e, 0x8f, 0xe9, 0x6e, 0x77, 0x2a, 0xd2, 0xed, 0xee, 0x75,
	0xba, 0x5d, 0x27, 0x34, 0xb9, 0x65, 0xbd, 0x7b, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x97, 0x6e, 0x16, 0x07, 0x0c, 0x00, 0x00,
}
