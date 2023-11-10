-- INSERT INTO groups (name, address, email, role, created_at, updated_at)
-- VALUES
--     ('techvify', 'planner address', 'planner@planner.com1', 'planner', NOW(), NOW()),
--     ('fsoft', 'vendor address', 'vendor@vendor.com1', 'vendor', NOW(), NOW()),
--     ('vng', 'contractor address', 'contractor@contractor.com1', 'contractor', NOW(), NOW());
--
-- INSERT INTO users (first_name, last_name, email, group_id, created_at, updated_at)
-- VALUES
--     ('---> planner1', 'planner1 -->', 'planner1@planner1.com', ?, 1, NOW(), NOW()),
--     ('---> planner2', 'planner2 -->', 'planner2@planner2.com', ?, 1, NOW(), NOW()),
--     ('---> planner3', 'planner3 -->', 'planner3@planner3.com', ?, 1, NOW(), NOW()),
--     ('---> vendor1', 'vendor1 -->', 'vendor1@vendor1.com', ?, 2, NOW(), NOW()),
--     ('---> vendor2', 'vendor2 -->', 'vendor2@vendor2.com', ?, 2, NOW(), NOW()),
--     ('---> vendor3', 'vendor3 -->', 'vendor3@vendor3.com', ?, 2, NOW(), NOW()),
--     ('---> contractor1', 'contractor1 -->', 'contractor1@contractor1.com', ?, 3, NOW(), NOW()),
--     ('---> contractor2', 'contractor2 -->', 'contractor2@contractor1.com', ?, 3, NOW(), NOW()),
--     ('---> contractor3', 'contractor3 -->', 'contractor3@contractor1.com', ?, 3, NOW(), NOW());

-- Create a temporary table to store the IDs of the inserted groups
CREATE TEMP TABLE tmp_group_ids (id INT, idx INT);

-- Insert groups and return their IDs
WITH ins AS (
INSERT INTO groups (name, address, email, role, created_at, updated_at)
VALUES
    ('techvify', 'planner address', 'planner@planner.com1', 'planner', NOW(), NOW()),
    ('fsoft', 'vendor address', 'vendor@vendor.com1', 'vendor', NOW(), NOW()),
    ('vng', 'contractor address', 'contractor@contractor.com1', 'contractor', NOW(), NOW())
    RETURNING id
    )
INSERT INTO tmp_group_ids
SELECT id, ROW_NUMBER() OVER() AS idx FROM ins;

-- Insert users linked to their groups using the temporary table
INSERT INTO users (first_name, last_name, email, group_users, created_at, updated_at)
SELECT
    CASE
        WHEN idx = 1 THEN '---> planner' || usr.row_num
        WHEN idx = 2 THEN '---> vendor' || usr.row_num
        WHEN idx = 3 THEN '---> contractor' || usr.row_num
        END,
    CASE
        WHEN idx = 1 THEN 'planner' || usr.row_num || ' -->'
        WHEN idx = 2 THEN 'vendor' || usr.row_num || ' -->'
        WHEN idx = 3 THEN 'contractor' || usr.row_num || ' -->'
        END,
    CASE
        WHEN idx = 1 THEN 'planner' || usr.row_num || '@planner.com'
        WHEN idx = 2 THEN 'vendor' || usr.row_num || '@vendor.com'
        WHEN idx = 3 THEN 'contractor' || usr.row_num || '@contractor.com'
        END,
    id, NOW(), NOW()
FROM tmp_group_ids, LATERAL (SELECT generate_series(1,5) AS row_num) usr;

-- Drop the temporary table
DROP TABLE tmp_group_ids;
