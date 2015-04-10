-- +migrate Up
SET client_min_messages = 'warning';

CREATE TABLE IF NOT EXISTS public.fiscal_periods (
  id            SERIAL      PRIMARY KEY,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  business_days integer     NOT NULL,
  starts_at     date        NOT NULL,
  ends_at       date        NOT NULL
);

-- +migrate Down
DROP TABLE public.fiscal_periods;

