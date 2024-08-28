package query

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

type ConsigneeFormBaseModel struct {
	ConsigneeFormID    int64           `db:"consignee_form_id" json:"-"`
	CustomerID         int64           `db:"customer_id" json:"-"`
	ShipmentCategoryID int64           `db:"shipment_category_id" json:"-"`
	SenderName         string          `db:"sender_name" json:"-"`
	SenderPhoneNumber  string          `db:"sender_phone_number" json:"-"`
	SenderAddress      string          `db:"sender_address" json:"-"`
	GoodsValue         float64         `db:"goods_value" json:"-"`
	Pieces             int64           `db:"pieces" json:"-"`
	GrossWeight        float64         `db:"gross_weight" json:"-"`
	DimensionWidth     float64         `db:"dimension_width" json:"-"`
	DimensionHeight    float64         `db:"dimension_height" json:"-"`
	DimensionLength    float64         `db:"dimension_length" json:"-"`
	Origin             string          `db:"origin" json:"-"`
	Latitude           sql.NullFloat64 `db:"latitude" json:"-"`
	Longitude          sql.NullFloat64 `db:"longitude" json:"-"`
	IsCOD              bool            `db:"is_cod" json:"-"`
	IsDFOD             sql.NullBool    `db:"is_dfod" json:"-"`
	IsFragile          bool            `db:"is_fragile" json:"-"`
	FormIdentifier     string          `db:"form_identifier" json:"-"`
	CreatedAt          time.Time       `db:"created_at" json:"-"`
	CreatedBy          string          `db:"created_by" json:"-"`
	UpdatedAt          sql.NullTime    `db:"updated_at" json:"-"`
	UpdatedBy          sql.NullString  `db:"updated_by" json:"-"`
	DeletedAt          sql.NullTime    `db:"deleted_at" json:"-"`
	DeletedBy          sql.NullString  `db:"deleted_by" json:"-"`
	ReferencePoint     sql.NullString  `db:"reference_point" json:"-"`
}

func QueryForm() bool {
	fmt.Println("UpdateReferencePointByConsigneeFormID")
	data := ConsigneeFormBaseModel{}
	queryU, argsU, errU := squirrel.Update("consignee_form").
		Set("reference_point", data.ReferencePoint).
		Set("updated_at", data.UpdatedAt).
		Set("updated_by", data.UpdatedBy).
		Where("consignee_form_id = ?", data.ConsigneeFormID).
		ToSql()
	fmt.Println(queryU)
	fmt.Println(argsU, errU)

	fmt.Println("create")
	formIdentifier := ""
	queryI, argsI, errI := squirrel.Insert("consignee_form").
		Columns("customer_id", "shipment_category_id", "sender_name", "sender_phone_number", "sender_address",
			"goods_value", "pieces", "gross_weight", "dimension_width", "dimension_height", "dimension_length", "origin",
			"latitude", "longitude", "is_cod", "is_dfod", "is_fragile", "form_identifier", "created_at", "created_by",
			"reference_point").
		Values(data.CustomerID, data.ShipmentCategoryID, data.SenderName, data.SenderPhoneNumber, data.SenderName,
			data.GoodsValue, data.Pieces, data.GrossWeight, data.DimensionWidth, data.DimensionHeight, data.DimensionLength,
			data.Origin, data.Latitude, data.Longitude, data.IsCOD, data.IsDFOD, data.IsFragile, formIdentifier,
			time.Now().UTC(), data.CreatedBy, data.ReferencePoint).
		Suffix("RETURNING \"consignee_form_id\"").
		ToSql()
	fmt.Println(queryI)
	fmt.Println(argsI, errI)

	fmt.Println("FindByID")
	ID := int64(1)
	query, args, err := squirrel.Select("cf.consignee_form_id", "cf.customer_id", "cf.shipment_category_id", "cf.sender_name",
		"cf.sender_phone_number", "cf.sender_address", "cf.goods_value", "cf.pieces", "cf.gross_weight", "cf.dimension_width",
		"cf.dimension_height", "cf.dimension_length", "cf.origin", "cf.latitude", "cf.longitude", "cf.is_cod", "cf.is_dfod",
		"cf.is_fragile", "cf.form_identifier", "cf.created_at", "cf.created_by", "cf.reference_point", "sc.shipment_category_name").
		From("consignee_form cf").
		Join("shipment_category sc ON sc.shipment_category_id = cf.shipment_category_id").
		Where("cf.consignee_form_id = ?", ID).
		Where("cf.deleted_at IS NULL").
		ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	fmt.Println("FindByFormIdentifier")
	query, args, err = squirrel.Select("cf.consignee_form_id", "cf.customer_id", "cf.shipment_category_id", "cf.sender_name",
		"cf.sender_phone_number", "cf.sender_address", "cf.goods_value", "cf.pieces", "cf.gross_weight", "cf.dimension_width",
		"cf.dimension_height", "cf.dimension_length", "cf.origin", "cf.latitude", "cf.longitude", "cf.is_cod", "cf.is_dfod",
		"cf.is_fragile", "cf.form_identifier", "cf.created_at", "cf.created_by", "cf.reference_point", "sc.shipment_category_name").
		From("consignee_form cf").
		Join("shipment_category sc ON sc.shipment_category_id = cf.shipment_category_id").
		Where("cf.form_identifier = ?", formIdentifier).
		Where("cf.deleted_at IS NULL").ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	fmt.Println("DeleteByConsigneeFormID")
	actorName := "tes"
	consigneeFormID := int64(100)
	now := time.Now().UTC()
	queryU, argsU, errU = squirrel.Update("consignee_form").
		Set("updated_by", actorName).
		Set("deleted_by", actorName).
		Set("updated_at", now).
		Set("deleted_at", now).
		Where("consignee_form_id = ?", consigneeFormID).
		ToSql()
	fmt.Println(queryU)
	fmt.Println(argsU, errU)

	fmt.Println("FindByConsigneeFormID")
	query, args, err = squirrel.Select("consignee_form_id", "customer_id", "shipment_category_id", "sender_name",
		"sender_phone_number", "sender_address", "goods_value", "pieces", "gross_weight", "dimension_width",
		"dimension_height", "dimension_length", "origin", "latitude", "longitude", "is_cod", "is_dfod",
		"is_fragile", "form_identifier", "created_at", "created_by", "updated_at", "updated_by", "deleted_at",
		"deleted_by", "reference_point").
		From("consignee_form").
		Where("consignee_form_id = ?", consigneeFormID).
		ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	fmt.Println("UpdateByConsigneeFormID")
	queryU, argsU, errU = squirrel.Update("consignee_form").
		Set("shipment_category_id", data.ShipmentCategoryID).
		Set("sender_name", data.SenderName).
		Set("sender_phone_number", data.SenderPhoneNumber).
		Set("sender_address", data.SenderAddress).
		Set("goods_value", data.GoodsValue).
		Set("pieces", data.Pieces).
		Set("gross_weight", data.GrossWeight).
		Set("dimension_width", data.DimensionWidth).
		Set("dimension_height", data.DimensionHeight).
		Set("dimension_length", data.DimensionLength).
		Set("origin", data.Origin).
		Set("latitude", data.Latitude.Float64).
		Set("longitude", data.Longitude.Float64).
		Set("is_cod", data.IsCOD).
		Set("is_dfod", data.IsDFOD).
		Set("is_fragile", data.IsFragile).
		Set("updated_at", data.UpdatedAt).
		Set("updated_by", data.UpdatedBy).
		Set("reference_point", data.ReferencePoint).
		Where("consignee_form_id = ?", data.ConsigneeFormID).
		ToSql()
	fmt.Println(queryU)
	fmt.Println(argsU, errU)
	return true
}

