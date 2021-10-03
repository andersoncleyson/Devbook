insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "Usuario_1", "usuario@gmail.com", "2a$10$RTmMT6z8ELZ9.9ybOqYasOU6P/kzbgy2vHkBH9SkFylVc2LGoqpbm"),
("Usuario 2", "Usuario_2", "usuario@gmail.com", "2a$10$RTmMT6z8ELZ9.9ybOqYasOU6P/kzbgy2vHkBH9SkFylVc2LGoqpbm"),
("Usuario 3", "Usuario_3", "usuario@gmail.com", "2a$10$RTmMT6z8ELZ9.9ybOqYasOU6P/kzbgy2vHkBH9SkFylVc2LGoqpbm");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);
