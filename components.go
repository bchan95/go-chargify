package chargify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Component struct {
	Component *ComponentBody `json:"component,omitempty"`
}

type Allocation struct {
	Allocation *ComponentBody `json:"allocation,omitempty"`
}

type ComponentBody struct {
	ID                        int64    `json:"id,omitempty"`
	Name                      string   `json:"name,omitempty"`
	PricingScheme             string   `json:"pricing_scheme,omitempty"`
	UnitName                  string   `json:"unit_name,omitempty"`
	UnitPrice                 string   `json:"unit_price,omitempty"`
	ProductFamilyID           int64    `json:"product_family_id,omitempty"`
	PricePerUnitInCents       int64    `json:"price_per_unit_in_cents,omitempty"`
	Kind                      string   `json:"kind,omitempty"`
	Archived                  bool     `json:"archived,omitempty"`
	Taxable                   bool     `json:"taxable,omitempty"`
	Description               string   `json:"description,omitempty"`
	DefaultPricePointID       int64    `json:"default_price_point_id,omitempty"`
	PricePointCount           int64    `json:"price_point_count,omitempty"`
	PricePointsURL            string   `json:"price_points_url,omitempty"`
	TaxCode                   string   `json:"tax_code,omitempty"`
	Recurring                 bool     `json:"recurring,omitempty"`
	UpgradeCharge             string   `json:"upgrade_charge,omitempty"`
	DowngradeCharge           string   `json:"downgrade_charge,omitempty"`
	CreatedAt                 string   `json:"created_at,omitempty"`
	Prices                    []*Price `json:"prices,omitempty"`
	Quantity                  int64    `json:"quantity,omitempty"`
	Timestamp                 string   `json:"timestamp,omitempty"`
	ProrationUpgradeScheme    string   `json:"proration_upgrade_scheme, omitempty"`
	ProrationDowngradeScheme  string   `json:"proration_downgrade_scheme,omitempty"`
	ProrationCollectionMethod string   `json:"proration_collection_method,omitempty"`
	// Response
	ComponentID       int64 `json:"component_id,omitempty"`
	SubscriptionID    int64 `json:"subscription_id,omitempty"`
	AllocatedQuantity int64 `json:"allocated_quantity, omitempty"`
	PricePointID      int64 `json:"price_point_id,omitempty"`
}

type Price struct {
	ID                  int64  `json:"id,omitempty"`
	StartingQuantity    int64  `json:"starting_quantity,omitempty"`
	EndingQuantity      int64  `json:"ending_quantity,omitempty"`
	UnitPrice           string `json:"unit_price,omitempty"`
	ComponentID         int64  `json:"component_id,omitempty"`
	PricePointID        int64  `json:"price_point_id,omitempty"`
	FormattedPricePoint string `json:"formatted_price_point,omitempty"`
}

type PricePoint struct {
	PricePoint []*ComponentBody `json:"price_points,omitempty"`
}

func GetComponentAllocation(client Client, subscriptionID int64, componentID int64) (component *Component, err error) {
	uri := fmt.Sprintf("/subscriptions/%d/components/%d.json", subscriptionID, componentID)
	var res *http.Response
	res, err = client.Get(uri)
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	component = new(Component)
	err = json.Unmarshal(body, component)
	return
}

func UpdateComponentQuantity(client Client, subscriptionID int64, componentID int64, quantity int64, upgradeCharge string) (component *Component, err error) {
	uri := fmt.Sprintf("/subscriptions/%d/components/%d/allocations.json", subscriptionID, componentID)
	var jsonReq []byte
	jsonReq, err = json.Marshal(&Allocation{
		Allocation: &ComponentBody{
			Quantity:      quantity,
			UpgradeCharge: upgradeCharge,
		},
	})
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Post(jsonReq, uri)
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	component = new(Component)
	err = json.Unmarshal(body, component)
	return
}

func GetComponentPricePoints(client Client, componentID int64) (pricePoint *PricePoint, err error) {
	uri := fmt.Sprintf("/components/%d/price_points.json", componentID)
	var res *http.Response
	res, err = client.Get(uri)
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	pricePoint = new(PricePoint)
	err = json.Unmarshal(body, pricePoint)
	return
}
