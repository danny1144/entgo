ent describe ./ent/schema


Car:
+---------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+
|     Field     |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators |
+---------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+
| id            | int       | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |
| model         | string    | false  | false    | false    | false   | false         | false     | json:"model,omitempty"         |          0 |
| registered_at | time.Time | false  | false    | false    | false   | false         | false     | json:"registered_at,omitempty" |          0 |
+---------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+
+-------+------+---------+---------+----------+--------+----------+
| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional |
+-------+------+---------+---------+----------+--------+----------+
| owner | User | true    | cars    | M2O      | true   | true     |
+-------+------+---------+---------+----------+--------+----------+

Group:
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
| name  | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
+-------+------+---------+---------+----------+--------+----------+
| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional |
+-------+------+---------+---------+----------+--------+----------+
| users | User | false   |         | M2M      | false  | true     |
+-------+------+---------+---------+----------+--------+----------+

User:
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
| age   | int    | false  | false    | false    | false   | false         | false     | json:"age,omitempty"  |          1 |
| name  | string | false  | false    | false    | true    | false         | false     | json:"name,omitempty" |          0 |
+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
+--------+-------+---------+---------+----------+--------+----------+
|  Edge  | Type  | Inverse | BackRef | Relation | Unique | Optional |
+--------+-------+---------+---------+----------+--------+----------+
| cars   | Car   | false   |         | O2M      | false  | true     |
| groups | Group | true    | users   | M2M      | false  | true     |
+--------+-------+---------+---------+----------+--------+----------+
