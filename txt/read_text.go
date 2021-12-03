v:=`
INSERT INTO statistics (id, created, player_id, action_id, amount, test, info) VALUES ('1-2',        '2021-12-03T17:40:00.111111Z', 1, 2, 100.00, true, 'del');
INSERT INTO statistics (id, created, player_id, action_id, amount, test, info) VALUES ('170-112',    '2021-12-03T17:40:00.111111Z', 170, 112, 100.00, true, 'del');
INSERT INTO statistics (id, created, player_id, action_id, amount, test, info) VALUES ('7-2',        '2021-12-03T17:40:00.111111Z', 7,    2,  10.00, true, 'del')
INSERT INTO statistics (id, created, player_id, action_id, amount, test, info) VALUES ('11170-11115','2021-12-03T18:01:00.111111Z',11170, 11115, 100.00, true, 'del')`

    
// for _, line := range strings.Split(strings.TrimRight(x, "\n"), "\n") {
//       println(line)
//    }

// https://go.dev/play/p/U9_B4TsH6M
scanner := bufio.NewScanner(strings.NewReader(v))

for scanner.Scan() {
	sql:=scanner.Text()
    Exec(sql)
 }
