// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testBlocks(t *testing.T) {
	t.Parallel()

	query := Blocks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBlocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = block.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blocks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{block}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBlocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BlockExists(tx, block.Hash)
	if err != nil {
		t.Errorf("Unable to check if Block exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlockExistsG to return true, but got false.")
	}
}
func testBlocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	blockFound, err := FindBlock(tx, block.Hash)
	if err != nil {
		t.Error(err)
	}

	if blockFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBlocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blocks(tx).Bind(block); err != nil {
		t.Error(err)
	}
}

func testBlocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Blocks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBlocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx, blockColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlockToManyNextBlockBlocks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, blockDBTypes, false, blockColumnsWithDefault...)
	randomize.Struct(seed, &c, blockDBTypes, false, blockColumnsWithDefault...)

	b.NextBlockID.Valid = true
	c.NextBlockID.Valid = true
	b.NextBlockID.String = a.Hash
	c.NextBlockID.String = a.Hash
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	block, err := a.NextBlockBlocks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range block {
		if v.NextBlockID.String == b.NextBlockID.String {
			bFound = true
		}
		if v.NextBlockID.String == c.NextBlockID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlockSlice{&a}
	if err = a.L.LoadNextBlockBlocks(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NextBlockBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.NextBlockBlocks = nil
	if err = a.L.LoadNextBlockBlocks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NextBlockBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", block)
	}
}

func testBlockToManyPreviousBlockBlocks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, blockDBTypes, false, blockColumnsWithDefault...)
	randomize.Struct(seed, &c, blockDBTypes, false, blockColumnsWithDefault...)

	b.PreviousBlockID.Valid = true
	c.PreviousBlockID.Valid = true
	b.PreviousBlockID.String = a.Hash
	c.PreviousBlockID.String = a.Hash
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	block, err := a.PreviousBlockBlocks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range block {
		if v.PreviousBlockID.String == b.PreviousBlockID.String {
			bFound = true
		}
		if v.PreviousBlockID.String == c.PreviousBlockID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlockSlice{&a}
	if err = a.L.LoadPreviousBlockBlocks(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PreviousBlockBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PreviousBlockBlocks = nil
	if err = a.L.LoadPreviousBlockBlocks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PreviousBlockBlocks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", block)
	}
}

