-- +migrate Up
SET client_min_messages = 'warning';

CREATE TABLE IF NOT EXISTS public.plan_years (
  id                            SERIAL      PRIMARY KEY,
  created_at                    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at                    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  fiscal_year_id                integer     NOT NULL,
	average_days_of_illness       numeric,
	average_days_of_children_care numeric,
	default_vacation_interest     numeric
);

-- +migrate Down
DROP TABLE public.plan_years;
