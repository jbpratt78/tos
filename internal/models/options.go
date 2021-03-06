// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Option is an object representing the database table.
type Option struct {
	ID         null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	Deleted    null.Int64 `boil:"deleted" json:"deleted,omitempty" toml:"deleted" yaml:"deleted,omitempty"`
	Available  null.Int64 `boil:"available" json:"available,omitempty" toml:"available" yaml:"available,omitempty"`
	KindID     int64      `boil:"kind_id" json:"kindID" toml:"kindID" yaml:"kindID"`
	Name       string     `boil:"name" json:"name" toml:"name" yaml:"name"`
	Price      int64      `boil:"price" json:"price" toml:"price" yaml:"price"`
	LightPrice null.Int64 `boil:"light_price" json:"lightPrice,omitempty" toml:"lightPrice" yaml:"lightPrice,omitempty"`
	HeavyPrice null.Int64 `boil:"heavy_price" json:"heavyPrice,omitempty" toml:"heavyPrice" yaml:"heavyPrice,omitempty"`

	R *optionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L optionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OptionColumns = struct {
	ID         string
	Deleted    string
	Available  string
	KindID     string
	Name       string
	Price      string
	LightPrice string
	HeavyPrice string
}{
	ID:         "id",
	Deleted:    "deleted",
	Available:  "available",
	KindID:     "kind_id",
	Name:       "name",
	Price:      "price",
	LightPrice: "light_price",
	HeavyPrice: "heavy_price",
}

// Generated where

var OptionWhere = struct {
	ID         whereHelpernull_Int64
	Deleted    whereHelpernull_Int64
	Available  whereHelpernull_Int64
	KindID     whereHelperint64
	Name       whereHelperstring
	Price      whereHelperint64
	LightPrice whereHelpernull_Int64
	HeavyPrice whereHelpernull_Int64
}{
	ID:         whereHelpernull_Int64{field: "\"options\".\"id\""},
	Deleted:    whereHelpernull_Int64{field: "\"options\".\"deleted\""},
	Available:  whereHelpernull_Int64{field: "\"options\".\"available\""},
	KindID:     whereHelperint64{field: "\"options\".\"kind_id\""},
	Name:       whereHelperstring{field: "\"options\".\"name\""},
	Price:      whereHelperint64{field: "\"options\".\"price\""},
	LightPrice: whereHelpernull_Int64{field: "\"options\".\"light_price\""},
	HeavyPrice: whereHelpernull_Int64{field: "\"options\".\"heavy_price\""},
}

// OptionRels is where relationship names are stored.
var OptionRels = struct {
	Kind             string
	ItemOption       string
	OrderItemOptions string
}{
	Kind:             "Kind",
	ItemOption:       "ItemOption",
	OrderItemOptions: "OrderItemOptions",
}

// optionR is where relationships are stored.
type optionR struct {
	Kind             *OptionKind          `boil:"Kind" json:"Kind" toml:"Kind" yaml:"Kind"`
	ItemOption       *ItemOption          `boil:"ItemOption" json:"ItemOption" toml:"ItemOption" yaml:"ItemOption"`
	OrderItemOptions OrderItemOptionSlice `boil:"OrderItemOptions" json:"OrderItemOptions" toml:"OrderItemOptions" yaml:"OrderItemOptions"`
}

// NewStruct creates a new relationship struct
func (*optionR) NewStruct() *optionR {
	return &optionR{}
}

// optionL is where Load methods for each relationship are stored.
type optionL struct{}

var (
	optionAllColumns            = []string{"id", "deleted", "available", "kind_id", "name", "price", "light_price", "heavy_price"}
	optionColumnsWithoutDefault = []string{}
	optionColumnsWithDefault    = []string{"id", "deleted", "available", "kind_id", "name", "price", "light_price", "heavy_price"}
	optionPrimaryKeyColumns     = []string{"id"}
)

type (
	// OptionSlice is an alias for a slice of pointers to Option.
	// This should generally be used opposed to []Option.
	OptionSlice []*Option

	optionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	optionType                 = reflect.TypeOf(&Option{})
	optionMapping              = queries.MakeStructMapping(optionType)
	optionPrimaryKeyMapping, _ = queries.BindMapping(optionType, optionMapping, optionPrimaryKeyColumns)
	optionInsertCacheMut       sync.RWMutex
	optionInsertCache          = make(map[string]insertCache)
	optionUpdateCacheMut       sync.RWMutex
	optionUpdateCache          = make(map[string]updateCache)
	optionUpsertCacheMut       sync.RWMutex
	optionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single option record from the query using the global executor.
func (q optionQuery) OneG(ctx context.Context) (*Option, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single option record from the query.
func (q optionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Option, error) {
	o := &Option{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for options")
	}

	return o, nil
}

// AllG returns all Option records from the query using the global executor.
func (q optionQuery) AllG(ctx context.Context) (OptionSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Option records from the query.
func (q optionQuery) All(ctx context.Context, exec boil.ContextExecutor) (OptionSlice, error) {
	var o []*Option

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Option slice")
	}

	return o, nil
}

// CountG returns the count of all Option records in the query, and panics on error.
func (q optionQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Option records in the query.
func (q optionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count options rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q optionQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q optionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if options exists")
	}

	return count > 0, nil
}

// Kind pointed to by the foreign key.
func (o *Option) Kind(mods ...qm.QueryMod) optionKindQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.KindID),
	}

	queryMods = append(queryMods, mods...)

	query := OptionKinds(queryMods...)
	queries.SetFrom(query.Query, "\"option_kinds\"")

	return query
}

