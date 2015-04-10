-- +migrate Up
SET client_min_messages = 'warning';

CREATE TABLE IF NOT EXISTS public.fiscal_years (
  id                          SERIAL      PRIMARY KEY,
  created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  year                        integer     NOT NULL,
  business_days               integer     NOT NULL,
  business_days_first_quarter integer     NOT NULL,
  calendar_weeks              integer     NOT NULL
);

-- +migrate Down
DROP TABLE public.fiscal_years;
