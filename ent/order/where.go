// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mcbebu/ALTER-TABLE/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldName, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDescription, v))
}

// MobileNumber applies equality check predicate on the "mobileNumber" field. It's identical to MobileNumberEQ.
func MobileNumber(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldMobileNumber, v))
}

// AltMobileNumber applies equality check predicate on the "altMobileNumber" field. It's identical to AltMobileNumberEQ.
func AltMobileNumber(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldAltMobileNumber, v))
}

// LeaveParcel applies equality check predicate on the "leaveParcel" field. It's identical to LeaveParcelEQ.
func LeaveParcel(v bool) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldLeaveParcel, v))
}

// StopsUntilDelivery applies equality check predicate on the "stopsUntilDelivery" field. It's identical to StopsUntilDeliveryEQ.
func StopsUntilDelivery(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStopsUntilDelivery, v))
}

// EstimatedArrivalTime applies equality check predicate on the "estimatedArrivalTime" field. It's identical to EstimatedArrivalTimeEQ.
func EstimatedArrivalTime(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldEstimatedArrivalTime, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldName, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldDescription, v))
}

// MobileNumberEQ applies the EQ predicate on the "mobileNumber" field.
func MobileNumberEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldMobileNumber, v))
}

// MobileNumberNEQ applies the NEQ predicate on the "mobileNumber" field.
func MobileNumberNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldMobileNumber, v))
}

// MobileNumberIn applies the In predicate on the "mobileNumber" field.
func MobileNumberIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldMobileNumber, vs...))
}

// MobileNumberNotIn applies the NotIn predicate on the "mobileNumber" field.
func MobileNumberNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldMobileNumber, vs...))
}

// MobileNumberGT applies the GT predicate on the "mobileNumber" field.
func MobileNumberGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldMobileNumber, v))
}

// MobileNumberGTE applies the GTE predicate on the "mobileNumber" field.
func MobileNumberGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldMobileNumber, v))
}

// MobileNumberLT applies the LT predicate on the "mobileNumber" field.
func MobileNumberLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldMobileNumber, v))
}

// MobileNumberLTE applies the LTE predicate on the "mobileNumber" field.
func MobileNumberLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldMobileNumber, v))
}

// MobileNumberContains applies the Contains predicate on the "mobileNumber" field.
func MobileNumberContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldMobileNumber, v))
}

// MobileNumberHasPrefix applies the HasPrefix predicate on the "mobileNumber" field.
func MobileNumberHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldMobileNumber, v))
}

// MobileNumberHasSuffix applies the HasSuffix predicate on the "mobileNumber" field.
func MobileNumberHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldMobileNumber, v))
}

// MobileNumberEqualFold applies the EqualFold predicate on the "mobileNumber" field.
func MobileNumberEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldMobileNumber, v))
}

// MobileNumberContainsFold applies the ContainsFold predicate on the "mobileNumber" field.
func MobileNumberContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldMobileNumber, v))
}

// AltMobileNumberEQ applies the EQ predicate on the "altMobileNumber" field.
func AltMobileNumberEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldAltMobileNumber, v))
}

// AltMobileNumberNEQ applies the NEQ predicate on the "altMobileNumber" field.
func AltMobileNumberNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldAltMobileNumber, v))
}

// AltMobileNumberIn applies the In predicate on the "altMobileNumber" field.
func AltMobileNumberIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldAltMobileNumber, vs...))
}

// AltMobileNumberNotIn applies the NotIn predicate on the "altMobileNumber" field.
func AltMobileNumberNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldAltMobileNumber, vs...))
}

// AltMobileNumberGT applies the GT predicate on the "altMobileNumber" field.
func AltMobileNumberGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldAltMobileNumber, v))
}

// AltMobileNumberGTE applies the GTE predicate on the "altMobileNumber" field.
func AltMobileNumberGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldAltMobileNumber, v))
}

// AltMobileNumberLT applies the LT predicate on the "altMobileNumber" field.
func AltMobileNumberLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldAltMobileNumber, v))
}

// AltMobileNumberLTE applies the LTE predicate on the "altMobileNumber" field.
func AltMobileNumberLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldAltMobileNumber, v))
}

// AltMobileNumberContains applies the Contains predicate on the "altMobileNumber" field.
func AltMobileNumberContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldAltMobileNumber, v))
}

// AltMobileNumberHasPrefix applies the HasPrefix predicate on the "altMobileNumber" field.
func AltMobileNumberHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldAltMobileNumber, v))
}

// AltMobileNumberHasSuffix applies the HasSuffix predicate on the "altMobileNumber" field.
func AltMobileNumberHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldAltMobileNumber, v))
}

// AltMobileNumberIsNil applies the IsNil predicate on the "altMobileNumber" field.
func AltMobileNumberIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldAltMobileNumber))
}

// AltMobileNumberNotNil applies the NotNil predicate on the "altMobileNumber" field.
func AltMobileNumberNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldAltMobileNumber))
}

// AltMobileNumberEqualFold applies the EqualFold predicate on the "altMobileNumber" field.
func AltMobileNumberEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldAltMobileNumber, v))
}

// AltMobileNumberContainsFold applies the ContainsFold predicate on the "altMobileNumber" field.
func AltMobileNumberContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldAltMobileNumber, v))
}

// LeaveParcelEQ applies the EQ predicate on the "leaveParcel" field.
func LeaveParcelEQ(v bool) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldLeaveParcel, v))
}

// LeaveParcelNEQ applies the NEQ predicate on the "leaveParcel" field.
func LeaveParcelNEQ(v bool) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldLeaveParcel, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// StopsUntilDeliveryEQ applies the EQ predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryNEQ applies the NEQ predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryNEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryIn applies the In predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStopsUntilDelivery, vs...))
}

// StopsUntilDeliveryNotIn applies the NotIn predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryNotIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStopsUntilDelivery, vs...))
}

// StopsUntilDeliveryGT applies the GT predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryGT(v int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryGTE applies the GTE predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryGTE(v int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryLT applies the LT predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryLT(v int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryLTE applies the LTE predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryLTE(v int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldStopsUntilDelivery, v))
}

// StopsUntilDeliveryIsNil applies the IsNil predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldStopsUntilDelivery))
}

// StopsUntilDeliveryNotNil applies the NotNil predicate on the "stopsUntilDelivery" field.
func StopsUntilDeliveryNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldStopsUntilDelivery))
}

// EstimatedArrivalTimeEQ applies the EQ predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeNEQ applies the NEQ predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeNEQ(v int) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeIn applies the In predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldEstimatedArrivalTime, vs...))
}

// EstimatedArrivalTimeNotIn applies the NotIn predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeNotIn(vs ...int) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldEstimatedArrivalTime, vs...))
}

// EstimatedArrivalTimeGT applies the GT predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeGT(v int) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeGTE applies the GTE predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeGTE(v int) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeLT applies the LT predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeLT(v int) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeLTE applies the LTE predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeLTE(v int) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldEstimatedArrivalTime, v))
}

// EstimatedArrivalTimeIsNil applies the IsNil predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldEstimatedArrivalTime))
}

// EstimatedArrivalTimeNotNil applies the NotNil predicate on the "estimatedArrivalTime" field.
func EstimatedArrivalTimeNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldEstimatedArrivalTime))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreatedAt, v))
}

// HasShippers applies the HasEdge predicate on the "shippers" edge.
func HasShippers() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShippersTable, ShippersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShippersWith applies the HasEdge predicate on the "shippers" edge with a given conditions (other predicates).
func HasShippersWith(preds ...predicate.Shipper) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShippersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShippersTable, ShippersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		p(s.Not())
	})
}
