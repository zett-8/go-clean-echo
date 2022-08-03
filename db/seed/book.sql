INSERT INTO books (id, name, author_id)
VALUES
    (1, 'Buried Giant', 1),
    (2, 'Klara and the sun', 1),
    (3, 'Never let me go', 1),
    (4, '1Q84', 2)
ON CONFLICT do nothing;
SELECT setval('authors_id_seq', nextval('authors_id_seq')-1);