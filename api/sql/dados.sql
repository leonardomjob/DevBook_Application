insert into usuarios (nome, nick, email, senha) -- Senha 123 o hash é : $2a$10$xdV7RnR77KgfIMclJSiGDOQ6YgnB8Q0VVQGnqQaq1lylBFBiosGUW
values
("Leonardo", "leonardo", "meuemail@gmail.com", "$2a$10$xdV7RnR77KgfIMclJSiGDOQ6YgnB8Q0VVQGnqQaq1lylBFBiosGUW"), -- Usuario 1
("DevJornal", "dev_jornal", "devjornal@gmail.com", "$2a$10$xdV7RnR77KgfIMclJSiGDOQ6YgnB8Q0VVQGnqQaq1lylBFBiosGUW"),
("Pedro", "pedro", "pedro@gmail.com", "$2a$10$xdV7RnR77KgfIMclJSiGDOQ6YgnB8Q0VVQGnqQaq1lylBFBiosGUW");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2), -- 1 é seguido pelo 2
(3, 1),
(1, 3),
(2, 1);


insert into publicacoes(titulo, conteudo, autor_id)
values
("Hoje Fui a Paris!", "Essa é o dia que realizei meu sonho!", 1),
("Notícia Urgente!!", "Estabelecimento pega fogo!", 2),
("Olá Gente!", "Olá Rede! Primeira postagem minha no DevBook", 3);