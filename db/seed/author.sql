INSERT INTO authors (id, name, country)
values
    (1, 'Kazuo Ishiguro', 'England'),
    (2, 'Haruki Murakami', 'Japan')
ON CONFLICT do nothing;
SELECT setval('authors_id_seq', 3, true);


