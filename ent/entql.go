// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/TcMits/ent-clean-template/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateTime:  {Type: field.TypeTime, Column: user.FieldCreateTime},
			user.FieldUpdateTime:  {Type: field.TypeTime, Column: user.FieldUpdateTime},
			user.FieldJwtTokenKey: {Type: field.TypeString, Column: user.FieldJwtTokenKey},
			user.FieldPassword:    {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldUsername:    {Type: field.TypeString, Column: user.FieldUsername},
			user.FieldFirstName:   {Type: field.TypeString, Column: user.FieldFirstName},
			user.FieldLastName:    {Type: field.TypeString, Column: user.FieldLastName},
			user.FieldEmail:       {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldIsStaff:     {Type: field.TypeBool, Column: user.FieldIsStaff},
			user.FieldIsSuperuser: {Type: field.TypeBool, Column: user.FieldIsSuperuser},
			user.FieldIsActive:    {Type: field.TypeBool, Column: user.FieldIsActive},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *UserFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereJwtTokenKey applies the entql string predicate on the jwt_token_key field.
func (f *UserFilter) WhereJwtTokenKey(p entql.StringP) {
	f.Where(p.Field(user.FieldJwtTokenKey))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *UserFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(user.FieldUsername))
}

// WhereFirstName applies the entql string predicate on the first_name field.
func (f *UserFilter) WhereFirstName(p entql.StringP) {
	f.Where(p.Field(user.FieldFirstName))
}

// WhereLastName applies the entql string predicate on the last_name field.
func (f *UserFilter) WhereLastName(p entql.StringP) {
	f.Where(p.Field(user.FieldLastName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereIsStaff applies the entql bool predicate on the is_staff field.
func (f *UserFilter) WhereIsStaff(p entql.BoolP) {
	f.Where(p.Field(user.FieldIsStaff))
}

// WhereIsSuperuser applies the entql bool predicate on the is_superuser field.
func (f *UserFilter) WhereIsSuperuser(p entql.BoolP) {
	f.Where(p.Field(user.FieldIsSuperuser))
}

// WhereIsActive applies the entql bool predicate on the is_active field.
func (f *UserFilter) WhereIsActive(p entql.BoolP) {
	f.Where(p.Field(user.FieldIsActive))
}