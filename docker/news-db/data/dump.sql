CREATE USER IF NOT EXISTS 'news_user' IDENTIFIED BY 'default_password';

GRANT ALL ON socksdb.* TO 'news_user';

CREATE TABLE IF NOT EXISTS recommender (
	news_id int NOT NULL AUTO_INCREMENT,
	title varchar(255),
	contents varchar(32768),
	PRIMARY KEY(news_id)
);


INSERT INTO `recommender` (title, contents) VALUES ('Weave special', 'Limited issue Weave socks.');
INSERT INTO `recommender` (title, contents) VALUES ('Nerd leg', 'For all those leg lovers out there. A perfect example of a swivel chair trained calf. Meticulously trained on a diet of sitting and Pina Coladas. Phwarr...');
INSERT INTO `recommender` (title, contents) VALUES ('Crossed', 'A mature sock, crossed, with an air of nonchalance.');
INSERT INTO `recommender` (title, contents) VALUES ('SuperSport XL', 'Ready for action. Engineers: be ready to smash that next bug! Be ready, with these super-action-sport-masterpieces. This particular engineer was chased away from the office with a stick.');
INSERT INTO `recommender` (title, contents) VALUES ('Holy', 'Socks fit for a Messiah. You too can experience walking in water with these special edition beauties. Each hole is lovingly proggled to leave smooth edges. The only sock approved by a higher power.');
INSERT INTO `recommender` (title, contents) VALUES ('YouTube.sock', 'We were not paid to sell this sock. It\'s just a bit geeky.');
INSERT INTO `recommender` (title, contents) VALUES ('Figueroa', 'enim officia aliqua excepteur esse deserunt quis aliquip nostrud anim');
INSERT INTO `recommender` (title, contents) VALUES ('Classic', 'Keep it simple.');
INSERT INTO `recommender` (title, contents) VALUES ('Colourful', 'proident occaecat irure et excepteur labore minim nisi amet irure');
INSERT INTO `recommender` (title, contents) VALUES ('Cat socks', 'consequat amet cupidatat minim laborum tempor elit ex consequat in');
