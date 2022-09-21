-- users
-- $2a$04$n.yc29qrd7AIFksxjfMOOeBWxJW//ZnRpPEfjf/VCCxvTRm5601Ry = Bcrypt(secret123)
INSERT INTO players 
  VALUES (1, 'Player 1', 'player1@gmail.com', '$2a$04$n.yc29qrd7AIFksxjfMOOeBWxJW//ZnRpPEfjf/VCCxvTRm5601Ry', 0, NOW(), 'system', NOW(),'system');

INSERT INTO missions (title, description, gold_bounty, created_by, updated_at, updated_by) VALUES 
  ('Mission 1', 'mission 1 description', 20, 'system',NOW(),'system'),
  ('Mission 2', 'mission 2 description', 20, 'system',NOW(),'system'),
  ('Mission 3', 'mission 3 description', 20, 'system',NOW(),'system'),
  ('Mission 4', 'mission 4 description', 20, 'system',NOW(),'system'),
  ('Mission 5', 'mission 5 description', 20, 'system',NOW(),'system'),
  ('Mission 6', 'mission 6 description', 20, 'system',NOW(),'system'),
  ('Mission 7', 'mission 7 description', 20, 'system',NOW(),'system'),
  ('Mission 8', 'mission 8 description', 20, 'system',NOW(),'system'),
  ('Mission 9', 'mission 9 description', 20, 'system',NOW(),'system'),
  ('Mission 10', 'mission 10 description', 20, 'system', NOW(),'system');
