INSERT INTO blog.comments (created_at,updated_at,deleted_at,content,user_id,post_id) VALUES
	 ('2025-06-13 01:49:29.800','2025-06-13 01:49:29.800',NULL,'comment post3 by user3',3,3),
	 ('2025-06-13 01:49:56.948','2025-06-13 01:49:56.948',NULL,'comment post2 by user3',3,2),
	 ('2025-06-13 09:38:30.052','2025-06-13 09:38:30.052',NULL,'comment by user3',3,3);
INSERT INTO blog.posts (created_at,updated_at,deleted_at,title,content,user_id) VALUES
	 ('2025-06-13 01:47:07.263','2025-06-13 01:47:07.263',NULL,'post 1','post 1 conent',1),
	 ('2025-06-13 01:47:49.906','2025-06-13 01:47:49.906',NULL,'post 2','post 2 conent',2),
	 ('2025-06-13 01:48:35.094','2025-06-13 09:49:33.380','2025-06-13 09:49:58.326','changed by user 3111111','changed by user 3',3),
	 ('2025-06-13 09:57:07.217','2025-06-13 09:57:07.217',NULL,'post 3','post 3 conent',3);
INSERT INTO blog.users (created_at,updated_at,deleted_at,username,password,email) VALUES
	 ('2025-06-13 01:45:41.562','2025-06-13 01:45:41.562',NULL,'test1','$2a$10$4RBK.Kl5QtY5dUvXZcfeFe34OUhd2niHk1d39pNEAJ.GDdGJpcrUa','test1@test.com'),
	 ('2025-06-13 01:46:46.204','2025-06-13 01:46:46.204',NULL,'test2','$2a$10$qfIaKTBuNWnZHhT/L/7kaOFoQXdjEYMNOtQoGi6YM9bv8/o8DbsGS','test2@test.com'),
	 ('2025-06-13 01:46:53.571','2025-06-13 01:46:53.571',NULL,'test3','$2a$10$HWp5MmPwVxKLOqY3.wzp9eqONqx2wxoxvKRPp0NGFRaAxsR.6zjpm','test3@test.com');
