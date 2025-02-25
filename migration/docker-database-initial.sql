create table players(
  id serial primary key,
  nome varchar,
  historia varchar
);

INSERT INTO players(nome, historia) VALUES
('Murilo', 'Revelado pelo Cruzeiro, onde chegou aos 13 anos, Murilo foi figura constante em convocações da Seleção Brasileira de base e estreou profissionalmente pelo time celeste em 2017. Integrou o elenco campeão mineiro em 2018 e 2019 e da Copa do Brasil em 2017 e 2018, somando, ao todo, 60 partidas pelo clube. Na Raposa, jogou ao lado de Rony, ainda na base, e de Mayke.'),
('Richard Ríos', 'Richard Ríos jogou futsal até os 18 anos e, após se destacar pela Seleção Colombiana em um campeonato no Rio de Janeiro, foi contratado para o Sub-20 do Flamengo, em 2018. Disputou sete partidas pela equipe profissional em 2020 e, no ano seguinte, foi emprestado ao Mazatlán, do México. Em 2022, transferiu-se para o Guarani, pelo qual disputou a Série B do Campeonato Brasileiro.');