// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gomicro/ent/accounts"
	"gomicro/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccounts = "Accounts"
)

// AccountsMutation represents an operation that mutates the Accounts nodes in the graph.
type AccountsMutation struct {
	config
	op            Op
	typ           string
	id            *int
	owner         *string
	balance       *int8
	addbalance    *int8
	currency      *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Accounts, error)
	predicates    []predicate.Accounts
}

var _ ent.Mutation = (*AccountsMutation)(nil)

// accountsOption allows management of the mutation configuration using functional options.
type accountsOption func(*AccountsMutation)

// newAccountsMutation creates new mutation for the Accounts entity.
func newAccountsMutation(c config, op Op, opts ...accountsOption) *AccountsMutation {
	m := &AccountsMutation{
		config:        c,
		op:            op,
		typ:           TypeAccounts,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccountsID sets the ID field of the mutation.
func withAccountsID(id int) accountsOption {
	return func(m *AccountsMutation) {
		var (
			err   error
			once  sync.Once
			value *Accounts
		)
		m.oldValue = func(ctx context.Context) (*Accounts, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Accounts.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccounts sets the old Accounts of the mutation.
func withAccounts(node *Accounts) accountsOption {
	return func(m *AccountsMutation) {
		m.oldValue = func(context.Context) (*Accounts, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccountsMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccountsMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccountsMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetOwner sets the "owner" field.
func (m *AccountsMutation) SetOwner(s string) {
	m.owner = &s
}

// Owner returns the value of the "owner" field in the mutation.
func (m *AccountsMutation) Owner() (r string, exists bool) {
	v := m.owner
	if v == nil {
		return
	}
	return *v, true
}

// OldOwner returns the old "owner" field's value of the Accounts entity.
// If the Accounts object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountsMutation) OldOwner(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOwner is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOwner requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwner: %w", err)
	}
	return oldValue.Owner, nil
}

// ResetOwner resets all changes to the "owner" field.
func (m *AccountsMutation) ResetOwner() {
	m.owner = nil
}

// SetBalance sets the "balance" field.
func (m *AccountsMutation) SetBalance(i int8) {
	m.balance = &i
	m.addbalance = nil
}

// Balance returns the value of the "balance" field in the mutation.
func (m *AccountsMutation) Balance() (r int8, exists bool) {
	v := m.balance
	if v == nil {
		return
	}
	return *v, true
}

// OldBalance returns the old "balance" field's value of the Accounts entity.
// If the Accounts object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountsMutation) OldBalance(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldBalance is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldBalance requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBalance: %w", err)
	}
	return oldValue.Balance, nil
}

// AddBalance adds i to the "balance" field.
func (m *AccountsMutation) AddBalance(i int8) {
	if m.addbalance != nil {
		*m.addbalance += i
	} else {
		m.addbalance = &i
	}
}

// AddedBalance returns the value that was added to the "balance" field in this mutation.
func (m *AccountsMutation) AddedBalance() (r int8, exists bool) {
	v := m.addbalance
	if v == nil {
		return
	}
	return *v, true
}

// ResetBalance resets all changes to the "balance" field.
func (m *AccountsMutation) ResetBalance() {
	m.balance = nil
	m.addbalance = nil
}

// SetCurrency sets the "currency" field.
func (m *AccountsMutation) SetCurrency(s string) {
	m.currency = &s
}

// Currency returns the value of the "currency" field in the mutation.
func (m *AccountsMutation) Currency() (r string, exists bool) {
	v := m.currency
	if v == nil {
		return
	}
	return *v, true
}

// OldCurrency returns the old "currency" field's value of the Accounts entity.
// If the Accounts object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountsMutation) OldCurrency(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCurrency is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCurrency requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCurrency: %w", err)
	}
	return oldValue.Currency, nil
}

// ResetCurrency resets all changes to the "currency" field.
func (m *AccountsMutation) ResetCurrency() {
	m.currency = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *AccountsMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AccountsMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Accounts entity.
// If the Accounts object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccountsMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AccountsMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the AccountsMutation builder.
func (m *AccountsMutation) Where(ps ...predicate.Accounts) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *AccountsMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Accounts).
func (m *AccountsMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccountsMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.owner != nil {
		fields = append(fields, accounts.FieldOwner)
	}
	if m.balance != nil {
		fields = append(fields, accounts.FieldBalance)
	}
	if m.currency != nil {
		fields = append(fields, accounts.FieldCurrency)
	}
	if m.created_at != nil {
		fields = append(fields, accounts.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccountsMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case accounts.FieldOwner:
		return m.Owner()
	case accounts.FieldBalance:
		return m.Balance()
	case accounts.FieldCurrency:
		return m.Currency()
	case accounts.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccountsMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case accounts.FieldOwner:
		return m.OldOwner(ctx)
	case accounts.FieldBalance:
		return m.OldBalance(ctx)
	case accounts.FieldCurrency:
		return m.OldCurrency(ctx)
	case accounts.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Accounts field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountsMutation) SetField(name string, value ent.Value) error {
	switch name {
	case accounts.FieldOwner:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwner(v)
		return nil
	case accounts.FieldBalance:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBalance(v)
		return nil
	case accounts.FieldCurrency:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCurrency(v)
		return nil
	case accounts.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Accounts field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccountsMutation) AddedFields() []string {
	var fields []string
	if m.addbalance != nil {
		fields = append(fields, accounts.FieldBalance)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccountsMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case accounts.FieldBalance:
		return m.AddedBalance()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccountsMutation) AddField(name string, value ent.Value) error {
	switch name {
	case accounts.FieldBalance:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddBalance(v)
		return nil
	}
	return fmt.Errorf("unknown Accounts numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccountsMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccountsMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccountsMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Accounts nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccountsMutation) ResetField(name string) error {
	switch name {
	case accounts.FieldOwner:
		m.ResetOwner()
		return nil
	case accounts.FieldBalance:
		m.ResetBalance()
		return nil
	case accounts.FieldCurrency:
		m.ResetCurrency()
		return nil
	case accounts.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Accounts field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccountsMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccountsMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccountsMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccountsMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccountsMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccountsMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccountsMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Accounts unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccountsMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Accounts edge %s", name)
}
