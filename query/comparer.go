package query

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
)

type CustomerPackageListCriteria struct {
	Limit  int64
	Offset int64
	Search string
}

type CustomerPackageBaseModel struct {
	CustomerPackageID    int64           `db:"customer_package_detail_id" json:"-"`
	AccountID            int64           `db:"customer_id" json:"-"`
	ShipmentCategoryID   int64           `db:"shipment_category_id" json:"-"`
	ShipmentCategoryName string          `db:"shipment_category_name" json:"-"`
	OtherCommodities     sql.NullString  `db:"other_commodities" json:"-"`
	ProductType          string          `db:"product_type" json:"-"`
	LabelName            string          `db:"label_name" json:"-"`
	NumberOfPackage      int             `db:"number_of_package" json:"-"`
	TotalWeight          float32         `db:"total_weight" json:"-"`
	DimensionWidth       sql.NullFloat64 `db:"dimension_width" json:"-"`
	DimensionHeight      sql.NullFloat64 `db:"dimension_height" json:"-"`
	DimensionLength      sql.NullFloat64 `db:"dimension_length" json:"-"`
	LastUsed             time.Time       `db:"last_used" json:"-"`
	CreatedAt            time.Time       `db:"created_at" json:"-"`
	UpdatedAt            time.Time       `db:"updated_at" json:"-"`
	GoodValues           sql.NullFloat64 `db:"good_values" json:"-"`
	IsFragile            sql.NullBool    `db:"is_fragile" json:"-"`
}

func CompareQuery() bool {
	customerID := int64(1)
	criteria := &CustomerPackageListCriteria{
		Search: "a",
		Offset: 0,
		Limit:  10,
	}
	// Building the q_data part
	qData := squirrel.Select("cpd.*, sc.shipment_category_name").
		From("customer_package_detail AS cpd").
		Join("shipment_category AS sc ON cpd.shipment_category_id = sc.shipment_category_id").
		Where(squirrel.And{
			squirrel.Expr("LOWER(cpd.label_name) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(criteria.Search))),
			squirrel.Expr("LOWER(sc.shipment_category_name) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(criteria.Search))),
			squirrel.Eq{"cpd.customer_id": customerID},
		})

	// Building the q_result part
	qResult := squirrel.Select("*").
		FromSelect(qData, "q_data").
		OrderBy("last_used DESC NULLS LAST, updated_at DESC").
		Limit(uint64(criteria.Limit)).
		Offset(uint64(criteria.Offset))

	// Building the q_total part
	qTotal := squirrel.Select("count(customer_package_detail_id) AS __total__").
		FromSelect(qData, "q_data")

	// Combining the results
	qBuilder := squirrel.Select("*").
		PlaceholderFormat(squirrel.Dollar).
		FromSelect(qResult, "q_result").
		SuffixExpr(squirrel.ConcatExpr(", (", qTotal, ") as q_total"))

	result, args, err := qBuilder.ToSql()

	fmt.Println(result)
	fmt.Println(err, args)

	return true
}

func QueryDelete() bool {
	customerPackageID := int64(1)
	customerID := int64(1)
	query := squirrel.Delete("customer_package_detail").
		Where(squirrel.Eq{"customer_package_detail_id": customerPackageID, "customer_id": customerID})

	resultd, argsd, errd := query.ToSql()

	fmt.Println(resultd)
	fmt.Println(errd, argsd)

	currentTime := time.Now().UTC()
	custPackageID := 1
	queryUpdate := squirrel.Update("customer_package_detail").
		Set("last_used", currentTime).
		Where(squirrel.Eq{"customer_package_detail_id": custPackageID, "customer_id": customerID})

	resultu, argsu, erru := queryUpdate.ToSql()

	fmt.Println(resultu)
	fmt.Println(erru, argsu)

	data := []CustomerPackageBaseModel{
		{
			CustomerPackageID: 0,
		},
		{
			CustomerPackageID: 1,
		},
	}

	for _, m := range data {
		if m.CustomerPackageID == 0 {
			query, args, err := squirrel.Insert("customer_package_detail").
				Columns(
					"customer_id",
					"shipment_category_id",
					"product_type",
					"label_name",
					"number_of_package",
					"total_weight",
					"dimension_width",
					"dimension_height",
					"dimension_length",
					"other_commodities",
					"good_values",
					"is_fragile").
				Values(
					m.AccountID,
					m.ShipmentCategoryID,
					m.ProductType,
					m.LabelName,
					m.NumberOfPackage,
					m.TotalWeight,
					m.DimensionWidth,
					m.DimensionHeight,
					m.DimensionLength,
					m.OtherCommodities,
					m.GoodValues,
					m.IsFragile).
				Suffix("RETURNING customer_package_detail_id").
				ToSql()
			fmt.Println(query)
			fmt.Println(args, err)
		} else {
			query, args, err := squirrel.Update("customer_package_detail").
				Set("shipment_category_id", m.ShipmentCategoryID).
				Set("product_type", m.ProductType).
				Set("label_name", m.LabelName).
				Set("number_of_package", m.NumberOfPackage).
				Set("total_weight", m.TotalWeight).
				Set("dimension_width", m.DimensionWidth).
				Set("dimension_height", m.DimensionHeight).
				Set("dimension_length", m.DimensionLength).
				Set("updated_at", currentTime).
				Set("other_commodities", m.OtherCommodities).
				Set("good_values", m.GoodValues).
				Set("is_fragile", m.IsFragile).
				Where(squirrel.Eq{"customer_package_detail_id": m.CustomerPackageID, "customer_id": m.AccountID}).
				ToSql()
			fmt.Println(query)
			fmt.Println(args, err)
		}
	}

	querys, argss, errs := squirrel.Select("COUNT(customer_package_detail_id)").
		From("customer_package_detail").
		Where(squirrel.Eq{"customer_id": customerID}).ToSql()
	fmt.Println(querys)
	fmt.Println(argss, errs)

	customerPackageIDs := []string{"1", "2"}
	joinedIDs := strings.Join(customerPackageIDs, ",")
	querys, argss, errs = squirrel.Select(fmt.Sprintf("COUNT(DISTINCT customer_package_detail_id) = %d", len(customerPackageIDs))).
		From("customer_package_detail").
		Where(squirrel.Expr("customer_package_detail_id = ANY(string_to_array(?, ',')::int[])", joinedIDs)).
		Where(squirrel.Eq{"customer_id": customerID}).ToSql()
	fmt.Println(querys)
	fmt.Println(argss, errs)

	querys, argss, errs = squirrel.Select("cpd.*, sc.shipment_category_name").
		From("customer_package_detail AS cpd").
		Join("shipment_category AS sc ON cpd.shipment_category_id = sc.shipment_category_id").
		Where(squirrel.Eq{"cpd.customer_package_detail_id": customerPackageID, "cpd.customer_id": customerID}).
		Limit(1).ToSql()
	fmt.Println(querys)
	fmt.Println(argss, errs)
	return true
}