type PaymentMethodUpdateModel struct {
	Method             string         `db:"payment_method" json:"-"`
	Status             string         `db:"status" json:"-"`
	Type               sql.NullString `db:"type" json:"-"`
	PaymentExpiration  int64          `db:"payment_expiration" json:"-"`
	PaymentMethodName  sql.NullString `db:"payment_method_name" json:"-"`
	PaymentMethodOrder sql.NullInt64  `db:"payment_method_order" json:"-"`
	Note               sql.NullString `db:"note" json:"-"`
}

func TestQuery() bool {
	paymentMethod := ""
	query, args, err := squirrel.
		Select("payment_method", "logo", "status", "type").
		From("payment_method").
		Where("status = 'ACTIVE'").
		OrderBy("payment_method ASC").
		ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	query, args, err = squirrel.Select("pm.payment_method", "pm.logo", "pm.status", "pm.type", "pm.payment_method_name", "pm.note").
		From("payment_method pm").
		InnerJoin("payment_method_type pt on pm.type = pt.type").
		Where(squirrel.NotEq{"status": []string{"INACTIVE"}}).
		OrderBy("payment_method_order, payment_method ASC").
		ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	query, args, err = squirrel.Select("payment_method", "status", "type", "payment_expiration", "payment_method_name", "payment_method_order", "note", "type").
		From("payment_method").
		Where("payment_method = ? ", paymentMethod).
		ToSql()
	fmt.Println(query)
	fmt.Println(args, err)

	data := PaymentMethodUpdateModel{}
	queryU, argsU, errU := squirrel.
		Update("payment_method").
		Set("status", data.Status).
		Set("payment_expiration", data.PaymentExpiration).
		Set("payment_method_name", data.PaymentMethodName).
		Set("payment_method_order", data.PaymentMethodOrder).
		Set("note", data.Note.String).
		Where("payment_method = ?", paymentMethod).
		ToSql()
	fmt.Println(queryU)
	fmt.Println(argsU, errU)
	return true
}
