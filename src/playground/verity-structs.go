package main

import (
	"time"

	"github.com/google/uuid"
)

type BulkCSVField struct {
	ID                         string                       `json:"id"`
	DataSource                 string                       `json:"data_source"`
	GeoJSONFeatureID           uuid.UUID                    `json:"geo_json_feature_id"`
	Name                       string                       `json:"field_name" gorm:"column:field_name"`
	CropType                   types.CropType               `json:"crop_type"`
	GrowingCycle               string                       `json:"growing_cycle"`
	Yield                      *float64                     `json:"yield"`
	MoistureAtHarvest          float64                      `json:"moisture_at_harvest"`
	Diesel                     *float64                     `json:"diesel"`
	DieselUnit                 types.EnergyUnit             `json:"diesel_unit"`
	Gasoline                   *float64                     `json:"gasoline"`
	GasolineUnit               types.EnergyUnit             `json:"gas_unit"`
	NaturalGas                 *float64                     `json:"natural_gas"`
	NaturalGasUnit             types.EnergyUnit             `json:"ng_unit"`
	LiquefiedPetroleumGas      *float64                     `json:"liquefied_petroleum_gas"`
	LiquefiedPetroleumGasUnit  types.EnergyUnit             `json:"lpg_unit"`
	Electricity                *float64                     `json:"electricity"`
	ElectricityUnit            types.EnergyUnit             `json:"electricity_unit"`
	DataQualityEnergy          float64                      `json:"data_quality_energy"`
	OperationName              string                       `json:"operation_name"`
	OperationType              types.OperationType          `json:"operation_type"`
	CustomApp                  bool                         `json:"custom_app"`
	DataQualityOperation       types.DataQualityType        `json:"data_quality_operation"`
	OperationStart             time.Time                    `json:"operation_start"`
	OperationEnd               time.Time                    `json:"operation_end"`
	ImplementIDs               []string                     `json:"implement_id"`
	NumTillPasses              int                          `json:"num_till_passes"`
	TillDepth                  float64                      `json:"till_depth"`
	TillPractice               types.Tillage                `json:"till_practice"`
	InputFormula               string                       `json:"input_formula"`
	InputName                  string                       `json:"input_name"`
	InputRate                  float64                      `json:"input_rate"`
	InputUnit                  types.ApplicationUnit        `json:"input_unit"`
	InputAcres                 float64                      `json:"input_acres"`
	DilutionFactor             float64                      `json:"dilution_factor"`
	CustomBlend                bool                         `json:"custom_blend"`
	GreenAmmonia               bool                         `json:"green_ammonia"`
	NitrogenManagementPractice types.NitrogenFertTechnology `json:"nitrogen_management_practice"`

	Operations []*Operation `json:"operations"`
}

type FarmField struct {
	Base             `gorm:"embedded"`
	FieldName        string          `json:"field_name"`
	ZipCode          string          `json:"zip_code"`
	StateProvince    string          `json:"state_province"`
	County           string          `json:"county"`
	City             string          `json:"city"`
	Acreage          float64         `json:"acreage"`
	FieldOwnership   types.Ownership `json:"field_ownership" gorm:"type:ownership"`
	FarmID           uuid.UUID       `json:"farm_id"`
	GeoJSONFeatureID uuid.UUID       `json:"geojson_feature_id" gorm:"column:geojson_feature_id"`
	Audited          bool            `json:"audited"`
	FieldCycles      []*FieldCycle   `json:"field_cycles,omitempty"`
}

type FieldCycle struct {
	Base                  `gorm:"embedded"`
	FarmID                uuid.UUID              `json:"farm_id"`
	FarmFieldID           uuid.UUID              `json:"field_id" gorm:"column:field_id"`
	Name                  string                 `json:"name"`
	DataSource            string                 `json:"data_source"`
	CropType              types.CropType         `json:"crop_type"`
	CoverCrop             bool                   `json:"cover_crop"`
	Year                  string                 `json:"year" gorm:"column:cycle_year"`
	FeedstockCiScore      *FeedstockCiResult     `json:"feedstock_ci_result,omitempty"`
	FeedstockCiCategories *FeedstockCiCategories `json:"feedstock_ci_categories,omitempty"`
	FieldOperations       []*FieldOperation      `json:"field_operations,omitempty"`
}

type FieldOperation struct {
	Base                  `gorm:"embedded"`
	FieldCycleID          uuid.UUID              `json:"field_cycle_id"`
	TaskName              string                 `json:"task_name"`
	Type                  types.OperationType    `json:"type"`
	DataQualityType       types.DataQualityType  `json:"data_quality_type"`
	ApplicationStart      time.Time              `json:"application_start"`
	ApplicationEnd        time.Time              `json:"application_end"`
	AcresApplied          float64                `json:"acres_applied"`
	Yield                 *float64               `json:"yield"`
	TillageDepth          *float64               `json:"tillage_depth"`
	TillagePasses         *float64               `json:"tillage_passes"`
	CustomApplication     bool                   `json:"custom_application"`
	Diesel                *Diesel                `json:"diesel,omitempty" gorm:"embedded"`
	Gasoline              *Gasoline              `json:"gasoline,omitempty" gorm:"embedded"`
	NaturalGas            *NaturalGas            `json:"natural_gas,omitempty" gorm:"embedded"`
	LiquefiedPetroleumGas *LiquefiedPetroleumGas `json:"liquefied_petroleum_gas,omitempty" gorm:"embedded"`
	Electricity           *Electricity           `json:"electricity,omitempty" gorm:"embedded"`
	ApplicationProducts   []*ApplicationProduct  `json:"application_products,omitempty"`
}

