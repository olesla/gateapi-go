/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Statistical data
type FuturesAccountHistory struct {
	// total amount of deposit and withdraw
	Dnw string `json:"dnw,omitempty"`
	// total amount of trading profit and loss
	Pnl string `json:"pnl,omitempty"`
	// total amount of fee
	Fee string `json:"fee,omitempty"`
	// total amount of referrer rebates
	Refr string `json:"refr,omitempty"`
	// total amount of funding costs
	Fund string `json:"fund,omitempty"`
	// total amount of point deposit and withdraw
	PointDnw string `json:"point_dnw,omitempty"`
	// total amount of point fee
	PointFee string `json:"point_fee,omitempty"`
	// total amount of referrer rebates of point fee
	PointRefr string `json:"point_refr,omitempty"`
	// total amount of perpetual contract bonus transfer
	BonusDnw string `json:"bonus_dnw,omitempty"`
	// total amount of perpetual contract bonus deduction
	BonusOffset string `json:"bonus_offset,omitempty"`
}
