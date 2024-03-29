// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ent-demo/ent/car"
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// RegisteredAt holds the value of the "registered_at" field.
	RegisteredAt time.Time `json:"registered_at,omitempty"`
	user_cars    *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // model
		&sql.NullTime{},   // registered_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Car) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_cars
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(values ...interface{}) error {
	if m, n := len(values), len(car.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field model", values[0])
	} else if value.Valid {
		c.Model = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field registered_at", values[1])
	} else if value.Valid {
		c.RegisteredAt = value.Time
	}
	values = values[2:]
	if len(values) == len(car.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_cars", value)
		} else if value.Valid {
			c.user_cars = new(int)
			*c.user_cars = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Car.
// Note that, you need to call Car.Unwrap() before calling this method, if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", model=")
	builder.WriteString(c.Model)
	builder.WriteString(", registered_at=")
	builder.WriteString(c.RegisteredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
