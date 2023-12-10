package meta

import (
	"encoding/json"
	"fmt"
)

// A meta set, represents an ensemble of associative relationships.
type MetaSet struct {
	Id            string                   // The id of the meta set.
	Tables        map[string]*Table        // The tables in the meta set.
	TableFields   map[string]*TableField   // The table fields in the meta set.
	Relationships map[string]*Relationship // The relationships in the meta set.
}

// Create a new meta set.
func NewMetaSet(id string, tables []*Table, fields []*TableField, relationships []*Relationship) (*MetaSet, error) {
	m := &MetaSet{}

	m.Id = id
	m.Tables = make(map[string]*Table, len(tables))
	m.TableFields = make(map[string]*TableField, len(fields))
	m.Relationships = make(map[string]*Relationship, len(relationships))

	for _, t := range tables {
		t.MetaId = m.Id
		m.Tables[t.Id] = t
	}
	for _, f := range fields {
		t, err := m.GetTable(f.TableId)
		if err != nil {
			return nil, err
		}

		f.Table = t
		f.MetaId = m.Id
		t.TableFields = append(t.TableFields, f)

		m.TableFields[f.Id] = f
	}
	for _, r := range relationships {
		leftField, err := m.GetTableField(r.ConditionLeft)
		if err != nil {
			return nil, err
		}

		leftTable := leftField.Table

		r.LeftTable = leftTable
		r.LeftField = leftField
		r.MetaId = m.Id
		leftTable.Relationships = append(leftTable.Relationships, r)
		leftTable.RightRelationships = append(leftTable.RightRelationships, r)

		rightField, err := m.GetTableField(r.ConditionRight)
		if err != nil {
			return nil, err
		}

		rightTable := rightField.Table

		r.RightTable = rightTable
		r.RightField = rightField
		r.MetaId = m.Id
		rightTable.Relationships = append(rightTable.Relationships, r)
		rightTable.LeftRelationships = append(rightTable.LeftRelationships, r)

		m.Relationships[r.Id] = r
	}

	return m, nil
}

func (m *MetaSet) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Id            string          `json:"id"`
		Tables        []*Table        `json:"tables"`
		TableFields   []*TableField   `json:"tableFields"`
		Relationships []*Relationship `json:"relationships"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	n, err := NewMetaSet(aux.Id, aux.Tables, aux.TableFields, aux.Relationships)
	if err != nil {
		return err
	}

	m.Id = n.Id
	m.Tables = n.Tables
	m.TableFields = n.TableFields
	m.Relationships = n.Relationships
	n = nil

	return nil
}

func (m *MetaSet) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// Get a table by id.
func (m *MetaSet) GetTable(id string) (*Table, error) {
	t, ok := m.Tables[id]
	if !ok { // not found
		return nil, fmt.Errorf("not found table: %s", id)
	}

	return t, nil
}

// Get a table field by id.
func (m *MetaSet) GetTableField(id string) (*TableField, error) {
	f, ok := m.TableFields[id]
	if !ok { // not found
		return nil, fmt.Errorf("not found table field: %s", id)
	}

	return f, nil
}

// Get a relationship by id.
func (m *MetaSet) GetRelationship(id string) (*Relationship, error) {
	r, ok := m.Relationships[id]
	if !ok { // not found
		return nil, fmt.Errorf("not found relationship: %s", id)
	}

	return r, nil
}
