// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/AlecAivazis/jeeves/db/guild"
	"github.com/AlecAivazis/jeeves/db/guildbank"
	"github.com/AlecAivazis/jeeves/db/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GuildQuery is the builder for querying Guild entities.
type GuildQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Guild
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (gq *GuildQuery) Where(ps ...predicate.Guild) *GuildQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GuildQuery) Limit(limit int) *GuildQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GuildQuery) Offset(offset int) *GuildQuery {
	gq.offset = &offset
	return gq
}

// Order adds an order step to the query.
func (gq *GuildQuery) Order(o ...Order) *GuildQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryBank chains the current query on the bank edge.
func (gq *GuildQuery) QueryBank() *GuildBankQuery {
	query := &GuildBankQuery{config: gq.config}

	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(guildbank.Table)
	t2 := gq.sqlQuery()
	t2.Select(t2.C(guild.FieldID))
	query.sql = builder.Select().
		From(t1).
		Join(t2).
		On(t1.C(guild.BankColumn), t2.C(guild.FieldID))
	return query
}

// First returns the first Guild entity in the query. Returns *ErrNotFound when no guild was found.
func (gq *GuildQuery) First(ctx context.Context) (*Guild, error) {
	gus, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gus) == 0 {
		return nil, &ErrNotFound{guild.Label}
	}
	return gus[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GuildQuery) FirstX(ctx context.Context) *Guild {
	gu, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gu
}

// FirstID returns the first Guild id in the query. Returns *ErrNotFound when no id was found.
func (gq *GuildQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{guild.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gq *GuildQuery) FirstXID(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Guild entity in the query, returns an error if not exactly one entity was returned.
func (gq *GuildQuery) Only(ctx context.Context) (*Guild, error) {
	gus, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gus) {
	case 1:
		return gus[0], nil
	case 0:
		return nil, &ErrNotFound{guild.Label}
	default:
		return nil, &ErrNotSingular{guild.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GuildQuery) OnlyX(ctx context.Context) *Guild {
	gu, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gu
}

// OnlyID returns the only Guild id in the query, returns an error if not exactly one id was returned.
func (gq *GuildQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{guild.Label}
	default:
		err = &ErrNotSingular{guild.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (gq *GuildQuery) OnlyXID(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Guilds.
func (gq *GuildQuery) All(ctx context.Context) ([]*Guild, error) {
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GuildQuery) AllX(ctx context.Context) []*Guild {
	gus, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gus
}

// IDs executes the query and returns a list of Guild ids.
func (gq *GuildQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(guild.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GuildQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GuildQuery) Count(ctx context.Context) (int, error) {
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GuildQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GuildQuery) Exist(ctx context.Context) (bool, error) {
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GuildQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GuildQuery) Clone() *GuildQuery {
	return &GuildQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		order:      append([]Order{}, gq.order...),
		unique:     append([]string{}, gq.unique...),
		predicates: append([]predicate.Guild{}, gq.predicates...),
		// clone intermediate queries.
		sql: gq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discordID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Guild.Query().
//		GroupBy(guild.FieldDiscordID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (gq *GuildQuery) GroupBy(field string, fields ...string) *GuildGroupBy {
	group := &GuildGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = gq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		DiscordID string `json:"discordID,omitempty"`
//	}
//
//	client.Guild.Query().
//		Select(guild.FieldDiscordID).
//		Scan(ctx, &v)
//
func (gq *GuildQuery) Select(field string, fields ...string) *GuildSelect {
	selector := &GuildSelect{config: gq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = gq.sqlQuery()
	return selector
}

func (gq *GuildQuery) sqlAll(ctx context.Context) ([]*Guild, error) {
	rows := &sql.Rows{}
	selector := gq.sqlQuery()
	if unique := gq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := gq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var gus Guilds
	if err := gus.FromRows(rows); err != nil {
		return nil, err
	}
	gus.config(gq.config)
	return gus, nil
}

func (gq *GuildQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := gq.sqlQuery()
	unique := []string{guild.FieldID}
	if len(gq.unique) > 0 {
		unique = gq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := gq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("db: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("db: failed reading count: %v", err)
	}
	return n, nil
}

func (gq *GuildQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %v", err)
	}
	return n > 0, nil
}

func (gq *GuildQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(guild.Table)
	selector := builder.Select(t1.Columns(guild.Columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(guild.Columns...)...)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GuildGroupBy is the builder for group-by Guild entities.
type GuildGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GuildGroupBy) Aggregate(fns ...Aggregate) *GuildGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scan the result into the given value.
func (ggb *GuildGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GuildGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ggb *GuildGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("db: GuildGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GuildGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ggb *GuildGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("db: GuildGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GuildGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ggb *GuildGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("db: GuildGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GuildGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ggb *GuildGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("db: GuildGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GuildGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GuildGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ggb.sqlQuery().Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GuildGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql
	columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
	columns = append(columns, ggb.fields...)
	for _, fn := range ggb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(ggb.fields...)
}

// GuildSelect is the builder for select fields of Guild entities.
type GuildSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (gs *GuildSelect) Scan(ctx context.Context, v interface{}) error {
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GuildSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("db: GuildSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GuildSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("db: GuildSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GuildSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("db: GuildSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GuildSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gs *GuildSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("db: GuildSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GuildSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GuildSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sqlQuery().Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gs *GuildSelect) sqlQuery() sql.Querier {
	view := "guild_view"
	return sql.Dialect(gs.driver.Dialect()).
		Select(gs.fields...).From(gs.sql.As(view))
}
