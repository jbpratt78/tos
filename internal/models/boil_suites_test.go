// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ItemKinds", testItemKinds)
	t.Run("ItemOptions", testItemOptions)
	t.Run("ItemSides", testItemSides)
	t.Run("Items", testItems)
	t.Run("OptionKinds", testOptionKinds)
	t.Run("Options", testOptions)
	t.Run("OrderItemOptions", testOrderItemOptions)
	t.Run("OrderItems", testOrderItems)
	t.Run("Orders", testOrders)
}

func TestDelete(t *testing.T) {
	t.Run("ItemKinds", testItemKindsDelete)
	t.Run("ItemOptions", testItemOptionsDelete)
	t.Run("ItemSides", testItemSidesDelete)
	t.Run("Items", testItemsDelete)
	t.Run("OptionKinds", testOptionKindsDelete)
	t.Run("Options", testOptionsDelete)
	t.Run("OrderItemOptions", testOrderItemOptionsDelete)
	t.Run("OrderItems", testOrderItemsDelete)
	t.Run("Orders", testOrdersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ItemKinds", testItemKindsQueryDeleteAll)
	t.Run("ItemOptions", testItemOptionsQueryDeleteAll)
	t.Run("ItemSides", testItemSidesQueryDeleteAll)
	t.Run("Items", testItemsQueryDeleteAll)
	t.Run("OptionKinds", testOptionKindsQueryDeleteAll)
	t.Run("Options", testOptionsQueryDeleteAll)
	t.Run("OrderItemOptions", testOrderItemOptionsQueryDeleteAll)
	t.Run("OrderItems", testOrderItemsQueryDeleteAll)
	t.Run("Orders", testOrdersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ItemKinds", testItemKindsSliceDeleteAll)
	t.Run("ItemOptions", testItemOptionsSliceDeleteAll)
	t.Run("ItemSides", testItemSidesSliceDeleteAll)
	t.Run("Items", testItemsSliceDeleteAll)
	t.Run("OptionKinds", testOptionKindsSliceDeleteAll)
	t.Run("Options", testOptionsSliceDeleteAll)
	t.Run("OrderItemOptions", testOrderItemOptionsSliceDeleteAll)
	t.Run("OrderItems", testOrderItemsSliceDeleteAll)
	t.Run("Orders", testOrdersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ItemKinds", testItemKindsExists)
	t.Run("ItemOptions", testItemOptionsExists)
	t.Run("ItemSides", testItemSidesExists)
	t.Run("Items", testItemsExists)
	t.Run("OptionKinds", testOptionKindsExists)
	t.Run("Options", testOptionsExists)
	t.Run("OrderItemOptions", testOrderItemOptionsExists)
	t.Run("OrderItems", testOrderItemsExists)
	t.Run("Orders", testOrdersExists)
}

func TestFind(t *testing.T) {
	t.Run("ItemKinds", testItemKindsFind)
	t.Run("ItemOptions", testItemOptionsFind)
	t.Run("ItemSides", testItemSidesFind)
	t.Run("Items", testItemsFind)
	t.Run("OptionKinds", testOptionKindsFind)
	t.Run("Options", testOptionsFind)
	t.Run("OrderItemOptions", testOrderItemOptionsFind)
	t.Run("OrderItems", testOrderItemsFind)
	t.Run("Orders", testOrdersFind)
}

func TestBind(t *testing.T) {
	t.Run("ItemKinds", testItemKindsBind)
	t.Run("ItemOptions", testItemOptionsBind)
	t.Run("ItemSides", testItemSidesBind)
	t.Run("Items", testItemsBind)
	t.Run("OptionKinds", testOptionKindsBind)
	t.Run("Options", testOptionsBind)
	t.Run("OrderItemOptions", testOrderItemOptionsBind)
	t.Run("OrderItems", testOrderItemsBind)
	t.Run("Orders", testOrdersBind)
}

func TestOne(t *testing.T) {
	t.Run("ItemKinds", testItemKindsOne)
	t.Run("ItemOptions", testItemOptionsOne)
	t.Run("ItemSides", testItemSidesOne)
	t.Run("Items", testItemsOne)
	t.Run("OptionKinds", testOptionKindsOne)
	t.Run("Options", testOptionsOne)
	t.Run("OrderItemOptions", testOrderItemOptionsOne)
	t.Run("OrderItems", testOrderItemsOne)
	t.Run("Orders", testOrdersOne)
}

func TestAll(t *testing.T) {
	t.Run("ItemKinds", testItemKindsAll)
	t.Run("ItemOptions", testItemOptionsAll)
	t.Run("ItemSides", testItemSidesAll)
	t.Run("Items", testItemsAll)
	t.Run("OptionKinds", testOptionKindsAll)
	t.Run("Options", testOptionsAll)
	t.Run("OrderItemOptions", testOrderItemOptionsAll)
	t.Run("OrderItems", testOrderItemsAll)
	t.Run("Orders", testOrdersAll)
}

func TestCount(t *testing.T) {
	t.Run("ItemKinds", testItemKindsCount)
	t.Run("ItemOptions", testItemOptionsCount)
	t.Run("ItemSides", testItemSidesCount)
	t.Run("Items", testItemsCount)
	t.Run("OptionKinds", testOptionKindsCount)
	t.Run("Options", testOptionsCount)
	t.Run("OrderItemOptions", testOrderItemOptionsCount)
	t.Run("OrderItems", testOrderItemsCount)
	t.Run("Orders", testOrdersCount)
}

func TestInsert(t *testing.T) {
	t.Run("ItemKinds", testItemKindsInsert)
	t.Run("ItemKinds", testItemKindsInsertWhitelist)
	t.Run("ItemOptions", testItemOptionsInsert)
	t.Run("ItemOptions", testItemOptionsInsertWhitelist)
	t.Run("ItemSides", testItemSidesInsert)
	t.Run("ItemSides", testItemSidesInsertWhitelist)
	t.Run("Items", testItemsInsert)
	t.Run("Items", testItemsInsertWhitelist)
	t.Run("OptionKinds", testOptionKindsInsert)
	t.Run("OptionKinds", testOptionKindsInsertWhitelist)
	t.Run("Options", testOptionsInsert)
	t.Run("Options", testOptionsInsertWhitelist)
	t.Run("OrderItemOptions", testOrderItemOptionsInsert)
	t.Run("OrderItemOptions", testOrderItemOptionsInsertWhitelist)
	t.Run("OrderItems", testOrderItemsInsert)
	t.Run("OrderItems", testOrderItemsInsertWhitelist)
	t.Run("Orders", testOrdersInsert)
	t.Run("Orders", testOrdersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ItemOptionToOptionUsingOption", testItemOptionToOneOptionUsingOption)
	t.Run("ItemOptionToItemUsingItem", testItemOptionToOneItemUsingItem)
	t.Run("ItemSideToItemUsingSideItem", testItemSideToOneItemUsingSideItem)
	t.Run("ItemSideToItemUsingItem", testItemSideToOneItemUsingItem)
	t.Run("ItemToItemKindUsingKind", testItemToOneItemKindUsingKind)
	t.Run("OptionToOptionKindUsingKind", testOptionToOneOptionKindUsingKind)
	t.Run("OrderItemOptionToOrderItemUsingOrderItem", testOrderItemOptionToOneOrderItemUsingOrderItem)
	t.Run("OrderItemOptionToOptionUsingOption", testOrderItemOptionToOneOptionUsingOption)
	t.Run("OrderItemOptionToOrderUsingOrder", testOrderItemOptionToOneOrderUsingOrder)
	t.Run("OrderItemToOrderItemUsingMainOrderItem", testOrderItemToOneOrderItemUsingMainOrderItem)
	t.Run("OrderItemToItemUsingItem", testOrderItemToOneItemUsingItem)
	t.Run("OrderItemToOrderUsingOrder", testOrderItemToOneOrderUsingOrder)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("ItemToItemOptionUsingItemOption", testItemOneToOneItemOptionUsingItemOption)
	t.Run("ItemToItemSideUsingSideItemItemSide", testItemOneToOneItemSideUsingSideItemItemSide)
	t.Run("ItemToItemSideUsingItemSide", testItemOneToOneItemSideUsingItemSide)
	t.Run("OptionToItemOptionUsingItemOption", testOptionOneToOneItemOptionUsingItemOption)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ItemKindToKindItems", testItemKindToManyKindItems)
	t.Run("ItemToOrderItems", testItemToManyOrderItems)
	t.Run("OptionKindToKindOptions", testOptionKindToManyKindOptions)
	t.Run("OptionToOrderItemOptions", testOptionToManyOrderItemOptions)
	t.Run("OrderItemToOrderItemOptions", testOrderItemToManyOrderItemOptions)
	t.Run("OrderItemToMainOrderItemOrderItems", testOrderItemToManyMainOrderItemOrderItems)
	t.Run("OrderToOrderItemOptions", testOrderToManyOrderItemOptions)
	t.Run("OrderToOrderItems", testOrderToManyOrderItems)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ItemOptionToOptionUsingItemOption", testItemOptionToOneSetOpOptionUsingOption)
	t.Run("ItemOptionToItemUsingItemOption", testItemOptionToOneSetOpItemUsingItem)
	t.Run("ItemSideToItemUsingSideItemItemSide", testItemSideToOneSetOpItemUsingSideItem)
	t.Run("ItemSideToItemUsingItemSide", testItemSideToOneSetOpItemUsingItem)
	t.Run("ItemToItemKindUsingKindItems", testItemToOneSetOpItemKindUsingKind)
	t.Run("OptionToOptionKindUsingKindOptions", testOptionToOneSetOpOptionKindUsingKind)
	t.Run("OrderItemOptionToOrderItemUsingOrderItemOptions", testOrderItemOptionToOneSetOpOrderItemUsingOrderItem)
	t.Run("OrderItemOptionToOptionUsingOrderItemOptions", testOrderItemOptionToOneSetOpOptionUsingOption)
	t.Run("OrderItemOptionToOrderUsingOrderItemOptions", testOrderItemOptionToOneSetOpOrderUsingOrder)
	t.Run("OrderItemToOrderItemUsingMainOrderItemOrderItems", testOrderItemToOneSetOpOrderItemUsingMainOrderItem)
	t.Run("OrderItemToItemUsingOrderItems", testOrderItemToOneSetOpItemUsingItem)
	t.Run("OrderItemToOrderUsingOrderItems", testOrderItemToOneSetOpOrderUsingOrder)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("OrderItemToOrderItemUsingMainOrderItemOrderItems", testOrderItemToOneRemoveOpOrderItemUsingMainOrderItem)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("ItemToItemOptionUsingItemOption", testItemOneToOneSetOpItemOptionUsingItemOption)
	t.Run("ItemToItemSideUsingSideItemItemSide", testItemOneToOneSetOpItemSideUsingSideItemItemSide)
	t.Run("ItemToItemSideUsingItemSide", testItemOneToOneSetOpItemSideUsingItemSide)
	t.Run("OptionToItemOptionUsingItemOption", testOptionOneToOneSetOpItemOptionUsingItemOption)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ItemKindToKindItems", testItemKindToManyAddOpKindItems)
	t.Run("ItemToOrderItems", testItemToManyAddOpOrderItems)
	t.Run("OptionKindToKindOptions", testOptionKindToManyAddOpKindOptions)
	t.Run("OptionToOrderItemOptions", testOptionToManyAddOpOrderItemOptions)
	t.Run("OrderItemToOrderItemOptions", testOrderItemToManyAddOpOrderItemOptions)
	t.Run("OrderItemToMainOrderItemOrderItems", testOrderItemToManyAddOpMainOrderItemOrderItems)
	t.Run("OrderToOrderItemOptions", testOrderToManyAddOpOrderItemOptions)
	t.Run("OrderToOrderItems", testOrderToManyAddOpOrderItems)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("OrderItemToMainOrderItemOrderItems", testOrderItemToManySetOpMainOrderItemOrderItems)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("OrderItemToMainOrderItemOrderItems", testOrderItemToManyRemoveOpMainOrderItemOrderItems)
}

func TestReload(t *testing.T) {
	t.Run("ItemKinds", testItemKindsReload)
	t.Run("ItemOptions", testItemOptionsReload)
	t.Run("ItemSides", testItemSidesReload)
	t.Run("Items", testItemsReload)
	t.Run("OptionKinds", testOptionKindsReload)
	t.Run("Options", testOptionsReload)
	t.Run("OrderItemOptions", testOrderItemOptionsReload)
	t.Run("OrderItems", testOrderItemsReload)
	t.Run("Orders", testOrdersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ItemKinds", testItemKindsReloadAll)
	t.Run("ItemOptions", testItemOptionsReloadAll)
	t.Run("ItemSides", testItemSidesReloadAll)
	t.Run("Items", testItemsReloadAll)
	t.Run("OptionKinds", testOptionKindsReloadAll)
	t.Run("Options", testOptionsReloadAll)
	t.Run("OrderItemOptions", testOrderItemOptionsReloadAll)
	t.Run("OrderItems", testOrderItemsReloadAll)
	t.Run("Orders", testOrdersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ItemKinds", testItemKindsSelect)
	t.Run("ItemOptions", testItemOptionsSelect)
	t.Run("ItemSides", testItemSidesSelect)
	t.Run("Items", testItemsSelect)
	t.Run("OptionKinds", testOptionKindsSelect)
	t.Run("Options", testOptionsSelect)
	t.Run("OrderItemOptions", testOrderItemOptionsSelect)
	t.Run("OrderItems", testOrderItemsSelect)
	t.Run("Orders", testOrdersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ItemKinds", testItemKindsUpdate)
	t.Run("ItemOptions", testItemOptionsUpdate)
	t.Run("ItemSides", testItemSidesUpdate)
	t.Run("Items", testItemsUpdate)
	t.Run("OptionKinds", testOptionKindsUpdate)
	t.Run("Options", testOptionsUpdate)
	t.Run("OrderItemOptions", testOrderItemOptionsUpdate)
	t.Run("OrderItems", testOrderItemsUpdate)
	t.Run("Orders", testOrdersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ItemKinds", testItemKindsSliceUpdateAll)
	t.Run("ItemOptions", testItemOptionsSliceUpdateAll)
	t.Run("ItemSides", testItemSidesSliceUpdateAll)
	t.Run("Items", testItemsSliceUpdateAll)
	t.Run("OptionKinds", testOptionKindsSliceUpdateAll)
	t.Run("Options", testOptionsSliceUpdateAll)
	t.Run("OrderItemOptions", testOrderItemOptionsSliceUpdateAll)
	t.Run("OrderItems", testOrderItemsSliceUpdateAll)
	t.Run("Orders", testOrdersSliceUpdateAll)
}