type ApplicationProduct struct {
	Base               `gorm:"embedded"`
	FieldOperationID   uuid.UUID                    `json:"field_operation_id"`
	Name               string                       `json:"name"`
	Type               types.ApplicationType        `json:"type"`
	DataQualityType    types.DataQualityType        `json:"data_quality_type"`
	AmountTotal        float64                      `json:"amount_total"`
	Unit               types.ApplicationUnit        `json:"unit"`
	NitrogenTechnology types.NitrogenFertTechnology `json:"nitrogen_technology"`
	GreenAmmonia       string                       `json:"green_ammonia"`
	DilutionFactor     float64                      `json:"dilution_factor"`
}


type FarmingInput struct {
	Base                    `swaggerignore:"true"`
	FarmId                  uuid.UUID               `swaggerignore:"true" gorm:"type:uuid"` // Foreign key (belongs to)
	Yield                   float64                 `json:"yield" `
	StateName               string                  `json:"state_name"`
	CountyName              string                  `json:"county_name"`
	Product                 Product                 `json:"product" gorm:"embedded"`
	Energy                  Energy                  `json:"energy" gorm:"embedded"`
	NitrogenFertilizer      NitrogenFertilizer      `json:"nitrogenFertilizer" gorm:"embedded"`
	PhosphorusFertilizer    PhosphorusFertilizer    `json:"phosphorusFertilizer" gorm:"embedded"`
	PotashFertilizer        PotashFertilizer        `json:"potashFertilizer" gorm:"embedded"`
	Lime                    Lime                    `json:"lime" gorm:"embedded"`
	Herbicide               Herbicide               `json:"herbicide" gorm:"embedded"`
	Insecticide             Insecticide             `json:"insecticide" gorm:"embedded"`
	Year                    string                  `swaggerignore:"true" json:"year,omitempty" validate:"required,len=4"`
	Audited                 bool                    `swaggerignore:"true" json:"audited"`
	FieldName               string                  `swaggerignore:"true" json:"fieldName"`
	GeoJsonFeatureID        uuid.UUID               `swaggerignore:"true" json:"geoJsonFeatureID"`
	SoilOrganicCarbonLookup SoilOrganicCarbonLookup `json:"soilOrganicCarbonLookup" gorm:"embedded"`
}

type FieldCycleFdcicInput struct {
	Base                        `gorm:"embedded"`
	FieldCycleID                uuid.UUID                    `json:"field_cycle_id"`
	Diesel                      float64                      `json:"diesel"`
	Gasoline                    float64                      `json:"gasoline"`
	NaturalGas                  float64                      `json:"natural_gas"`
	LiquefiedPetroleumGas       float64                      `json:"liquefied_petroleum_gas"`
	Electricity                 float64                      `json:"electricity"`
	Ammonia                     float64                      `json:"ammonia"`
	Urea                        float64                      `json:"urea"`
	AmmoniumNitrate             float64                      `json:"ammonium_nitrate"`
	AmmoniumSulfate             float64                      `json:"ammonium_sulfate"`
	UreaAmmoniumNitrateSolution float64                      `json:"urea_ammonium_nitrate_solution"`
	MonoammoniumPhosphateNFert  float64                      `json:"monoammonium_phosphate_n_fert"`
	DiammoniumPhosphateNFert    float64                      `json:"diammonium_phosphate_n_fert"`
	MonoammoniumPhosphatePFert  float64                      `json:"monoammonium_phosphate_p_fert"`
	DiammoniumPhosphatePFert    float64                      `json:"diammonium_phosphate_p_fert"`
	PotashK2O                   float64                      `json:"potash_k2o" gorm:"column:potash_k2o"`
	LimeMinus0                  float64                      `json:"lime_minus_0"`
	LimeMinus1                  float64                      `json:"lime_minus_1"`
	LimeMinus2                  float64                      `json:"lime_minus_2"`
	LimeMinus3                  float64                      `json:"lime_minus_3"`
	Herbicide                   float64                      `json:"herbicide"`
	Insecticide                 float64                      `json:"insecticide"`
	Type                        types.CropType               `json:"type"`
	Acres                       float64                      `json:"acres"`
	Yield                       float64                      `json:"yield"`
	FipsCode                    string                       `json:"fips_code"`
	LocationState               string                       `json:"location_state"`
	LocationCity                string                       `json:"location_city"`
	CoverCropsUsed              bool                         `json:"cover_crops_used"`
	ManureUsed                  bool                         `json:"manure_used"`
	Tillage                     types.Tillage                `json:"tillage"`
	SourceAmmonia               types.SourceAmmonia          `json:"source_ammonia"`
	TechnologyNitrogen          types.NitrogenFertTechnology `json:"technology_nitrogen"`
	ClimateZone                 types.ClimateZone            `json:"climate_zone"`
	MoistureAtHarvest           float64                      `json:"moisture_at_harvest"`
	MeasuredSOC                 float64                      `json:"measured_soc"`
	Audited                     bool                         `json:"audited"`
}
/*
1 lbs / acre