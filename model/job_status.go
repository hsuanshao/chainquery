// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// JobStatus is an object representing the database table.
type JobStatus struct {
	JobName      string      `boil:"job_name" json:"job_name" toml:"job_name" yaml:"job_name"`
	LastSync     time.Time   `boil:"last_sync" json:"last_sync" toml:"last_sync" yaml:"last_sync"`
	IsSuccess    bool        `boil:"is_success" json:"is_success" toml:"is_success" yaml:"is_success"`
	ErrorMessage null.String `boil:"error_message" json:"error_message,omitempty" toml:"error_message" yaml:"error_message,omitempty"`

	R *jobStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jobStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JobStatusColumns = struct {
	JobName      string
	LastSync     string
	IsSuccess    string
	ErrorMessage string
}{
	JobName:      "job_name",
	LastSync:     "last_sync",
	IsSuccess:    "is_success",
	ErrorMessage: "error_message",
}

// jobStatusR is where relationships are stored.
type jobStatusR struct {
}

// jobStatusL is where Load methods for each relationship are stored.
type jobStatusL struct{}

var (
	jobStatusColumns               = []string{"job_name", "last_sync", "is_success", "error_message"}
	jobStatusColumnsWithoutDefault = []string{"job_name", "error_message"}
	jobStatusColumnsWithDefault    = []string{"last_sync", "is_success"}
	jobStatusPrimaryKeyColumns     = []string{"job_name"}
)

type (
	// JobStatusSlice is an alias for a slice of pointers to JobStatus.
	// This should generally be used opposed to []JobStatus.
	JobStatusSlice []*JobStatus

	jobStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jobStatusType                 = reflect.TypeOf(&JobStatus{})
	jobStatusMapping              = queries.MakeStructMapping(jobStatusType)
	jobStatusPrimaryKeyMapping, _ = queries.BindMapping(jobStatusType, jobStatusMapping, jobStatusPrimaryKeyColumns)
	jobStatusInsertCacheMut       sync.RWMutex
	jobStatusInsertCache          = make(map[string]insertCache)
	jobStatusUpdateCacheMut       sync.RWMutex
	jobStatusUpdateCache          = make(map[string]updateCache)
	jobStatusUpsertCacheMut       sync.RWMutex
	jobStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single jobStatus record from the query, and panics on error.
func (q jobStatusQuery) OneP() *JobStatus {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single jobStatus record from the query.
func (q jobStatusQuery) One() (*JobStatus, error) {
	o := &JobStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for job_status")
	}

	return o, nil
}

// AllP returns all JobStatus records from the query, and panics on error.
func (q jobStatusQuery) AllP() JobStatusSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all JobStatus records from the query.
func (q jobStatusQuery) All() (JobStatusSlice, error) {
	var o []*JobStatus

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to JobStatus slice")
	}

	return o, nil
}

// CountP returns the count of all JobStatus records in the query, and panics on error.
func (q jobStatusQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all JobStatus records in the query.
func (q jobStatusQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count job_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q jobStatusQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q jobStatusQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if job_status exists")
	}

	return count > 0, nil
}

// JobStatusesG retrieves all records.
func JobStatusesG(mods ...qm.QueryMod) jobStatusQuery {
	return JobStatuses(boil.GetDB(), mods...)
}

// JobStatuses retrieves all the records using an executor.
func JobStatuses(exec boil.Executor, mods ...qm.QueryMod) jobStatusQuery {
	mods = append(mods, qm.From("`job_status`"))
	return jobStatusQuery{NewQuery(exec, mods...)}
}

// FindJobStatusG retrieves a single record by ID.
func FindJobStatusG(jobName string, selectCols ...string) (*JobStatus, error) {
	return FindJobStatus(boil.GetDB(), jobName, selectCols...)
}

// FindJobStatusGP retrieves a single record by ID, and panics on error.
func FindJobStatusGP(jobName string, selectCols ...string) *JobStatus {
	retobj, err := FindJobStatus(boil.GetDB(), jobName, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindJobStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJobStatus(exec boil.Executor, jobName string, selectCols ...string) (*JobStatus, error) {
	jobStatusObj := &JobStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `job_status` where `job_name`=?", sel,
	)

	q := queries.Raw(exec, query, jobName)

	err := q.Bind(jobStatusObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from job_status")
	}

	return jobStatusObj, nil
}

// FindJobStatusP retrieves a single record by ID with an executor, and panics on error.
func FindJobStatusP(exec boil.Executor, jobName string, selectCols ...string) *JobStatus {
	retobj, err := FindJobStatus(exec, jobName, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *JobStatus) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *JobStatus) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *JobStatus) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *JobStatus) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no job_status provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(jobStatusColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	jobStatusInsertCacheMut.RLock()
	cache, cached := jobStatusInsertCache[key]
	jobStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			jobStatusColumns,
			jobStatusColumnsWithDefault,
			jobStatusColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(jobStatusType, jobStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jobStatusType, jobStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `job_status` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `job_status` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `job_status` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, jobStatusPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "model: unable to insert into job_status")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.JobName,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for job_status")
	}

