// Code generated by ent, DO NOT EDIT.

package contract

import (
	"golang-boilerplate/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// VendorID applies equality check predicate on the "vendor_id" field. It's identical to VendorIDEQ.
func VendorID(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVendorID), v))
	})
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// TotalAmount applies equality check predicate on the "total_amount" field. It's identical to TotalAmountEQ.
func TotalAmount(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalAmount), v))
	})
}

// RemainingAmount applies equality check predicate on the "remaining_amount" field. It's identical to RemainingAmountEQ.
func RemainingAmount(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemainingAmount), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// VendorIDEQ applies the EQ predicate on the "vendor_id" field.
func VendorIDEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVendorID), v))
	})
}

// VendorIDNEQ applies the NEQ predicate on the "vendor_id" field.
func VendorIDNEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVendorID), v))
	})
}

// VendorIDIn applies the In predicate on the "vendor_id" field.
func VendorIDIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldVendorID), v...))
	})
}

// VendorIDNotIn applies the NotIn predicate on the "vendor_id" field.
func VendorIDNotIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldVendorID), v...))
	})
}

// VendorIDGT applies the GT predicate on the "vendor_id" field.
func VendorIDGT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVendorID), v))
	})
}

// VendorIDGTE applies the GTE predicate on the "vendor_id" field.
func VendorIDGTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVendorID), v))
	})
}

// VendorIDLT applies the LT predicate on the "vendor_id" field.
func VendorIDLT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVendorID), v))
	})
}

// VendorIDLTE applies the LTE predicate on the "vendor_id" field.
func VendorIDLTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVendorID), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// StartDateIsNil applies the IsNil predicate on the "start_date" field.
func StartDateIsNil() predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartDate)))
	})
}

// StartDateNotNil applies the NotNil predicate on the "start_date" field.
func StartDateNotNil() predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartDate)))
	})
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndDate), v))
	})
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndDate), v...))
	})
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndDate), v...))
	})
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndDate), v))
	})
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndDate), v))
	})
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndDate), v))
	})
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndDate), v))
	})
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndDate)))
	})
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndDate)))
	})
}

// TotalAmountEQ applies the EQ predicate on the "total_amount" field.
func TotalAmountEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountNEQ applies the NEQ predicate on the "total_amount" field.
func TotalAmountNEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountIn applies the In predicate on the "total_amount" field.
func TotalAmountIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTotalAmount), v...))
	})
}

// TotalAmountNotIn applies the NotIn predicate on the "total_amount" field.
func TotalAmountNotIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTotalAmount), v...))
	})
}

// TotalAmountGT applies the GT predicate on the "total_amount" field.
func TotalAmountGT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountGTE applies the GTE predicate on the "total_amount" field.
func TotalAmountGTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountLT applies the LT predicate on the "total_amount" field.
func TotalAmountLT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTotalAmount), v))
	})
}

// TotalAmountLTE applies the LTE predicate on the "total_amount" field.
func TotalAmountLTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTotalAmount), v))
	})
}

// RemainingAmountEQ applies the EQ predicate on the "remaining_amount" field.
func RemainingAmountEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemainingAmount), v))
	})
}

// RemainingAmountNEQ applies the NEQ predicate on the "remaining_amount" field.
func RemainingAmountNEQ(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemainingAmount), v))
	})
}

// RemainingAmountIn applies the In predicate on the "remaining_amount" field.
func RemainingAmountIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemainingAmount), v...))
	})
}

// RemainingAmountNotIn applies the NotIn predicate on the "remaining_amount" field.
func RemainingAmountNotIn(vs ...int) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemainingAmount), v...))
	})
}

// RemainingAmountGT applies the GT predicate on the "remaining_amount" field.
func RemainingAmountGT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemainingAmount), v))
	})
}

// RemainingAmountGTE applies the GTE predicate on the "remaining_amount" field.
func RemainingAmountGTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemainingAmount), v))
	})
}

// RemainingAmountLT applies the LT predicate on the "remaining_amount" field.
func RemainingAmountLT(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemainingAmount), v))
	})
}

// RemainingAmountLTE applies the LTE predicate on the "remaining_amount" field.
func RemainingAmountLTE(v int) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemainingAmount), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Contract {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contract) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contract) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
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
func Not(p predicate.Contract) predicate.Contract {
	return predicate.Contract(func(s *sql.Selector) {
		p(s.Not())
	})
}
