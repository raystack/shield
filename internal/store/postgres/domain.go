package postgres

import (
	"database/sql"
	"time"

	"github.com/raystack/frontier/core/domain"
)

type Domain struct {
	ID        string       `db:"id"`
	OrgID     string       `db:"org_id"`
	Name      string       `db:"name"`
	Token     string       `db:"token"`
	State     string       `db:"state"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	CreatedAt time.Time    `db:"created_at"`
}

func (d Domain) transform() domain.Domain {
	return domain.Domain{
		ID:        d.ID,
		OrgID:     d.OrgID,
		Name:      d.Name,
		Token:     d.Token,
		State:     d.State,
		UpdatedAt: d.UpdatedAt.Time,
		CreatedAt: d.CreatedAt,
	}
}