// ItemOption pointed to by the foreign key.
func (o *Option) ItemOption(mods ...qm.QueryMod) itemOptionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"option_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := ItemOptions(queryMods...)
	queries.SetFrom(query.Query, "\"item_options\"")

	return query
}

// OrderItemOptions retrieves all the order_item_option's OrderItemOptions with an executor.
func (o *Option) OrderItemOptions(mods ...qm.QueryMod) orderItemOptionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"order_item_options\".\"option_id\"=?", o.ID),
	)

	query := OrderItemOptions(queryMods...)
	queries.SetFrom(query.Query, "\"order_item_options\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"order_item_options\".*"})
	}

	return query
}

// LoadKind allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (optionL) LoadKind(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOption interface{}, mods queries.Applicator) error {
	var slice []*Option
	var object *Option

	if singular {
		object = maybeOption.(*Option)
	} else {
		slice = *maybeOption.(*[]*Option)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionR{}
		}
		if !queries.IsNil(object.KindID) {
			args = append(args, object.KindID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.KindID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.KindID) {
				args = append(args, obj.KindID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`option_kinds`),
		qm.WhereIn(`option_kinds.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OptionKind")
	}

	var resultSlice []*OptionKind
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OptionKind")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for option_kinds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for option_kinds")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Kind = foreign
		if foreign.R == nil {
			foreign.R = &optionKindR{}
		}
		foreign.R.KindOptions = append(foreign.R.KindOptions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.KindID, foreign.ID) {
				local.R.Kind = foreign
				if foreign.R == nil {
					foreign.R = &optionKindR{}
				}
				foreign.R.KindOptions = append(foreign.R.KindOptions, local)
				break
			}
		}
	}

	return nil
}

// LoadItemOption allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (optionL) LoadItemOption(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOption interface{}, mods queries.Applicator) error {
	var slice []*Option
	var object *Option

	if singular {
		object = maybeOption.(*Option)
	} else {
		slice = *maybeOption.(*[]*Option)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`item_options`),
		qm.WhereIn(`item_options.option_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ItemOption")
	}

	var resultSlice []*ItemOption
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ItemOption")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for item_options")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for item_options")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ItemOption = foreign
		if foreign.R == nil {
			foreign.R = &itemOptionR{}
		}
		foreign.R.Option = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ID, foreign.OptionID) {
				local.R.ItemOption = foreign
				if foreign.R == nil {
					foreign.R = &itemOptionR{}
				}
				foreign.R.Option = local
				break
			}
		}
	}

	return nil
}

// LoadOrderItemOptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (optionL) LoadOrderItemOptions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOption interface{}, mods queries.Applicator) error {
	var slice []*Option
	var object *Option

	if singular {
		object = maybeOption.(*Option)
	} else {
		slice = *maybeOption.(*[]*Option)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`order_item_options`),
		qm.WhereIn(`order_item_options.option_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load order_item_options")
	}

	var resultSlice []*OrderItemOption
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice order_item_options")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on order_item_options")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for order_item_options")
	}

	if singular {
		object.R.OrderItemOptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderItemOptionR{}
			}
			foreign.R.Option = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.OptionID) {
				local.R.OrderItemOptions = append(local.R.OrderItemOptions, foreign)
				if foreign.R == nil {
					foreign.R = &orderItemOptionR{}
				}
				foreign.R.Option = local
				break
			}
		}
	}

	return nil
}

// SetKindG of the option to the related item.
// Sets o.R.Kind to related.
// Adds o to related.R.KindOptions.
// Uses the global database handle.
func (o *Option) SetKindG(ctx context.Context, insert bool, related *OptionKind) error {
	return o.SetKind(ctx, boil.GetContextDB(), insert, related)
}

// SetKind of the option to the related item.
// Sets o.R.Kind to related.
// Adds o to related.R.KindOptions.
func (o *Option) SetKind(ctx context.Context, exec boil.ContextExecutor, insert bool, related *OptionKind) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"options\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"kind_id"}),
		strmangle.WhereClause("\"", "\"", 0, optionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.KindID, related.ID)
	if o.R == nil {
		o.R = &optionR{
			Kind: related,
		}
	} else {
		o.R.Kind = related
	}

	if related.R == nil {
		related.R = &optionKindR{
			KindOptions: OptionSlice{o},
		}
	} else {
		related.R.KindOptions = append(related.R.KindOptions, o)
	}

	return nil
}

// SetItemOptionG of the option to the related item.
// Sets o.R.ItemOption to related.
// Adds o to related.R.Option.
// Uses the global database handle.
func (o *Option) SetItemOptionG(ctx context.Context, insert bool, related *ItemOption) error {
	return o.SetItemOption(ctx, boil.GetContextDB(), insert, related)
}

// SetItemOption of the option to the related item.
// Sets o.R.ItemOption to related.
// Adds o to related.R.Option.
func (o *Option) SetItemOption(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ItemOption) error {
	var err error

	if insert {
		queries.Assign(&related.OptionID, o.ID)

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"item_options\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, []string{"option_id"}),
			strmangle.WhereClause("\"", "\"", 0, itemOptionPrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ItemID, related.OptionID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		queries.Assign(&related.OptionID, o.ID)
	}

	if o.R == nil {
		o.R = &optionR{
			ItemOption: related,
		}
	} else {
		o.R.ItemOption = related
	}

	if related.R == nil {
		related.R = &itemOptionR{
			Option: o,
		}
	} else {
		related.R.Option = o
	}
	return nil
}

// AddOrderItemOptionsG adds the given related objects to the existing relationships
// of the option, optionally inserting them as new records.
// Appends related to o.R.OrderItemOptions.
// Sets related.R.Option appropriately.
// Uses the global database handle.
func (o *Option) AddOrderItemOptionsG(ctx context.Context, insert bool, related ...*OrderItemOption) error {
	return o.AddOrderItemOptions(ctx, boil.GetContextDB(), insert, related...)
}

// AddOrderItemOptions adds the given related objects to the existing relationships
// of the option, optionally inserting them as new records.
// Appends related to o.R.OrderItemOptions.
// Sets related.R.Option appropriately.
func (o *Option) AddOrderItemOptions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OrderItemOption) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.OptionID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"order_item_options\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"option_id"}),
				strmangle.WhereClause("\"", "\"", 0, orderItemOptionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.OptionID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &optionR{
			OrderItemOptions: related,
		}
	} else {
		o.R.OrderItemOptions = append(o.R.OrderItemOptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderItemOptionR{
				Option: o,
			}
		} else {
			rel.R.Option = o
		}
	}
	return nil
}

// Options retrieves all the records using an executor.
func Options(mods ...qm.QueryMod) optionQuery {
	mods = append(mods, qm.From("\"options\""))
	return optionQuery{NewQuery(mods...)}
}

// FindOptionG retrieves a single record by ID.
func FindOptionG(ctx context.Context, iD null.Int64, selectCols ...string) (*Option, error) {
	return FindOption(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOption retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOption(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Option, error) {
	optionObj := &Option{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"options\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, optionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from options")
	}

	return optionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Option) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Option) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no options provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(optionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	optionInsertCacheMut.RLock()
	cache, cached := optionInsertCache[key]
	optionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			optionAllColumns,
			optionColumnsWithDefault,
			optionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(optionType, optionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(optionType, optionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"options\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"options\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"options\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, optionPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into options")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for options")
	}

CacheNoHooks:
	if !cached {
		optionInsertCacheMut.Lock()
		optionInsertCache[key] = cache
		optionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Option record using the global executor.
// See Update for more documentation.
func (o *Option) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Option.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Option) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	optionUpdateCacheMut.RLock()
	cache, cached := optionUpdateCache[key]
	optionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			optionAllColumns,
			optionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update options, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"options\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, optionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(optionType, optionMapping, append(wl, optionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update options row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for options")
	}

	if !cached {
		optionUpdateCacheMut.Lock()
		optionUpdateCache[key] = cache
		optionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q optionQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q optionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for options")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OptionSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OptionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"options\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in option slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all option")
	}
	return rowsAff, nil
}

// DeleteG deletes a single Option record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Option) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Option record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Option) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Option provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), optionPrimaryKeyMapping)
	sql := "DELETE FROM \"options\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for options")
	}

	return rowsAff, nil
}

func (q optionQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q optionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no optionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from options")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for options")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OptionSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OptionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"options\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from option slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for options")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Option) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Option provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Option) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOption(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OptionSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty OptionSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OptionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OptionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"options\".* FROM \"options\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OptionSlice")
	}

	*o = slice

	return nil
}

// OptionExistsG checks if the Option row exists.
func OptionExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return OptionExists(ctx, boil.GetContextDB(), iD)
}

// OptionExists checks if the Option row exists.
func OptionExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"options\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if options exists")
	}

	return exists, nil
}