CacheNoHooks:
	if !cached {
		jobStatusInsertCacheMut.Lock()
		jobStatusInsertCache[key] = cache
		jobStatusInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single JobStatus record. See Update for
// whitelist behavior description.
func (o *JobStatus) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single JobStatus record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *JobStatus) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the JobStatus, and panics on error.
// See Update for whitelist behavior description.
func (o *JobStatus) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the JobStatus.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *JobStatus) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	jobStatusUpdateCacheMut.RLock()
	cache, cached := jobStatusUpdateCache[key]
	jobStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			jobStatusColumns,
			jobStatusPrimaryKeyColumns,
			whitelist,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update job_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `job_status` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, jobStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jobStatusType, jobStatusMapping, append(wl, jobStatusPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update job_status row")
	}

	if !cached {
		jobStatusUpdateCacheMut.Lock()
		jobStatusUpdateCache[key] = cache
		jobStatusUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q jobStatusQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q jobStatusQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for job_status")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JobStatusSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o JobStatusSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o JobStatusSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JobStatusSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `job_status` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobStatusPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in jobStatus slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *JobStatus) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *JobStatus) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *JobStatus) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *JobStatus) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no job_status provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(jobStatusColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	jobStatusUpsertCacheMut.RLock()
	cache, cached := jobStatusUpsertCache[key]
	jobStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			jobStatusColumns,
			jobStatusColumnsWithDefault,
			jobStatusColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			jobStatusColumns,
			jobStatusPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("model: unable to upsert job_status, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "job_status", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `job_status` WHERE `job_name`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(jobStatusType, jobStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jobStatusType, jobStatusMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for job_status")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.JobName,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for job_status")
	}

CacheNoHooks:
	if !cached {
		jobStatusUpsertCacheMut.Lock()
		jobStatusUpsertCache[key] = cache
		jobStatusUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single JobStatus record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JobStatus) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single JobStatus record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *JobStatus) DeleteG() error {
	if o == nil {
		return errors.New("model: no JobStatus provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single JobStatus record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *JobStatus) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single JobStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *JobStatus) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no JobStatus provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jobStatusPrimaryKeyMapping)
	sql := "DELETE FROM `job_status` WHERE `job_name`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from job_status")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q jobStatusQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q jobStatusQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("model: no jobStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from job_status")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o JobStatusSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o JobStatusSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("model: no JobStatus slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o JobStatusSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JobStatusSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no JobStatus slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `job_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobStatusPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from jobStatus slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *JobStatus) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *JobStatus) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *JobStatus) ReloadG() error {
	if o == nil {
		return errors.New("model: no JobStatus provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *JobStatus) Reload(exec boil.Executor) error {
	ret, err := FindJobStatus(exec, o.JobName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JobStatusSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *JobStatusSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JobStatusSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty JobStatusSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JobStatusSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	jobStatuses := JobStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `job_status`.* FROM `job_status` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&jobStatuses)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in JobStatusSlice")
	}

	*o = jobStatuses

	return nil
}

// JobStatusExists checks if the JobStatus row exists.
func JobStatusExists(exec boil.Executor, jobName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `job_status` where `job_name`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, jobName)
	}

	row := exec.QueryRow(sql, jobName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if job_status exists")
	}

	return exists, nil
}

// JobStatusExistsG checks if the JobStatus row exists.
func JobStatusExistsG(jobName string) (bool, error) {
	return JobStatusExists(boil.GetDB(), jobName)
}

// JobStatusExistsGP checks if the JobStatus row exists. Panics on error.
func JobStatusExistsGP(jobName string) bool {
	e, err := JobStatusExists(boil.GetDB(), jobName)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// JobStatusExistsP checks if the JobStatus row exists. Panics on error.
func JobStatusExistsP(exec boil.Executor, jobName string) bool {
	e, err := JobStatusExists(exec, jobName)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}