func testBlockToManyTransactions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, transactionDBTypes, false, transactionColumnsWithDefault...)
	randomize.Struct(seed, &c, transactionDBTypes, false, transactionColumnsWithDefault...)

	b.BlockID = a.Hash
	c.BlockID = a.Hash
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	transaction, err := a.Transactions(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range transaction {
		if v.BlockID == b.BlockID {
			bFound = true
		}
		if v.BlockID == c.BlockID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlockSlice{&a}
	if err = a.L.LoadTransactions(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Transactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Transactions = nil
	if err = a.L.LoadTransactions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Transactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", transaction)
	}
}

func testBlockToManyAddOpNextBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Block{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNextBlockBlocks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Hash != first.NextBlockID.String {
			t.Error("foreign key was wrong value", a.Hash, first.NextBlockID.String)
		}
		if a.Hash != second.NextBlockID.String {
			t.Error("foreign key was wrong value", a.Hash, second.NextBlockID.String)
		}

		if first.R.NextBlock != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.NextBlock != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.NextBlockBlocks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.NextBlockBlocks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.NextBlockBlocks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBlockToManySetOpNextBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetNextBlockBlocks(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.NextBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetNextBlockBlocks(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.NextBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.NextBlockID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.NextBlockID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.Hash != d.NextBlockID.String {
		t.Error("foreign key was wrong value", a.Hash, d.NextBlockID.String)
	}
	if a.Hash != e.NextBlockID.String {
		t.Error("foreign key was wrong value", a.Hash, e.NextBlockID.String)
	}

	if b.R.NextBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.NextBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.NextBlock != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.NextBlock != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.NextBlockBlocks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.NextBlockBlocks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBlockToManyRemoveOpNextBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddNextBlockBlocks(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.NextBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveNextBlockBlocks(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.NextBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.NextBlockID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.NextBlockID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.NextBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.NextBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.NextBlock != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.NextBlock != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.NextBlockBlocks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.NextBlockBlocks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.NextBlockBlocks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBlockToManyAddOpPreviousBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Block{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPreviousBlockBlocks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Hash != first.PreviousBlockID.String {
			t.Error("foreign key was wrong value", a.Hash, first.PreviousBlockID.String)
		}
		if a.Hash != second.PreviousBlockID.String {
			t.Error("foreign key was wrong value", a.Hash, second.PreviousBlockID.String)
		}

		if first.R.PreviousBlock != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.PreviousBlock != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PreviousBlockBlocks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PreviousBlockBlocks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PreviousBlockBlocks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBlockToManySetOpPreviousBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetPreviousBlockBlocks(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PreviousBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetPreviousBlockBlocks(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PreviousBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PreviousBlockID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PreviousBlockID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.Hash != d.PreviousBlockID.String {
		t.Error("foreign key was wrong value", a.Hash, d.PreviousBlockID.String)
	}
	if a.Hash != e.PreviousBlockID.String {
		t.Error("foreign key was wrong value", a.Hash, e.PreviousBlockID.String)
	}

	if b.R.PreviousBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.PreviousBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.PreviousBlock != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.PreviousBlock != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.PreviousBlockBlocks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.PreviousBlockBlocks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBlockToManyRemoveOpPreviousBlockBlocks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Block{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddPreviousBlockBlocks(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.PreviousBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemovePreviousBlockBlocks(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.PreviousBlockBlocks(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.PreviousBlockID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.PreviousBlockID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.PreviousBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.PreviousBlock != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.PreviousBlock != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.PreviousBlock != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.PreviousBlockBlocks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.PreviousBlockBlocks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.PreviousBlockBlocks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBlockToManyAddOpTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Transaction{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Hash != first.BlockID {
			t.Error("foreign key was wrong value", a.Hash, first.BlockID)
		}
		if a.Hash != second.BlockID {
			t.Error("foreign key was wrong value", a.Hash, second.BlockID)
		}

		if first.R.Block != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Block != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Transactions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Transactions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Transactions(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBlockToOneBlockUsingNextBlock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Block
	var foreign Block

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	local.NextBlockID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.NextBlockID.String = foreign.Hash
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.NextBlock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Hash != foreign.Hash {
		t.Errorf("want: %v, got %v", foreign.Hash, check.Hash)
	}

	slice := BlockSlice{&local}
	if err = local.L.LoadNextBlock(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.NextBlock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.NextBlock = nil
	if err = local.L.LoadNextBlock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.NextBlock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlockToOneBlockUsingPreviousBlock(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Block
	var foreign Block

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	local.PreviousBlockID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.PreviousBlockID.String = foreign.Hash
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.PreviousBlock(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Hash != foreign.Hash {
		t.Errorf("want: %v, got %v", foreign.Hash, check.Hash)
	}

	slice := BlockSlice{&local}
	if err = local.L.LoadPreviousBlock(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.PreviousBlock == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PreviousBlock = nil
	if err = local.L.LoadPreviousBlock(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.PreviousBlock == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlockToOneSetOpBlockUsingNextBlock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Block{&b, &c} {
		err = a.SetNextBlock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.NextBlock != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.NextBlockBlocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.NextBlockID.String != x.Hash {
			t.Error("foreign key was wrong value", a.NextBlockID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.NextBlockID.String))
		reflect.Indirect(reflect.ValueOf(&a.NextBlockID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.NextBlockID.String != x.Hash {
			t.Error("foreign key was wrong value", a.NextBlockID.String, x.Hash)
		}
	}
}

func testBlockToOneRemoveOpBlockUsingNextBlock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetNextBlock(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveNextBlock(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.NextBlock(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.NextBlock != nil {
		t.Error("R struct entry should be nil")
	}

	if a.NextBlockID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.NextBlockBlocks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBlockToOneSetOpBlockUsingPreviousBlock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Block{&b, &c} {
		err = a.SetPreviousBlock(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PreviousBlock != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PreviousBlockBlocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PreviousBlockID.String != x.Hash {
			t.Error("foreign key was wrong value", a.PreviousBlockID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PreviousBlockID.String))
		reflect.Indirect(reflect.ValueOf(&a.PreviousBlockID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PreviousBlockID.String != x.Hash {
			t.Error("foreign key was wrong value", a.PreviousBlockID.String, x.Hash)
		}
	}
}

func testBlockToOneRemoveOpBlockUsingPreviousBlock(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetPreviousBlock(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemovePreviousBlock(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.PreviousBlock(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.PreviousBlock != nil {
		t.Error("R struct entry should be nil")
	}

	if a.PreviousBlockID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PreviousBlockBlocks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBlocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = block.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBlocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{block}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBlocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blockDBTypes = map[string]string{`Bits`: `varchar`, `BlockSize`: `bigint`, `BlockTime`: `bigint`, `Chainwork`: `varchar`, `Confirmations`: `int`, `Difficulty`: `decimal`, `Hash`: `varchar`, `Height`: `bigint`, `MedianTime`: `bigint`, `MerkleRoot`: `varchar`, `NameClaimRoot`: `varchar`, `NextBlockID`: `varchar`, `Nonce`: `bigint`, `PreviousBlockID`: `varchar`, `Target`: `varchar`, `TransactionHashes`: `text`, `TransactionsProcessed`: `tinyint`, `Version`: `bigint`, `VersionHex`: `varchar`}
	_            = bytes.MinRead
)

func testBlocksUpdate(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = block.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBlocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, block, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blockColumns, blockPrimaryKeyColumns) {
		fields = blockColumns
	} else {
		fields = strmangle.SetComplement(
			blockColumns,
			blockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(block))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BlockSlice{block}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBlocksUpsert(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	block := Block{}
	if err = randomize.Struct(seed, &block, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &block, blockDBTypes, false, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = block.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err = Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
