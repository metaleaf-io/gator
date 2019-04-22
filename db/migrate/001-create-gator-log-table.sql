CREATE TABLE gator_log (
  time    TIMESTAMP,
  source  VARCHAR(255),
  level   VARCHAR(8),
  fields  JSON,
  message TEXT
);

CREATE PROCEDURE gator_save(BIGINT, VARCHAR(255), VARCHAR(8), TEXT, TEXT)
  LANGUAGE SQL
  AS $$
  INSERT INTO gator_log (time, source, level, fields, message) VALUES (to_timestamp($1::float/1000000000), $2, $3, $4::json, $5);
  $$;
