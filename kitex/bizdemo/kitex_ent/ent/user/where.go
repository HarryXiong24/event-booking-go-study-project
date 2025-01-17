// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGender, v))
}

// Introduce applies equality check predicate on the "introduce" field. It's identical to IntroduceEQ.
func Introduce(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIntroduce, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAge, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGender, vs...))
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldGender, v))
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGender, v))
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldGender, v))
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGender, v))
}

// IntroduceEQ applies the EQ predicate on the "introduce" field.
func IntroduceEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIntroduce, v))
}

// IntroduceNEQ applies the NEQ predicate on the "introduce" field.
func IntroduceNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIntroduce, v))
}

// IntroduceIn applies the In predicate on the "introduce" field.
func IntroduceIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldIntroduce, vs...))
}

// IntroduceNotIn applies the NotIn predicate on the "introduce" field.
func IntroduceNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldIntroduce, vs...))
}

// IntroduceGT applies the GT predicate on the "introduce" field.
func IntroduceGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldIntroduce, v))
}

// IntroduceGTE applies the GTE predicate on the "introduce" field.
func IntroduceGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldIntroduce, v))
}

// IntroduceLT applies the LT predicate on the "introduce" field.
func IntroduceLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldIntroduce, v))
}

// IntroduceLTE applies the LTE predicate on the "introduce" field.
func IntroduceLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldIntroduce, v))
}

// IntroduceContains applies the Contains predicate on the "introduce" field.
func IntroduceContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldIntroduce, v))
}

// IntroduceHasPrefix applies the HasPrefix predicate on the "introduce" field.
func IntroduceHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldIntroduce, v))
}

// IntroduceHasSuffix applies the HasSuffix predicate on the "introduce" field.
func IntroduceHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldIntroduce, v))
}

// IntroduceEqualFold applies the EqualFold predicate on the "introduce" field.
func IntroduceEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldIntroduce, v))
}

// IntroduceContainsFold applies the ContainsFold predicate on the "introduce" field.
func IntroduceContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldIntroduce, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
