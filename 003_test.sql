CREATE TEMP TABLE temp_posts_info AS
SELECT * FROM PRAGMA_TABLE_INFO('posts');

SELECT name, type FROM temp_posts_info;
