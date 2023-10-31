create table presidents(
    id serial primary key,
    name varchar,
    biography varchar,
    birth date
);

INSERT INTO  presidents (name, biography, birth) VALUES
('Luiz Inacio Lula da Silva ', 'Luiz Inacio Lula da Silva also known as Lula da Silva or simply Lula, is a Brazilian politician who is the 39th and current president of Brazil. A member of the Workers Party, he previously served as the 35th president of Brazil from 2003 to 2010.', '1945-10-27'),

('Jair Messias Bolsonaro', 'Jair Messias Bolsonaro is a Brazilian politician and retired military officer who served as the 38th president of Brazil from 2019 to 2022. He previously served in the Brazil Chamber of Deputies from 1991 to 2018.', '1955-03-21');