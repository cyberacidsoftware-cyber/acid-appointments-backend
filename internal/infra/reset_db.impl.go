package infra

import (
	"context"
)

func (q *Queries) ResetDb(ctx context.Context) error {
	var err error

	_, err = q.db.Exec(ctx, "DELETE FROM admins")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM appointments")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM customers")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM secretaries")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM services")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM service_names")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM specialists")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM specialist_hours")
	if err != nil {
		return err
	}

	_, err = q.db.Exec(ctx, "DELETE FROM specializations")
	if err != nil {
		return err
	}

	return nil
}
