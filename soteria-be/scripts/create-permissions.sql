INSERT INTO Permission
VALUES
  /* Facility Service Permissions */
  (gen_random_uuid(), 'facility', 'list'),
  (gen_random_uuid(), 'facility', 'get'),
  (gen_random_uuid(), 'facility', 'create'),
  (gen_random_uuid(), 'facility', 'update'),
  (gen_random_uuid(), 'facility', 'remove'),
  (gen_random_uuid(), 'location', 'list'),
  (gen_random_uuid(), 'location', 'get'),
  (gen_random_uuid(), 'location', 'create'),
  (gen_random_uuid(), 'location', 'update'),
  (gen_random_uuid(), 'location', 'remove'),
  (gen_random_uuid(), 'locationType', 'list'),
  (gen_random_uuid(), 'locationType', 'get'),
  (gen_random_uuid(), 'locationType', 'create'),
  (gen_random_uuid(), 'locationType', 'update'),
  (gen_random_uuid(), 'locationType', 'remove'),
  /* IAM Service Permissions */
  (gen_random_uuid(), 'user', 'list'),
  (gen_random_uuid(), 'user', 'get'),
  (gen_random_uuid(), 'user', 'create'),
  (gen_random_uuid(), 'user', 'update'),
  (gen_random_uuid(), 'user', 'remove'),
  (gen_random_uuid(), 'permission', 'list'),
  (gen_random_uuid(), 'permission', 'get'),
  (gen_random_uuid(), 'permission', 'create'),
  (gen_random_uuid(), 'permission', 'update'),
  (gen_random_uuid(), 'permission', 'remove'),
  (gen_random_uuid(), 'api_key', 'create'),
  (gen_random_uuid(), 'api_key', 'remove');

SELECT * FROM Permission;